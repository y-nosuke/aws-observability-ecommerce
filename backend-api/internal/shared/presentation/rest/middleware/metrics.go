package middleware

import (
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// MetricsMiddleware はHTTPメトリクス収集のミドルウェアを作成します
func MetricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// リクエストサイズを取得
			reqBodyBytes, ok := c.Get(ContextKeyRequestBody).([]byte)
			if !ok {
				reqBodyBytes = []byte{}
			}

			// ハンドラー実行
			err := next(c)

			// メトリクス記録（グローバル関数を使用）
			duration := time.Since(start)

			resBodyBytes, ok := c.Get(ContextKeyResponseBody).([]byte)
			if !ok {
				resBodyBytes = []byte{}
			}
			ctx := c.Request().Context()
			otel.CountRequestsTotal(ctx)
			otel.RecordRequestDuration(ctx, duration)
			otel.RecordRequestSizeBytes(ctx, int64(len(reqBodyBytes)))
			otel.RecordResponseSizeBytes(ctx, int64(len(resBodyBytes)))
			otel.CountErrorsTotal(ctx, c.Response().Status)

			return err
		}
	}
}
