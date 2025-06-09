package observability

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// OTelManager はOpenTelemetryの初期化と管理を行う構造体
type OTelManager struct {
	loggerProvider *sdklog.LoggerProvider
	resource       *resource.Resource
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
	loggerProvider, err := initLogging(ctx, otelConfig, res)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %w", err)
	}

	// グローバルLoggerProviderを設定
	global.SetLoggerProvider(loggerProvider)

	log.Printf("OpenTelemetry initialized for service: %s", otelConfig.ServiceName)

	return &OTelManager{
		loggerProvider: loggerProvider,
		resource:       res,
	}, nil
}

// GetLoggerProvider はLoggerProviderを返します
func (m *OTelManager) GetLoggerProvider() *sdklog.LoggerProvider {
	return m.loggerProvider
}

// GetResource はResourceを返します
func (m *OTelManager) GetResource() *resource.Resource {
	return m.resource
}

// Shutdown はOpenTelemetryリソースをクリーンアップします
func (m *OTelManager) Shutdown() error {
	if m.loggerProvider != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := m.loggerProvider.Shutdown(ctx); err != nil {
			return fmt.Errorf("failed to shutdown logger provider: %w", err)
		}
		log.Println("OpenTelemetry shutdown completed")
	}
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
