package logger

import (
	"context"
	"time"
)

type Logger interface {
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Debug(ctx context.Context, msg string, args ...any)
	InfoF(ctx context.Context, format string, args ...any)
	ErrorF(ctx context.Context, format string, args ...any)
	WithError(ctx context.Context, msg string, err error, args ...any)
	LogOperation(ctx context.Context, operation string, duration time.Duration, success bool, args ...any)
	LogHTTPRequest(ctx context.Context, method string, path string, status int, duration time.Duration, args ...any)
	LogBusinessEvent(ctx context.Context, event string, entityType string, id string, args ...any)
}
