package tracer

import (
	"context"
	"fmt"
	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
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
		return &fallbackTracer{tracer: otel.Tracer("aws-observability-ecommerce")}
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

// === ドメイン固有の便利関数 ===

// StartHandler はハンドラー層のスパンを開始
func StartHandler(ctx context.Context, operation string, domain string) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("handler.%s", operation)
	return Start(ctx, spanName, trace.WithAttributes(
		attribute.String("app.layer", "handler"),
		attribute.String("app.domain", domain),
		attribute.String("app.operation", operation),
	))
}

// StartUseCase はユースケース層のスパンを開始
func StartUseCase(ctx context.Context, operation string, domain string) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("usecase.%s", operation)
	return Start(ctx, spanName, trace.WithAttributes(
		attribute.String("app.layer", "usecase"),
		attribute.String("app.domain", domain),
		attribute.String("app.operation_name", operation),
	))
}

// StartRepository はリポジトリ層のスパンを開始
func StartRepository(ctx context.Context, operation string, domain string) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("repository.%s", operation)
	return Start(ctx, spanName, trace.WithAttributes(
		attribute.String("app.layer", "repository"),
		attribute.String("app.domain", domain),
		attribute.String("app.operation", operation),
	))
}

// StartDatabase はデータベース操作のスパンを開始
func StartDatabase(ctx context.Context, operation string, tableName string) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("db.%s", operation)
	return Start(ctx, spanName, trace.WithAttributes(
		attribute.String("app.layer", "database"),
		attribute.String("app.operation", operation),
		attribute.String("db.table", tableName),
	))
}

// StartExternalAPI は外部API呼び出しのスパンを開始
func StartExternalAPI(ctx context.Context, service string, operation string) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("external.%s.%s", service, operation)
	return Start(ctx, spanName, trace.WithAttributes(
		attribute.String("app.layer", "external"),
		attribute.String("app.service", service),
		attribute.String("app.operation", operation),
	))
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
