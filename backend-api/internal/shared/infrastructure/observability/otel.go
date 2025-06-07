package observability

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// InitOpenTelemetry はOpenTelemetryを初期化します
func InitOpenTelemetry(cfg config.OTelConfig) (func(), error) {
	ctx := context.Background()

	// リソース情報を作成
	res, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceName(cfg.ServiceName),
			semconv.ServiceVersion(cfg.ServiceVersion),
			semconv.ServiceNamespace(cfg.ServiceNamespace),
			semconv.DeploymentEnvironment(cfg.DeploymentEnvironment),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create resource: %w", err)
	}

	var shutdownFuncs []func()

	// ログ初期化
	loggerShutdown, err := initLogging(ctx, cfg, res)
	if err != nil {
		// 既に初期化されたものをクリーンアップ
		for _, shutdown := range shutdownFuncs {
			shutdown()
		}
		return nil, fmt.Errorf("failed to initialize logging: %w", err)
	}
	shutdownFuncs = append(shutdownFuncs, loggerShutdown)

	// シャットダウン関数を返す
	return func() {
		for i := len(shutdownFuncs) - 1; i >= 0; i-- {
			shutdownFuncs[i]()
		}
	}, nil
}

// initLogging はログを初期化します
func initLogging(ctx context.Context, cfg config.OTelConfig, res *resource.Resource) (func(), error) {
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
	lp := log.NewLoggerProvider(
		log.WithProcessor(log.NewBatchProcessor(logExporter,
			log.WithExportTimeout(cfg.Logging.BatchTimeout),
			log.WithMaxQueueSize(cfg.Logging.MaxQueueSize),
			log.WithExportMaxBatchSize(cfg.Logging.MaxExportBatchSize),
			log.WithExportTimeout(cfg.Logging.ExportTimeout),
		)),
		log.WithResource(res),
	)

	global.SetLoggerProvider(lp)

	return func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := lp.Shutdown(ctx); err != nil {
			// ログに記録するがエラーは握りつぶす
			fmt.Printf("Error shutting down logger provider: %v\n", err)
		}
	}, nil
}
