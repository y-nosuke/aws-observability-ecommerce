package metrics

import (
	"context"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

// DefaultHTTPMetricsRecorder はHTTPリクエストのメトリクスを管理する構造体
type DefaultHTTPMetricsRecorder struct {
	requestsTotal     metric.Int64Counter
	errorsTotal       metric.Int64Counter
	requestDuration   metric.Float64Histogram
	requestSizeBytes  metric.Int64Histogram
	responseSizeBytes metric.Int64Histogram
}

// NewDefaultHTTPMetricsRecorder はHTTPMetricsの新しいインスタンスを作成します
func NewDefaultHTTPMetricsRecorder(meter metric.Meter) (*DefaultHTTPMetricsRecorder, error) {
	// HTTP リクエスト総数カウンター
	requestsTotal, err := meter.Int64Counter(
		"http_requests_total",
		metric.WithDescription("Total number of HTTP requests"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_requests_total counter: %w", err)
	}

	// HTTP エラー総数カウンター
	errorsTotal, err := meter.Int64Counter(
		"http_request_errors_total",
		metric.WithDescription("Total number of HTTP error responses (4xx, 5xx)"),
		metric.WithUnit("1"),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_request_errors_total counter: %w", err)
	}

	// HTTP リクエスト処理時間ヒストグラム
	requestDuration, err := meter.Float64Histogram(
		"http_request_duration_seconds",
		metric.WithDescription("HTTP request duration in seconds"),
		metric.WithUnit("s"),
		metric.WithExplicitBucketBoundaries(0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_request_duration_seconds histogram: %w", err)
	}

	// HTTP リクエストサイズヒストグラム
	requestSizeBytes, err := meter.Int64Histogram(
		"http_request_size_bytes",
		metric.WithDescription("HTTP request body size in bytes"),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(64, 256, 1024, 4096, 16384, 65536, 262144, 1048576),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_request_size_bytes histogram: %w", err)
	}

	// HTTP レスポンスサイズヒストグラム
	responseSizeBytes, err := meter.Int64Histogram(
		"http_response_size_bytes",
		metric.WithDescription("HTTP response body size in bytes"),
		metric.WithUnit("By"),
		metric.WithExplicitBucketBoundaries(64, 256, 1024, 4096, 16384, 65536, 262144, 1048576),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create http_response_size_bytes histogram: %w", err)
	}

	return &DefaultHTTPMetricsRecorder{
		requestsTotal:     requestsTotal,
		errorsTotal:       errorsTotal,
		requestDuration:   requestDuration,
		requestSizeBytes:  requestSizeBytes,
		responseSizeBytes: responseSizeBytes,
	}, nil
}

// RecordRequest はHTTPリクエストのメトリクスを記録します
func (m *DefaultHTTPMetricsRecorder) RecordRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	// 共通のラベル属性
	commonAttrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", route),
		attribute.Int("http.status_code", statusCode),
	}

	// リクエスト総数を記録
	m.requestsTotal.Add(context.Background(), 1, metric.WithAttributes(commonAttrs...))

	// エラーの場合はエラーカウンターも増加
	if statusCode >= 400 {
		errorType := "client_error"
		if statusCode >= 500 {
			errorType = "server_error"
		}

		errorAttrs := append(commonAttrs, attribute.String("error.type", errorType))
		m.errorsTotal.Add(
			context.Background(),
			1,
			metric.WithAttributes(errorAttrs...),
		)
	}

	// 処理時間を記録（レスポンス時間はステータスコードを含まない）
	durationAttrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", route),
	}
	m.requestDuration.Record(
		context.Background(),
		duration.Seconds(),
		metric.WithAttributes(durationAttrs...),
	)

	// リクエストサイズを記録
	if requestSize > 0 {
		m.requestSizeBytes.Record(
			context.Background(),
			requestSize,
			metric.WithAttributes(durationAttrs...),
		)
	}

	// レスポンスサイズを記録
	if responseSize > 0 {
		m.responseSizeBytes.Record(
			context.Background(),
			responseSize,
			metric.WithAttributes(durationAttrs...),
		)
	}
}

// RecordRequestWithContext はコンテキスト付きでHTTPリクエストのメトリクスを記録します
func (m *DefaultHTTPMetricsRecorder) RecordRequestWithContext(ctx context.Context, method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	// 共通のラベル属性
	commonAttrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", route),
		attribute.Int("http.status_code", statusCode),
	}

	// リクエスト総数を記録
	m.requestsTotal.Add(ctx, 1, metric.WithAttributes(commonAttrs...))

	// エラーの場合はエラーカウンターも増加
	if statusCode >= 400 {
		errorType := "client_error"
		if statusCode >= 500 {
			errorType = "server_error"
		}

		errorAttrs := append(commonAttrs, attribute.String("error.type", errorType))
		m.errorsTotal.Add(
			ctx,
			1,
			metric.WithAttributes(errorAttrs...),
		)
	}

	// 処理時間を記録（レスポンス時間はステータスコードを含まない）
	durationAttrs := []attribute.KeyValue{
		attribute.String("http.method", method),
		attribute.String("http.route", route),
	}
	m.requestDuration.Record(
		ctx,
		duration.Seconds(),
		metric.WithAttributes(durationAttrs...),
	)

	// リクエストサイズを記録
	if requestSize > 0 {
		m.requestSizeBytes.Record(
			ctx,
			requestSize,
			metric.WithAttributes(durationAttrs...),
		)
	}

	// レスポンスサイズを記録
	if responseSize > 0 {
		m.responseSizeBytes.Record(
			ctx,
			responseSize,
			metric.WithAttributes(durationAttrs...),
		)
	}
}
