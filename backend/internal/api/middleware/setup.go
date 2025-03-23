package middleware

import (
	"github.com/labstack/echo/v4"
	echomw "github.com/labstack/echo/v4/middleware"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/config"
)

// SetupMiddleware はすべてのミドルウェアを設定する
func SetupMiddleware(e *echo.Echo) {
	// 基本的なミドルウェア
	e.Use(echomw.Recover()) // パニック回復
	e.Use(echomw.CORS())    // CORS対応

	// カスタムミドルウェア
	e.Use(RequestIDMiddleware()) // リクエストID生成

	// リクエスト/レスポンスのログ記録
	loggerConfig := DefaultLoggerConfig
	// 開発環境では本文も記録
	if config.App.Environment == "development" {
		loggerConfig.LogRequestBody = true
		loggerConfig.LogResponseBody = true
	}
	e.Use(LoggerMiddleware(loggerConfig))

	// 認証情報のログ付与（認証後に実行）
	e.Use(AuthLoggerMiddleware())
}
