package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Config はアプリケーション全体の設定を格納する構造体
type Config struct {
	App struct {
		Name        string
		Version     string
		Environment string
	}
	Server struct {
		Port string
	}
	AWS struct {
		UseLocalStack bool
		Region        string
		Endpoint      string
		AccessKey     string
		SecretKey     string
		Token         string
	}
}

// アプリケーション設定インスタンス
var (
	config Config
	App    = &config.App
	Server = &config.Server
	AWS    = &config.AWS
)

// Load は環境変数と設定ファイルから設定をロードします
func Load() error {
	// 環境変数のデフォルト値の設定
	viper.SetDefault("app.name", "aws-observability-ecommerce")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")

	viper.SetDefault("server.port", "8080")

	// AWS関連のデフォルト設定
	viper.SetDefault("aws.useLocalStack", false)               // デフォルトはLocalStack無効
	viper.SetDefault("aws.region", "ap-northeast-1")           // デフォルトリージョン
	viper.SetDefault("aws.endpoint", "http://localstack:4566") // LocalStack使用時のエンドポイント
	viper.SetDefault("aws.accessKey", "")                      // デフォルトは空文字列
	viper.SetDefault("aws.secretKey", "")                      // デフォルトは空文字列
	viper.SetDefault("aws.token", "")                          // デフォルトは空文字列

	// 環境変数のバインド
	if err := viper.BindEnv("app.name", "APP_NAME"); err != nil {
		return err
	}
	if err := viper.BindEnv("app.version", "APP_VERSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("app.environment", "APP_ENV"); err != nil {
		return err
	}
	if err := viper.BindEnv("server.port", "PORT"); err != nil {
		return err
	}

	// AWS関連の環境変数のバインド
	if err := viper.BindEnv("aws.useLocalStack", "AWS_USE_LOCALSTACK"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.region", "AWS_REGION"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.endpoint", "AWS_ENDPOINT"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.accessKey", "AWS_ACCESS_KEY_ID"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.secretKey", "AWS_SECRET_ACCESS_KEY"); err != nil {
		return err
	}
	if err := viper.BindEnv("aws.token", "AWS_SESSION_TOKEN"); err != nil {
		return err
	}

	// 設定ファイルの読み込み（存在する場合）
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 設定ファイルの読み込み（存在しなくてもエラーとしない）
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
	}

	// viper.Unmarshalを使って設定を一括で読み込む
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}
