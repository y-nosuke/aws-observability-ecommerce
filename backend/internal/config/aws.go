package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

// NewAWSConfig はAWS SDKの設定を生成します
func NewAWSConfig(ctx context.Context, cfg *Config) (aws.Config, error) {
	var options []func(*config.LoadOptions) error

	// カスタムエンドポイントが設定されている場合（LocalStackなど）
	if cfg.AWS.Endpoint != "" {
		options = append(options, config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{
						URL:           cfg.AWS.Endpoint,
						SigningRegion: cfg.AWS.Region,
					}, nil
				},
			),
		))
	}

	// 認証情報の設定
	options = append(options, config.WithRegion(cfg.AWS.Region))

	// LocalStackでは静的な認証情報を使用
	if cfg.AWS.AccessKey != "" && cfg.AWS.SecretKey != "" {
		options = append(options, config.WithCredentialsProvider(
			aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
				return aws.Credentials{
					AccessKeyID:     cfg.AWS.AccessKey,
					SecretAccessKey: cfg.AWS.SecretKey,
				}, nil
			}),
		))
	}

	return config.LoadDefaultConfig(ctx, options...)
}
