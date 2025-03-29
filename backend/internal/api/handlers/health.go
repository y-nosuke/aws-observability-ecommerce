package handlers

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend/internal/config"
)

// HealthResponse はヘルスチェックの応答を表す構造体
type HealthResponse struct {
	Status    string                 `json:"status"`
	Timestamp string                 `json:"timestamp"`
	Version   string                 `json:"version"`
	Uptime    int64                  `json:"uptime"`
	Resources map[string]interface{} `json:"resources"`
	Services  map[string]interface{} `json:"services"`
}

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

// HandleHealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HandleHealthCheck(c echo.Context) error {
	// リクエストの処理開始をログに記録
	log.Println("Health check request received",
		"method", c.Request().Method,
		"path", c.Path(),
		"remote_ip", c.RealIP(),
	)

	// サービスの状態をチェック（ここでは簡易的にすべて稼働中とする）
	services := map[string]interface{}{
		"api": map[string]string{
			"name":   config.App.Name,
			"status": "up",
		},
		// 実際のアプリケーションでは、データベース接続などをチェックする
		// "database": checkDatabaseConnection(),
	}

	// システムリソースの状態を取得
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	resources := map[string]any{
		"memory": map[string]any{
			"allocated": memStats.Alloc,
			"total":     memStats.TotalAlloc,
			"system":    memStats.Sys,
		},
		"goroutines": runtime.NumGoroutine(),
	}

	// レスポンスを構築
	response := &HealthResponse{
		Status:    "ok",
		Timestamp: time.Now().Format(time.RFC3339),
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
