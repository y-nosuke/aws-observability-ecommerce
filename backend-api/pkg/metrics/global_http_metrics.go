package metrics

import (
	"sync"
	"time"

	"go.opentelemetry.io/otel/metric"
)

// グローバルメトリクスインスタンス
var (
	globalHTTPMetrics HTTPMetricsRecorder
	initOnce          sync.Once
	initError         error
)

// Init はグローバルHTTPメトリクスを初期化します
func Init(meter metric.Meter) error {
	initOnce.Do(func() {
		if meter == nil {
			globalHTTPMetrics = &NoopHTTPMetrics{}
			return
		}

		httpMetrics, err := NewHTTPMetrics(meter)
		if err != nil {
			initError = err
			globalHTTPMetrics = &NoopHTTPMetrics{}
			return
		}
		globalHTTPMetrics = httpMetrics
	})
	return initError
}

// SetHTTPMetrics はグローバルHTTPメトリクスを直接設定します（テスト用）
func SetHTTPMetrics(metrics HTTPMetricsRecorder) {
	globalHTTPMetrics = metrics
}

// getGlobalHTTPMetrics はグローバルHTTPメトリクスを取得します（内部用）
func getGlobalHTTPMetrics() HTTPMetricsRecorder {
	if globalHTTPMetrics == nil {
		return &NoopHTTPMetrics{}
	}
	return globalHTTPMetrics
}

// === パッケージレベルの便利関数 ===

// RecordHTTPRequest はHTTPリクエストのメトリクスを記録します
func RecordHTTPRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64) {
	getGlobalHTTPMetrics().RecordRequest(method, route, statusCode, duration, requestSize, responseSize)
}

// RecordHTTPRequestWithContext はコンテキスト付きでHTTPリクエストのメトリクスを記録します
func RecordHTTPRequestWithContext(method, route, statusCode string, duration time.Duration, requestSize, responseSize int64) {
	getGlobalHTTPMetrics().RecordRequestWithContext(method, route, statusCode, duration, requestSize, responseSize)
}
