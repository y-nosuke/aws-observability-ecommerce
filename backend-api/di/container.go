package di

import (
	"database/sql"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/presentation/rest/handler"
	queryHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	systemHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/system/presentation/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
)

// AppContainer はアプリケーション全体の依存関係を管理
type AppContainer struct {
	// Database
	DB        *sql.DB
	DBManager *database.DBManager

	// Observability
	ProviderFactory                observability.ProviderFactory
	GlobalObservabilityInitializer *observability.GlobalObservabilityInitializer

	// AWS Services
	AWSServiceRegistry *aws.ServiceRegistry
	ClientFactory      *aws.ClientFactory
	S3ClientWrapper    *aws.S3ClientWrapper

	// Handlers
	ProductHandler        *handler.ProductHandler
	CategoryListHandler   *queryHandler.CategoryListHandler
	ProductCatalogHandler *queryHandler.ProductCatalogHandler
	ProductDetailHandler  *queryHandler.ProductDetailHandler
	HealthHandler         *systemHandler.HealthHandler
}

// NewAppContainer は新しいAppContainerを作成
func NewAppContainer(
	db *sql.DB,
	dbManager *database.DBManager,
	providerFactory observability.ProviderFactory,
	globalObservabilityInitializer *observability.GlobalObservabilityInitializer,
	awsServiceRegistry *aws.ServiceRegistry,
	clientFactory *aws.ClientFactory,
	s3ClientWrapper *aws.S3ClientWrapper,
	productHandler *handler.ProductHandler,
	categoryListHandler *queryHandler.CategoryListHandler,
	productCatalogHandler *queryHandler.ProductCatalogHandler,
	productDetailHandler *queryHandler.ProductDetailHandler,
	healthHandler *systemHandler.HealthHandler,
) *AppContainer {
	return &AppContainer{
		DB:                             db,
		DBManager:                      dbManager,
		ProviderFactory:                providerFactory,
		GlobalObservabilityInitializer: globalObservabilityInitializer,
		AWSServiceRegistry:             awsServiceRegistry,
		ClientFactory:                  clientFactory,
		S3ClientWrapper:                s3ClientWrapper,
		ProductHandler:                 productHandler,
		CategoryListHandler:            categoryListHandler,
		ProductCatalogHandler:          productCatalogHandler,
		ProductDetailHandler:           productDetailHandler,
		HealthHandler:                  healthHandler,
	}
}

// Cleanup はリソースをクリーンアップします
func (c *AppContainer) Cleanup() error {
	// データベース接続を閉じる
	if c.DBManager != nil {
		if err := c.DBManager.Close(); err != nil {
			return err
		}
	}

	// OpenTelemetryをシャットダウン
	if c.ProviderFactory != nil {
		if err := c.ProviderFactory.Shutdown(); err != nil {
			return err
		}
	}

	return nil
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

// GetDB はデータベース接続を取得
func (c *AppContainer) GetDB() *sql.DB {
	return c.DB
}

// GetDBManager はDBManagerを取得
func (c *AppContainer) GetDBManager() *database.DBManager {
	return c.DBManager
}

// GetProviderFactory はProviderFactoryを取得
func (c *AppContainer) GetProviderFactory() observability.ProviderFactory {
	return c.ProviderFactory
}

// GetAWSServiceRegistry はAWSサービスレジストリを取得
func (c *AppContainer) GetAWSServiceRegistry() *aws.ServiceRegistry {
	return c.AWSServiceRegistry
}

// GetGlobalObservabilityInitializer はグローバルオブザーバビリティ初期化サービスを取得
func (c *AppContainer) GetGlobalObservabilityInitializer() *observability.GlobalObservabilityInitializer {
	return c.GlobalObservabilityInitializer
}
