package logging

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	slogmulti "github.com/samber/slog-multi"
	"go.opentelemetry.io/contrib/bridges/otelslog"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// StructuredLogger ログ機能を提供
type StructuredLogger struct {
	slogger *slog.Logger
	config  config.LoggingConfig
}

// RequestIDKey はコンテキストキー
type contextKey string

const RequestIDKey contextKey = "request_id"

// NewLogger は新しいLoggerを作成
func NewLogger(cfg config.ObservabilityConfig) *StructuredLogger {
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

	return logger
}

// Info はInfoレベルのログを出力（シンプル版）
func (l *StructuredLogger) Info(ctx context.Context, msg string, args ...any) {
	l.logWithContext(ctx, slog.LevelInfo, msg, args...)
}

// Warn はWarnレベルのログを出力（シンプル版）
func (l *StructuredLogger) Warn(ctx context.Context, msg string, args ...any) {
	l.logWithContext(ctx, slog.LevelWarn, msg, args...)
}

// Error はErrorレベルのログを出力（シンプル版）
func (l *StructuredLogger) Error(ctx context.Context, msg string, args ...any) {
	l.logWithContext(ctx, slog.LevelError, msg, args...)
}

// Debug はDebugレベルのログを出力（シンプル版）
func (l *StructuredLogger) Debug(ctx context.Context, msg string, args ...any) {
	l.logWithContext(ctx, slog.LevelDebug, msg, args...)
}

// logWithContext はコンテキスト情報を付加してログ出力
func (l *StructuredLogger) logWithContext(ctx context.Context, level slog.Level, msg string, args ...any) {
	// 基本的な共通フィールドを自動で追加
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

	// トレース情報を追加（コンテキストから自動取得）
	if traceID := extractTraceID(ctx); traceID != "" {
		attrs = append(attrs, slog.Group("trace",
			slog.String("id", traceID),
		))
	}

	// ホスト情報を追加
	if hostname, err := os.Hostname(); err == nil && hostname != "" {
		attrs = append(attrs, slog.Group("host",
			slog.String("name", hostname),
		))
	}

	// ユーザー提供の引数をattrsに変換
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			key := fmt.Sprintf("%v", args[i])
			value := args[i+1]
			attrs = append(attrs, slog.Any(key, value))
		}
	}

	l.slogger.LogAttrs(ctx, level, msg, attrs...)
}

// extractTraceID はコンテキストからトレースIDを取得（OpenTelemetry対応）
func extractTraceID(ctx context.Context) string {
	// TODO: OpenTelemetryのSpanからトレースIDを取得する実装を追加
	// 現在は簡易実装
	if traceID, ok := ctx.Value("trace_id").(string); ok {
		return traceID
	}
	return ""
}

// === 便利なヘルパーメソッド ===

// InfoF はフォーマット付きでInfoログを出力
func (l *StructuredLogger) InfoF(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Info(ctx, msg)
}

// ErrorF はフォーマット付きでErrorログを出力
func (l *StructuredLogger) ErrorF(ctx context.Context, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.Error(ctx, msg)
}

// WithError はエラー情報を含むログを出力
func (l *StructuredLogger) WithError(ctx context.Context, msg string, err error, args ...any) {
	allArgs := append(args, "error", err.Error())
	if err != nil {
		allArgs = append(allArgs, "error_type", fmt.Sprintf("%T", err))
	}
	l.Error(ctx, msg, allArgs...)
}

// === ビジネスロジック用のヘルパー ===

// LogOperation は操作の開始/完了を記録
func (l *StructuredLogger) LogOperation(ctx context.Context, operation string, duration time.Duration, success bool, args ...any) {
	level := slog.LevelInfo
	if !success {
		level = slog.LevelError
	}

	allArgs := append(args,
		"operation", operation,
		"duration_ms", float64(duration.Nanoseconds())/1e6,
		"success", success,
		"log_type", "operation",
	)

	l.logWithContext(ctx, level, fmt.Sprintf("Operation %s completed", operation), allArgs...)
}

// LogHTTPRequest はHTTPリクエストをログ出力
func (l *StructuredLogger) LogHTTPRequest(ctx context.Context, method, path string, status int, duration time.Duration, args ...any) {
	allArgs := append(args,
		"http_method", method,
		"http_path", path,
		"http_status", status,
		"duration_ms", float64(duration.Nanoseconds())/1e6,
		"log_type", "http_request",
	)

	l.Info(ctx, "HTTP request processed", allArgs...)
}

// LogBusinessEvent はビジネスイベントをログ出力
func (l *StructuredLogger) LogBusinessEvent(ctx context.Context, event string, entityType string, entityID string, args ...any) {
	allArgs := append(args,
		"event", event,
		"entity_type", entityType,
		"entity_id", entityID,
		"log_type", "business_event",
	)

	l.Info(ctx, fmt.Sprintf("Business event: %s", event), allArgs...)
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
