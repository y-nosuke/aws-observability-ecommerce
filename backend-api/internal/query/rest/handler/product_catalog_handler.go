package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// ProductCatalogHandler は商品カタログAPIのハンドラー
type ProductCatalogHandler struct {
	reader *reader.ProductCatalogReader
	mapper *mapper.ProductCatalogMapper
}

// NewProductCatalogHandler は新しいProductCatalogHandlerを作成
func NewProductCatalogHandler(db boil.ContextExecutor) *ProductCatalogHandler {
	return &ProductCatalogHandler{
		reader: reader.NewProductCatalogReader(db),
		mapper: mapper.NewProductCatalogMapper(),
	}
}

// ListProducts は商品一覧を取得する（OpenAPI仕様に準拠）
func (h *ProductCatalogHandler) ListProducts(ctx echo.Context, params openapi.ListProductsParams) error {
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
	products, total, err := h.reader.FindProductsWithDetails(ctx.Request().Context(), readerParams)
	if err != nil {
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch products", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンス変換
	response := h.mapper.ToProductListResponse(products, total, readerParams.Page, readerParams.PageSize)

	return ctx.JSON(http.StatusOK, response)
}

// ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
func (h *ProductCatalogHandler) ListProductsByCategory(ctx echo.Context, id openapi.CategoryIdPathParam, params openapi.ListProductsByCategoryParams) error {
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
	products, total, err := h.reader.FindProductsWithDetails(ctx.Request().Context(), readerParams)
	if err != nil {
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch products by category", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンス変換
	response := h.mapper.ToProductListResponse(products, total, readerParams.Page, readerParams.PageSize)

	return ctx.JSON(http.StatusOK, response)
}
