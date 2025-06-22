package observability

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// HandlerHelper は Handler 層でのトレース処理を簡素化するヘルパー
type HandlerHelper struct {
	BaseObservabilityHelper // 共通機能を埋め込み
}

// StartHandler は Handler のトレースを開始
func StartHandler(ctx context.Context, operationName string, method, path string, statusCode int, userAgent, remoteAddr string, contentLength int64) *HandlerHelper {
	helper := &HandlerHelper{}

	// HTTP固有の追加属性を準備
	httpAttrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", path),
		attribute.Int("http.status_code", statusCode),
		attribute.String("http.user_agent", userAgent),
		attribute.String("http.remote_addr", remoteAddr),
	}

	if contentLength > 0 {
		httpAttrs = append(httpAttrs, attribute.Int64("http.request.content_length", contentLength))
	}

	// span作成から初期化まで一貫実行
	helper.initializeWithSpan(ctx, operationName, "handler", httpAttrs...)

	return helper
}

// RecordValidationError はバリデーションエラーを記録
func (h *HandlerHelper) RecordValidationError(err error, field string, value any) {
	// エラーログを記録
	logger.WithError(h.ctx, "Validation error", err,
		"layer", "handler",
		"validation_error", "request_validation",
		"field", field,
		"value", value,
	)

	// スパンにバリデーションエラー情報を記録
	h.span.RecordError(err)
	h.span.SetStatus(codes.Error, err.Error())
	h.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", "validation_error"),
		attribute.String("validation.field", field),
		attribute.String("validation.value", fmt.Sprintf("%v", value)),
	)
}

// FinishWithHTTPStatus はHTTPステータスコードで処理を完了
func (h *HandlerHelper) FinishWithHTTPStatus(statusCode int, args ...any) {
	defer h.span.End()

	success := statusCode >= 200 && statusCode < 400
	h.span.SetAttributes(
		attribute.Bool("app.success", success),
		attribute.Int("http.status_code", statusCode),
	)

	// ステータスコードに応じたスパンステータス設定
	if statusCode >= 400 {
		if statusCode >= 500 {
			h.span.SetStatus(codes.Error, http.StatusText(statusCode))
		} else {
			// 4xx は通常エラーとして扱わない（クライアントエラー）
			h.span.SetStatus(codes.Unset, "")
		}
	}

	// 操作ログを完了
	statusArgs := append([]any{"http_status", statusCode}, args...)
	h.operationLog(success, statusArgs...)
}

// FinishWithError はエラーで Handler の処理を完了（オーバーライド）
func (h *HandlerHelper) FinishWithError(err error, message string, statusCode int, args ...any) {
	defer h.span.End()

	// エラーログとスパン情報を記録
	h.LogError(message, err, args...)

	// HTTPステータス情報も記録
	h.span.SetAttributes(attribute.Int("http.status_code", statusCode))

	// 操作ログを失敗で完了
	errorArgs := append([]any{"error_type", "handler_failure", "http_status", statusCode}, args...)
	h.operationLog(false, errorArgs...)
}
