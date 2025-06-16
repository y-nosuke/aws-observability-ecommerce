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
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/metrics"
)

func main() {
	// 設定をロード
	if err := config.LoadConfig(); err != nil {
		log.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// グローバルロガーの初期化（早期初期化でどこでも使用可能に）
	logger.Init(config.Observability)

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

	// アプリケーション開始ログ
	logger.LogBusinessEvent(ctx, "application_startup", "system", "main",
		"config_loaded", true,
		"di_container_ready", true,
		"database_connected", true,
		"aws_services_ready", true,
		"opentelemetry_enabled", true,
		"stage", "initialization",
		"action", "start")

	// グローバルHTTPメトリクスの初期化
	meter := container.OTelManager.GetMeter()
	if err := metrics.Init(meter); err != nil {
		logger.WithError(ctx, "HTTPメトリクスの初期化に失敗", err,
			"operation", "init_http_metrics",
			"severity", "medium",
			"business_impact", "metrics_collection_disabled")
		// メトリクス初期化失敗は致命的エラーではないため、アプリケーションは継続
	} else {
		logger.Info(ctx, "HTTPメトリクスを初期化しました")
	}

	// ルーターの初期化とセットアップ
	r := router.NewRouter()
	if err := r.SetupRoutes(container); err != nil {
		logger.WithError(ctx, "ルーティングのセットアップに失敗", err,
			"operation", "setup_routes",
			"severity", "critical",
			"business_impact", "service_startup_failure")
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
		// サーバー起動ログ
		logger.Info(ctx, "HTTPサーバーを起動中",
			"address", address,
			"environment", config.App.Environment,
			"layer", "main")

		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// サーバー起動エラーログ
			logger.WithError(ctx, "HTTPサーバーの起動に失敗", err,
				"address", address,
				"operation", "start_server",
				"severity", "critical",
				"business_impact", "service_unavailable")
			log.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	// シャットダウン開始ログ
	logger.Info(ctx, "シャットダウンシグナルを受信、グレースフルシャットダウン開始")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		// シャットダウンエラーログ
		logger.WithError(shutdownCtx, "グレースフルシャットダウンに失敗", err,
			"operation", "shutdown_server",
			"severity", "medium",
			"business_impact", "graceful_shutdown_failed")
		log.Printf("Failed to shutdown server gracefully: %v\n", err)
	} else {
		// シャットダウン成功ログ
		logger.Info(shutdownCtx, "サーバーが正常にシャットダウンしました")
	}

	// アプリケーション終了ログ
	logger.LogBusinessEvent(shutdownCtx, "application_shutdown", "system", "main",
		"graceful_shutdown", true,
		"cleanup_executed", true,
		"stage", "completion",
		"action", "stop")
}
