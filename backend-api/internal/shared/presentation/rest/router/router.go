package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/handler"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"

	customMiddleware "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/middleware"
)

// Router はアプリケーションのルーティングを管理する
type Router struct {
	echo   *echo.Echo
	logger logging.Logger
}

// NewRouter は新しいRouterインスタンスを作成
func NewRouter(logger logging.Logger) *Router {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	return &Router{
		echo:   e,
		logger: logger,
	}
}

// SetupRoutes は全てのルーティングを設定
func (r *Router) SetupRoutes(awsServiceRegistry *aws.ServiceRegistry) error {
	// 1. ミドルウェアの設定
	r.setupMiddleware()

	// 2. 静的ファイル配信
	r.setupStaticRoutes()

	// 3. OpenAPI仕様に基づくAPIルーティング
	api := r.echo.Group("/api")
	return r.setupAPIRoutes(api, awsServiceRegistry)
}

// setupMiddleware は共通ミドルウェアを設定
func (r *Router) setupMiddleware() {
	// 基本的なミドルウェア
	r.echo.Use(middleware.Recover())
	r.echo.Use(middleware.CORS())

	// 構造化ログ関連のミドルウェア（順序重要）
	r.echo.Use(customMiddleware.RequestIDMiddleware())
	r.echo.Use(customMiddleware.StructuredLoggingMiddleware(r.logger))
	r.echo.Use(customMiddleware.ErrorHandlingMiddleware(r.logger))

	// 将来実装予定のミドルウェア
	r.echo.Use(customMiddleware.NewMetricsMiddleware())
	r.echo.Use(customMiddleware.NewRateLimitMiddleware())
}

// setupAPIRoutes はoapi-codegenを使用してAPIルーティングを設定
func (r *Router) setupAPIRoutes(api *echo.Group, awsServiceRegistry *aws.ServiceRegistry) error {
	// ハンドラーの初期化（ロガーを渡す）
	h, err := handler.NewHandler(awsServiceRegistry, r.logger)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}

	// oapi-codegenの自動ルーティング機能を使用
	openapi.RegisterHandlers(api, h)

	return nil
}

// setupStaticRoutes は静的ファイル配信を設定
func (r *Router) setupStaticRoutes() {
	r.echo.Static("/swagger", "static/swagger-ui")
	r.echo.File("/swagger", "static/swagger-ui/index.html")
	r.echo.File("/openapi.yaml", "openapi.yaml")
}

// GetEcho はEchoインスタンスを返す
func (r *Router) GetEcho() *echo.Echo {
	return r.echo
}
