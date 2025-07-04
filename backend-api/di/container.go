package di

import (
	"log/slog"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/presentation/rest/handler"
	queryHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/handler"
	systemHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/system/presentation/rest/handler"
)

// AppContainer はアプリケーション全体の依存関係を管理
type AppContainer struct {
	// Observability
	Logger *slog.Logger
	Tracer trace.Tracer
	Meter  metric.Meter

	// Handlers
	ProductHandler        *handler.ProductHandler
	CategoryListHandler   *queryHandler.CategoryListHandler
	ProductCatalogHandler *queryHandler.ProductCatalogHandler
	ProductDetailHandler  *queryHandler.ProductDetailHandler
	HealthHandler         *systemHandler.HealthHandler
}

// NewAppContainer は新しいAppContainerを作成
func NewAppContainer(
	logger *slog.Logger,
	tracer trace.Tracer,
	meter metric.Meter,
	productHandler *handler.ProductHandler,
	categoryListHandler *queryHandler.CategoryListHandler,
	productCatalogHandler *queryHandler.ProductCatalogHandler,
	productDetailHandler *queryHandler.ProductDetailHandler,
	healthHandler *systemHandler.HealthHandler,
) *AppContainer {
	return &AppContainer{
		Logger:                logger,
		Tracer:                tracer,
		Meter:                 meter,
		ProductHandler:        productHandler,
		CategoryListHandler:   categoryListHandler,
		ProductCatalogHandler: productCatalogHandler,
		ProductDetailHandler:  productDetailHandler,
		HealthHandler:         healthHandler,
	}
}

// GetProductHandler は商品ハンドラーを取得
func (c *AppContainer) GetProductHandler() *handler.ProductHandler {
	return c.ProductHandler
}

// GetCategoryListHandler はカテゴリリストハンドラーを取得
func (c *AppContainer) GetCategoryListHandler() *queryHandler.CategoryListHandler {
	return c.CategoryListHandler
}

// GetProductCatalogHandler は商品カタログハンドラーを取得
func (c *AppContainer) GetProductCatalogHandler() *queryHandler.ProductCatalogHandler {
	return c.ProductCatalogHandler
}

// GetProductDetailHandler は商品詳細ハンドラーを取得
func (c *AppContainer) GetProductDetailHandler() *queryHandler.ProductDetailHandler {
	return c.ProductDetailHandler
}

// GetHealthHandler はヘルスハンドラーを取得
func (c *AppContainer) GetHealthHandler() *systemHandler.HealthHandler {
	return c.HealthHandler
}
