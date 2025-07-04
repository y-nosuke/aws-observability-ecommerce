package otel

import (
	"context"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"go.opentelemetry.io/contrib/bridges/otelslog"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// TraceHandler はslog.Handlerをラップして、トレースIDとスパンIDを追加します
type TraceHandler struct {
	slog.Handler
}

// NewTraceHandler はTraceHandlerを初期化します
func NewTraceHandler(h slog.Handler) *TraceHandler {
	return &TraceHandler{h}
}

// Handle はslog.Recordを処理します
func (h *TraceHandler) Handle(ctx context.Context, r slog.Record) error {
	span := trace.SpanFromContext(ctx)
	sc := span.SpanContext()
	if sc.IsValid() {
		r.AddAttrs(slog.String("trace_id", sc.TraceID().String()))
		r.AddAttrs(slog.String("span_id", sc.SpanID().String()))
	}
	return h.Handler.Handle(ctx, r)
}

// NewLogger はロガーを初期化します
func NewLogger(provider *sdklog.LoggerProvider, cfg config.LoggingConfig) *slog.Logger {
	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     parseLogLevel(cfg.Level),
	}
	var handler slog.Handler
	if cfg.Format == "json" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	handler = NewTraceHandler(handler)
	// OpenTelemetry ブリッジを使用
	if cfg.EnableOTel {
		handler = slogmulti.Fanout(
			handler,
			otelslog.NewHandler(utils.GetModulePath(), otelslog.WithLoggerProvider(provider)),
		)
	}

	logger := slog.New(handler)

	slog.SetDefault(logger)

	return logger
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
