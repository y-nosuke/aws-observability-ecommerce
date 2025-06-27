package handler

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/usecase"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
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
func (h *ProductHandler) UploadProductImage(ctx echo.Context, id openapi.ProductIdParam) (err error) {
	spanCtx, o := otel.Start(ctx.Request().Context())
	defer func() {
		o.End(err)
	}()

	// ファイル取得
	file, err := ctx.FormFile("image")
	if err != nil {
		return fmt.Errorf("failed to get uploaded file: %w", err)
	}

	// ファイル読み込み
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("failed to open uploaded file: %w", err)
	}
	defer func(src multipart.File) {
		if closeErr := src.Close(); closeErr != nil {
			err = fmt.Errorf("original error: %v, failed to close uploaded file: %w", err, closeErr)
			return
		}
	}(src)

	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return fmt.Errorf("failed to read uploaded file: %w", err)
	}

	// UseCase実行
	req := dto.NewUploadImageRequest(id, fileBytes, file.Filename)
	response, err := h.uploadProductImageUseCase.Execute(spanCtx, req)
	if err != nil {
		return fmt.Errorf("failed to upload product image: %w", err)
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetProductImage は商品画像を取得する
func (h *ProductHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) (err error) {
	spanCtx, o := otel.Start(ctx.Request().Context())
	defer func() {
		o.End(err)
	}()

	// サイズパラメータの取得（デフォルトはmedium）
	size := service.Medium
	if params.Size != nil {
		size = service.SizeType(*params.Size)
	}

	// UseCase実行
	response, err := h.getProductImageUseCase.Execute(spanCtx, id, size)
	if err != nil {
		return fmt.Errorf("failed to get product image: %w", err)
	}

	// 画像データを返却
	return ctx.Blob(http.StatusOK, response.ContentType, response.ImageData)
}
