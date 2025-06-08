package errors

import (
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
)

// AppError は構造化されたアプリケーションエラー
type AppError struct {
	Type       string                 `json:"type"`
	Message    string                 `json:"message"`
	Code       string                 `json:"code"`
	Context    map[string]interface{} `json:"context,omitempty"`
	Underlying error                  `json:"-"`
}

func (e *AppError) Error() string {
	if e.Underlying != nil {
		return fmt.Sprintf("%s: %s (underlying: %v)", e.Type, e.Message, e.Underlying)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// LogFields はログ出力用のフィールドを返す
func (e *AppError) LogFields() []logging.Field {
	fields := []logging.Field{
		{Key: "error_type", Value: e.Type},
		{Key: "error_code", Value: e.Code},
		{Key: "error_message", Value: e.Message},
	}
	if len(e.Context) > 0 {
		fields = append(fields, logging.Field{Key: "error_context", Value: e.Context})
	}
	return fields
}

// NewAppError は新しいAppErrorを作成
func NewAppError(errorType, message, code string) *AppError {
	return &AppError{
		Type:    errorType,
		Message: message,
		Code:    code,
		Context: make(map[string]interface{}),
	}
}

// WithContext はコンテキスト情報を追加
func (e *AppError) WithContext(key string, value interface{}) *AppError {
	if e.Context == nil {
		e.Context = make(map[string]interface{})
	}
	e.Context[key] = value
	return e
}

// WithUnderlying は根本原因のエラーを設定
func (e *AppError) WithUnderlying(err error) *AppError {
	e.Underlying = err
	return e
}
