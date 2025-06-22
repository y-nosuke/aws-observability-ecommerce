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

// RepositoryHelper は Repository 層でのトレース処理を簡素化するヘルパー
type RepositoryHelper struct {
	ctx          context.Context
	span         trace.Span
	operationLog func(success bool, args ...any)
}

// StartRepository は Repository のトレースを開始
func StartRepository(ctx context.Context, operationName string) *RepositoryHelper {
	// contextからdomainを自動取得
	domain := GetDomainFromContext(ctx)

	// 既存のStartRepository関数を使用
	spanCtx, span := tracer.StartRepository(ctx, operationName, domain)

	// contextからentityIDを自動取得
	if id := GetEntityIDFromContext(ctx); id > 0 {
		span.SetAttributes(attribute.Int("app.entity_id", id))
	}

	// 操作ログを開始
	operationLog := logger.StartOperation(spanCtx, operationName,
		"layer", "repository",
		"domain", domain,
	)

	return &RepositoryHelper{
		ctx:          spanCtx,
		span:         span,
		operationLog: operationLog,
	}
}

// Context は現在のコンテキストを返す
func (r *RepositoryHelper) Context() context.Context {
	return r.ctx
}

// SetAttributes はスパンに属性を設定
func (r *RepositoryHelper) SetAttributes(attrs ...attribute.KeyValue) {
	r.span.SetAttributes(attrs...)
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
	stepCtx, stepSpan := tracer.StartDatabase(r.ctx, stepName, tableName)
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

// LogInfo は情報ログを記録
func (r *RepositoryHelper) LogInfo(message string, args ...any) {
	logger.Info(r.ctx, message, args...)
}

// LogError はエラーログを記録し、スパンにエラー情報を設定
func (r *RepositoryHelper) LogError(message string, err error, args ...any) {
	// エラーログを記録
	allArgs := append([]any{"layer", "repository"}, args...)
	logger.WithError(r.ctx, message, err, allArgs...)

	// スパンにエラー情報を記録
	r.span.RecordError(err)
	r.span.SetStatus(codes.Error, err.Error())
	r.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", "repository_error"),
	)
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

// Finish は Repository の処理を完了
func (r *RepositoryHelper) Finish(success bool, args ...any) {
	defer r.span.End()

	// 成功/失敗をスパンに記録
	r.span.SetAttributes(attribute.Bool("app.success", success))

	// 操作ログを完了
	r.operationLog(success, args...)
}

// FinishWithError はエラーで Repository の処理を完了
func (r *RepositoryHelper) FinishWithError(err error, message string, args ...any) {
	defer r.span.End()

	// エラーログとスパン情報を記録
	r.LogError(message, err, args...)

	// 操作ログを失敗で完了
	errorArgs := append([]any{"error_type", "repository_failure"}, args...)
	r.operationLog(false, errorArgs...)
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
