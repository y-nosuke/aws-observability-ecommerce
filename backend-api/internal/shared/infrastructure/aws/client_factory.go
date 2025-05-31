package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	configPkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// ClientFactory はAWSクライアントを生成するファクトリ
type ClientFactory struct {
	config *aws.Config
}

// NewClientFactory は新しいAWSクライアントファクトリを作成します
func NewClientFactory(ctx context.Context, awsConfig configPkg.AWSConfig) (*ClientFactory, error) {
	cfg, err := createAWSConfig(ctx, awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS config: %w", err)
	}

	return &ClientFactory{
		config: cfg,
	}, nil
}

// GetS3Client はS3クライアントを取得します
func (f *ClientFactory) GetS3Client() *s3.Client {
	return s3.NewFromConfig(*f.config, func(o *s3.Options) {
		o.UsePathStyle = true // LocalStack対応
	})
}

// GetConfig はAWS設定を取得します
func (f *ClientFactory) GetConfig() *aws.Config {
	return f.config
}

// createAWSConfig はAWS設定を作成します
func createAWSConfig(ctx context.Context, awsConfig configPkg.AWSConfig) (*aws.Config, error) {
	var cfg aws.Config
	var err error

	if awsConfig.UseLocalStack {
		// LocalStack用の設定
		opts := []func(*config.LoadOptions) error{
			config.WithRegion(awsConfig.Region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				awsConfig.AccessKey,
				awsConfig.SecretKey,
				awsConfig.Token,
			)),
			config.WithBaseEndpoint(awsConfig.Endpoint),
		}

		if cfg, err = config.LoadDefaultConfig(ctx, opts...); err != nil {
			return nil, fmt.Errorf("failed to load LocalStack AWS config: %w", err)
		}

		fmt.Println("AWS config loaded for LocalStack environment")
	} else {
		// 本番環境用の標準設定
		if cfg, err = config.LoadDefaultConfig(ctx); err != nil {
			return nil, fmt.Errorf("failed to load production AWS config: %w", err)
		}

		fmt.Println("AWS config loaded for production environment")
	}

	return &cfg, nil
}
