package observability

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
)

// HandlerHelper は Handler 層でのトレース処理を簡素化するヘルパー
type HandlerHelper struct {
	ctx          context.Context
	span         trace.Span
	operationLog func(success bool, args ...any)
}

// StartHandler は Handler のトレースを開始
func StartHandler(ctx context.Context, operationName string, method, path string, statusCode int, userAgent, remoteAddr string, contentLength int64) *HandlerHelper {
	// contextからdomainを自動取得
	domain := GetDomainFromContext(ctx)

	// 既存のStartHandler関数を使用
	spanCtx, span := tracer.StartHandler(ctx, operationName, domain)

	// contextからentityIDを自動取得
	if id := GetEntityIDFromContext(ctx); id > 0 {
		span.SetAttributes(attribute.Int("app.entity_id", id))
	}

	// HTTPリクエスト情報を記録
	span.SetAttributes(
		attribute.String("http.method", method),
		attribute.String("http.route", path),
		attribute.Int("http.status_code", statusCode),
	)

	// リクエスト情報を記録
	attrs := []attribute.KeyValue{
		attribute.String("http.user_agent", userAgent),
		attribute.String("http.remote_addr", remoteAddr),
	}
	if contentLength > 0 {
		attrs = append(attrs, attribute.Int64("http.request.content_length", contentLength))
	}
	span.SetAttributes(attrs...)

	// 操作ログを開始
	operationLog := logger.StartOperation(spanCtx, operationName,
		"layer", "handler",
		"domain", domain,
	)

	return &HandlerHelper{
		ctx:          spanCtx,
		span:         span,
		operationLog: operationLog,
	}
}

// Context は現在のコンテキストを返す
func (h *HandlerHelper) Context() context.Context {
	return h.ctx
}

// SetAttributes はスパンに属性を設定
func (h *HandlerHelper) SetAttributes(attrs ...attribute.KeyValue) {
	h.span.SetAttributes(attrs...)
}

// LogInfo は情報ログを記録
func (h *HandlerHelper) LogInfo(message string, args ...any) {
	logger.Info(h.ctx, message, args...)
}

// LogError はエラーログを記録し、スパンにエラー情報を設定
func (h *HandlerHelper) LogError(message string, err error, args ...any) {
	// エラーログを記録
	allArgs := append([]any{"layer", "handler"}, args...)
	logger.WithError(h.ctx, message, err, allArgs...)

	// スパンにエラー情報を記録
	h.span.RecordError(err)
	h.span.SetStatus(codes.Error, err.Error())
	h.span.SetAttributes(
		attribute.Bool("error", true),
		attribute.String("error.type", "handler_error"),
	)
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

// Finish は Handler の処理を完了
func (h *HandlerHelper) Finish(success bool, args ...any) {
	defer h.span.End()

	// 成功/失敗をスパンに記録
	h.span.SetAttributes(attribute.Bool("app.success", success))

	// 操作ログを完了
	h.operationLog(success, args...)
}

// FinishWithError はエラーで Handler の処理を完了
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
