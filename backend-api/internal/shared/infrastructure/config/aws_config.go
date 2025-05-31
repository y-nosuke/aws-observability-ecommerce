package config

import (
	"github.com/spf13/viper"
)

// AWSConfig はAWS設定
type AWSConfig struct {
	UseLocalStack bool   `mapstructure:"use_localstack"`
	Region        string `mapstructure:"region"`
	Endpoint      string `mapstructure:"endpoint"`
	AccessKey     string `mapstructure:"access_key"`
	SecretKey     string `mapstructure:"secret_key"`
	Token         string `mapstructure:"token"`

	// S3設定
	S3 S3Config `mapstructure:"s3"`
}

// S3Config はS3固有の設定
type S3Config struct {
	BucketName      string `mapstructure:"bucket_name"`
	PresignedURLTTL int    `mapstructure:"presigned_url_ttl"` // seconds
	UsePathStyle    bool   `mapstructure:"use_path_style"`
}

// SetDefaults はデフォルト値を設定します
func (c *AWSConfig) SetDefaults() {
	viper.SetDefault("aws.use_localstack", false)
	viper.SetDefault("aws.region", "ap-northeast-1")
	viper.SetDefault("aws.endpoint", "http://localstack:4566")
	viper.SetDefault("aws.access_key", "")
	viper.SetDefault("aws.secret_key", "")
	viper.SetDefault("aws.token", "")

	// S3設定
	viper.SetDefault("aws.s3.bucket_name", "product-images")
	viper.SetDefault("aws.s3.presigned_url_ttl", 3600) // 1 hour
	viper.SetDefault("aws.s3.use_path_style", true)
}

// BindEnvironmentVariables は環境変数をバインドします
func (c *AWSConfig) BindEnvironmentVariables() error {
	awsEnvBindings := map[string]string{
		"aws.use_localstack": "AWS_USE_LOCALSTACK",
		"aws.region":         "AWS_REGION",
		"aws.endpoint":       "AWS_ENDPOINT",
		"aws.access_key":     "AWS_ACCESS_KEY_ID",
		"aws.secret_key":     "AWS_SECRET_ACCESS_KEY",
		"aws.token":          "AWS_SESSION_TOKEN",
		"aws.s3.bucket_name": "AWS_S3_BUCKET_NAME",
	}

	for viperKey, envKey := range awsEnvBindings {
		if err := viper.BindEnv(viperKey, envKey); err != nil {
			return err
		}
	}

	return nil
}
