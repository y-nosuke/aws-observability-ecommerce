package di

import (
	"database/sql"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/presentation/rest/handler"
	queryHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	systemHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/system/presentation/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
)

// AppContainer はアプリケーション全体の依存関係を管理
type AppContainer struct {
	// Database
	DB        *sql.DB
	DBManager *database.DBManager

	// Observability
	OTelManager *observability.OTelManager

	// AWS Services
	AWSServiceRegistry *aws.ServiceRegistry
	ClientFactory      *aws.ClientFactory
	S3ClientWrapper    *aws.S3ClientWrapper

	// Logger
	Logger logging.Logger

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
	otelManager *observability.OTelManager,
	awsServiceRegistry *aws.ServiceRegistry,
	clientFactory *aws.ClientFactory,
	s3ClientWrapper *aws.S3ClientWrapper,
	logger logging.Logger,
	productHandler *handler.ProductHandler,
	categoryListHandler *queryHandler.CategoryListHandler,
	productCatalogHandler *queryHandler.ProductCatalogHandler,
	productDetailHandler *queryHandler.ProductDetailHandler,
	healthHandler *systemHandler.HealthHandler,
) *AppContainer {
	return &AppContainer{
		DB:                    db,
		DBManager:             dbManager,
		OTelManager:           otelManager,
		AWSServiceRegistry:    awsServiceRegistry,
		ClientFactory:         clientFactory,
		S3ClientWrapper:       s3ClientWrapper,
		Logger:                logger,
		ProductHandler:        productHandler,
		CategoryListHandler:   categoryListHandler,
		ProductCatalogHandler: productCatalogHandler,
		ProductDetailHandler:  productDetailHandler,
		HealthHandler:         healthHandler,
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
	if c.OTelManager != nil {
		if err := c.OTelManager.Shutdown(); err != nil {
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

// GetLogger はロガーを取得
func (c *AppContainer) GetLogger() logging.Logger {
	return c.Logger
}

// GetDB はデータベース接続を取得
func (c *AppContainer) GetDB() *sql.DB {
	return c.DB
}

// GetDBManager はDBManagerを取得
func (c *AppContainer) GetDBManager() *database.DBManager {
	return c.DBManager
}

// GetOTelManager はOTelManagerを取得
func (c *AppContainer) GetOTelManager() *observability.OTelManager {
	return c.OTelManager
}

// GetAWSServiceRegistry はAWSサービスレジストリを取得
func (c *AppContainer) GetAWSServiceRegistry() *aws.ServiceRegistry {
	return c.AWSServiceRegistry
}
