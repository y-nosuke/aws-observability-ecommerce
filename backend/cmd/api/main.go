package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/router"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
)

func main() {
	// コンテキストの初期化（シグナルハンドリング）
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// アプリケーションの起動をログに記録
	log.Printf("Starting application (version: %s, environment: %s)\n",
		config.Config.App.Version,
		config.Config.App.Environment)

	// Echoインスタンスの作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Logger()) // 標準のロガーミドルウェアを使用

	// ハンドラーの作成
	healthHandler := handlers.NewHealthHandler()
	productHandler := handlers.NewProductHandler()

	// ルーターの設定
	router.SetupRoutes(e, healthHandler, productHandler)

	// サーバーの起動（非同期）
	go func() {
		log.Printf("Server starting on port %s\n", config.Config.Server.Port)
		if err := e.Start(":" + config.Config.Server.Port); err != nil {
			log.Printf("Server shutdown: %v\n", err)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	log.Println("Shutdown signal received, gracefully shutting down...")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("Server shutdown failed: %v\n", err)
	}

	log.Println("Server has been shutdown gracefully")
}
