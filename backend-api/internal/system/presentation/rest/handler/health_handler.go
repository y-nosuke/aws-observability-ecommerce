package handler

import (
	"context"
	"errors"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// HealthHandler はヘルスチェックのハンドラーを表す構造体
type HealthHandler struct {
	startTime  time.Time
	version    string
	awsFactory *aws.ClientFactory
}

// NewHealthHandler は新しいヘルスハンドラーを作成します
func NewHealthHandler(awsFactory *aws.ClientFactory) *HealthHandler {
	return &HealthHandler{
		startTime:  time.Now(),
		version:    config.App.Version,
		awsFactory: awsFactory,
	}
}

// HealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HealthCheck(c echo.Context, params openapi.HealthCheckParams) error {
	log.Println("Health check request received",
		"method", c.Request().Method,
		"path", c.Path(),
		"remote_ip", c.RealIP(),
	)

	ctx, cancel := context.WithTimeout(c.Request().Context(), 5*time.Second)
	defer cancel()

	var checks []string
	if params.Checks != nil {
		checks = strings.Split(*params.Checks, ",")
	}

	response := &openapi.HealthResponse{
		Status:     "ok",
		Timestamp:  time.Now(),
		Version:    h.version,
		Uptime:     time.Since(h.startTime).Milliseconds(),
		Resources:  h.createResources(),
		Components: h.createComponents(ctx, checks),
	}

	log.Println("Health check completed",
		"status", response.Status,
		"uptime", response.Uptime,
	)

	return c.JSON(http.StatusOK, response)
}

func (h *HealthHandler) createResources() openapi.SystemResources {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	return openapi.SystemResources{
		"system": {
			Memory: openapi.MemoryStats{
				Allocated: memStats.Alloc,
				Total:     memStats.TotalAlloc,
				System:    memStats.Sys,
			},
			Goroutines: runtime.NumGoroutine(),
		},
	}
}

func (h *HealthHandler) createComponents(ctx context.Context, checks []string) map[string]string {
	components := map[string]string{}

	healthCheckers := NewHealthCheckers(database.DB, h.awsFactory, checks)
	results := healthCheckers.Check(ctx)
	for n, e := range results {
		if e != nil {
			var clientMsg string
			var hcErr *HealthCheckError
			if errors.As(e, &hcErr) {
				clientMsg = hcErr.Msg
			} else {
				clientMsg = "unknown error"
			}
			components[n] = "ng: " + clientMsg
			log.Println("[health]", n, "error:", e.Error())
		} else {
			components[n] = "ok"
		}
	}

	return components
}
