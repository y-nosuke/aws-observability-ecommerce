package errors

import (
	"fmt"
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

// LogArgs はパッケージレベルログ用のkey-value引数を返す
// 新しいログシステム用（logging.Info(ctx, msg, args...)で使用）
func (e *AppError) LogArgs() []any {
	args := []any{
		"error_type", e.Type,
		"error_code", e.Code,
		"error_message", e.Message,
	}

	if len(e.Context) > 0 {
		args = append(args, "error_context", e.Context)
	}

	if e.Underlying != nil {
		args = append(args, "underlying_error", e.Underlying.Error())
	}

	return args
}

// LogArgsMap はマップ形式でエラー情報を返す（デバッグやテスト用）
func (e *AppError) LogArgsMap() map[string]interface{} {
	result := map[string]interface{}{
		"error_type":    e.Type,
		"error_code":    e.Code,
		"error_message": e.Message,
	}

	if len(e.Context) > 0 {
		result["error_context"] = e.Context
	}

	if e.Underlying != nil {
		result["underlying_error"] = e.Underlying.Error()
	}

	return result
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
