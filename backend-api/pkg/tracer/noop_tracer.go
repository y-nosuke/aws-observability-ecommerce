package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

// NoopTracer は何もしないトレーサー実装
type NoopTracer struct{}

// NewNoopTracer は新しいNoopTracerを作成
func NewNoopTracer() *NoopTracer {
	return &NoopTracer{}
}

// Start は何もしないスパンを返します
func (n *NoopTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return ctx, trace.SpanFromContext(ctx)
}

// GetTracer はNoopTracerを返します
func (n *NoopTracer) GetTracer() trace.Tracer {
	return noop.NewTracerProvider().Tracer("noop")
}
