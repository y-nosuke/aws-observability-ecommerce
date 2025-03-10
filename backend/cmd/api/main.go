package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/exp/slog"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/observability/logging"
)

func main() {
	// 設定ファイルのパスを取得
	var configPath string
	flag.StringVar(&configPath, "config", "config/config.yaml", "path to config file")
	flag.Parse()

	// 設定を読み込む
	cfg, err := config.Load(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// ロガーの初期化
	logger := logging.NewLogger(cfg.LogLevel)
	slog.SetDefault(logger)

	// Echoインスタンスの作成
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// ミドルウェアの設定
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(logging.Middleware())

	// ルートの設定
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"status":  "ok",
			"version": "0.1.0",
		})
	})

	// サーバーを起動
	go func() {
		address := fmt.Sprintf(":%d", cfg.Server.Port)
		if err := e.Start(address); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// シグナルを待機
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// グレースフルシャットダウン
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		slog.Error("Failed to shutdown server", "error", err)
	}
}
