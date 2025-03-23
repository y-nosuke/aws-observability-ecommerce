package logger

import (
	"context"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/y-nosuke/aws-observability-ecommerce/internal/aws"
)

// LoggerBuilder はログ設定をビルドするためのインターフェース
type LoggerBuilder interface {
	// 基本設定

	WithAppInfo(appName, environment string) LoggerBuilder
	WithLogLevel(level string) LoggerBuilder

	// ハンドラー設定

	WithConsoleHandler(enabled bool) LoggerBuilder
	WithFileHandler(enabled bool, filePath string) LoggerBuilder
	WithCloudWatchHandler(enabled bool, createLogGroup bool, logGroupName string, awsOptions aws.Options) LoggerBuilder
	WithCloudWatchOptions(flushIntervalSecs int, batchSize int) LoggerBuilder

	// ビルド処理

	Build(ctx context.Context) (*Logger, error)
}

// defaultLoggerBuilder はLoggerBuilderの標準実装
type defaultLoggerBuilder struct {
	appName     string
	environment string
	logLevel    slog.Level

	// ハンドラー設定
	useConsole              bool
	useFile                 bool
	filePath                string
	useCloudWatch           bool
	createLogGroup          bool
	logGroupName            string
	awsOptions              aws.Options
	cloudWatchFlushInterval time.Duration
	cloudWatchBatchSize     int
}

// NewLoggerBuilder は新しいLoggerBuilderを作成します
func NewLoggerBuilder() LoggerBuilder {
	return &defaultLoggerBuilder{
		logLevel:                slog.LevelInfo,  // デフォルト値
		useConsole:              true,            // デフォルトでコンソール出力有効
		createLogGroup:          true,            // デフォルト: は自動作成を有効化
		cloudWatchFlushInterval: 5 * time.Second, // デフォルト値: 5秒
		cloudWatchBatchSize:     100,             // デフォルト値: 100件
	}
}

// WithAppInfo はアプリケーション情報を設定します
func (b *defaultLoggerBuilder) WithAppInfo(appName, environment string) LoggerBuilder {
	b.appName = appName
	b.environment = environment
	return b
}

// WithLogLevel はログレベルを設定します
func (b *defaultLoggerBuilder) WithLogLevel(level string) LoggerBuilder {
	switch level {
	case "debug":
		b.logLevel = slog.LevelDebug
	case "info":
		b.logLevel = slog.LevelInfo
	case "warn":
		b.logLevel = slog.LevelWarn
	case "error":
		b.logLevel = slog.LevelError
	default:
		b.logLevel = slog.LevelInfo
	}
	return b
}

// WithConsoleHandler はコンソール出力の設定を行います
func (b *defaultLoggerBuilder) WithConsoleHandler(enabled bool) LoggerBuilder {
	b.useConsole = enabled
	return b
}

// WithFileHandler はファイル出力の設定を行います
func (b *defaultLoggerBuilder) WithFileHandler(enabled bool, filePath string) LoggerBuilder {
	b.useFile = enabled
	b.filePath = filePath
	return b
}

// WithCloudWatchHandler はCloudWatch Logs出力の設定を行います
func (b *defaultLoggerBuilder) WithCloudWatchHandler(enabled bool, createLogGroup bool, logGroupName string, awsOptions aws.Options) LoggerBuilder {
	b.useCloudWatch = enabled
	b.createLogGroup = createLogGroup
	b.logGroupName = logGroupName
	b.awsOptions = awsOptions
	return b
}

// WithCloudWatchOptions はCloudWatch Logsの詳細設定を行います
func (b *defaultLoggerBuilder) WithCloudWatchOptions(flushIntervalSecs int, batchSize int) LoggerBuilder {
	if flushIntervalSecs > 0 {
		b.cloudWatchFlushInterval = time.Duration(flushIntervalSecs) * time.Second
	}
	if batchSize > 0 {
		b.cloudWatchBatchSize = batchSize
	}
	return b
}

// Build はロガーを構築します
func (b *defaultLoggerBuilder) Build(ctx context.Context) (*Logger, error) {
	var handlers []slog.Handler
	var closers []io.Closer

	// コンソール出力ハンドラーの追加
	if b.useConsole {
		handler, err := createConsoleHandler(b.logLevel)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	// ファイル出力ハンドラーの追加
	if b.useFile && b.filePath != "" {
		handler, closer, err := createFileHandler(b.filePath, b.logLevel)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
		closers = append(closers, closer)
	}

	// CloudWatch Logsハンドラーの追加
	var cloudWatchHandler *CloudWatchHandler
	if b.useCloudWatch && b.logGroupName != "" {
		handler, err := createCloudWatchHandler(
			ctx,
			b.logLevel,
			b.createLogGroup,
			b.logGroupName,
			b.cloudWatchFlushInterval,
			b.cloudWatchBatchSize,
			b.awsOptions,
		)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
		closers = append(closers, handler)
		cloudWatchHandler = handler
	}

	// ハンドラーが何も設定されていない場合はデフォルトでコンソール出力を追加
	if len(handlers) == 0 {
		handler, err := createConsoleHandler(b.logLevel)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, handler)
	}

	// マルチハンドラーの作成
	var mainHandler slog.Handler
	if len(handlers) == 1 {
		mainHandler = handlers[0]
	} else {
		mainHandler = newMultiHandler(handlers...)
	}

	// 共通属性の追加
	mainHandler = mainHandler.WithAttrs([]slog.Attr{
		slog.String("app", b.appName),
		slog.String("env", b.environment),
	})

	// ロガーの作成
	slogger := slog.New(mainHandler)

	// グローバルロガーとして設定
	slog.SetDefault(slogger)

	logger := &Logger{
		slogger:           slogger,
		cloudWatchHandler: cloudWatchHandler,
		closers:           closers,
	}

	// グローバルインスタンスとして設定
	defaultLogger = logger

	return logger, nil
}

// createConsoleHandler はコンソール出力ハンドラーを作成します
func createConsoleHandler(level slog.Level) (slog.Handler, error) {
	return slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}), nil
}

// createFileHandler はファイル出力ハンドラーを作成します
func createFileHandler(filePath string, level slog.Level) (slog.Handler, io.Closer, error) {
	// ディレクトリが存在することを確認
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, nil, err
	}

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, nil, err
	}

	handler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: level,
	})

	return handler, file, nil
}

// createCloudWatchHandler はCloudWatch Logsハンドラーを作成します
func createCloudWatchHandler(
	ctx context.Context,
	level slog.Level,
	createLogGroup bool,
	logGroupName string,
	flushInterval time.Duration,
	batchSize int,
	awsOptions aws.Options,
) (*CloudWatchHandler, error) {
	// AWS設定の取得
	awsConfig, err := aws.NewAWSConfig(ctx, awsOptions)
	if err != nil {
		return nil, err
	}

	// CloudWatch Logsハンドラーの作成
	return NewCloudWatchHandler(
		awsConfig.CloudWatchLogs,
		logGroupName,
		WithLevel(level),
		WithCreateLogGroup(createLogGroup),
		WithFlushInterval(flushInterval),
		WithBatchSize(batchSize),
	)
}
