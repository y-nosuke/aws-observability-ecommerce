package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
)

// ProductDetailHandler は商品詳細APIのハンドラー
type ProductDetailHandler struct {
	reader *reader.ProductDetailReader
	mapper *mapper.ProductDetailMapper
}

// NewProductDetailHandler は新しいProductDetailHandlerを作成
func NewProductDetailHandler(db boil.ContextExecutor) *ProductDetailHandler {
	return &ProductDetailHandler{
		reader: reader.NewProductDetailReader(db),
		mapper: mapper.NewProductDetailMapper(),
	}
}

// GetProductById は指定されたIDの商品を取得する
func (h *ProductDetailHandler) GetProductById(ctx echo.Context, id openapi.ProductIdParam) error {
	// トレーシングスパンを開始
	requestCtx, span := tracer.Start(ctx.Request().Context(), "handler.get_product_by_id", trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", "product"),
		attribute.String("app.operation", "get_product_by_id"),
		attribute.String("http.method", ctx.Request().Method),
		attribute.String("http.route", ctx.Path()),
		attribute.Int("app.product_id", int(id)),
	))
	defer span.End()

	// IDの整合性チェック
	if id <= 0 {
		span.SetAttributes(
			attribute.Bool("app.validation_failed", true),
			attribute.String("app.validation_error", "invalid_product_id"),
			attribute.Int("http.response.status_code", http.StatusBadRequest),
		)
		errorResponse := h.mapper.PresentInvalidParameter("Invalid product ID")
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	// 子スパンで商品詳細取得
	dataCtx, dataSpan := tracer.Start(requestCtx, "handler.fetch_product_detail_data")
	product, err := h.reader.FindProductByID(dataCtx, int(id))
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if err.Error() == "product not found: "+strconv.Itoa(int(id)) {
			dataSpan.SetAttributes(
				attribute.Bool("app.product_not_found", true),
				attribute.String("app.error_type", "not_found"),
			)
			dataSpan.SetStatus(codes.Error, "product not found")
			dataSpan.End()
			span.SetAttributes(
				attribute.Bool("app.product_not_found", true),
				attribute.Int("http.response.status_code", http.StatusNotFound),
			)
			span.SetStatus(codes.Error, "product not found")

			errorResponse := h.mapper.PresentProductNotFound("Product not found", int(id))
			return ctx.JSON(http.StatusNotFound, errorResponse)
		}

		// その他のエラー
		dataSpan.RecordError(err)
		dataSpan.SetStatus(codes.Error, err.Error())
		dataSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(attribute.Int("http.response.status_code", http.StatusInternalServerError))

		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch product details", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// 商品情報をスパンに記録
	dataSpan.SetAttributes(
		attribute.String("app.product_name", product.Name),
		attribute.String("app.product_sku", product.Sku),
		attribute.Float64("app.product_price", product.Price.InexactFloat64()),
	)
	dataSpan.End()

	// 子スパンでレスポンス構築
	_, mapSpan := tracer.Start(requestCtx, "handler.map_product_detail_response")
	response := h.mapper.ToProductResponse(product)
	mapSpan.SetAttributes(attribute.Bool("app.response_mapped", true))
	mapSpan.End()

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("http.response.status_code", http.StatusOK),
		attribute.Bool("app.product_found", true),
		attribute.String("app.product_name", product.Name),
		attribute.String("app.product_sku", product.Sku),
		attribute.Float64("app.product_price", product.Price.InexactFloat64()),
	)

	return ctx.JSON(http.StatusOK, response)
}
