package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Credentials はAWS認証情報を保持する構造体
type Credentials struct {
	AccessKey string
	SecretKey string
	Token     string
}

// Options はAWS設定オプションを保持する構造体
type Options struct {
	UseLocalStack bool
	Region        string
	Endpoint      string
	Credentials   Credentials
}

// Config はAWS設定を保持する構造体
type Config struct {
	Config *aws.Config
	S3     *s3.Client
}

// NewAWSConfig は新しいAWS設定を作成します
func NewAWSConfig(ctx context.Context, options Options) (*Config, error) {
	var cfg aws.Config
	var err error

	fmt.Printf("Initializing AWS Config (UseLocalStack: %v, Region: %s, Endpoint: %s)\n",
		options.UseLocalStack, options.Region, options.Endpoint)

	if options.UseLocalStack {
		// LocalStack用の設定
		// 設定オプションを準備
		opts := []func(*config.LoadOptions) error{
			config.WithRegion(options.Region),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
				options.Credentials.AccessKey,
				options.Credentials.SecretKey,
				options.Credentials.Token,
			)),
			config.WithBaseEndpoint(options.Endpoint),
		}

		// 設定のロード
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

	// S3クライアントの作成
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true // S3 Path Style を有効化
	})

	return &Config{
		Config: &cfg,
		S3:     s3Client,
	}, nil
}
