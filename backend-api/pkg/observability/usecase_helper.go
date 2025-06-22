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

// UseCaseHelper は UseCase 層でのトレース処理を簡素化するヘルパー
type UseCaseHelper struct {
	ctx          context.Context
	span         trace.Span
	operationLog func(success bool, args ...any)
}

// StartUseCase は UseCase のトレースを開始
func StartUseCase(ctx context.Context, operationName string) *UseCaseHelper {
	// contextからdomainを自動取得
	domain := GetDomainFromContext(ctx)

	// 既存のStartUseCase関数を使用
	spanCtx, span := tracer.StartUseCase(ctx, operationName, domain)

	// contextからentityIDを自動取得
	span.SetAttributes(attribute.Int("app.entity_id", GetEntityIDFromContext(ctx)))

	// 操作ログを開始
	operationLog := logger.StartOperation(spanCtx, operationName,
		"layer", "usecase",
		"domain", domain,
	)

	return &UseCaseHelper{
		ctx:          spanCtx,
		span:         span,
		operationLog: operationLog,
	}
}

// Context は現在のコンテキストを返す
func (u *UseCaseHelper) Context() context.Context {
	return u.ctx
}

// SetAttributes はスパンに属性を設定
func (u *UseCaseHelper) SetAttributes(attrs ...attribute.KeyValue) {
	u.span.SetAttributes(attrs...)
}

// AddStep は処理ステップを記録
func (u *UseCaseHelper) AddStep(stepName string, fn func(context.Context) error) error {
	stepCtx, stepSpan := tracer.Start(u.ctx, fmt.Sprintf("usecase.%s", stepName))
	defer stepSpan.End()

	err := fn(stepCtx)
	if err != nil {
		stepSpan.RecordError(err)
		stepSpan.SetStatus(codes.Error, err.Error())
		stepSpan.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.step", stepName),
		)
	}

	return err
}

// LogInfo は情報ログを記録
func (u *UseCaseHelper) LogInfo(message string, args ...any) {
	logger.Info(u.ctx, message, args...)
}

// LogError はエラーログを記録し、スパンにエラー情報を設定
func (u *UseCaseHelper) LogError(message string, err error, args ...any) {
	// エラーログを記録
	allArgs := append([]any{"layer", "usecase"}, args...)
	logger.WithError(u.ctx, message, err, allArgs...)

	// スパンにエラー情報を記録
	u.span.RecordError(err)
	u.span.SetStatus(codes.Error, err.Error())
	u.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", "usecase_error"),
	)
}

// Finish は UseCase の処理を完了し、ログとスパンを閉じる
func (u *UseCaseHelper) Finish(success bool, args ...any) {
	defer u.span.End()

	// 成功/失敗をスパンに記録
	u.span.SetAttributes(attribute.Bool("app.success", success))

	// 操作ログを完了
	u.operationLog(success, args...)
}

// FinishWithError はエラーで UseCase の処理を完了
func (u *UseCaseHelper) FinishWithError(err error, message string, args ...any) {
	defer u.span.End()

	// エラーログとスパン情報を記録
	u.LogError(message, err, args...)

	// 操作ログを失敗で完了
	errorArgs := append([]any{"error_type", "usecase_failure"}, args...)
	u.operationLog(false, errorArgs...)
}

// ValidationError はバリデーションエラーを記録
func (u *UseCaseHelper) ValidationError(message string, field string, value any) error {
	err := fmt.Errorf("validation error: %s", message)

	// エラーログを記録
	logger.WithError(u.ctx, message, err,
		"layer", "usecase",
		"validation_error", "field_validation",
		"field", field,
		"value", value,
	)

	// スパンにバリデーションエラー情報を記録
	u.span.RecordError(err)
	u.span.SetStatus(codes.Error, err.Error())
	u.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", "validation_error"),
		attribute.String("validation.field", field),
		attribute.String("validation.value", fmt.Sprintf("%v", value)),
	)

	return err
}

// BusinessEvent はビジネスイベントを記録
func (u *UseCaseHelper) BusinessEvent(eventName, entityType, entityID string, args ...any) {
	logger.LogBusinessEvent(u.ctx, eventName, entityType, entityID, args...)
	u.span.SetAttributes(
		attribute.String("app.business_event", eventName),
		attribute.String("app.business_entity_type", entityType),
		attribute.String("app.business_entity_id", entityID),
	)
}
