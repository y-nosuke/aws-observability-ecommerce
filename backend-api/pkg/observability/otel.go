package observability

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.32.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// ProviderFactory はOpenTelemetryのプロバイダー生成を管理する構造体
type ProviderFactory struct {
	config         config.OTelConfig
	resource       *resource.Resource
	loggerProvider *sdklog.LoggerProvider
	meterProvider  *sdkmetric.MeterProvider
	tracerProvider *sdktrace.TracerProvider
}

// NewProviderFactory はProviderFactoryのコンストラクタ
func NewProviderFactory(otelConfig config.OTelConfig) (*ProviderFactory, error) {
	ctx := context.Background()

	// リソース情報を作成
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(otelConfig.ServiceName),
			semconv.ServiceVersion(otelConfig.ServiceVersion),
			semconv.ServiceNamespace(otelConfig.ServiceNamespace),
			semconv.DeploymentEnvironmentName(otelConfig.DeploymentEnvironment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	return &ProviderFactory{
		config:   otelConfig,
		resource: res,
	}, nil
}

// CreateLoggerProvider はLoggerProviderを作成します
func (f *ProviderFactory) CreateLoggerProvider() (*sdklog.LoggerProvider, error) {
	if !f.config.Logging.Enabled {
		return nil, nil
	}

	if f.loggerProvider != nil {
		return f.loggerProvider, nil
	}

	ctx := context.Background()
	provider, err := f.initLogging(ctx)
	if err != nil {
		return nil, err
	}

	f.loggerProvider = provider
	return provider, nil
}

// CreateMeterProvider はMeterProviderを作成します
func (f *ProviderFactory) CreateMeterProvider() (*sdkmetric.MeterProvider, error) {
	if !f.config.Metrics.Enabled {
		return nil, nil
	}

	if f.meterProvider != nil {
		return f.meterProvider, nil
	}

	ctx := context.Background()
	provider, err := f.initMetrics(ctx)
	if err != nil {
		return nil, err
	}

	f.meterProvider = provider
	return provider, nil
}

// CreateTracerProvider はTracerProviderを作成します
func (f *ProviderFactory) CreateTracerProvider() (*sdktrace.TracerProvider, error) {
	if !f.config.Tracing.Enabled {
		return nil, nil
	}

	if f.tracerProvider != nil {
		return f.tracerProvider, nil
	}

	ctx := context.Background()
	provider, err := f.initTracing(ctx)
	if err != nil {
		return nil, err
	}

	f.tracerProvider = provider
	return provider, nil
}

// GetResource はResourceを返します
func (f *ProviderFactory) GetResource() *resource.Resource {
	return f.resource
}

// GetConfig は設定を返します
func (f *ProviderFactory) GetConfig() config.OTelConfig {
	return f.config
}

// Shutdown はOpenTelemetryプロバイダーをクリーンアップします
func (f *ProviderFactory) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ShutdownProvider(ctx, f.loggerProvider); err != nil {
		return fmt.Errorf("failed to shutdown logger provider: %w", err)
	}

	if err := ShutdownProvider(ctx, f.meterProvider); err != nil {
		return fmt.Errorf("failed to shutdown meter provider: %w", err)
	}

	if err := ShutdownProvider(ctx, f.tracerProvider); err != nil {
		return fmt.Errorf("failed to shutdown tracer provider: %w", err)
	}

	return nil
}

// ShutdownProvider はOpenTelemetryプロバイダーをクリーンアップします
func ShutdownProvider(ctx context.Context, provider interface{}) error {
	switch p := provider.(type) {
	case *sdklog.LoggerProvider:
		if p != nil {
			return p.Shutdown(ctx)
		}
	case *sdkmetric.MeterProvider:
		if p != nil {
			return p.Shutdown(ctx)
		}
	case *sdktrace.TracerProvider:
		if p != nil {
			return p.Shutdown(ctx)
		}
	}
	return nil
}

// initLogging はログを初期化します
func (f *ProviderFactory) initLogging(ctx context.Context) (*sdklog.LoggerProvider, error) {
	// OTLP Log Exporter
	logExporter, err := otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(f.config.Collector.Endpoint),
		otlploghttp.WithTimeout(f.config.Logging.ExportTimeout),
		otlploghttp.WithCompression(otlploghttp.GzipCompression),
		otlploghttp.WithInsecure(), // 開発環境用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}

	// Log Provider
	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter,
			sdklog.WithExportTimeout(f.config.Logging.BatchTimeout),
			sdklog.WithMaxQueueSize(f.config.Logging.MaxQueueSize),
			sdklog.WithExportMaxBatchSize(f.config.Logging.MaxExportBatchSize),
		)),
		sdklog.WithResource(f.resource),
	)

	return lp, nil
}

// initMetrics はメトリクスを初期化します
func (f *ProviderFactory) initMetrics(ctx context.Context) (*sdkmetric.MeterProvider, error) {
	// OTLP Metric Exporter
	metricExporter, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(f.config.Collector.Endpoint),
		otlpmetrichttp.WithTimeout(f.config.Metrics.ExportTimeout),
		otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression),
		otlpmetrichttp.WithInsecure(), // 開発環境用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	// Metric Provider
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter,
			sdkmetric.WithInterval(f.config.Metrics.BatchTimeout),
		)),
		sdkmetric.WithResource(f.resource),
	)

	return mp, nil
}

// initTracing はトレースを初期化します
func (f *ProviderFactory) initTracing(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// OTLP Trace Exporter
	traceExporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(f.config.Collector.Endpoint),
		otlptracehttp.WithTimeout(f.config.Tracing.ExportTimeout),
		otlptracehttp.WithCompression(otlptracehttp.GzipCompression),
		otlptracehttp.WithInsecure(), // 開発環境用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create tracer exporter: %w", err)
	}

	// サンプラーの設定
	sampler := sdktrace.TraceIDRatioBased(f.config.Tracing.SamplingRatio)

	// Trace Provider
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(traceExporter,
			sdktrace.WithBatchTimeout(f.config.Tracing.BatchTimeout),
			sdktrace.WithMaxQueueSize(f.config.Tracing.MaxQueueSize),
			sdktrace.WithMaxExportBatchSize(f.config.Tracing.MaxExportBatchSize),
		),
		sdktrace.WithResource(f.resource),
		sdktrace.WithSampler(sampler),
	)

	return tp, nil
}
