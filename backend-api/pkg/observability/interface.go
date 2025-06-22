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

// ProviderFactoryOption はProviderFactoryのオプション関数型
//
// 使用例:
//
//	// デフォルト設定（すべて有効）
//	factory, err := NewOTelProviderFactory(config)
//
//	// 特定の機能のみ無効化
//	factory, err := NewOTelProviderFactory(config, WithMetricsDisabled())
//
//	// 複数のオプションを組み合わせ
//	factory, err := NewOTelProviderFactory(config,
//	    WithLoggingDisabled(),
//	    WithTracingEnabled(true),
//	)
//
//	// テスト環境での使用例
//	factory, err := NewOTelProviderFactory(config,
//	    WithMetricsDisabled(),
//	    WithLoggingDisabled(),
//	)
type ProviderFactoryOption func(*ProviderFactoryOptions)

// DefaultProviderFactoryOptions はデフォルトのオプションを返します
func DefaultProviderFactoryOptions() ProviderFactoryOptions {
	return ProviderFactoryOptions{
		EnableLogging: true,
		EnableMetrics: true,
		EnableTracing: true,
	}
}

// WithLoggingEnabled はログ機能を有効にするオプション
func WithLoggingEnabled(enabled bool) ProviderFactoryOption {
	return func(opts *ProviderFactoryOptions) {
		opts.EnableLogging = enabled
	}
}

// WithLoggingDisabled はログ機能を無効にするオプション
func WithLoggingDisabled() ProviderFactoryOption {
	return func(opts *ProviderFactoryOptions) {
		opts.EnableLogging = false
	}
}

// WithMetricsEnabled はメトリクス機能を有効にするオプション
func WithMetricsEnabled(enabled bool) ProviderFactoryOption {
	return func(opts *ProviderFactoryOptions) {
		opts.EnableMetrics = enabled
	}
}

// WithMetricsDisabled はメトリクス機能を無効にするオプション
func WithMetricsDisabled() ProviderFactoryOption {
	return func(opts *ProviderFactoryOptions) {
		opts.EnableMetrics = false
	}
}

// WithTracingEnabled はトレース機能を有効にするオプション
func WithTracingEnabled(enabled bool) ProviderFactoryOption {
	return func(opts *ProviderFactoryOptions) {
		opts.EnableTracing = enabled
	}
}

// WithTracingDisabled はトレース機能を無効にするオプション
func WithTracingDisabled() ProviderFactoryOption {
	return func(opts *ProviderFactoryOptions) {
		opts.EnableTracing = false
	}
}
