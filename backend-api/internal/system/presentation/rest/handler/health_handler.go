package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// HealthHandler はヘルスチェックのハンドラーを表す構造体
type HealthHandler struct {
	startTime  time.Time
	version    string
	db         *sql.DB
	awsFactory *aws.ClientFactory
}

// NewHealthHandler は新しいヘルスハンドラーを作成します
func NewHealthHandler(db *sql.DB, awsFactory *aws.ClientFactory) *HealthHandler {
	return &HealthHandler{
		startTime:  time.Now(),
		version:    config.App.Version,
		db:         db,
		awsFactory: awsFactory,
	}
}

// HealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HealthCheck(c echo.Context, params openapi.HealthCheckParams) error {
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

	healthCheckers := NewHealthCheckers(h.db, h.awsFactory, checks)
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
			logger.WithError(ctx, "ヘルスチェック失敗", e, "component", n, "layer", "health_check")
		} else {
			components[n] = "ok"
		}
	}

	return components
}
