package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/handlers"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/middleware"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/router"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

func main() {
	// コンテキストの初期化（シグナルハンドリング）
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// AWS設定オプションの準備
	awsOptions := aws.Options{
		UseLocalStack: config.AWS.UseLocalStack,
		Region:        config.AWS.Region,
		Endpoint:      config.AWS.Endpoint,
		Credentials: aws.Credentials{
			AccessKey: config.AWS.AccessKey,
			SecretKey: config.AWS.SecretKey,
			Token:     config.AWS.Token,
		},
	}

	// ロガーの初期化
	appLogger, err := logger.New(ctx, logger.InitConfig{
		AppName:             config.App.Name,
		Environment:         config.App.Environment,
		LogLevel:            config.Log.Level,
		UseConsole:          config.Log.UseConsole,
		UseFile:             config.Log.UseFile,
		FilePath:            config.Log.LogFilePath,
		UseCloudWatch:       config.Log.UseCloudWatch,
		CreateLogGroup:      config.Log.CreateLogGroup,
		LogGroupName:        config.Log.CloudWatchLogGroup,
		CloudWatchFlushSecs: config.Log.CloudWatchFlushSecs,
		CloudWatchBatchSize: config.Log.CloudWatchBatchSize,
		AWSOptions:          awsOptions,
	})
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer func(appLogger *logger.Logger) {
		if closeErr := appLogger.Close(); closeErr != nil {
			log.Fatalf("Failed to close logger: %v", closeErr)
		}
	}(appLogger)

	slogger := appLogger.Logger()

	// アプリケーションの起動をログに記録
	slogger.Info("Starting application",
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
		address := ":" + config.Server.Port
		slogger.Info("Server starting", "address", address)
		if startErr := e.Start(address); startErr != nil {
			slogger.Error("Server shutdown unexpectedly", "error", startErr)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	slogger.Info("Shutdown signal received, gracefully shutting down...")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		slogger.Error("Server shutdown failed", "error", err)
	}

	slogger.Info("Server has been shutdown gracefully")
}
