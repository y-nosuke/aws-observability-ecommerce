package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/middleware"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/router"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

func main() {
	// コンテキストの初期化（シグナルハンドリング）
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// ロガーの初期化
	log := logger.Init(logger.Config{
		Environment: config.App.Environment,
		LogLevel:    config.Log.Level,
		ServiceName: config.App.Name,
		Version:     config.App.Version,
	})

	// アプリケーションの起動をログに記録
	log.Info("Starting application",
		"version", config.App.Version,
		"environment", config.App.Environment)

	// Echoインスタンスの作成
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// カスタムエラーハンドラーの設定
	e.HTTPErrorHandler = middleware.ErrorHandler

	// ミドルウェアの設定（ロガーも含む）
	middleware.SetupMiddleware(e)

	// ハンドラーの作成
	healthHandler := handlers.NewHealthHandler()
	productHandler := handlers.NewProductHandler()

	// ルーターの設定
	router.SetupRoutes(e, healthHandler, productHandler)

	// サーバーの起動（非同期）
	go func() {
		log.Info("Server starting", "port", config.Server.Port)
		if err := e.Start(":" + config.Server.Port); err != nil {
			log.Error("Server shutdown", "error", err)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	log.Info("Shutdown signal received, gracefully shutting down...")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Error("Server shutdown failed", "error", err)
	}

	log.Info("Server has been shutdown gracefully")
}
