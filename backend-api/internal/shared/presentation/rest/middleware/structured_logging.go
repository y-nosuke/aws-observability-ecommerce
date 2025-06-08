package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
)

// StructuredLoggingMiddleware は構造化ログミドルウェアを作成
func StructuredLoggingMiddleware(logger logging.Logger) echo.MiddlewareFunc {
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

			// ログデータを構築
			duration := time.Since(start)
			logData := logging.RequestLogData{
				Method:        c.Request().Method,
				Path:          c.Request().URL.Path,
				Query:         c.Request().URL.RawQuery,
				StatusCode:    c.Response().Status,
				RequestSize:   requestSize,
				ResponseSize:  resWrapper.size,
				Duration:      duration,
				UserAgent:     c.Request().UserAgent(),
				RemoteIP:      c.RealIP(),
				XForwardedFor: c.Request().Header.Get("X-Forwarded-For"),
				Referer:       c.Request().Referer(),
				ContentType:   c.Request().Header.Get("Content-Type"),
				Accept:        c.Request().Header.Get("Accept"),
				// ユーザー情報は認証実装後に追加
				CacheHit:         false, // キャッシュ実装後に追加
				DatabaseQueries:  0,     // DB監視実装後に追加
				ExternalAPICalls: 0,     // 外部API監視実装後に追加
			}

			// リクエストログを出力
			logger.LogRequest(c.Request().Context(), logData)

			return err
		}
	}
}

// responseWriter はレスポンスサイズを追跡するためのラッパー
// 他のmiddlewareで必要ならtypes.go等に共通化も検討
// 今回はこのファイル内に配置

type responseWriter struct {
	http.ResponseWriter
	size int64
}

func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += int64(size)
	return size, err
}
