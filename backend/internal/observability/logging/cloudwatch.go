package logging

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"golang.org/x/exp/slog"
)

// CloudWatchHandler はslogのハンドラーとしてCloudWatch Logsにログを送信します
type CloudWatchHandler struct {
	client        *cloudwatchlogs.Client
	logGroupName  string
	logStreamName string
	minLevel      slog.Level
	nextHandler   slog.Handler
}

// NewCloudWatchHandler はCloudWatch Logs向けのslogハンドラーを作成します
func NewCloudWatchHandler(ctx context.Context, awsCfg aws.Config, logGroupName, logStreamName string, minLevel slog.Level, nextHandler slog.Handler) *CloudWatchHandler {
	client := cloudwatchlogs.NewFromConfig(awsCfg)

	return &CloudWatchHandler{
		client:        client,
		logGroupName:  logGroupName,
		logStreamName: logStreamName,
		minLevel:      minLevel,
		nextHandler:   nextHandler,
	}
}

// 簡略化のため、実装の詳細は省略しています
// 実際の実装では、ログを収集してバッチで送信するロジックを実装します
