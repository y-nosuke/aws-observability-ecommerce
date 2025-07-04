package middleware

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/labstack/echo/v4"
)

const ContextKeyRequestBody = "request.body"
const ContextKeyResponseBody = "response.body"

func RequestResponseCaptureMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ===== リクエスト情報の取得 =====
			req := c.Request()

			var reqBodyBytes []byte
			var readErr error
			if req.Body != nil {
				reqBodyBytes, readErr = io.ReadAll(req.Body)
				if readErr != nil {
					return fmt.Errorf("reading body: %w", readErr)
				}

				// 読み直しできるように戻す
				req.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes))
			}

			c.Set(ContextKeyRequestBody, &reqBodyBytes)

			// ===== レスポンスボディをキャプチャするための Writer ラッパー =====
			res := c.Response()

			var resBodyBytes []byte
			bodyWriter := &responseWriter{
				ResponseWriter: res.Writer,
				body:           &resBodyBytes,
			}
			res.Writer = bodyWriter

			c.Set(ContextKeyResponseBody, &resBodyBytes)

			// ハンドラー実行
			err := next(c)

			return err
		}
	}
}

// responseWriter はレスポンスを追跡するためのラッパー
type responseWriter struct {
	http.ResponseWriter
	body       *[]byte
	size       int64
	statusCode int
}

// Write はレスポンスボディを追跡します
func (w *responseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	*w.body = append(*w.body, b...)
	w.size += int64(size)
	return size, err
}

// WriteHeader はレスポンスのステータスコードを追跡します
func (w *responseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

// Flush はバッファをフラッシュします（SSEやストリーミングで必要）
func (w *responseWriter) Flush() {
	if flusher, ok := w.ResponseWriter.(http.Flusher); ok {
		flusher.Flush()
	}
}

// Hijack はコネクションをハイジャックします（WebSocketやgRPCで必要）
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hijacker, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hijacker.Hijack()
	}
	return nil, nil, fmt.Errorf("response writer does not support hijacking")
}
