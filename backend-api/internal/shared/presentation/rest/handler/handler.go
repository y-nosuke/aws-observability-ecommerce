package handler

import (
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/infrastructure/di"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/presentation/rest/handler"
	queryHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	systemHandler "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/system/presentation/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
)

// Handler は全てのAPIハンドラーを統合する構造体
// oapi-codegenのServerInterfaceを実装する
type Handler struct {
	*systemHandler.HealthHandler
	*queryHandler.CategoryListHandler
	*queryHandler.ProductCatalogHandler
	*queryHandler.ProductDetailHandler
	*handler.ProductHandler
	logger logging.Logger
}

// NewHandler は新しいServiceRegistryを使用してハンドラーを作成
func NewHandler(awsServiceRegistry *aws.ServiceRegistry, logger logging.Logger) (*Handler, error) {
	if database.DB == nil {
		return nil, fmt.Errorf("database connection is not initialized")
	}

	return &Handler{
		HealthHandler:         systemHandler.NewHealthHandler(awsServiceRegistry.GetClientFactory()),
		CategoryListHandler:   queryHandler.NewCategoryListHandler(database.DB),
		ProductCatalogHandler: queryHandler.NewProductCatalogHandler(database.DB),
		ProductDetailHandler:  queryHandler.NewProductDetailHandler(database.DB),
		ProductHandler:        di.InitializeProductHandler(awsServiceRegistry.GetS3ClientWrapper(), logger),
		logger:                logger,
	}, nil
}
