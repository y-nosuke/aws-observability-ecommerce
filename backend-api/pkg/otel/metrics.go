package otel

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

var (
	meter             metric.Meter
	requestsTotal     metric.Int64Counter
	errorsTotal       metric.Int64Counter
	requestDuration   metric.Float64Histogram
	requestSizeBytes  metric.Int64Histogram
	responseSizeBytes metric.Int64Histogram
)

// NewHTTPMetricsRecorder はHTTPMetricsの新しいインスタンスを作成します
func NewHTTPMetricsRecorder(provider *sdkmetric.MeterProvider, cfg config.MetricsConfig) (metric.Meter, error) {
	meter = provider.Meter(utils.GetModulePath())

	var err error

	// HTTP リクエスト総数カウンター
	requestsTotal, err = meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_requests_total counter: %w", err)
	}

	// HTTP エラー総数カウンター
	errorsTotal, err = meter.Int64Counter(
		"http_request_errors_total",
		metric.WithDescription("Total number of HTTP error responses (4xx, 5xx)"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_request_errors_total counter: %w", err)
	}

	// HTTP リクエスト処理時間ヒストグラム
	requestDuration, err = meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("HTTP request duration in seconds"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(cfg.RequestTimeHistogramBoundaries...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_request_duration_seconds histogram: %w", err)
	}

	// HTTP リクエストサイズヒストグラム
	requestSizeBytes, err = meter.Int64Histogram(
		"http_request_size_bytes",
		metric.WithDescription("HTTP request body size in bytes"),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(cfg.RequestSizeHistogramBoundaries...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_request_size_bytes histogram: %w", err)
	}

	// HTTP レスポンスサイズヒストグラム
	responseSizeBytes, err = meter.Int64Histogram(
		"http_response_size_bytes",
		metric.WithDescription("HTTP response body size in bytes"),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(cfg.RequestSizeHistogramBoundaries...),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_response_size_bytes histogram: %w", err)
	}

	return nil, nil
}

// CountRequestsTotal はHTTPリクエスト総数をカウントします
func CountRequestsTotal(ctx context.Context, attributes ...attribute.KeyValue) {
	attrs := append(getCommonAttrs(ctx), attributes...)
	requestsTotal.Add(ctx, 1, metric.WithAttributes(attrs...))
}

// RecordRequestDuration はHTTPリクエスト処理時間を記録します
func RecordRequestDuration(ctx context.Context, duration time.Duration, attributes ...attribute.KeyValue) {
	attrs := append(getCommonAttrs(ctx), attributes...)
	requestDuration.Record(ctx, duration.Seconds(), metric.WithAttributes(attrs...))
}

// RecordRequestSizeBytes はHTTPリクエストサイズを記録します
func RecordRequestSizeBytes(ctx context.Context, requestSize int64, attributes ...attribute.KeyValue) {
	attrs := append(getCommonAttrs(ctx), attributes...)
	requestSizeBytes.Record(ctx, requestSize, metric.WithAttributes(attrs...))
}

// RecordResponseSizeBytes はHTTPレスポンスサイズを記録します
func RecordResponseSizeBytes(ctx context.Context, responseSize int64, attributes ...attribute.KeyValue) {
	attrs := append(getCommonAttrs(ctx), attributes...)
	responseSizeBytes.Record(ctx, responseSize, metric.WithAttributes(attrs...))
}

// CountErrorsTotal はHTTPエラー総数をカウントします
func CountErrorsTotal(ctx context.Context, statusCode int, attributes ...attribute.KeyValue) {
	if statusCode >= 400 {
		errorType := "client_error"
		if statusCode >= 500 {
			errorType = "server_error"
		}
		attrs := append(getCommonAttrs(ctx), attribute.String("error.type", errorType))
		attrs = append(attrs, attributes...)
		errorsTotal.Add(ctx, 1, metric.WithAttributes(attrs...))
	}
}

// getCommonAttrs は共通の属性を取得します
func getCommonAttrs(ctx context.Context) []attribute.KeyValue {
	attrs, ok := ctx.Value(ContextKeyAttrs).(map[string]string)
	if !ok {
		attrs = map[string]string{}
	}

	return []attribute.KeyValue{
		attribute.String("http.method", attrs["http.method"]),
		attribute.String("http.route", attrs["http.route"]),
		attribute.String("http.host", attrs["http.host"]),
	}
}
