package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// LoggingMiddleware 構造化ログミドルウェアを作成
func LoggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// リクエストボディのサイズを取得
			var requestSize int64
			if c.Request().Body != nil {
				bodyBytes, err := io.ReadAll(c.Request().Body)
				if err == nil {
					requestSize = int64(len(bodyBytes))
					c.Request().Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				}
			}

			// レスポンスライターをラップしてサイズを追跡
			resWrapper := &responseWriter{ResponseWriter: c.Response().Writer}
			c.Response().Writer = resWrapper

			// 次のハンドラーを実行
			err := next(c)

			// ログ出力
			duration := time.Since(start)
			ctx := c.Request().Context()

			// HTTPリクエストログを出力（パッケージレベル関数を使用）
			logger.LogHTTPRequest(ctx,
				c.Request().Method,
				c.Request().URL.Path,
				c.Response().Status,
				duration,
				"request_size_bytes", requestSize,
				"response_size_bytes", resWrapper.size,
				"user_agent", c.Request().UserAgent(),
				"remote_ip", c.RealIP(),
				"x_forwarded_for", c.Request().Header.Get("X-Forwarded-For"),
				"referer", c.Request().Referer(),
				"content_type", c.Request().Header.Get("Content-Type"),
				"accept", c.Request().Header.Get("Accept"),
				"query", c.Request().URL.RawQuery,
				"cache_hit", false, // TODO: キャッシュ実装後に動的に設定
				"database_queries", 0, // TODO: DB監視実装後に動的に設定
				"external_api_calls", 0) // TODO: 外部API監視実装後に動的に設定

			return err
		}
	}
}

// responseWriter はレスポンスサイズを追跡するためのラッパー
type responseWriter struct {
	http.ResponseWriter
	size int64
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += int64(size)
	return size, err
}
