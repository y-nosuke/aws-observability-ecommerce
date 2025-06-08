package provider

import (
	"github.com/google/wire"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/usecase"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/infrastructure/external/storage"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/presentation/rest/handler"
)

// ProductProviderSet は商品ドメインのProvider Set
var ProductProviderSet = wire.NewSet(
	// Infrastructure層 - 直接コンストラクタを使用
	storage.NewS3ImageStorageImpl,

	// UseCase層 - 設定値注入が必要なため残す
	usecase.NewUploadProductImageUseCase,
	usecase.NewGetProductImageUseCase,

	// Presentation層 - 直接コンストラクタを使用
	handler.NewProductHandler,
)
