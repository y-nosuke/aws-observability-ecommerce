package middleware

import (
	"github.com/labstack/echo/v4"
)

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
