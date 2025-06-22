package logger

import (
	"context"
	"sync"
	"time"

	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// グローバルロガーインスタンス
var (
	globalLogger Logger
	initOnce     sync.Once
)

// InitWithProvider はOpenTelemetryプロバイダーを使用してグローバルロガーを初期化します
func InitWithProvider(provider *sdklog.LoggerProvider, cfg config.ObservabilityConfig) error {
	var initError error
	initOnce.Do(func() {
		if provider == nil {
			// プロバイダーがnilの場合は通常のロガーを使用
			cfg.Logging.EnableOTel = false
			globalLogger = NewDefaultLogger(cfg)
			return
		}

		// OpenTelemetryプロバイダーをグローバルに設定
		global.SetLoggerProvider(provider)

		// OpenTelemetry対応のロガーを作成
		cfg.Logging.EnableOTel = true
		globalLogger = NewDefaultLogger(cfg)
	})
	return initError
}

// SetGlobalLogger はグローバルロガーを直接設定します（テスト用）
func SetGlobalLogger(logger Logger) {
	globalLogger = logger
}

// getGlobalLogger はグローバルロガーを取得します（内部用）
func getGlobalLogger() Logger {
	if globalLogger == nil {
		// フォールバック：デフォルト設定でロガーを作成
		cfg := config.ObservabilityConfig{
			Logging: config.LoggingConfig{
				Level:      "info",
				Format:     "json",
				EnableOTel: false,
			},
		}
		globalLogger = NewDefaultLogger(cfg)
	}
	return globalLogger
}

// === パッケージレベルの便利関数 ===

// Info はInfoレベルのログを出力します
func Info(ctx context.Context, msg string, args ...any) {
	getGlobalLogger().Info(ctx, msg, args...)
}

// Warn はWarnレベルのログを出力します
func Warn(ctx context.Context, msg string, args ...any) {
	getGlobalLogger().Warn(ctx, msg, args...)
}

// Error はErrorレベルのログを出力します
func Error(ctx context.Context, msg string, args ...any) {
	getGlobalLogger().Error(ctx, msg, args...)
}

// Debug はDebugレベルのログを出力します
func Debug(ctx context.Context, msg string, args ...any) {
	getGlobalLogger().Debug(ctx, msg, args...)
}

// InfoF はフォーマット付きでInfoログを出力します
func InfoF(ctx context.Context, format string, args ...any) {
	getGlobalLogger().InfoF(ctx, format, args...)
}

// ErrorF はフォーマット付きでErrorログを出力します
func ErrorF(ctx context.Context, format string, args ...any) {
	getGlobalLogger().ErrorF(ctx, format, args...)
}

// WithError はエラー情報を含むログを出力します
func WithError(ctx context.Context, msg string, err error, args ...any) {
	getGlobalLogger().WithError(ctx, msg, err, args...)
}

// LogOperation は操作の開始/完了を記録します
func LogOperation(ctx context.Context, operation string, duration time.Duration, success bool, args ...any) {
	getGlobalLogger().LogOperation(ctx, operation, duration, success, args...)
}

// LogHTTPRequest はHTTPリクエストをログ出力します
func LogHTTPRequest(ctx context.Context, method, path string, status int, duration time.Duration, args ...any) {
	getGlobalLogger().LogHTTPRequest(ctx, method, path, status, duration, args...)
}

// LogBusinessEvent はビジネスイベントをログ出力します
func LogBusinessEvent(ctx context.Context, event string, entityType string, entityID string, args ...any) {
	getGlobalLogger().LogBusinessEvent(ctx, event, entityType, entityID, args...)
}

// === より簡単な使用のためのヘルパー ===

// StartOperation は操作を開始し、完了時に呼び出す関数を返します
func StartOperation(ctx context.Context, operation string, args ...any) func(success bool, additionalArgs ...any) {
	start := time.Now()

	// 開始ログ
	startArgs := append(args, "stage", "started")
	Info(ctx, "Operation started: "+operation, startArgs...)

	return func(success bool, additionalArgs ...any) {
		duration := time.Since(start)
		allArgs := append(args, additionalArgs...)
		LogOperation(ctx, operation, duration, success, allArgs...)
	}
}
