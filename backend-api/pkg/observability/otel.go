package observability

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/metric"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// OTelManager はOpenTelemetryの初期化と管理を行う構造体
type OTelManager struct {
	loggerProvider *sdklog.LoggerProvider
	meterProvider  *sdkmetric.MeterProvider
	resource       *resource.Resource
	meter          metric.Meter // TODO metricsパッケージで初期化したい。
}

// NewOTelManager はOTelManagerのコンストラクタ（wireプロバイダー）
func NewOTelManager(otelConfig config.OTelConfig) (*OTelManager, error) {
	ctx := context.Background()

	// リソース情報を作成
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(otelConfig.ServiceName),
			semconv.ServiceVersion(otelConfig.ServiceVersion),
			semconv.ServiceNamespace(otelConfig.ServiceNamespace),
			semconv.DeploymentEnvironment(otelConfig.DeploymentEnvironment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	// ログ初期化
	var loggerProvider *sdklog.LoggerProvider

	if otelConfig.Logging.Enabled {
		loggerProvider, err = initLogging(ctx, otelConfig, res)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize logger: %w", err)
		}
		// グローバルLoggerProviderを設定
		global.SetLoggerProvider(loggerProvider)
	}

	// メトリクス初期化
	var meterProvider *sdkmetric.MeterProvider
	var meter metric.Meter

	if otelConfig.Metrics.Enabled {
		meterProvider, err = initMetrics(ctx, otelConfig, res)
		if err != nil {
			return nil, fmt.Errorf("failed to initialize metrics: %w", err)
		}
		otel.SetMeterProvider(meterProvider)
		meter = meterProvider.Meter(otelConfig.ServiceName)
	}

	log.Printf("OpenTelemetry initialized for service: %s (metrics enabled: %v)", otelConfig.ServiceName, otelConfig.Metrics.Enabled)

	return &OTelManager{
		loggerProvider: loggerProvider,
		meterProvider:  meterProvider,
		resource:       res,
		meter:          meter,
	}, nil
}

// GetLoggerProvider はLoggerProviderを返します
func (m *OTelManager) GetLoggerProvider() *sdklog.LoggerProvider {
	return m.loggerProvider
}

// GetMeterProvider はMeterProviderを返します
func (m *OTelManager) GetMeterProvider() *sdkmetric.MeterProvider {
	return m.meterProvider
}

// GetMeter はMeterを返します
func (m *OTelManager) GetMeter() metric.Meter {
	return m.meter
}

// GetResource はResourceを返します
func (m *OTelManager) GetResource() *resource.Resource {
	return m.resource
}

// Shutdown はOpenTelemetryリソースをクリーンアップします
func (m *OTelManager) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// ログプロバイダーのシャットダウン
	if m.loggerProvider != nil {
		if err := m.loggerProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown logger provider: %w", err)
		}
	}

	// メトリクスプロバイダーのシャットダウン
	if m.meterProvider != nil {
		if err := m.meterProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown meter provider: %w", err)
		}
	}

	log.Println("OpenTelemetry shutdown completed")
	return nil
}

// initLogging はログを初期化します
func initLogging(ctx context.Context, cfg config.OTelConfig, res *resource.Resource) (*sdklog.LoggerProvider, error) {
	// OTLP Log Exporter
	logExporter, err := otlploghttp.New(ctx,
		otlploghttp.WithEndpoint(cfg.Collector.Endpoint),
		otlploghttp.WithTimeout(cfg.Logging.ExportTimeout),
		otlploghttp.WithCompression(otlploghttp.GzipCompression),
		otlploghttp.WithInsecure(), // 開発環境用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create log exporter: %w", err)
	}

	// Log Provider
	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter,
			sdklog.WithExportTimeout(cfg.Logging.BatchTimeout),
			sdklog.WithMaxQueueSize(cfg.Logging.MaxQueueSize),
			sdklog.WithExportMaxBatchSize(cfg.Logging.MaxExportBatchSize),
		)),
		sdklog.WithResource(res),
	)

	return lp, nil
}

// initMetrics はメトリクスを初期化します
func initMetrics(ctx context.Context, cfg config.OTelConfig, res *resource.Resource) (*sdkmetric.MeterProvider, error) {
	// OTLP Metric Exporter
	metricExporter, err := otlpmetrichttp.New(ctx,
		otlpmetrichttp.WithEndpoint(cfg.Collector.Endpoint),
		otlpmetrichttp.WithTimeout(cfg.Metrics.ExportTimeout),
		otlpmetrichttp.WithCompression(otlpmetrichttp.GzipCompression),
		otlpmetrichttp.WithInsecure(), // 開発環境用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create metric exporter: %w", err)
	}

	// Metric Provider
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter,
			sdkmetric.WithInterval(cfg.Metrics.BatchTimeout),
		)),
		sdkmetric.WithResource(res),
	)

	return mp, nil
}
