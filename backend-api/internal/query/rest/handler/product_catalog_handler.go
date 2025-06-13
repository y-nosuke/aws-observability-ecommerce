package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

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
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	requestCtx, span := tracer.Start(ctx.Request().Context(), "handler.list_products", trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", "product_catalog"),
		attribute.String("app.operation", "list_products"),
		attribute.String("http.method", ctx.Request().Method),
		attribute.String("http.route", ctx.Path()),
	))
	defer span.End()

	// パラメータ変換
	readerParams := &reader.ProductListParams{}

	if params.Page != nil {
		readerParams.Page = *params.Page
		span.SetAttributes(attribute.Int("app.request.page", *params.Page))
	}
	if params.PageSize != nil {
		readerParams.PageSize = *params.PageSize
		span.SetAttributes(attribute.Int("app.request.page_size", *params.PageSize))
	}
	if params.CategoryId != nil {
		categoryID := int(*params.CategoryId)
		readerParams.CategoryID = &categoryID
		span.SetAttributes(attribute.Int("app.request.category_id", categoryID))
	}
	if params.Keyword != nil && *params.Keyword != "" {
		readerParams.Keyword = params.Keyword
		span.SetAttributes(attribute.String("app.request.keyword", *params.Keyword))
	}

	// 子スパンでデータ取得
	dataCtx, dataSpan := tracer.Start(requestCtx, "handler.fetch_products_data")
	products, total, err := h.reader.FindProductsWithDetails(dataCtx, readerParams)
	if err != nil {
		dataSpan.RecordError(err)
		dataSpan.SetStatus(codes.Error, err.Error())
		dataSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch products", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	dataSpan.SetAttributes(
		attribute.Int("app.products_fetched", len(products)),
		attribute.Int64("app.total_count", total),
	)
	dataSpan.End()

	// 子スパンでレスポンス変換
	_, mapSpan := tracer.Start(requestCtx, "handler.map_products_response")
	response := h.mapper.ToProductListResponse(products, total, readerParams.Page, readerParams.PageSize)
	mapSpan.SetAttributes(
		attribute.Int("app.response_items", len(response.Items)),
		attribute.Int("app.response_page", response.Page),
		attribute.Int("app.response_page_size", response.PageSize),
		attribute.Int("app.response_total", response.Total),
	)
	mapSpan.End()

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("http.response.status_code", http.StatusOK),
		attribute.Int("app.products_returned", len(response.Items)),
		attribute.Int("app.total_products", response.Total),
		attribute.Int("app.current_page", response.Page),
		attribute.Bool("app.has_more_pages", response.Page*response.PageSize < int(response.Total)),
	)

	return ctx.JSON(http.StatusOK, response)
}

// ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
func (h *ProductCatalogHandler) ListProductsByCategory(ctx echo.Context, id openapi.CategoryIdPathParam, params openapi.ListProductsByCategoryParams) error {
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	requestCtx, span := tracer.Start(ctx.Request().Context(), "handler.list_products_by_category", trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", "product_catalog"),
		attribute.String("app.operation", "list_products_by_category"),
		attribute.String("http.method", ctx.Request().Method),
		attribute.String("http.route", ctx.Path()),
		attribute.Int("app.category_id", int(id)),
	))
	defer span.End()

	// パラメータ変換
	readerParams := &reader.ProductListParams{}

	if params.Page != nil {
		readerParams.Page = *params.Page
		span.SetAttributes(attribute.Int("app.request.page", *params.Page))
	}
	if params.PageSize != nil {
		readerParams.PageSize = *params.PageSize
		span.SetAttributes(attribute.Int("app.request.page_size", *params.PageSize))
	}

	categoryID := int(id)
	readerParams.CategoryID = &categoryID

	// 子スパンでデータ取得
	dataCtx, dataSpan := tracer.Start(requestCtx, "handler.fetch_products_by_category_data")
	products, total, err := h.reader.FindProductsWithDetails(dataCtx, readerParams)
	if err != nil {
		dataSpan.RecordError(err)
		dataSpan.SetStatus(codes.Error, err.Error())
		dataSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch products by category", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}
	dataSpan.SetAttributes(
		attribute.Int("app.products_fetched", len(products)),
		attribute.Int64("app.total_count", total),
	)
	dataSpan.End()

	// 子スパンでレスポンス変換
	_, mapSpan := tracer.Start(requestCtx, "handler.map_products_by_category_response")
	response := h.mapper.ToProductListResponse(products, total, readerParams.Page, readerParams.PageSize)
	mapSpan.SetAttributes(
		attribute.Int("app.response_items", len(response.Items)),
		attribute.Int("app.response_page", response.Page),
		attribute.Int("app.response_page_size", response.PageSize),
		attribute.Int("app.response_total", response.Total),
	)
	mapSpan.End()

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("http.response.status_code", http.StatusOK),
		attribute.Int("app.products_returned", len(response.Items)),
		attribute.Int("app.total_products", response.Total),
		attribute.Int("app.current_page", response.Page),
		attribute.Bool("app.has_more_pages", response.Page*response.PageSize < int(response.Total)),
	)

	return ctx.JSON(http.StatusOK, response)
}
