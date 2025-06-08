package logging

import (
	"context"
	"time"
)

// ApplicationOperation はアプリケーション操作のデータ構造
type ApplicationOperation struct {
	Name            string
	Category        string
	Duration        time.Duration
	Success         bool
	Stage           string
	EntityType      string
	EntityID        string
	Action          string
	Source          string
	Data            map[string]interface{}
	PerformanceData map[string]interface{}
}

// LogApplication はアプリケーションログを出力します
func (l *StructuredLogger) LogApplication(ctx context.Context, op ApplicationOperation) {
	fields := []Field{
		{Key: "log_type", Value: "application"},
		{Key: "operation", Value: map[string]interface{}{
			"name":        op.Name,
			"category":    op.Category,
			"duration_ms": float64(op.Duration.Nanoseconds()) / 1e6,
			"success":     op.Success,
			"stage":       op.Stage,
		}},
		{Key: "business", Value: map[string]interface{}{
			"entity_type": op.EntityType,
			"entity_id":   op.EntityID,
			"action":      op.Action,
			"source":      op.Source,
		}},
	}

	if len(op.Data) > 0 {
		fields = append(fields, Field{Key: "data", Value: op.Data})
	}

	if len(op.PerformanceData) > 0 {
		fields = append(fields, Field{Key: "performance", Value: op.PerformanceData})
	}

	message := "Application operation completed"
	if !op.Success {
		message = "Application operation failed"
	}

	l.Info(ctx, message, fields...)
}

// NewApplicationOperation はApplicationOperationの新しいインスタンスを作成します
func NewApplicationOperation(name, category string) *ApplicationOperation {
	return &ApplicationOperation{
		Name:            name,
		Category:        category,
		Success:         true,
		Data:            make(map[string]interface{}),
		PerformanceData: make(map[string]interface{}),
	}
}

// SetEntity はエンティティ情報を設定します
func (op *ApplicationOperation) SetEntity(entityType, entityID string) *ApplicationOperation {
	op.EntityType = entityType
	op.EntityID = entityID
	return op
}

// SetAction はアクション情報を設定します
func (op *ApplicationOperation) SetAction(action, source string) *ApplicationOperation {
	op.Action = action
	op.Source = source
	return op
}

// SetStage はステージ情報を設定します
func (op *ApplicationOperation) SetStage(stage string) *ApplicationOperation {
	op.Stage = stage
	return op
}

// SetSuccess は成功/失敗を設定します
func (op *ApplicationOperation) SetSuccess(success bool) *ApplicationOperation {
	op.Success = success
	return op
}

// SetDuration は実行時間を設定します
func (op *ApplicationOperation) SetDuration(duration time.Duration) *ApplicationOperation {
	op.Duration = duration
	return op
}

// AddData はデータフィールドを追加します
func (op *ApplicationOperation) AddData(key string, value interface{}) *ApplicationOperation {
	op.Data[key] = value
	return op
}

// AddPerformanceData はパフォーマンスデータを追加します
func (op *ApplicationOperation) AddPerformanceData(key string, value interface{}) *ApplicationOperation {
	op.PerformanceData[key] = value
	return op
}

// MeasureExecution は実行時間を計測してDurationに設定します
func (op *ApplicationOperation) MeasureExecution(fn func() error) error {
	start := time.Now()
	err := fn()
	op.Duration = time.Since(start)
	op.Success = (err == nil)
	return err
}
