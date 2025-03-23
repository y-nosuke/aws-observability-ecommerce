package middleware

import (
	"bytes"
	"io"
	"net/http"
	"time"

	"slices"

	"github.com/labstack/echo/v4"
)

// LoggerConfig はロギングミドルウェアの設定
type LoggerConfig struct {
	// ログに含めないURLパス（ヘルスチェックなど）
	SkipPaths []string
	// リクエスト本文をログに含めるかどうか
	LogRequestBody bool
	// レスポンス本文をログに含めるかどうか
	LogResponseBody bool
	// 最大本文サイズ（バイト単位）
	MaxBodySize int
}

// DefaultLoggerConfig はLoggerConfigのデフォルト値
var DefaultLoggerConfig = LoggerConfig{
	SkipPaths:       []string{"/api/health", "/api/metrics"},
	LogRequestBody:  false,
	LogResponseBody: false,
	MaxBodySize:     1024, // 1KB
}

// LoggerMiddleware はリクエストとレスポンスをログに記録するミドルウェア
func LoggerMiddleware(config LoggerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ヘルスチェックなど特定のパスはスキップ
			path := c.Request().URL.Path
			if slices.Contains(config.SkipPaths, path) {
				return next(c)
			}

			start := time.Now()

			// リクエスト情報を記録
			log := GetLogger(c)
			reqBody := ""
			if config.LogRequestBody && c.Request().Body != nil {
				// リクエスト本文を読み取り、バッファに保存
				buf, err := io.ReadAll(io.LimitReader(c.Request().Body, int64(config.MaxBodySize)))
				if err != nil {
					return err
				}
				reqBody = string(buf)
				// 本文を復元
				c.Request().Body = io.NopCloser(bytes.NewBuffer(buf))
			}

			req := c.Request()
			res := c.Response()

			// リクエスト情報のログ記録
			log.Info("API request received",
				"method", req.Method,
				"path", req.URL.Path,
				"query", req.URL.RawQuery,
				"remote_ip", c.RealIP(),
				"user_agent", req.UserAgent())

			if config.LogRequestBody {
				log.Debug("Request body", "body", reqBody)
			}

			// レスポンスをキャプチャするためのレスポンスライター
			resBody := new(bytes.Buffer)
			mw := io.MultiWriter(res.Writer, resBody)
			writer := &bodyDumpResponseWriter{
				ResponseWriter: res.Writer,
				Writer:         mw,
			}
			res.Writer = writer

			// 次のハンドラーを呼び出し
			err := next(c)

			// レスポンス情報を記録
			elapsed := time.Since(start)
			statusCode := res.Status
			responseSize := res.Size

			// ログレベルをステータスコードに基づいて決定
			var logFunc func(msg string, args ...interface{})
			if statusCode >= 500 {
				logFunc = log.Error
			} else if statusCode >= 400 {
				logFunc = log.Warn
			} else {
				logFunc = log.Info
			}

			logFunc("API request completed",
				"method", req.Method,
				"path", req.URL.Path,
				"status", statusCode,
				"elapsed_ms", elapsed.Milliseconds(),
				"size", responseSize)

			// レスポンス本文のログ記録（必要な場合）
			if config.LogResponseBody {
				var body string
				if resBody.Len() <= config.MaxBodySize {
					body = resBody.String()
				} else {
					body = resBody.String()[:config.MaxBodySize]
				}
				log.Debug("Response body", "body", body)
			}

			return err
		}
	}
}

// bodyDumpResponseWriter はレスポンス本文をキャプチャするためのレスポンスライター
type bodyDumpResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}
