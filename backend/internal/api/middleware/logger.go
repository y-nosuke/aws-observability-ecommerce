package middleware

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/logger"
)

// コンテキストキー
const (
	RequestIDKey = "request_id"
	LoggerKey    = "logger"
)

// GetRequestID はEchoコンテキストからリクエストIDを取得する
func GetRequestID(c echo.Context) string {
	if requestID, ok := c.Get(RequestIDKey).(string); ok {
		return requestID
	}
	return ""
}

// SetRequestID はEchoコンテキストからリクエストIDを設定する
func SetRequestID(c echo.Context, r string) {
	c.Set(RequestIDKey, r)
}

// GetLogger はEchoコンテキストからロガーを取得する
func GetLogger(c echo.Context) *slog.Logger {
	if l, ok := c.Get(LoggerKey).(*slog.Logger); ok {
		return l
	}
	return logger.Logger(c.Request().Context()) // デフォルトロガーを返す
}

// SetLogger はEchoコンテキストにロガーを設定する
func SetLogger(c echo.Context, l *slog.Logger) {
	c.Set(LoggerKey, l)
}
