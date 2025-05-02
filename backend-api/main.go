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

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/handlers"
	awsconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/config"
	dbconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/config"
)

func main() {
	// 設定をロード
	if err := config.Load(); err != nil {
		log.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// AWS設定オプションの準備
	awsOptions := awsconfig.Options{
		UseLocalStack: config.AWS.UseLocalStack,
		Region:        config.AWS.Region,
		Endpoint:      config.AWS.Endpoint,
		Credentials: awsconfig.Credentials{
			AccessKey: config.AWS.AccessKey,
			SecretKey: config.AWS.SecretKey,
			Token:     config.AWS.Token,
		},
	}

	awsConfig, err := awsconfig.NewAWSConfig(context.Background(), awsOptions)
	if err != nil {
		log.Printf("AWS設定エラー: %v", err)
		os.Exit(1)
	}

	// データベース接続の初期化 - 追加
	if err := dbconfig.InitDatabase(); err != nil {
		log.Printf("Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	// defer でアプリケーション終了時にデータベース接続をクローズ - 追加
	defer func() {
		if err := dbconfig.CloseDatabase(); err != nil {
			log.Printf("Failed to close database: %v\n", err)
		}
	}()

	// Echoインスタンスを作成
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// ミドルウェアの設定
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger()) // 標準のロガーミドルウェアを使用
	e.Use(middleware.CORS())

	// APIグループ
	api := e.Group("/api")

	if err := handlers.RegisterHandlers(api, awsConfig); err != nil {
		log.Fatalf("Failed to register handlers: %v", err)
	}

	e.Static("/swagger", "static/swagger-ui")
	e.File("/swagger", "static/swagger-ui/index.html")
	e.File("/openapi.yaml", "openapi.yaml")

	// コンテキストの初期化（シグナルハンドリング）
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// サーバーを起動
	go func() {
		address := fmt.Sprintf(":%d", config.Server.Port)
		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	log.Println("Shutdown signal received, gracefully shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Printf("Failed to shutdown server gracefully: %v\n", err)
	} else {
		log.Printf("Server shutdown gracefully")
	}
}
