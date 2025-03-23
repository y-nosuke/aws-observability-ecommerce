package handlers

import (
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/api/middleware"
	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
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
	// コンテキストからロガーを取得
	log := middleware.GetLogger(c)

	// リクエストの処理開始をログに記録
	log.Debug("Health check request received",
		"method", c.Request().Method,
		"path", c.Path(),
		"remote_ip", c.RealIP(),
		"request_id", middleware.GetRequestID(c),
	)

	// サービスの状態をチェック（ここでは簡易的にすべて稼働中とする）
	services := map[string]interface{}{
		"api": map[string]string{
			"status": "up",
		},
		// 実際のアプリケーションでは、データベース接続などをチェックする
		// "database": checkDatabaseConnection(),
	}

	// システムリソースの状態を取得
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	resources := map[string]interface{}{
		"memory": map[string]interface{}{
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
	log.Info("Health check completed",
		"status", response.Status,
		"uptime_ms", response.Uptime,
		"goroutines", resources["goroutines"],
	)

	return c.JSON(http.StatusOK, response)
}
