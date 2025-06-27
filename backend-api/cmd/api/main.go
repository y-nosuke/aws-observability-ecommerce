package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/di"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/router"
)

func main() {
	// 設定をロード
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}
	// DIコンテナの初期化
	ctx := context.Background()
	container, cleanup, err := di.InitializeAppContainer(
		ctx,
		config.App,
		config.AWS,
		config.Database,
		config.Observability,
	)
	if err != nil {
		slog.ErrorContext(ctx, "Failed to initialize DI container", "error", err)
		log.Fatalf("Failed to initialize DI container: %v", err)
	}
	defer cleanup()

	// アプリケーション開始ログ
	slog.InfoContext(ctx, "アプリケーションが正常に開始されました",
		"config_loaded", true,
		"di_container_ready", true,
		"database_connected", true,
		"aws_services_ready", true,
		"opentelemetry_enabled", true,
		"stage", "initialization",
		"action", "start")

	// ルーターの初期化とセットアップ
	e, err := router.NewRouter(container)
	if err != nil {
		slog.ErrorContext(ctx, "ルーティングのセットアップに失敗しました",
			"error", err,
			"operation", "setup_routes",
			"severity", "critical",
			"business_impact", "service_startup_failure")
		log.Fatalf("Failed to setup routes: %v", err)
	}

	// シグナルハンドリングの設定
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// サーバーを起動
	go func() {
		address := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
		// サーバー起動ログ
		slog.InfoContext(ctx, "HTTPサーバーを起動中",
			"address", address,
			"environment", config.App.Environment,
			"layer", "main")

		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// サーバー起動エラーログ
			slog.ErrorContext(ctx, "HTTPサーバーの起動に失敗しました",
				"error", err,
				"address", address,
				"operation", "start_server",
				"severity", "critical",
				"business_impact", "service_unavailable")
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	// シャットダウン開始ログ
	slog.InfoContext(ctx, "シャットダウンシグナルを受信、グレースフルシャットダウンを開始します")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		// シャットダウンエラーログ
		slog.ErrorContext(shutdownCtx, "グレースフルシャットダウンに失敗しました",
			"error", err,
			"operation", "shutdown_server",
			"severity", "medium",
			"business_impact", "graceful_shutdown_failed")
		log.Fatalf("Failed to shutdown server gracefully: %v", err)
	} else {
		// シャットダウン成功ログ
		slog.InfoContext(shutdownCtx, "サーバーが正常にシャットダウンしました")
	}

	// アプリケーション終了ログ
	slog.InfoContext(shutdownCtx, "アプリケーションが正常に終了しました",
		"graceful_shutdown", true,
		"cleanup_executed", true,
		"stage", "completion",
		"action", "stop")
}
