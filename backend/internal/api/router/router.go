package router

import (
	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
)

// SetupRoutes はAPIルートをセットアップします
func SetupRoutes(e *echo.Echo, healthHandler *handlers.HealthHandler, productHandler *handlers.ProductHandler) {
	// APIグループ
	api := e.Group("/api")

	// ヘルスチェックエンドポイント
	api.GET("/health", healthHandler.HandleHealthCheck)

	// 商品関連エンドポイント
	products := api.Group("/products")
	products.GET("", productHandler.HandleGetProducts)
	products.GET("/categories", productHandler.HandleGetCategories)
}
