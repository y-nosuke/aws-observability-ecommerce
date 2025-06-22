package observability

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
)

// BaseObservabilityHelper は各層のヘルパーで共通する機能を提供する基底構造体
type BaseObservabilityHelper struct {
	ctx          context.Context
	span         trace.Span
	operationLog func(success bool, args ...any)
	layer        string // "handler", "usecase", "repository"
}

// startSpan は共通のspan作成処理
func (b *BaseObservabilityHelper) startSpan(ctx context.Context, operationName, layer string, additionalAttrs ...attribute.KeyValue) (context.Context, trace.Span) {
	// contextからdomainを自動取得
	domain := GetDomainFromContext(ctx)

	// spanNameを作成
	spanName := fmt.Sprintf("%s.%s", layer, operationName)

	// 基本属性を設定
	baseAttrs := []attribute.KeyValue{
		attribute.String("app.layer", layer),
		attribute.String("app.domain", domain),
		attribute.String("app.operation", operationName),
	}

	// 追加属性があれば結合
	if len(additionalAttrs) > 0 {
		baseAttrs = append(baseAttrs, additionalAttrs...)
	}

	// spanを作成
	spanCtx, span := tracer.Start(ctx, spanName, trace.WithAttributes(baseAttrs...))

	return spanCtx, span
}

// startSubSpan はサブspan作成処理（Step系で使用）
func (b *BaseObservabilityHelper) startSubSpan(stepName, subLayer string, additionalAttrs ...attribute.KeyValue) (context.Context, trace.Span) {
	// spanNameを作成
	spanName := fmt.Sprintf("%s.%s", subLayer, stepName)

	// 基本属性を設定
	baseAttrs := []attribute.KeyValue{
		attribute.String("app.layer", subLayer),
		attribute.String("app.operation", stepName),
	}

	// 追加属性があれば結合
	if len(additionalAttrs) > 0 {
		baseAttrs = append(baseAttrs, additionalAttrs...)
	}

	// spanを作成
	stepCtx, stepSpan := tracer.Start(b.ctx, spanName, trace.WithAttributes(baseAttrs...))

	return stepCtx, stepSpan
}

// initializeWithSpan はspan作成から初期化まで一貫して実行
func (b *BaseObservabilityHelper) initializeWithSpan(ctx context.Context, operationName, layer string, additionalAttrs ...attribute.KeyValue) {
	// span作成
	spanCtx, span := b.startSpan(ctx, operationName, layer, additionalAttrs...)

	// 基本初期化処理
	b.initializeBase(spanCtx, span, operationName, layer)
}

// initializeBase は各層のヘルパーの共通初期化処理
func (b *BaseObservabilityHelper) initializeBase(ctx context.Context, span trace.Span, operationName, layer string) {
	spanCtx := trace.ContextWithSpan(ctx, span)

	// contextからentityIDを自動取得
	if id := GetEntityIDFromContext(ctx); id > 0 {
		span.SetAttributes(attribute.Int("app.entity_id", id))
	}

	if entityType := GetEntityTypeFromContext(ctx); entityType != "" {
		span.SetAttributes(attribute.String("app.entity_type", entityType))
	}

	// contextからdomainを自動取得
	domain := GetDomainFromContext(ctx)

	// 操作ログを開始
	operationLog := logger.StartOperation(spanCtx, operationName,
		"layer", layer,
		"domain", domain,
	)

	b.ctx = spanCtx
	b.span = span
	b.operationLog = operationLog
	b.layer = layer
}

// Context は現在のコンテキストを返す
func (b *BaseObservabilityHelper) Context() context.Context {
	return b.ctx
}

// SetAttributes はスパンに属性を設定
func (b *BaseObservabilityHelper) SetAttributes(attrs ...attribute.KeyValue) {
	b.span.SetAttributes(attrs...)
}

// LogInfo は情報ログを記録
func (b *BaseObservabilityHelper) LogInfo(message string, args ...any) {
	logger.Info(b.ctx, message, args...)
}

// LogError はエラーログを記録し、スパンにエラー情報を設定
func (b *BaseObservabilityHelper) LogError(message string, err error, args ...any) {
	// エラーログを記録
	allArgs := append([]any{"layer", b.layer}, args...)
	logger.WithError(b.ctx, message, err, allArgs...)

	// スパンにエラー情報を記録
	b.span.RecordError(err)
	b.span.SetStatus(codes.Error, err.Error())
	b.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", b.layer+"_error"),
	)
}

// Finish は処理を完了
func (b *BaseObservabilityHelper) Finish(success bool, args ...any) {
	defer b.span.End()

	// 成功/失敗をスパンに記録
	b.span.SetAttributes(attribute.Bool("app.success", success))

	// 操作ログを完了
	b.operationLog(success, args...)
}

// FinishWithError はエラーで処理を完了
func (b *BaseObservabilityHelper) FinishWithError(err error, message string, args ...any) {
	defer b.span.End()

	// エラーログとスパン情報を記録
	b.LogError(message, err, args...)

	// 操作ログを失敗で完了
	errorArgs := append([]any{"error_type", b.layer + "_failure"}, args...)
	b.operationLog(false, errorArgs...)
}
