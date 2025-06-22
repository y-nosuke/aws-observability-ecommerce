package tracer

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
)

// グローバルトレーサーインスタンス
var (
	globalTracer Tracer
	initOnce     sync.Once
)

// InitWithProvider はOpenTelemetryプロバイダーを使用してグローバルトレーサーを初期化します
func InitWithProvider(provider *sdktrace.TracerProvider) error {
	var initError error
	initOnce.Do(func() {
		if provider == nil {
			globalTracer = NewNoopTracer()
			return
		}

		// グローバルTracerProviderを設定
		otel.SetTracerProvider(provider)

		// Tracerを取得
		tracer := provider.Tracer("github.com/y-nosuke/aws-observability-ecommerce/backend-api")

		globalTracer = NewDefaultTracer(tracer)
	})
	return initError
}

// SetGlobalTracer はグローバルトレーサーを直接設定します（テスト用）
func SetGlobalTracer(tracer Tracer) {
	globalTracer = tracer
}

// getGlobalTracer はグローバルトレーサーを取得します（内部用）
func getGlobalTracer() Tracer {
	if globalTracer == nil {
		// フォールバック：デフォルトのOTelトレーサー
		return &fallbackTracer{tracer: otel.Tracer("github.com/y-nosuke/aws-observability-ecommerce/backend-api")}
	}
	return globalTracer
}

// === パッケージレベルの便利関数 ===

// Start はスパンを開始します
func Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return getGlobalTracer().Start(ctx, spanName, opts...)
}

// GetTracer は実際のOpenTelemetryトレーサーを返します
func GetTracer() trace.Tracer {
	return getGlobalTracer().GetTracer()
}

// fallbackTracer はフォールバック用の実装
type fallbackTracer struct {
	tracer trace.Tracer
}

func (f *fallbackTracer) Start(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return f.tracer.Start(ctx, spanName, opts...)
}

func (f *fallbackTracer) GetTracer() trace.Tracer {
	return f.tracer
}
