package handler

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// HealthHandler はヘルスチェックのハンドラーを表す構造体
type HealthHandler struct {
	startTime time.Time
	version   string
	db        *sql.DB
	stsClient *sts.Client
	s3Client  *s3.Client
	config    config.S3Config
}

// NewHealthHandler は新しいヘルスハンドラーを作成します
func NewHealthHandler(db *sql.DB, stsClient *sts.Client, s3Client *s3.Client, cfg config.S3Config) *HealthHandler {
	return &HealthHandler{
		startTime: time.Now(),
		version:   config.App.Version,
		db:        db,
		stsClient: stsClient,
		s3Client:  s3Client,
		config:    cfg,
	}
}

// HealthCheck はヘルスチェックエンドポイントのハンドラー関数
func (h *HealthHandler) HealthCheck(c echo.Context, params openapi.HealthCheckParams) (err error) {
	spanCtx, o := otel.Start(c.Request().Context())
	defer func() {
		o.End(err)
	}()

	spanCtx, cancel := context.WithTimeout(spanCtx, 5*time.Second)
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
		Components: h.createComponents(spanCtx, checks),
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

	healthCheckers := NewHealthCheckers(h.db, h.stsClient, h.s3Client, h.config, checks)
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
		} else {
			components[n] = "ok"
		}
	}

	return components
}
