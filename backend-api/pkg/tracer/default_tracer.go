package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// DefaultTracer は実際のトレーサーを使用するデフォルト実装
type DefaultTracer struct {
	tracer trace.Tracer
}

// NewDefaultTracer は新しいDefaultTracerを作成
func NewDefaultTracer(tracer trace.Tracer) *DefaultTracer {
	return &DefaultTracer{
		tracer: tracer,
	}
}

// Start はスパンを開始します
func (t *DefaultTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return t.tracer.Start(ctx, spanName, opts...)
}

// GetTracer は実際のOpenTelemetryトレーサーを返します
func (t *DefaultTracer) GetTracer() trace.Tracer {
	return t.tracer
}
