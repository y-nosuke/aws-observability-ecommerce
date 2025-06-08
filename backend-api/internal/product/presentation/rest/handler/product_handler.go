package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"

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
	// 操作開始ログ
	completeOp := logging.StartOperation(ctx.Request().Context(), "upload_product_image",
		"product_id", id,
		"operation_type", "image_upload",
		"layer", "handler")

	// フォームからファイルを取得
	file, err := ctx.FormFile("image")
	if err != nil {
		// エラーログ
		logging.WithError(ctx.Request().Context(), "アップロードファイルの取得に失敗", err,
			"product_id", id,
			"layer", "handler",
			"operation", "get_form_file")
		completeOp(false, "error_type", "form_file_error")
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get uploaded file: " + err.Error(),
		})
	}

	// ファイル情報をログ
	logging.Info(ctx.Request().Context(), "ファイル情報を取得",
		"product_id", id,
		"file_name", file.Filename,
		"file_size_bytes", file.Size,
		"content_type", file.Header.Get("Content-Type"),
		"layer", "handler")

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		logging.WithError(ctx.Request().Context(), "アップロードファイルのオープンに失敗", err,
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
		logging.WithError(ctx.Request().Context(), "ファイル内容の読み込みに失敗", err,
			"product_id", id,
			"file_name", file.Filename,
			"layer", "handler")
		completeOp(false, "error_type", "file_read_error")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to read uploaded file: " + err.Error(),
		})
	}

	// リクエストDTOを作成
	req := dto.NewUploadImageRequest(id, fileBytes, file.Filename)

	// ユースケースを実行
	response, err := h.uploadProductImageUseCase.Execute(ctx.Request().Context(), req)
	if err != nil {
		logging.WithError(ctx.Request().Context(), "画像アップロード処理に失敗", err,
			"product_id", id,
			"file_name", file.Filename,
			"layer", "handler")
		completeOp(false, "error_type", "usecase_error")
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

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
	logging.Info(ctx.Request().Context(), "画像アップロードが完了",
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
	logging.LogBusinessEvent(ctx.Request().Context(), "product_image_uploaded", "product", fmt.Sprint(id),
		"image_url", imageURL,
		"s3_key", response.S3Key,
		"filename", response.Filename,
		"file_name", file.Filename,
		"file_size", file.Size,
		"uploaded_by", "admin") // 実際は認証情報から取得

	return ctx.JSON(http.StatusOK, response)
}

// GetProductImage は商品画像を取得する
func (h *ProductHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) error {
	// 操作開始ログ
	completeOp := logging.StartOperation(ctx.Request().Context(), "get_product_image",
		"product_id", id,
		"layer", "handler")

	// サイズパラメータの取得（デフォルトはmedium）
	size := "medium"
	if params.Size != nil {
		size = string(*params.Size)
	}

	// リクエスト情報をログ
	logging.Info(ctx.Request().Context(), "画像取得リクエストを開始",
		"product_id", id,
		"requested_size", size,
		"layer", "handler")

	// ユースケースを実行
	response, err := h.getProductImageUseCase.Execute(ctx.Request().Context(), id, size)
	if err != nil {
		// エラーログ
		logging.WithError(ctx.Request().Context(), "画像取得処理に失敗", err,
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

	// 成功ログ
	logging.Info(ctx.Request().Context(), "画像取得が完了",
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

	// 画像データを返却
	return ctx.Blob(http.StatusOK, response.ContentType, response.ImageData)
}
