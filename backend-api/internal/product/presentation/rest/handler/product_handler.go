package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"

	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/usecase"
)

// ProductHandler は商品APIのハンドラー
type ProductHandler struct {
	uploadProductImageUseCase *usecase.UploadProductImageUseCase
	getProductImageUseCase    *usecase.GetProductImageUseCase
}

// NewProductHandler は新しいProductHandlerを作成する
func NewProductHandler(
	uploadProductImageUseCase *usecase.UploadProductImageUseCase,
	getProductImageUseCase *usecase.GetProductImageUseCase,
) *ProductHandler {
	return &ProductHandler{
		uploadProductImageUseCase: uploadProductImageUseCase,
		getProductImageUseCase:    getProductImageUseCase,
	}
}

// UploadProductImage は商品画像をアップロードする
func (h *ProductHandler) UploadProductImage(ctx echo.Context, id openapi.ProductIdParam) error {
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	requestCtx, span := tracer.Start(ctx.Request().Context(), "handler.upload_product_image", trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", "product"),
		attribute.String("app.operation", "upload_product_image"),
		attribute.String("http.method", ctx.Request().Method),
		attribute.String("http.route", ctx.Path()),
		attribute.Int64("app.product_id", id),
	))
	defer span.End()

	// 操作開始ログ
	completeOp := logger.StartOperation(requestCtx, "upload_product_image",
		"product_id", id,
		"operation_type", "image_upload",
		"layer", "handler")

	// 子スパンでファイル取得処理
	_, fileSpan := tracer.Start(requestCtx, "handler.get_form_file")
	file, err := ctx.FormFile("image")
	if err != nil {
		fileSpan.RecordError(err)
		fileSpan.SetStatus(codes.Error, err.Error())
		fileSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusBadRequest))

		// エラーログ
		logger.WithError(requestCtx, "アップロードファイルの取得に失敗", err,
			"product_id", id,
			"layer", "handler",
			"operation", "get_form_file")
		completeOp(false, "error_type", "form_file_error")
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get uploaded file: " + err.Error(),
		})
	}
	fileSpan.SetAttributes(
		attribute.String("app.file_name", file.Filename),
		attribute.Int64("app.file_size", file.Size),
		attribute.String("app.content_type", file.Header.Get("Content-Type")),
	)
	fileSpan.End()

	// ファイル情報をログ
	logger.Info(requestCtx, "ファイル情報を取得",
		"product_id", id,
		"file_name", file.Filename,
		"file_size_bytes", file.Size,
		"content_type", file.Header.Get("Content-Type"),
		"layer", "handler")

	// 子スパンでファイル読み込み処理
	_, readSpan := tracer.Start(requestCtx, "handler.read_file_content")
	src, err := file.Open()
	if err != nil {
		readSpan.RecordError(err)
		readSpan.SetStatus(codes.Error, err.Error())
		readSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		logger.WithError(requestCtx, "アップロードファイルのオープンに失敗", err,
			"product_id", id,
			"file_name", file.Filename,
			"layer", "handler")
		completeOp(false, "error_type", "file_open_error")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to open uploaded file: " + err.Error(),
		})
	}
	defer src.Close()

	// ファイル内容を読み込む
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		readSpan.RecordError(err)
		readSpan.SetStatus(codes.Error, err.Error())
		readSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		logger.WithError(requestCtx, "ファイル内容の読み込みに失敗", err,
			"product_id", id,
			"file_name", file.Filename,
			"layer", "handler")
		completeOp(false, "error_type", "file_read_error")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to read uploaded file: " + err.Error(),
		})
	}
	readSpan.SetAttributes(attribute.Int("app.file_bytes_read", len(fileBytes)))
	readSpan.End()

	// 子スパンでリクエストDTO作成
	_, dtoSpan := tracer.Start(requestCtx, "handler.create_upload_request_dto")
	req := dto.NewUploadImageRequest(id, fileBytes, file.Filename)
	dtoSpan.End()

	// 子スパンでユースケース実行
	usecaseCtx, usecaseSpan := tracer.Start(requestCtx, "handler.execute_upload_usecase")
	response, err := h.uploadProductImageUseCase.Execute(usecaseCtx, req)
	if err != nil {
		usecaseSpan.RecordError(err)
		usecaseSpan.SetStatus(codes.Error, err.Error())
		usecaseSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		logger.WithError(requestCtx, "画像アップロード処理に失敗", err,
			"product_id", id,
			"file_name", file.Filename,
			"layer", "handler")
		completeOp(false, "error_type", "usecase_error")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	usecaseSpan.SetAttributes(
		attribute.String("app.s3_key", response.S3Key),
		attribute.Int("app.generated_urls", len(response.URLs)),
	)
	usecaseSpan.End()

	// 成功時の追加データ
	// URLsマップから適切なURLを取得（存在する場合）
	var imageURL string
	if len(response.URLs) > 0 {
		// 最初に利用可能なURLを使用（medium, large, originalの順で取得を試行）
		if url, exists := response.URLs["medium"]; exists {
			imageURL = url
		} else if url, exists := response.URLs["original"]; exists {
			imageURL = url
		} else {
			// いずれかのURLを使用
			for _, url := range response.URLs {
				imageURL = url
				break
			}
		}
	}

	// 成功ログ
	logger.Info(requestCtx, "画像アップロードが完了",
		"product_id", id,
		"file_name", file.Filename,
		"file_size_bytes", file.Size,
		"s3_key", response.S3Key,
		"image_url", imageURL,
		"layer", "handler")

	// 操作完了を記録
	completeOp(true,
		"s3_key", response.S3Key,
		"image_url", imageURL,
		"content_type", file.Header.Get("Content-Type"))

	// ビジネスイベントとして記録
	logger.LogBusinessEvent(requestCtx, "product_image_uploaded", "product", fmt.Sprint(id),
		"image_url", imageURL,
		"s3_key", response.S3Key,
		"filename", response.Filename,
		"file_name", file.Filename,
		"file_size", file.Size,
		"uploaded_by", "admin") // 実際は認証情報から取得

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("http.response.status_code", http.StatusOK),
		attribute.String("app.s3_key", response.S3Key),
		attribute.String("app.image_url", imageURL),
		attribute.String("app.file_name", file.Filename),
		attribute.Int64("app.file_size", file.Size),
		attribute.String("app.content_type", file.Header.Get("Content-Type")),
		attribute.Int("app.generated_urls", len(response.URLs)),
	)

	return ctx.JSON(http.StatusOK, response)
}

// GetProductImage は商品画像を取得する
func (h *ProductHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) error {
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	requestCtx, span := tracer.Start(ctx.Request().Context(), "handler.get_product_image", trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", "product"),
		attribute.String("app.operation", "get_product_image"),
		attribute.String("http.method", ctx.Request().Method),
		attribute.String("http.route", ctx.Path()),
		attribute.Int64("app.product_id", id),
	))
	defer span.End()

	// 操作開始ログ
	completeOp := logger.StartOperation(requestCtx, "get_product_image",
		"product_id", id,
		"layer", "handler")

	// サイズパラメータの取得（デフォルトはmedium）
	size := "medium"
	if params.Size != nil {
		size = string(*params.Size)
	}
	span.SetAttributes(attribute.String("app.image_size", size))

	// リクエスト情報をログ
	logger.Info(requestCtx, "画像取得リクエストを開始",
		"product_id", id,
		"requested_size", size,
		"layer", "handler")

	// 子スパンでユースケース実行
	usecaseCtx, usecaseSpan := tracer.Start(requestCtx, "handler.execute_get_image_usecase")
	response, err := h.getProductImageUseCase.Execute(usecaseCtx, id, size)
	if err != nil {
		usecaseSpan.RecordError(err)
		usecaseSpan.SetStatus(codes.Error, err.Error())
		usecaseSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		// エラーログ
		logger.WithError(requestCtx, "画像取得処理に失敗", err,
			"product_id", id,
			"requested_size", size,
			"layer", "handler")
		completeOp(false, "error_type", "usecase_error", "requested_size", size)
		return ctx.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Code:    "internal_error",
			Message: "Failed to get product image",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		})
	}
	usecaseSpan.SetAttributes(
		attribute.String("app.content_type", response.ContentType),
		attribute.Int("app.image_size_bytes", len(response.ImageData)),
	)
	usecaseSpan.End()

	// 成功ログ
	logger.Info(requestCtx, "画像取得が完了",
		"product_id", id,
		"requested_size", size,
		"content_type", response.ContentType,
		"image_size_bytes", len(response.ImageData),
		"layer", "handler")

	// 操作完了を記録
	completeOp(true,
		"content_type", response.ContentType,
		"image_size_bytes", len(response.ImageData),
		"requested_size", size)

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("http.response.status_code", http.StatusOK),
		attribute.String("app.content_type", response.ContentType),
		attribute.Int("app.image_size_bytes", len(response.ImageData)),
		attribute.String("app.requested_size", size),
	)

	// 画像データを返却
	return ctx.Blob(http.StatusOK, response.ContentType, response.ImageData)
}
