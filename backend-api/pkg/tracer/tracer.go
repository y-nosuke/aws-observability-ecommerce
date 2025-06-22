package tracer

import (
	"context"

	"go.opentelemetry.io/otel/trace"
)

// Tracer はトレーシング機能のインターフェース
type Tracer interface {
	// Start はスパンを開始します
	Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span)

	// GetTracer は実際のOpenTelemetryトレーサーを返します
	GetTracer() trace.Tracer
}
