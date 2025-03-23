package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"
)

// Config はロガーの設定を表す構造体
type Config struct {
	Environment string
	LogLevel    string
	ServiceName string
	Version     string
}

// ロガーのインスタンスを格納するグローバル変数
var defaultLogger *slog.Logger

// Init はロガーを初期化する
func Init(cfg Config) *slog.Logger {
	// ログレベルの設定
	var level slog.Level
	switch cfg.LogLevel {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	// 出力先の設定（本番環境では別の書き込み先を設定することもある）
	var w io.Writer = os.Stdout

	// JSONハンドラーのオプション設定
	opts := &slog.HandlerOptions{
		Level: level,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// タイムスタンプのフォーマット変更
			if a.Key == "time" {
				if t, ok := a.Value.Any().(time.Time); ok {
					a.Value = slog.StringValue(t.Format(time.RFC3339))
				}
			}
			return a
		},
	}

	// JSONハンドラーの作成
	var handler slog.Handler
	handler = slog.NewJSONHandler(w, opts)

	// アプリケーション全体の共通属性を持つハンドラーをラップ
	handler = NewContextHandler(handler, map[string]any{
		"service":     cfg.ServiceName,
		"environment": cfg.Environment,
		"version":     cfg.Version,
	})

	// ロガーの作成と設定
	logger := slog.New(handler)
	slog.SetDefault(logger)
	defaultLogger = logger

	return logger
}

// Logger は現在のコンテキストに基づいてロガーを返す
func Logger(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return defaultLogger
	}

	// コンテキストからロガーを取得
	if logger, ok := ctx.Value(loggerKey{}).(*slog.Logger); ok {
		return logger
	}

	return defaultLogger
}

// WithLogger は指定されたロガーをコンテキストに追加する
func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// コンテキストキー
type loggerKey struct{}

// カスタムコンテキストハンドラー
type contextHandler struct {
	handler slog.Handler
	attrs   []slog.Attr
}

// NewContextHandler は共通属性を持つハンドラーを作成する
func NewContextHandler(handler slog.Handler, attrs map[string]interface{}) slog.Handler {
	slogAttrs := make([]slog.Attr, 0, len(attrs))
	for k, v := range attrs {
		slogAttrs = append(slogAttrs, slog.Any(k, v))
	}
	return &contextHandler{
		handler: handler,
		attrs:   slogAttrs,
	}
}

// Enabled はハンドラーのEnabled関数を呼び出す
func (h *contextHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

// Handle はすべてのログレコードに共通属性を追加する
func (h *contextHandler) Handle(ctx context.Context, r slog.Record) error {
	// 共通属性をレコードに追加
	for _, attr := range h.attrs {
		r.AddAttrs(attr)
	}
	return h.handler.Handle(ctx, r)
}

// WithAttrs は新しい属性を持つハンドラーを返す
func (h *contextHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &contextHandler{
		handler: h.handler.WithAttrs(attrs),
		attrs:   h.attrs,
	}
}

// WithGroup はグループを持つハンドラーを返す
func (h *contextHandler) WithGroup(name string) slog.Handler {
	return &contextHandler{
		handler: h.handler.WithGroup(name),
		attrs:   h.attrs,
	}
}
