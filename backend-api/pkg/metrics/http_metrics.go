package metrics

import "time"

// HTTPMetricsRecorder はHTTPメトリクス記録のインターフェース
type HTTPMetricsRecorder interface {
	RecordRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64)
	RecordRequestWithContext(method, route, statusCode string, duration time.Duration, requestSize, responseSize int64)
}
