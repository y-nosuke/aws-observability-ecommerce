package middleware

import (
	"github.com/labstack/echo/v4"
)

// NewMetricsMiddleware はメトリクスミドルウェアを作成
// 将来的にPrometheusやその他のメトリクス収集ツールと統合可能
func NewMetricsMiddleware() echo.MiddlewareFunc {
	// TODO: メトリクス収集の実装
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// メトリクス収集のロジックをここに実装
			return next(c)
		}
	}
}

// NewRateLimitMiddleware はレート制限ミドルウェアを作成
// 将来的にRedisベースのレート制限を実装可能
func NewRateLimitMiddleware() echo.MiddlewareFunc {
	// TODO: レート制限の実装
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// レート制限のロジックをここに実装
			return next(c)
		}
	}
}
