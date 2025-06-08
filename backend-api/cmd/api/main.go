package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/di"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/router"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
)

func main() {
	// 設定をロード
	if err := config.LoadConfig(); err != nil {
		log.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// DIコンテナの初期化（OpenTelemetryも含めて一括初期化）
	ctx := context.Background()
	container, err := di.InitializeAppContainer(
		ctx,
		config.App,
		config.AWS,
		config.Database,
		config.Observability,
	)
	if err != nil {
		log.Printf("Failed to initialize DI container: %v\n", err)
		os.Exit(1)
	}

	// アプリケーション終了時のクリーンアップを設定
	defer func() {
		if cleanupErr := container.Cleanup(); cleanupErr != nil {
			log.Printf("Error during cleanup: %v\n", cleanupErr)
		}
	}()

	// ログ出力用のloggerを取得
	logger := container.GetLogger()

	// 構造化ログでアプリケーション開始をログ出力
	logger.LogApplication(ctx, logging.ApplicationOperation{
		Name:     "application_startup",
		Category: "system",
		Duration: 0,
		Success:  true,
		Stage:    "initialization",
		Action:   "start",
		Source:   "main",
		Data: map[string]interface{}{
			"config_loaded":         true,
			"di_container_ready":    true,
			"database_connected":    true,
			"aws_services_ready":    true,
			"opentelemetry_enabled": true,
		},
	})

	// ルーターの初期化とセットアップ
	r := router.NewRouter(logger)
	if err := r.SetupRoutes(container); err != nil {
		logger.LogError(ctx, err, logging.ErrorContext{
			Operation:      "setup_routes",
			Severity:       "critical",
			BusinessImpact: "service_startup_failure",
		})
		log.Fatalf("Failed to setup routes: %v", err)
	}

	// Echoインスタンスを取得
	e := r.GetEcho()

	// シグナルハンドリングの設定
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// サーバーを起動
	go func() {
		address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
		logger.Info(ctx, "Starting HTTP server",
			logging.Field{Key: "address", Value: address},
			logging.Field{Key: "environment", Value: config.App.Environment},
		)

		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.LogError(ctx, err, logging.ErrorContext{
				Operation:      "start_server",
				Severity:       "critical",
				BusinessImpact: "service_unavailable",
			})
			log.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	logger.Info(ctx, "Shutdown signal received, gracefully shutting down...")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		logger.LogError(shutdownCtx, err, logging.ErrorContext{
			Operation:      "shutdown_server",
			Severity:       "medium",
			BusinessImpact: "graceful_shutdown_failed",
		})
		log.Printf("Failed to shutdown server gracefully: %v\n", err)
	} else {
		logger.Info(shutdownCtx, "Server shutdown gracefully")
	}

	// アプリケーション終了をログ出力
	logger.LogApplication(shutdownCtx, logging.ApplicationOperation{
		Name:     "application_shutdown",
		Category: "system",
		Duration: 0,
		Success:  true,
		Stage:    "completion",
		Action:   "stop",
		Source:   "main",
		Data: map[string]interface{}{
			"graceful_shutdown": true,
			"cleanup_executed":  true,
		},
	})
}
