package metrics

import (
	"context"
	"time"
)

// HTTPMetricsRecorder はHTTPメトリクス記録のインターフェース
type HTTPMetricsRecorder interface {
	RecordRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64)
	RecordRequestWithContext(ctx context.Context, method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64)
}
