package provider

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
)

// SharedProviderSet は共通インフラのProvider Set
var SharedProviderSet = wire.NewSet(
	// AWS関連
	aws.NewClientFactory,
	ProvideS3ClientWrapper,
	aws.NewServiceRegistry,

	// データベース関連
	database.NewDBManager,
	ProvideDB, // *sql.DBを提供するプロバイダー

	// オブザーバビリティ関連
	ProvideOTelConfig, // ObservabilityConfigからOTelConfigを抽出
	ProvideOTelProviderFactory,
	wire.Bind(new(observability.ProviderFactory), new(*observability.OTelProviderFactory)),

	// SqlBoiler用のバインド
	wire.Bind(new(boil.ContextExecutor), new(*sql.DB)),
)

// ProvideS3ClientWrapper はS3ClientWrapperを提供する
func ProvideS3ClientWrapper(clientFactory *aws.ClientFactory, awsConfig config.AWSConfig) *aws.S3ClientWrapper {
	s3Client := clientFactory.GetS3Client()
	return aws.NewS3ClientWrapper(s3Client, awsConfig.S3)
}

// ProvideDB はDBManagerから*sql.DBを提供する
func ProvideDB(dbManager *database.DBManager) *sql.DB {
	return dbManager.DB()
}

// ProvideOTelConfig はObservabilityConfigからOTelConfigを抽出する
func ProvideOTelConfig(observabilityConfig config.ObservabilityConfig) config.OTelConfig {
	return observabilityConfig.OTel
}

// ProvideOTelProviderFactory はOTelProviderFactoryを提供する
func ProvideOTelProviderFactory(otelConfig config.OTelConfig) (*observability.OTelProviderFactory, error) {
	return observability.NewOTelProviderFactory(otelConfig, observability.DefaultProviderFactoryOptions())
}
