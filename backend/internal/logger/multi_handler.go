package logger

import (
	"context"
	"fmt"
	"log/slog"
)

// multiHandler は複数のハンドラーに同時にログを送信するslogハンドラー
type multiHandler struct {
	handlers []slog.Handler
}

// newMultiHandler は新しいマルチハンドラーを作成します
func newMultiHandler(handlers ...slog.Handler) *multiHandler {
	fmt.Printf("multiHandler created. handlers: %+v\n", handlers)
	return &multiHandler{handlers: handlers}
}

// Enabled は少なくとも1つのハンドラーが有効であるかどうかを返します
func (h *multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle は全てのハンドラーでログレコードを処理します
func (h *multiHandler) Handle(ctx context.Context, record slog.Record) error {
	for _, handler := range h.handlers {
		if err := handler.Handle(ctx, record); err != nil {
			return err
		}
	}
	return nil
}

// WithAttrs は属性を持つ新しいマルチハンドラーを返します
func (h *multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	handlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.WithAttrs(attrs)
	}
	return newMultiHandler(handlers...)
}

// WithGroup はグループ化された属性を持つ新しいマルチハンドラーを返します
func (h *multiHandler) WithGroup(name string) slog.Handler {
	handlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.WithGroup(name)
	}
	return newMultiHandler(handlers...)
}
