package handler

import (
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/di"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/presentation/rest/handler"
	queryHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/handler"
	systemHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/system/presentation/rest/handler"
)

// Handler は全てのAPIハンドラーを統合する構造体
// oapi-codegenのServerInterfaceを実装する
type Handler struct {
	*systemHandler.HealthHandler
	*queryHandler.CategoryListHandler
	*queryHandler.ProductCatalogHandler
	*queryHandler.ProductDetailHandler
	*handler.ProductHandler
}

// NewHandler はDIコンテナから各ハンドラーを取得してHandlerを作成
func NewHandler(container *di.AppContainer) (*Handler, error) {
	return &Handler{
		HealthHandler:         container.GetHealthHandler(),
		CategoryListHandler:   container.GetCategoryListHandler(),
		ProductCatalogHandler: container.GetProductCatalogHandler(),
		ProductDetailHandler:  container.GetProductDetailHandler(),
		ProductHandler:        container.GetProductHandler(),
	}, nil
}
