package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// ProductCatalogHandler は商品カタログAPIのハンドラー
type ProductCatalogHandler struct {
	reader *reader.ProductCatalogReader
	mapper *mapper.ProductCatalogMapper
}

// NewProductCatalogHandler は新しいProductCatalogHandlerを作成
func NewProductCatalogHandler() *ProductCatalogHandler {
	return &ProductCatalogHandler{
		reader: reader.NewProductCatalogReader(),
		mapper: mapper.NewProductCatalogMapper(),
	}
}

// ListProducts は商品一覧を取得する（OpenAPI仕様に準拠）
func (h *ProductCatalogHandler) ListProducts(ctx echo.Context, params openapi.ListProductsParams) error {
	return otel.WithSpan(ctx.Request().Context(), func(spanCtx context.Context, o *otel.Observer) error {
		// パラメータ変換
		readerParams := &reader.ProductListParams{}

		if params.Page != nil {
			readerParams.Page = *params.Page
		}
		if params.PageSize != nil {
			readerParams.PageSize = *params.PageSize
		}
		if params.CategoryId != nil {
			categoryID := *params.CategoryId
			readerParams.CategoryID = &categoryID
		}
		if params.Keyword != nil && *params.Keyword != "" {
			readerParams.Keyword = params.Keyword
		}

		// データ取得
		productsWithTotal, err := h.reader.FindProductsWithDetails(spanCtx, readerParams)
		if err != nil {
			return fmt.Errorf("failed to fetch products by category: %s", err)
		}

		// レスポンス変換
		response := h.mapper.ToProductListResponse(productsWithTotal.Products, productsWithTotal.Total, readerParams.Page, readerParams.PageSize)

		return ctx.JSON(http.StatusOK, response)
	})
}

// ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
func (h *ProductCatalogHandler) ListProductsByCategory(ctx echo.Context, id openapi.CategoryIdPathParam, params openapi.ListProductsByCategoryParams) error {
	return otel.WithSpan(ctx.Request().Context(), func(spanCtx context.Context, o *otel.Observer) error {
		// パラメータ変換
		readerParams := &reader.ProductListParams{}

		if params.Page != nil {
			readerParams.Page = *params.Page
		}
		if params.PageSize != nil {
			readerParams.PageSize = *params.PageSize
		}

		categoryID := id
		readerParams.CategoryID = &categoryID

		// データ取得
		productsWithTotal, err := h.reader.FindProductsWithDetails(spanCtx, readerParams)
		if err != nil {
			return fmt.Errorf("failed to fetch products by category: %s", err)
		}

		// レスポンス変換
		response := h.mapper.ToProductListResponse(productsWithTotal.Products, productsWithTotal.Total, readerParams.Page, readerParams.PageSize)

		return ctx.JSON(http.StatusOK, response)
	})
}
