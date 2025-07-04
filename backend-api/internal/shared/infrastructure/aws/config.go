package aws

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws"

	configPkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// NewAWSConfig は新しいAWS設定を作成します
func NewAWSConfig(ctx context.Context, awsConfig configPkg.AWSConfig) (aws.Config, error) {
	var opts []func(*config.LoadOptions) error
	if awsConfig.UseLocalStack {
		// LocalStack用の設定
		opts = []func(*config.LoadOptions) error{
			config.WithRegion(awsConfig.Region),
			config.WithCredentialsProvider(
				credentials.NewStaticCredentialsProvider(
					awsConfig.AccessKey,
					awsConfig.SecretKey,
					awsConfig.Token,
				),
			),
			config.WithBaseEndpoint(awsConfig.Endpoint),
		}
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return aws.Config{}, fmt.Errorf("failed to load aws config: %w", err)
	}

	// OpenTelemetryインストルメンテーションを追加
	otelaws.AppendMiddlewares(&cfg.APIOptions)

	slog.InfoContext(ctx, "AWS config loaded")

	return cfg, nil
}
