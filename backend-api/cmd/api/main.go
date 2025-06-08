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

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/router"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
)

func main() {
	// 設定をロード
	if err := config.LoadConfig(); err != nil {
		log.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// 構造化ロガーの初期化
	logger, err := logging.NewLogger(config.Observability)
	if err != nil {
		log.Printf("Failed to initialize logger: %v\n", err)
		os.Exit(1)
	}

	// OpenTelemetryの初期化
	otelShutdown, err := observability.InitOpenTelemetry(config.Observability.OTel)
	if err != nil {
		log.Printf("Failed to initialize OpenTelemetry: %v\n", err)
		os.Exit(1)
	}
	defer otelShutdown()

	// データベース接続の初期化
	if initErr := database.InitDatabase(); err != nil {
		log.Printf("Failed to initialize database: %v\n", initErr)
		os.Exit(1)
	}
	defer func() {
		if closeErr := database.CloseDatabase(); closeErr != nil {
			log.Printf("Failed to close database: %v\n", closeErr)
		}
	}()

	// AWSサービスレジストリの初期化
	ctx := context.Background()
	awsServiceRegistry, err := aws.NewServiceRegistry(ctx, config.AWS)
	if err != nil {
		log.Printf("Failed to initialize AWS services: %v", err)
		os.Exit(1)
	}

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
			"config_loaded":      true,
			"database_connected": true,
			"aws_services_ready": true,
		},
	})

	// ルーターの初期化とセットアップ
	r := router.NewRouter(logger)
	if err := r.SetupRoutes(awsServiceRegistry); err != nil {
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
	})
}
