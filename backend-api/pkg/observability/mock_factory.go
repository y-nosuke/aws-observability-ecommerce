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

// MockProviderFactory はテスト用のProviderFactory実装
type MockProviderFactory struct {
	config         config.OTelConfig
	resource       *resource.Resource
	loggerProvider *sdklog.LoggerProvider
	meterProvider  *sdkmetric.MeterProvider
	tracerProvider *sdktrace.TracerProvider
	shutdownCalled bool
	errorOnCreate  error
}

// NewMockProviderFactory はテスト用のMockProviderFactoryを作成します
func NewMockProviderFactory() (*MockProviderFactory, error) {
	ctx := context.Background()

	// テスト用のリソース
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName("test-service"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return nil, err
	}

	return &MockProviderFactory{
		config: config.OTelConfig{
			ServiceName:    "test-service",
			ServiceVersion: "1.0.0",
		},
		resource: res,
	}, nil
}

// SetErrorOnCreate はProvider作成時にエラーを発生させるように設定します
func (m *MockProviderFactory) SetErrorOnCreate(err error) {
	m.errorOnCreate = err
}

// CreateLoggerProvider はモックのLoggerProviderを作成します
func (m *MockProviderFactory) CreateLoggerProvider() (*sdklog.LoggerProvider, error) {
	if m.errorOnCreate != nil {
		return nil, m.errorOnCreate
	}

	if m.loggerProvider == nil {
		m.loggerProvider = sdklog.NewLoggerProvider(
			sdklog.WithResource(m.resource),
		)
	}
	return m.loggerProvider, nil
}

// CreateMeterProvider はモックのMeterProviderを作成します
func (m *MockProviderFactory) CreateMeterProvider() (*sdkmetric.MeterProvider, error) {
	if m.errorOnCreate != nil {
		return nil, m.errorOnCreate
	}

	if m.meterProvider == nil {
		m.meterProvider = sdkmetric.NewMeterProvider(
			sdkmetric.WithResource(m.resource),
		)
	}
	return m.meterProvider, nil
}

// CreateTracerProvider はモックのTracerProviderを作成します
func (m *MockProviderFactory) CreateTracerProvider() (*sdktrace.TracerProvider, error) {
	if m.errorOnCreate != nil {
		return nil, m.errorOnCreate
	}

	if m.tracerProvider == nil {
		m.tracerProvider = sdktrace.NewTracerProvider(
			sdktrace.WithResource(m.resource),
		)
	}
	return m.tracerProvider, nil
}

// GetResource はResourceを返します
func (m *MockProviderFactory) GetResource() *resource.Resource {
	return m.resource
}

// GetConfig は設定を返します
func (m *MockProviderFactory) GetConfig() config.OTelConfig {
	return m.config
}

// Shutdown はモックのShutdownを実行します
func (m *MockProviderFactory) Shutdown() error {
	m.shutdownCalled = true
	return nil
}

// IsShutdownCalled はShutdownが呼ばれたかどうかを返します
func (m *MockProviderFactory) IsShutdownCalled() bool {
	return m.shutdownCalled
}
