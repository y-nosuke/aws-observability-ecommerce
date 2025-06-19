package handler

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"

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
	// ファイル取得
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "failed to get uploaded file: " + err.Error(),
		})
	}

	// ファイル読み込み
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to open uploaded file: " + err.Error(),
		})
	}
	defer src.Close()

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": "failed to read uploaded file: " + err.Error(),
		})
	}

	// UseCase実行
	req := dto.NewUploadImageRequest(id, fileBytes, file.Filename)
	response, err := h.uploadProductImageUseCase.Execute(ctx.Request().Context(), req)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetProductImage は商品画像を取得する
func (h *ProductHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) error {
	// サイズパラメータの取得（デフォルトはmedium）
	size := "medium"
	if params.Size != nil {
		size = string(*params.Size)
	}

	// UseCase実行
	response, err := h.getProductImageUseCase.Execute(ctx.Request().Context(), id, size)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, openapi.ErrorResponse{
			Code:    "internal_error",
			Message: "Failed to get product image",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		})
	}

	// 画像データを返却
	return ctx.Blob(http.StatusOK, response.ContentType, response.ImageData)
}
