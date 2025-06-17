package observability

import (
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// ProviderFactory はOpenTelemetryのプロバイダー生成を管理するインターフェース
type ProviderFactory interface {
	// CreateLoggerProvider はLoggerProviderを作成します
	CreateLoggerProvider() (*sdklog.LoggerProvider, error)

	// CreateMeterProvider はMeterProviderを作成します
	CreateMeterProvider() (*sdkmetric.MeterProvider, error)

	// CreateTracerProvider はTracerProviderを作成します
	CreateTracerProvider() (*sdktrace.TracerProvider, error)

	// GetResource はResourceを返します
	GetResource() *resource.Resource

	// GetConfig は設定を返します
	GetConfig() config.OTelConfig

	// Shutdown はOpenTelemetryプロバイダーをクリーンアップします
	Shutdown() error
}

// ProviderFactoryOptions はProviderFactoryの初期化オプション
type ProviderFactoryOptions struct {
	// EnableLogging はログ機能を有効にするかどうか
	EnableLogging bool

	// EnableMetrics はメトリクス機能を有効にするかどうか
	EnableMetrics bool

	// EnableTracing はトレース機能を有効にするかどうか
	EnableTracing bool
}

// DefaultProviderFactoryOptions はデフォルトのオプションを返します
func DefaultProviderFactoryOptions() ProviderFactoryOptions {
	return ProviderFactoryOptions{
		EnableLogging: true,
		EnableMetrics: true,
		EnableTracing: true,
	}
}
