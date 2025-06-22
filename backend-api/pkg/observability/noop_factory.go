package observability

import (
	"context"

	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// NoopProviderFactory は何もしないProviderFactory実装
// 本番環境でObservabilityを無効化したい場合に使用
type NoopProviderFactory struct {
	config   config.OTelConfig
	resource *resource.Resource
}

// NewNoopProviderFactory は何もしないProviderFactoryを作成します
func NewNoopProviderFactory(otelConfig config.OTelConfig) (*NoopProviderFactory, error) {
	ctx := context.Background()

	// 最小限のリソース情報を作成
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(otelConfig.ServiceName),
			semconv.ServiceVersion(otelConfig.ServiceVersion),
		),
	)
	if err != nil {
		return nil, err
	}

	return &NoopProviderFactory{
		config:   otelConfig,
		resource: res,
	}, nil
}

// CreateLoggerProvider は何もしない（nilを返す）
func (n *NoopProviderFactory) CreateLoggerProvider() (*sdklog.LoggerProvider, error) {
	return nil, nil
}

// CreateMeterProvider は何もしない（nilを返す）
func (n *NoopProviderFactory) CreateMeterProvider() (*sdkmetric.MeterProvider, error) {
	return nil, nil
}

// CreateTracerProvider は何もしない（nilを返す）
func (n *NoopProviderFactory) CreateTracerProvider() (*sdktrace.TracerProvider, error) {
	return nil, nil
}

// GetResource はResourceを返します
func (n *NoopProviderFactory) GetResource() *resource.Resource {
	return n.resource
}

// GetConfig は設定を返します
func (n *NoopProviderFactory) GetConfig() config.OTelConfig {
	return n.config
}

// Shutdown は何もしない
func (n *NoopProviderFactory) Shutdown() error {
	return nil
}
