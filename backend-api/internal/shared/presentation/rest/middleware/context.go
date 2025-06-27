package middleware

import (
	"context"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// ContextMiddleware はビジネスコンテキスト情報を自動抽出するミドルウェア
func ContextMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			attrs := map[string]string{
				"http.method": c.Request().Method,
				"http.route":  c.Path(),
				"http.target": c.Request().URL.String(),
				"http.host":   c.Request().Host,
				// 必要なら追加
			}

			ctx := context.WithValue(c.Request().Context(), otel.ContextKeyAttrs, attrs)

			c.SetRequest(c.Request().WithContext(ctx))

			return next(c)
		}
	}
}
