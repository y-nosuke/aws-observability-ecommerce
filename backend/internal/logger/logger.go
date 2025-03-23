package logger

import (
	"context"
	"io"
	"log/slog"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/aws"
)

// Logger はアプリケーション全体で使用するロガー構造体
type Logger struct {
	slogger           *slog.Logger
	cloudWatchHandler *CloudWatchHandler
	closers           []io.Closer // クローズが必要なリソースを保存
}

// グローバルロガーインスタンス
var defaultLogger *Logger

// InitConfig はロガー初期化用の設定情報を表す構造体
type InitConfig struct {
	AppName             string
	Environment         string
	LogLevel            string
	UseConsole          bool   // コンソール出力を使用するか
	UseFile             bool   // ファイル出力を使用するか
	FilePath            string // ファイル出力のパス
	UseCloudWatch       bool   // CloudWatch Logsを使用するか
	CreateLogGroup      bool   // ロググループが存在しない場合に作成するかどうか
	LogGroupName        string // CloudWatch Logsのグループ名
	CloudWatchFlushSecs int
	CloudWatchBatchSize int
	AWSOptions          aws.Options // AWS接続オプション
}

// New は新しいロガーインスタンスを作成します
func New(ctx context.Context, config InitConfig) (*Logger, error) {
	// ビルダーを使用してロガーを構築
	return NewLoggerBuilder().
		WithAppInfo(config.AppName, config.Environment).
		WithLogLevel(config.LogLevel).
		WithConsoleHandler(config.UseConsole).
		WithFileHandler(config.UseFile, config.FilePath).
		WithCloudWatchHandler(config.UseCloudWatch, config.CreateLogGroup, config.LogGroupName, config.AWSOptions).
		WithCloudWatchOptions(config.CloudWatchFlushSecs, config.CloudWatchBatchSize).
		Build(ctx)
}

// Logger はslog.Loggerインスタンスを返します
func (l *Logger) Logger() *slog.Logger {
	return l.slogger
}

// Close はロガーリソースを解放します
func (l *Logger) Close() error {
	// 保存しているすべてのクローザーをクローズ
	for _, closer := range l.closers {
		if err := closer.Close(); err != nil {
			return err
		}
	}

	// CloudWatch Logsハンドラーのクローズ（互換性のために残す）
	if l.cloudWatchHandler != nil {
		return l.cloudWatchHandler.Close()
	}

	return nil
}

// WithLogger は指定されたロガーをコンテキストに追加する
func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, loggerKey{}, logger)
}

// FromContext は現在のコンテキストに基づいてロガーを返す
func FromContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return DefaultLogger()
	}

	// コンテキストからロガーを取得
	if logger, ok := ctx.Value(loggerKey{}).(*Logger); ok {
		return logger.slogger
	}

	return DefaultLogger()
}

// DefaultLogger はデフォルトのslog.Loggerインスタンスを返す
func DefaultLogger() *slog.Logger {
	if defaultLogger == nil {
		// デフォルトロガーが初期化されていない場合は標準のロガーを返す
		return slog.Default()
	}
	return defaultLogger.slogger
}

// コンテキストキー
type loggerKey struct{}
