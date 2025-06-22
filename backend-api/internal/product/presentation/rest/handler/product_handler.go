package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/usecase"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
	// Handler トレーサーを開始
	handler := observability.StartHandler(ctx.Request().Context(), "upload_product_image")
	defer handler.FinishWithHTTPStatus(http.StatusOK)

	// HTTPリクエスト情報を記録
	handler.RecordHTTPRequest(ctx.Request().Method, ctx.Request().URL.Path, http.StatusOK)
	handler.RecordRequestInfo(ctx.Request().UserAgent(), ctx.RealIP(), ctx.Request().ContentLength)

	// ファイル取得
	file, err := ctx.FormFile("image")
	if err != nil {
		handler.RecordValidationError(err, "image", "form file")
		handler.FinishWithHTTPStatus(http.StatusBadRequest, "validation_error", "missing_file")
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get uploaded file: " + err.Error(),
		})
	}

	// ファイル読み込み
	src, err := file.Open()
	if err != nil {
		handler.FinishWithError(err, "Failed to open uploaded file", http.StatusInternalServerError)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to open uploaded file: " + err.Error(),
		})
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		handler.FinishWithError(err, "Failed to read uploaded file", http.StatusInternalServerError)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to read uploaded file: " + err.Error(),
		})
	}

	handler.LogInfo("File upload processing started",
		"filename", file.Filename,
		"file_size", len(fileBytes),
		"product_id", id,
	)

	// UseCase実行
	req := dto.NewUploadImageRequest(id, fileBytes, file.Filename)
	response, err := h.uploadProductImageUseCase.Execute(handler.Context(), req)
	if err != nil {
		handler.FinishWithError(err, "UseCase execution failed", http.StatusInternalServerError)
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	handler.LogInfo("Image upload completed successfully",
		"product_id", id,
		"response_urls", len(response.URLs),
	)

	return ctx.JSON(http.StatusOK, response)
}

// GetProductImage は商品画像を取得する
func (h *ProductHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) error {
	// Handler トレーサーを開始
	handler := observability.StartHandler(ctx.Request().Context(), "get_product_image")
	defer handler.FinishWithHTTPStatus(http.StatusOK)

	// HTTPリクエスト情報を記録
	handler.RecordHTTPRequest(ctx.Request().Method, ctx.Request().URL.Path, http.StatusOK)

	// サイズパラメータの取得（デフォルトはmedium）
	size := "medium"
	if params.Size != nil {
		size = string(*params.Size)
	}

	handler.LogInfo("Get product image requested",
		"product_id", id,
		"requested_size", size,
	)

	// UseCase実行
	response, err := h.getProductImageUseCase.Execute(handler.Context(), id, size)
	if err != nil {
		handler.FinishWithError(err, "Failed to get product image", http.StatusInternalServerError)
		return ctx.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Code:    "internal_error",
			Message: "Failed to get product image",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	handler.LogInfo("Product image retrieved successfully",
		"product_id", id,
		"content_type", response.ContentType,
		"image_size_bytes", len(response.ImageData),
	)

	// 画像データを返却
	return ctx.Blob(http.StatusOK, response.ContentType, response.ImageData)
}
