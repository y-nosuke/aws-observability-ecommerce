package metrics

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

// グローバルメトリクスインスタンス
var (
	globalHTTPMetricsRecorder HTTPMetricsRecorder
	initOnce                  sync.Once
)

// InitWithProvider はOpenTelemetryプロバイダーを使用してグローバルHTTPメトリクスを初期化します
func InitWithProvider(provider *sdkmetric.MeterProvider) error {
	var initError error
	initOnce.Do(func() {
		if provider == nil {
			globalHTTPMetricsRecorder = NewNoopHTTPMetricsRecorder()
			return
		}

		// グローバルMeterProviderを設定
		otel.SetMeterProvider(provider)

		// Meterを取得
		meter := provider.Meter("github.com/y-nosuke/aws-observability-ecommerce/backend-api")

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
