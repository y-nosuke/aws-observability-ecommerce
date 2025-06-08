package logging

import (
	"context"
	"time"
)

// LogHelper は便利なログ機能を提供する
type LogHelper struct {
	logger Logger
}

// NewLogHelper は新しいLogHelperを作成する
func NewLogHelper(logger Logger) *LogHelper {
	return &LogHelper{logger: logger}
}

// StartOperation は操作の開始をログに記録し、完了時の記録用関数を返す
func (h *LogHelper) StartOperation(ctx context.Context, name, category string) *OperationLogger {
	start := time.Now()

	op := NewApplicationOperation(name, category).
		SetStage("started")

	h.logger.LogApplication(ctx, *op)

	return &OperationLogger{
		helper:    h,
		operation: op,
		startTime: start,
	}
}

// OperationLogger は操作の完了をログに記録するためのヘルパー
type OperationLogger struct {
	helper    *LogHelper
	operation *ApplicationOperation
	startTime time.Time
}

// Complete は操作の完了をログに記録する
func (ol *OperationLogger) Complete(ctx context.Context) {
	ol.operation.SetDuration(time.Since(ol.startTime)).
		SetStage("completed").
		SetSuccess(true)

	ol.helper.logger.LogApplication(ctx, *ol.operation)
}

// Fail は操作の失敗をログに記録する
func (ol *OperationLogger) Fail(ctx context.Context, err error) {
	ol.operation.SetDuration(time.Since(ol.startTime)).
		SetStage("failed").
		SetSuccess(false).
		AddData("error", err.Error())

	ol.helper.logger.LogApplication(ctx, *ol.operation)
}

// WithEntity はエンティティ情報を設定する
func (ol *OperationLogger) WithEntity(entityType, entityID string) *OperationLogger {
	ol.operation.SetEntity(entityType, entityID)
	return ol
}

// WithAction はアクション情報を設定する
func (ol *OperationLogger) WithAction(action, source string) *OperationLogger {
	ol.operation.SetAction(action, source)
	return ol
}

// WithData はデータフィールドを追加する
func (ol *OperationLogger) WithData(key string, value interface{}) *OperationLogger {
	ol.operation.AddData(key, value)
	return ol
}

// WithPerformanceData はパフォーマンスデータを追加する
func (ol *OperationLogger) WithPerformanceData(key string, value interface{}) *OperationLogger {
	ol.operation.AddPerformanceData(key, value)
	return ol
}

// LogBusinessEvent はビジネスイベントを記録する便利メソッド
func (h *LogHelper) LogBusinessEvent(ctx context.Context, event, entityType, entityID string, data map[string]interface{}) {
	op := NewApplicationOperation(event, "business_event").
		SetEntity(entityType, entityID).
		SetAction("event", "system").
		SetStage("completed").
		SetSuccess(true)

	for k, v := range data {
		op.AddData(k, v)
	}

	h.logger.LogApplication(ctx, *op)
}

// LogSystemEvent はシステムイベントを記録する便利メソッド
func (h *LogHelper) LogSystemEvent(ctx context.Context, event string, data map[string]interface{}) {
	op := NewApplicationOperation(event, "system_event").
		SetAction("event", "system").
		SetStage("completed").
		SetSuccess(true)

	for k, v := range data {
		op.AddData(k, v)
	}

	h.logger.LogApplication(ctx, *op)
}

// LogPerformanceMetric はパフォーマンスメトリクスを記録する便利メソッド
func (h *LogHelper) LogPerformanceMetric(ctx context.Context, metric string, value interface{}, unit string) {
	op := NewApplicationOperation("performance_metric", "monitoring").
		SetAction("measure", "system").
		SetStage("completed").
		SetSuccess(true).
		AddData("metric_name", metric).
		AddData("metric_value", value).
		AddData("metric_unit", unit)

	h.logger.LogApplication(ctx, *op)
}
