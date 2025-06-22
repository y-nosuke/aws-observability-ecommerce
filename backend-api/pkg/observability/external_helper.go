package observability

import (
	"context"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// ExternalHelper は External層（外部API/サービス）でのトレース処理を簡素化するヘルパー
type ExternalHelper struct {
	BaseObservabilityHelper // 共通機能を埋め込み
}

// StartExternal は External のトレースを開始
func StartExternal(ctx context.Context, operationName string) *ExternalHelper {
	helper := &ExternalHelper{}

	// span作成から初期化まで一貫実行
	helper.initializeWithSpan(ctx, operationName, "external")

	return helper
}

// RecordS3Operation はS3操作情報を記録
func (e *ExternalHelper) RecordS3Operation(bucket, key, operation string, objectSize int64) {
	e.span.SetAttributes(
		attribute.String("aws.s3.bucket", bucket),
		attribute.String("aws.s3.key", key),
		attribute.String("aws.s3.operation", operation),
		attribute.Int64("aws.s3.object.size", objectSize),
	)
}

// RecordAWSServiceInfo はAWSサービス情報を記録
func (e *ExternalHelper) RecordAWSServiceInfo(service, region, operation string) {
	e.span.SetAttributes(
		attribute.String("aws.service", service),
		attribute.String("aws.region", region),
		attribute.String("aws.operation", operation),
	)
}

// AddExternalStep は外部サービス呼び出しステップを記録
func (e *ExternalHelper) AddExternalStep(stepName, serviceType string, fn func(context.Context) error) error {
	// External固有の追加属性を準備
	externalAttrs := []attribute.KeyValue{
		attribute.String("external.service.type", serviceType),
		attribute.String("external.step.name", stepName),
	}

	// サブspanを作成
	stepCtx, stepSpan := e.startSubSpan(stepName, "external", externalAttrs...)
	defer stepSpan.End()

	// 開始時刻を記録
	startTime := time.Now()
	stepSpan.SetAttributes(attribute.String("external.step.start_time", startTime.Format(time.RFC3339)))

	err := fn(stepCtx)

	// 完了時刻と所要時間を記録
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	stepSpan.SetAttributes(
		attribute.String("external.step.end_time", endTime.Format(time.RFC3339)),
		attribute.Int64("external.step.duration_ms", duration.Milliseconds()),
	)

	if err != nil {
		stepSpan.RecordError(err)
		stepSpan.SetStatus(codes.Error, err.Error())
		stepSpan.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.step", stepName),
			attribute.String("external.service.type", serviceType),
		)
	}

	return err
}

// FinishWithBytesTransferred はデータ転送量と共に完了
func (e *ExternalHelper) FinishWithBytesTransferred(success bool, bytesTransferred int64, operation string, args ...any) {
	defer e.span.End()

	// 成功/失敗とデータ転送量をスパンに記録
	e.span.SetAttributes(
		attribute.Bool("app.success", success),
		attribute.Int64("external.bytes_transferred", bytesTransferred),
		attribute.String("external.operation", operation),
	)

	// 操作ログを完了
	transferArgs := append([]any{"bytes_transferred", bytesTransferred, "operation", operation}, args...)
	e.operationLog(success, transferArgs...)
}
