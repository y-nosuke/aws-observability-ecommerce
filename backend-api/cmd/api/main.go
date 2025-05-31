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
)

func main() {
	// 設定をロード
	if err := config.LoadConfig(); err != nil {
		log.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// データベース接続の初期化
	if err := database.InitDatabase(); err != nil {
		log.Printf("Failed to initialize database: %v\n", err)
		os.Exit(1)
	}
	defer func() {
		if err := database.CloseDatabase(); err != nil {
			log.Printf("Failed to close database: %v\n", err)
		}
	}()

	// AWSサービスレジストリの初期化
	ctx := context.Background()
	awsServiceRegistry, err := aws.NewServiceRegistry(ctx, config.AWS)
	if err != nil {
		log.Printf("Failed to initialize AWS services: %v", err)
		os.Exit(1)
	}

	// ルーターの初期化とセットアップ
	r := router.NewRouter()
	if err := r.SetupRoutes(awsServiceRegistry); err != nil {
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
		log.Printf("Starting server on %s", address)
		if err := e.Start(address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Failed to start server: %v\n", err)
			os.Exit(1)
		}
	}()

	// シグナルを待機
	<-ctx.Done()
	log.Println("Shutdown signal received, gracefully shutting down...")

	// グレースフルシャットダウン
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(shutdownCtx); err != nil {
		log.Printf("Failed to shutdown server gracefully: %v\n", err)
	} else {
		log.Printf("Server shutdown gracefully")
	}
}
