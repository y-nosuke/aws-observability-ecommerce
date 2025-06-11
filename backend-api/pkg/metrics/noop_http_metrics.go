package metrics

import "time"

// NoopHTTPMetrics はメトリクス無効時に使用するダミー実装
type NoopHTTPMetrics struct{}

// RecordRequest はダミー実装（何もしない）
func (n *NoopHTTPMetrics) RecordRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	// 何もしない
}

// RecordRequestWithContext はダミー実装（何もしない）
func (n *NoopHTTPMetrics) RecordRequestWithContext(method, route, statusCode string, duration time.Duration, requestSize, responseSize int64) {
	// 何もしない
}
