package handlers

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/config"
)

// HealthHandler はヘルスチェックのハンドラーを表す構造体
type HealthHandler struct {
	startTime time.Time
	version   string
}

// NewHealthHandler は新しいヘルスハンドラーを作成します
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
		version:   config.App.Version, // アプリケーションバージョン
	}
}

// HealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HealthCheck(c echo.Context) error {
	// リクエストの処理開始をログに記録
	log.Println("Health check request received",
		"method", c.Request().Method,
		"path", c.Path(),
		"remote_ip", c.RealIP(),
	)

	// サービスの状態をチェック（ここでは簡易的にすべて稼働中とする）
	services := openapi.DependentServices{
		"api": {
			Name:   config.App.Name,
			Status: "up",
		},
		// 実際のアプリケーションでは、データベース接続などをチェックする
		// "database": checkDatabaseConnection(),
	}

	// システムリソースの状態を取得
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	resources := openapi.SystemResources{
		"system": {
			Memory: openapi.MemoryStats{
				Allocated: memStats.Alloc,
				Total:     memStats.TotalAlloc,
				System:    memStats.Sys,
			},
			Goroutines: runtime.NumGoroutine(),
		},
	}

	// レスポンスを構築
	response := &openapi.HealthResponse{
		Status:    "ok",
		Timestamp: time.Now(),
		Version:   h.version,
		Uptime:    time.Since(h.startTime).Milliseconds(),
		Resources: resources,
		Services:  services,
	}

	// レスポンスの送信をログに記録
	log.Println("Health check completed",
		"status", response.Status,
		"uptime", response.Uptime,
	)

	return c.JSON(http.StatusOK, response)
}
