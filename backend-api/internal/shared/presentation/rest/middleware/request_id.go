package middleware

import (
	"context"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// RequestIDMiddleware はリクエストID生成ミドルウェアを作成
func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ヘッダーからリクエストIDを取得、なければ生成
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = "req_" + uuid.New().String()
			}

			// コンテキストにリクエストIDを設定
			ctx := context.WithValue(c.Request().Context(), logger.RequestIDKey, requestID)
			c.SetRequest(c.Request().WithContext(ctx))

			// レスポンスヘッダーにも設定
			c.Response().Header().Set("X-Request-ID", requestID)

			return next(c)
		}
	}
}
