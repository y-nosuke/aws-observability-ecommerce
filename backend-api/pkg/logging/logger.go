package logging

import (
	"context"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"

	slogmulti "github.com/samber/slog-multi"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// Logger は構造化ログ機能のインターフェース
type Logger interface {
	Info(ctx context.Context, msg string, fields ...Field)
	Warn(ctx context.Context, msg string, fields ...Field)
	Error(ctx context.Context, msg string, err error, fields ...Field)
	Debug(ctx context.Context, msg string, fields ...Field)

	// 特定ログタイプ向けヘルパー
	LogRequest(ctx context.Context, req RequestLogData)
	LogError(ctx context.Context, err error, errorCtx ErrorContext)
	LogApplication(ctx context.Context, op ApplicationOperation)
}

// Field はログフィールドの構造体
type Field struct {
	Key   string
	Value interface{}
}

// StructuredLogger は構造化ログの実装
type StructuredLogger struct {
	slogger *slog.Logger
	config  config.LoggingConfig
}

// RequestIDKey はコンテキストキー
// staticcheck対策: 独自型を使う
type contextKey string

const RequestIDKey contextKey = "request_id"

// NewLogger は新しいStructuredLoggerを作成します
func NewLogger(cfg config.ObservabilityConfig) (Logger, error) {
	var handler slog.Handler

	opts := &slog.HandlerOptions{
		Level:     parseLogLevel(cfg.Logging.Level),
		AddSource: true,
	}

	if cfg.Logging.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	// OpenTelemetry ブリッジを使用
	if cfg.Logging.EnableOTel {
		handler = slogmulti.Fanout(
			otelslog.NewHandler("aws-observability-ecommerce"),
			handler,
		)
	}

	logger := &StructuredLogger{
		slogger: slog.New(handler),
		config:  cfg.Logging,
	}

	return logger, nil
}

// Info はInfoレベルのログを出力します
func (l *StructuredLogger) Info(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, slog.LevelInfo, msg, fields...)
}

// Warn はWarnレベルのログを出力します
func (l *StructuredLogger) Warn(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, slog.LevelWarn, msg, fields...)
}

// Error はErrorレベルのログを出力します
func (l *StructuredLogger) Error(ctx context.Context, msg string, err error, fields ...Field) {
	allFields := append(fields, Field{Key: "error", Value: err.Error()})
	l.log(ctx, slog.LevelError, msg, allFields...)
}

// Debug はDebugレベルのログを出力します
func (l *StructuredLogger) Debug(ctx context.Context, msg string, fields ...Field) {
	l.log(ctx, slog.LevelDebug, msg, fields...)
}

// log は実際のログ出力を行う内部メソッド
func (l *StructuredLogger) log(ctx context.Context, level slog.Level, msg string, fields ...Field) {
	attrs := make([]slog.Attr, 0, len(fields)+10) // 余裕を持ったサイズ

	// 共通フィールドを追加
	attrs = append(attrs, l.buildCommonFields(ctx)...)

	// カスタムフィールドを追加
	for _, field := range fields {
		attrs = append(attrs, slog.Any(field.Key, field.Value))
	}

	l.slogger.LogAttrs(ctx, level, msg, attrs...)
}

// buildCommonFields は共通フィールドを構築します
func (l *StructuredLogger) buildCommonFields(ctx context.Context) []slog.Attr {
	attrs := []slog.Attr{
		slog.Group("service",
			slog.String("name", config.App.Name),
			slog.String("version", config.App.Version),
			slog.String("environment", config.App.Environment),
		),
	}

	// リクエストIDを追加
	if reqID := extractRequestID(ctx); reqID != "" {
		attrs = append(attrs, slog.Group("request",
			slog.String("id", reqID),
		))
	}

	// ホスト情報を追加
	hostname, err := os.Hostname()
	if err == nil && hostname != "" {
		attrs = append(attrs, slog.Group("host",
			slog.String("name", hostname),
		))
	}

	return attrs
}

// extractRequestID はコンテキストからリクエストIDを抽出します
func extractRequestID(ctx context.Context) string {
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}

// parseLogLevel は文字列からslog.Levelに変換します
func parseLogLevel(level string) slog.Level {
	switch level {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
