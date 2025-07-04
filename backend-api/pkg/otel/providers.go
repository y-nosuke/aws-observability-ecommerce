package otel

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/contrib/processors/minsev"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

func NewLoggerProvider(ctx context.Context, res *resource.Resource, config config.OTelLoggingConfig) (*sdklog.LoggerProvider, func(), error) {
	if !config.Enabled {
		return nil, nil, nil
	}
	exporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithEndpoint(config.Endpoint),
		otlploggrpc.WithTimeout(config.Timeout),
		otlploggrpc.WithCompressor(config.Compression),
		otlploggrpc.WithRetry(otlploggrpc.RetryConfig{
			Enabled:         config.RetryEnabled,
			InitialInterval: config.RetryInitialInterval,
			MaxInterval:     config.RetryMaxInterval,
			MaxElapsedTime:  config.RetryMaxElapsedTime,
		}),
		otlploggrpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("otlploggrpc.New: %w", err)
	}
	provider := sdklog.NewLoggerProvider(
		sdklog.WithResource(res),
		sdklog.WithProcessor(
			minsev.NewLogProcessor(
				sdklog.NewBatchProcessor(exporter,
					sdklog.WithExportTimeout(config.ExportTimeout),
					sdklog.WithMaxQueueSize(config.MaxQueueSize),
					sdklog.WithExportMaxBatchSize(config.MaxExportBatchSize),
				),
				minsev.SeverityDebug,
			),
		),
	)
	global.SetLoggerProvider(provider)
	shutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if provider != nil {
			if err := provider.Shutdown(shutdownCtx); err != nil {
				log.Printf("failed to shutdown logger provider: %v", err)
			}
		}
	}
	return provider, shutdown, nil
}
func NewMeterProvider(ctx context.Context, res *resource.Resource, config config.OTelMetricsConfig) (*sdkmetric.MeterProvider, func(), error) {
	if !config.Enabled {
		return nil, nil, nil
	}
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithEndpoint(config.Endpoint),
		otlpmetricgrpc.WithTimeout(config.Timeout),
		otlpmetricgrpc.WithCompressor(config.Compression),
		otlpmetricgrpc.WithRetry(otlpmetricgrpc.RetryConfig{
			Enabled:         config.RetryEnabled,
			InitialInterval: config.RetryInitialInterval,
			MaxInterval:     config.RetryMaxInterval,
			MaxElapsedTime:  config.RetryMaxElapsedTime,
		}),
		otlpmetricgrpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("otlpmetricgrpc.New: %w", err)
	}
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(exporter,
				sdkmetric.WithInterval(config.Interval),
			),
		),
	)
	otel.SetMeterProvider(provider)
	shutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if provider != nil {
			if err := provider.Shutdown(shutdownCtx); err != nil {
				log.Printf("failed to shutdown meter provider: %v", err)
			}
		}
	}
	return provider, shutdown, nil
}
func NewTracerProvider(ctx context.Context, res *resource.Resource, config config.OTelTracingConfig) (*sdktrace.TracerProvider, func(), error) {
	if !config.Enabled {
		return nil, nil, nil
	}
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithEndpoint(config.Endpoint),
		otlptracegrpc.WithTimeout(config.Timeout),
		otlptracegrpc.WithCompressor(config.Compression),
		otlptracegrpc.WithRetry(otlptracegrpc.RetryConfig{
			Enabled:         config.RetryEnabled,
			InitialInterval: config.RetryInitialInterval,
			MaxInterval:     config.RetryMaxInterval,
			MaxElapsedTime:  config.RetryMaxElapsedTime,
		}),
		otlptracegrpc.WithInsecure(),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("otlptracegrpc.New: %w", err)
	}
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter,
			sdktrace.WithBatchTimeout(config.BatchTimeout),
			sdktrace.WithMaxQueueSize(config.MaxQueueSize),
			sdktrace.WithMaxExportBatchSize(config.MaxExportBatchSize),
		),
		sdktrace.WithSampler(sdktrace.TraceIDRatioBased(config.SamplingRatio)),
	)
	otel.SetTracerProvider(provider)
	shutdown := func() {
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if provider != nil {
			if err := provider.Shutdown(shutdownCtx); err != nil {
				log.Printf("failed to shutdown tracer provider: %v", err)
			}
		}
	}
	return provider, shutdown, nil
}
