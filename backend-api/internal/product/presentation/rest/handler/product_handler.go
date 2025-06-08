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
	logHelper                 *logging.LogHelper
}

// NewProductHandler は新しいProductHandlerを作成する
func NewProductHandler(
	uploadProductImageUseCase *usecase.UploadProductImageUseCase,
	getProductImageUseCase *usecase.GetProductImageUseCase,
	logger logging.Logger,
) *ProductHandler {
	return &ProductHandler{
		uploadProductImageUseCase: uploadProductImageUseCase,
		getProductImageUseCase:    getProductImageUseCase,
		logHelper:                 logging.NewLogHelper(logger),
	}
}

// UploadProductImage は商品画像をアップロードする
func (h *ProductHandler) UploadProductImage(ctx echo.Context, id openapi.ProductIdParam) error {
	// 操作ログの開始
	opLogger := h.logHelper.StartOperation(ctx.Request().Context(), "upload_product_image", "product_management").
		WithEntity("product", fmt.Sprint(id)).
		WithAction("upload", "admin_ui").
		WithData("operation_type", "image_upload")

	// フォームからファイルを取得
	file, err := ctx.FormFile("image")
	if err != nil {
		opLogger.Fail(ctx.Request().Context(), err)
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get uploaded file: " + err.Error(),
		})
	}

	// ファイル情報をログに追加
	opLogger.WithData("file_name", file.Filename).
		WithData("file_size_bytes", file.Size).
		WithData("content_type", file.Header.Get("Content-Type"))

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		opLogger.Fail(ctx.Request().Context(), err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to open uploaded file: " + err.Error(),
		})
	}
	defer src.Close()

	// ファイル内容を読み込む
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		opLogger.Fail(ctx.Request().Context(), err)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to read uploaded file: " + err.Error(),
		})
	}

	// リクエストDTOを作成
	req := dto.NewUploadImageRequest(id, fileBytes, file.Filename)

	// ユースケースを実行
	response, err := h.uploadProductImageUseCase.Execute(ctx.Request().Context(), req)
	if err != nil {
		opLogger.Fail(ctx.Request().Context(), err)
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

	opLogger.WithData("upload_result", map[string]interface{}{
		"image_url":  imageURL,
		"s3_key":     response.S3Key,
		"filename":   response.Filename,
		"product_id": response.ProductID,
		"success":    true,
	})

	// 操作完了をログ
	opLogger.Complete(ctx.Request().Context())

	// ビジネスイベントとしても記録
	h.logHelper.LogBusinessEvent(ctx.Request().Context(), "product_image_uploaded", "product", fmt.Sprint(id), map[string]interface{}{
		"image_url":   imageURL,
		"s3_key":      response.S3Key,
		"filename":    response.Filename,
		"file_name":   file.Filename,
		"file_size":   file.Size,
		"uploaded_by": "admin", // 実際は認証情報から取得
	})

	return ctx.JSON(http.StatusOK, response)
}

// GetProductImage は商品画像を取得する
func (h *ProductHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) error {
	// 操作ログの開始
	opLogger := h.logHelper.StartOperation(ctx.Request().Context(), "get_product_image", "product_management").
		WithEntity("product", fmt.Sprint(id)).
		WithAction("view", "customer_ui")

	// サイズパラメータの取得（デフォルトはmedium）
	size := "medium"
	if params.Size != nil {
		size = string(*params.Size)
	}

	opLogger.WithData("requested_size", size)

	// ユースケースを実行
	response, err := h.getProductImageUseCase.Execute(ctx.Request().Context(), id, size)
	if err != nil {
		opLogger.Fail(ctx.Request().Context(), err)
		return ctx.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Code:    "internal_error",
			Message: "Failed to get product image",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	// 成功時の追加データ
	opLogger.WithData("response_info", map[string]interface{}{
		"content_type":  response.ContentType,
		"image_size":    len(response.ImageData),
		"cache_enabled": true, // キャッシュ実装後に実際の値を設定
	}).WithPerformanceData("response_size_bytes", len(response.ImageData))

	// 操作完了をログ
	opLogger.Complete(ctx.Request().Context())

	// 画像データを返却
	return ctx.Blob(http.StatusOK, response.ContentType, response.ImageData)
}
