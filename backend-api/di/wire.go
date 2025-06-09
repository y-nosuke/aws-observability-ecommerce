//go:build wireinject
// +build wireinject

package di

import (
	"context"

	"github.com/google/wire"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/di/provider"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

//go:generate go run -mod=mod github.com/google/wire/cmd/wire

// InitializeAppContainer は全体のアプリケーションコンテナを初期化
func InitializeAppContainer(
	ctx context.Context,
	appConfig config.AppConfig,
	awsConfig config.AWSConfig,
	dbConfig config.DatabaseConfig,
	observabilityConfig config.ObservabilityConfig,
) (*AppContainer, error) {
	wire.Build(
		// Provider sets
		provider.SharedProviderSet,
		provider.ProductProviderSet,
		provider.QueryProviderSet,
		provider.SystemProviderSet,

		// Container構築
		NewAppContainer,
	)
	return &AppContainer{}, nil
}
