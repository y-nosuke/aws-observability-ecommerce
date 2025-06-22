package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
	// Handler トレーサーを開始
	handler := observability.StartHandler(ctx.Request().Context(), "list_products")
	defer handler.FinishWithHTTPStatus(http.StatusOK)

	// HTTPリクエスト情報を記録
	handler.RecordHTTPRequest(ctx.Request().Method, ctx.Request().URL.Path, http.StatusOK)

	// パラメータ変換
	readerParams := &reader.ProductListParams{}

	if params.Page != nil {
		readerParams.Page = *params.Page
	}
	if params.PageSize != nil {
		readerParams.PageSize = *params.PageSize
	}
	if params.CategoryId != nil {
		categoryID := int(*params.CategoryId)
		readerParams.CategoryID = &categoryID
	}
	if params.Keyword != nil && *params.Keyword != "" {
		readerParams.Keyword = params.Keyword
	}

	handler.LogInfo("Product list request received",
		"page", readerParams.Page,
		"page_size", readerParams.PageSize,
		"category_id", readerParams.CategoryID,
		"keyword", readerParams.Keyword,
	)

	// データ取得
	products, total, err := h.reader.FindProductsWithDetails(handler.Context(), readerParams)
	if err != nil {
		handler.FinishWithError(err, "Failed to fetch products", http.StatusInternalServerError)
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch products", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	handler.LogInfo("Products fetched successfully",
		"total_products", total,
		"fetched_count", len(products),
		"page", readerParams.Page,
	)

	// レスポンス変換
	response := h.mapper.ToProductListResponse(products, total, readerParams.Page, readerParams.PageSize)

	return ctx.JSON(http.StatusOK, response)
}

// ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
func (h *ProductCatalogHandler) ListProductsByCategory(ctx echo.Context, id openapi.CategoryIdPathParam, params openapi.ListProductsByCategoryParams) error {
	// Handler トレーサーを開始
	handler := observability.StartHandler(ctx.Request().Context(), "list_products_by_category")
	defer handler.FinishWithHTTPStatus(http.StatusOK)

	// HTTPリクエスト情報を記録
	handler.RecordHTTPRequest(ctx.Request().Method, ctx.Request().URL.Path, http.StatusOK)

	// パラメータ変換
	readerParams := &reader.ProductListParams{}

	if params.Page != nil {
		readerParams.Page = *params.Page
	}
	if params.PageSize != nil {
		readerParams.PageSize = *params.PageSize
	}

	categoryID := int(id)
	readerParams.CategoryID = &categoryID

	handler.LogInfo("Products by category request received",
		"category_id", categoryID,
		"page", readerParams.Page,
		"page_size", readerParams.PageSize,
	)

	// データ取得
	products, total, err := h.reader.FindProductsWithDetails(handler.Context(), readerParams)
	if err != nil {
		handler.FinishWithError(err, "Failed to fetch products by category", http.StatusInternalServerError)
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch products by category", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	handler.LogInfo("Products by category fetched successfully",
		"category_id", categoryID,
		"total_products", total,
		"fetched_count", len(products),
	)

	// レスポンス変換
	response := h.mapper.ToProductListResponse(products, total, readerParams.Page, readerParams.PageSize)

	return ctx.JSON(http.StatusOK, response)
}
