package router

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/di"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/handler"
	customMiddleware "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/middleware"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// NewRouter は新しいEchoインスタンスを作成し、全てのルーティングを設定
func NewRouter(container *di.AppContainer) (*echo.Echo, error) {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.HTTPErrorHandler = CustomHTTPErrorHandler

	// 1. ミドルウェアの設定
	if err := setupMiddleware(e); err != nil {
		return nil, fmt.Errorf("setup middleware error: %w", err)
	}

	// 2. 静的ファイル配信
	setupStaticRoutes(e)

	// 3. OpenAPI仕様に基づくAPIルーティング
	api := e.Group("/api")

	if err := setupAPIRoutes(api, container); err != nil {
		return nil, err
	}

	return e, nil
}

// setupMiddleware は共通ミドルウェアを設定
func setupMiddleware(e *echo.Echo) error {
	// 基本的なミドルウェア
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.BodyLimit("10M"))

	// ビジネスコンテキスト抽出ミドルウェア
	e.Use(customMiddleware.ContextMiddleware())
	e.Use(customMiddleware.RequestResponseCaptureMiddleware())

	// OpenTelemetry公式のEchoインストゥルメンテーション
	e.Use(otelecho.Middleware("aws-observability-ecommerce-backend-api"))

	// メトリクス収集ミドルウェア
	e.Use(customMiddleware.MetricsMiddleware())

	// ログミドルウェア
	e.Use(customMiddleware.LoggingMiddleware())

	return nil
}

// setupAPIRoutes はoapi-codegenを使用してAPIルーティングを設定
func setupAPIRoutes(api *echo.Group, container *di.AppContainer) error {
	// ハンドラーの初期化（DIコンテナから取得）
	h, err := handler.NewHandler(container)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)
	}

	// oapi-codegenの自動ルーティング機能を使用
	openapi.RegisterHandlers(api, h)

	return nil
}

// setupStaticRoutes は静的ファイル配信を設定
func setupStaticRoutes(e *echo.Echo) {
	e.Static("/swagger", "static/swagger-ui")
	e.File("/swagger", "static/swagger-ui/index.html")
	e.File("/openapi.yaml", "openapi.yaml")
}
