package metrics

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// グローバルメトリクスインスタンス
var (
	globalHTTPMetricsRecorder HTTPMetricsRecorder
	initOnce                  sync.Once
)

// Init はグローバルHTTPメトリクスを初期化します
func Init(meter metric.Meter) error {
	var initError error
	initOnce.Do(func() {
		if meter == nil {
			globalHTTPMetricsRecorder = NewNoopHTTPMetricsRecorder()
			return
		}

		httpMetrics, err := NewDefaultHTTPMetricsRecorder(meter)
		if err != nil {
			initError = err
			globalHTTPMetricsRecorder = NewNoopHTTPMetricsRecorder()
			return
		}
		globalHTTPMetricsRecorder = httpMetrics
	})
	return initError
}

// SetGlobalHTTPMetricsRecorder はグローバルHTTPメトリクスを直接設定します（テスト用）
func SetGlobalHTTPMetricsRecorder(metrics HTTPMetricsRecorder) {
	globalHTTPMetricsRecorder = metrics
}

// getGlobalHTTPMetricsRecorder はグローバルHTTPメトリクスを取得します（内部用）
func getGlobalHTTPMetricsRecorder() HTTPMetricsRecorder {
	if globalHTTPMetricsRecorder == nil {
		return &NoopHTTPMetricsRecorder{}
	}
	return globalHTTPMetricsRecorder
}

// === パッケージレベルの便利関数 ===

// RecordHTTPRequest はHTTPリクエストのメトリクスを記録します
func RecordHTTPRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	getGlobalHTTPMetricsRecorder().RecordRequest(method, route, statusCode, duration, requestSize, responseSize)
}

// RecordHTTPRequestWithContext はコンテキスト付きでHTTPリクエストのメトリクスを記録します
func RecordHTTPRequestWithContext(ctx context.Context, method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	getGlobalHTTPMetricsRecorder().RecordRequestWithContext(ctx, method, route, statusCode, duration, requestSize, responseSize)
}
