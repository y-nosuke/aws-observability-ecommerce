package provider

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"go.opentelemetry.io/otel/sdk/resource"

	awsPkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	databasePkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// SharedProviderSet は共通インフラのProvider Set
var SharedProviderSet = wire.NewSet(
	// AWS関連
	awsPkg.NewAWSConfig,
	ProvideStsClient,
	ProvideS3Client,
	ProvideS3Config,
	// データベース関連
	databasePkg.NewDBConfig,
	// オブザーバビリティ関連
	ProvideOTelConfig,
	ProvideOTelResource,
	ProvideOTelLoggingConfig,
	ProvideLoggingConfig,
	ProvideOTelTracingConfig,
	ProvideOTelMetricsConfig,
	ProvideMetricsConfig,
	otel.NewLoggerProvider,
	otel.NewMeterProvider,
	otel.NewTracerProvider,
	otel.NewLogger,
	otel.NewHTTPMetricsRecorder,
	otel.NewTracer,
)

// ProvideStsClient はSTSクライアントを提供する
func ProvideStsClient(awsConfig aws.Config) *sts.Client {
	return sts.NewFromConfig(awsConfig)
}

// ProvideS3Client はS3クライアントを提供する
func ProvideS3Client(awsConfig aws.Config) *s3.Client {
	return s3.NewFromConfig(awsConfig, func(o *s3.Options) {
		o.UsePathStyle = true // LocalStack対応
	})
}

// ProvideS3Config はS3設定を提供する
func ProvideS3Config(awsConfig config.AWSConfig) config.S3Config {
	return awsConfig.S3
}

// ProvideOTelConfig はObservabilityConfigからOTelConfigを抽出する
func ProvideOTelConfig(observabilityConfig config.ObservabilityConfig) config.OTelConfig {
	return observabilityConfig.OTel
}

// ProvideOTelResource はOTelリソースを提供する
func ProvideOTelResource(oTelConfig config.OTelConfig) (*resource.Resource, error) {
	return otel.NewResource(oTelConfig)
}

// ProvideOTelLoggingConfig はOTelロギング設定を提供する
func ProvideOTelLoggingConfig(otelConfig config.OTelConfig) config.OTelLoggingConfig {
	return otelConfig.Logging
}

// ProvideLoggingConfig はロギング設定を提供する
func ProvideLoggingConfig(observabilityConfig config.ObservabilityConfig) config.LoggingConfig {
	return observabilityConfig.Logging
}

// ProvideOTelTracingConfig はOTelトレーシング設定を提供する
func ProvideOTelTracingConfig(otelConfig config.OTelConfig) config.OTelTracingConfig {
	return otelConfig.Tracing
}

// ProvideOTelMetricsConfig はOTelメトリクス設定を提供する
func ProvideOTelMetricsConfig(otelConfig config.OTelConfig) config.OTelMetricsConfig {
	return otelConfig.Metrics
}

// ProvideMetricsConfig はメトリクス設定を提供する
func ProvideMetricsConfig(observabilityConfig config.ObservabilityConfig) config.MetricsConfig {
	return observabilityConfig.Metrics
}
