package logging

import (
	"context"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/exp/slog"
)

// NewLogger 新しいslogロガーを作成
func NewLogger(level string) *slog.Logger {
	var logLevel slog.Level
	switch level {
	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError
	default:
		logLevel = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)

	return logger
}

// Middleware Echoフレームワーク用のロギングミドルウェア
func Middleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			res := c.Response()
			start := time.Now()

			// リクエストIDがない場合は生成
			requestID := req.Header.Get(echo.HeaderXRequestID)
			if requestID == "" {
				requestID = c.Response().Header().Get(echo.HeaderXRequestID)
			}

			// コンテキストにロガーを設定
			logger := slog.Default().With(
				"request_id", requestID,
				"method", req.Method,
				"path", req.URL.Path,
				"ip", c.RealIP(),
				"user_agent", req.UserAgent(),
			)

			c.SetRequest(req.WithContext(context.WithValue(req.Context(), "logger", logger)))

			// リクエストを処理
			err := next(c)

			// レスポンスのログを記録
			stop := time.Now()
			latency := stop.Sub(start).Milliseconds()

			logger.Info("Request completed",
				"status", res.Status,
				"latency_ms", latency,
				"bytes_in", req.ContentLength,
				"bytes_out", res.Size,
			)

			return err
		}
	}
}

// GetLogger コンテキストからロガーを取得
func GetLogger(ctx context.Context) *slog.Logger {
	if logger, ok := ctx.Value("logger").(*slog.Logger); ok {
		return logger
	}
	return slog.Default()
}
