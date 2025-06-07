package logging

import (
	"context"
	"crypto/sha256"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// ErrorContext はエラーログのコンテキスト情報
type ErrorContext struct {
	Operation      string
	ResourceType   string
	ResourceID     string
	RetryCount     int
	MaxRetries     int
	TimeoutMs      int
	Severity       string
	AffectedUsers  int
	BusinessImpact string
}

// LogError はエラーログを出力します
func (l *StructuredLogger) LogError(ctx context.Context, err error, errorCtx ErrorContext) {
	fields := []Field{
		{Key: "log_type", Value: "error"},
		{Key: "error", Value: map[string]interface{}{
			"type":        getErrorType(err),
			"message":     err.Error(),
			"code":        getErrorCode(err),
			"stack_trace": getStackTrace(err),
			"fingerprint": generateErrorFingerprint(err, errorCtx),
		}},
		{Key: "context", Value: map[string]interface{}{
			"operation":     errorCtx.Operation,
			"resource_type": errorCtx.ResourceType,
			"resource_id":   errorCtx.ResourceID,
			"retry_count":   errorCtx.RetryCount,
			"max_retries":   errorCtx.MaxRetries,
			"timeout_ms":    errorCtx.TimeoutMs,
		}},
		{Key: "impact", Value: map[string]interface{}{
			"severity":        errorCtx.Severity,
			"affected_users":  errorCtx.AffectedUsers,
			"business_impact": errorCtx.BusinessImpact,
		}},
	}

	l.Error(ctx, "Application error occurred", err, fields...)
}

// getErrorType はエラーのタイプを取得します
func getErrorType(err error) string {
	if err == nil {
		return "unknown"
	}

	// エラータイプを反映から取得
	errorType := reflect.TypeOf(err).String()

	// パッケージ名を除いてシンプルにする
	if idx := strings.LastIndex(errorType, "."); idx != -1 {
		return errorType[idx+1:]
	}

	return errorType
}

// getErrorCode はエラーコードを取得します
func getErrorCode(err error) string {
	// カスタムエラーインターフェースがあれば、それを実装
	type errorCoder interface {
		ErrorCode() string
	}

	if ec, ok := err.(errorCoder); ok {
		return ec.ErrorCode()
	}

	// デフォルトはエラータイプをコードとして使用
	return strings.ToUpper(strings.ReplaceAll(getErrorType(err), " ", "_"))
}

// getStackTrace はスタックトレースを取得します
func getStackTrace(_ error) string {
	// シンプルなスタックトレース生成
	// 本格的な実装では github.com/pkg/errors などを使用することを推奨

	var stack []string
	for i := 2; i < 10; i++ { // skip getStackTrace and LogError
		if pc, file, line, ok := runtime.Caller(i); ok {
			fn := runtime.FuncForPC(pc)
			if fn != nil {
				funcName := fn.Name()
				// パッケージ名を短縮
				if idx := strings.LastIndex(funcName, "/"); idx != -1 {
					funcName = funcName[idx+1:]
				}

				// ファイル名を短縮
				if idx := strings.LastIndex(file, "/"); idx != -1 {
					file = file[idx+1:]
				}

				stack = append(stack, fmt.Sprintf("%s:%d %s", file, line, funcName))
			}
		} else {
			break
		}
	}

	return strings.Join(stack, "\n")
}

// generateErrorFingerprint はエラーのフィンガープリントを生成します
func generateErrorFingerprint(err error, ctx ErrorContext) string {
	// エラーメッセージ、タイプ、操作を組み合わせてハッシュ化
	data := fmt.Sprintf("%s:%s:%s",
		getErrorType(err),
		err.Error(),
		ctx.Operation)

	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)[:16] // 最初の16文字を使用
}
