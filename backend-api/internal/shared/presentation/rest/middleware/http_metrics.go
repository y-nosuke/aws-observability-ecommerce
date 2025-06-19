package middleware

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/metrics"
)

// HTTPMetricsMiddleware はHTTPメトリクス収集のミドルウェアを作成します
func HTTPMetricsMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			// リクエストサイズを取得
			requestSize := getRequestSize(c.Request())

			// レスポンス監視用ラッパー
			resWrapper := &metricsResponseWriter{
				ResponseWriter: c.Response().Writer,
				statusCode:     200, // デフォルトは200
			}
			c.Response().Writer = resWrapper

			// ハンドラー実行
			err := next(c)

			// メトリクス記録（グローバル関数を使用）
			duration := time.Since(start)
			route := getRoutePattern(c)

			metrics.RecordHTTPRequest(
				c.Request().Method,
				route,
				resWrapper.statusCode,
				duration,
				requestSize,
				resWrapper.size,
			)

			return err
		}
	}
}

// getRequestSize はリクエストボディのサイズを取得します
func getRequestSize(req *http.Request) int64 {
	if req.Body == nil {
		return 0
	}

	// Content-Lengthヘッダーがある場合はそれを使用
	if req.ContentLength > 0 {
		return req.ContentLength
	}

	// Content-Lengthが不明な場合は実際に読み取る（小さいサイズのみ）
	const maxSize = 1024 * 1024 // 1MB制限
	bodyBytes, err := io.ReadAll(io.LimitReader(req.Body, maxSize))
	if err != nil {
		return 0
	}

	// ボディを復元
	req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	return int64(len(bodyBytes))
}

// metricsResponseWriter はレスポンスサイズとステータスコードを追跡するためのラッパー
type metricsResponseWriter struct {
	http.ResponseWriter
	size       int64
	statusCode int
}

func (w *metricsResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += int64(size)
	return size, err
}

func (w *metricsResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Flush はバッファをフラッシュします（http.Flusherインターフェース対応）
func (w *metricsResponseWriter) Flush() {
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

// Hijack はコネクションをハイジャックします（http.Hijackerインターフェース対応）
func (w *metricsResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, fmt.Errorf("response writer does not support hijacking")
}
