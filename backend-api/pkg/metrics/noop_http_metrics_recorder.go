package metrics

import (
	"context"
	"time"
)

// NoopHTTPMetricsRecorder はメトリクス無効時に使用するダミー実装
type NoopHTTPMetricsRecorder struct{}

func NewNoopHTTPMetricsRecorder() *NoopHTTPMetricsRecorder {
	return &NoopHTTPMetricsRecorder{}
}

// RecordRequest はダミー実装（何もしない）
func (n *NoopHTTPMetricsRecorder) RecordRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	// 何もしない
}

// RecordRequestWithContext はダミー実装（何もしない）
func (n *NoopHTTPMetricsRecorder) RecordRequestWithContext(ctx context.Context, method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	// 何もしない
}
