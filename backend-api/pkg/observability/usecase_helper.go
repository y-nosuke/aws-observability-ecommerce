package observability

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// UseCaseHelper は UseCase 層でのトレース処理を簡素化するヘルパー
type UseCaseHelper struct {
	BaseObservabilityHelper // 共通機能を埋め込み
}

// StartUseCase は UseCase のトレースを開始
func StartUseCase(ctx context.Context, operationName string) *UseCaseHelper {
	helper := &UseCaseHelper{}

	// span作成から初期化まで一貫実行
	helper.initializeWithSpan(ctx, operationName, "usecase")

	return helper
}

// AddStep は処理ステップを記録
func (u *UseCaseHelper) AddStep(stepName string, fn func(context.Context) error) error {
	// サブspanを作成（追加属性なし）
	stepCtx, stepSpan := u.startSubSpan(stepName, "usecase")
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
