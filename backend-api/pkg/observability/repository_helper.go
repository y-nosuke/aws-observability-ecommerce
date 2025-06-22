package observability

import (
	"context"
	"fmt"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// RepositoryHelper は Repository 層でのトレース処理を簡素化するヘルパー
type RepositoryHelper struct {
	BaseObservabilityHelper // 共通機能を埋め込み
}

// StartRepository は Repository のトレースを開始
func StartRepository(ctx context.Context, operationName string) *RepositoryHelper {
	helper := &RepositoryHelper{}

	// span作成から初期化まで一貫実行
	helper.initializeWithSpan(ctx, operationName, "repository")

	return helper
}

// RecordDatabaseOperation はデータベース操作情報を記録
func (r *RepositoryHelper) RecordDatabaseOperation(tableName, operation string, recordCount int) {
	r.span.SetAttributes(
		attribute.String("db.table", tableName),
		attribute.String("db.operation", operation),
		attribute.Int("db.record_count", recordCount),
	)
}

// RecordQuery はクエリ情報を記録（本番環境ではSQLログは慎重に）
func (r *RepositoryHelper) RecordQuery(operation string, affectedRows int64) {
	r.span.SetAttributes(
		attribute.String("db.operation.type", operation),
		attribute.Int64("db.rows_affected", affectedRows),
	)
}

// AddDatabaseStep はデータベース操作ステップを記録
func (r *RepositoryHelper) AddDatabaseStep(stepName, tableName string, fn func(context.Context) error) error {
	// DB固有の追加属性を準備
	dbAttrs := []attribute.KeyValue{
		attribute.String("db.table", tableName),
	}

	// サブspanを作成
	stepCtx, stepSpan := r.startSubSpan(stepName, "database", dbAttrs...)
	defer stepSpan.End()

	err := fn(stepCtx)
	if err != nil {
		stepSpan.RecordError(err)
		stepSpan.SetStatus(codes.Error, err.Error())
		stepSpan.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.step", stepName),
			attribute.String("db.table", tableName),
		)
	}

	return err
}

// RecordConstraintError は制約エラーを記録
func (r *RepositoryHelper) RecordConstraintError(err error, constraintType, field string) {
	// エラーログを記録
	logger.WithError(r.ctx, "Database constraint error", err,
		"layer", "repository",
		"constraint_error", constraintType,
		"field", field,
	)

	// スパンに制約エラー情報を記録
	r.span.RecordError(err)
	r.span.SetStatus(codes.Error, err.Error())
	r.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", "constraint_error"),
		attribute.String("db.constraint_type", constraintType),
		attribute.String("db.constraint_field", field),
	)
}

// RecordNotFoundError はレコード未発見エラーを記録
func (r *RepositoryHelper) RecordNotFoundError(entityType string, searchCriteria any) {
	// 通常のログレベルで記録（エラーではなく正常な状況）
	logger.Info(r.ctx, "Record not found",
		"layer", "repository",
		"entity_type", entityType,
		"search_criteria", fmt.Sprintf("%v", searchCriteria),
	)

	// スパンには属性として記録（エラーステータスは設定しない）
	r.span.SetAttributes(
		attribute.String("db.result", "not_found"),
		attribute.String("db.search_entity", entityType),
		attribute.String("db.search_criteria", fmt.Sprintf("%v", searchCriteria)),
	)
}

// FinishWithRecordCount は処理したレコード数と共に完了
func (r *RepositoryHelper) FinishWithRecordCount(success bool, recordCount int, args ...any) {
	defer r.span.End()

	// 成功/失敗とレコード数をスパンに記録
	r.span.SetAttributes(
		attribute.Bool("app.success", success),
		attribute.Int("db.record_count", recordCount),
	)

	// 操作ログを完了
	countArgs := append([]any{"record_count", recordCount}, args...)
	r.operationLog(success, countArgs...)
}
