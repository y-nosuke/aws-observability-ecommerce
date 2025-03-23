package config

import (
	"log"
	"os"

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
	Log struct {
		Level  string
		Format string
	}
}

// アプリケーション設定インスタンス
var (
	config Config
	App    = &config.App
	Server = &config.Server
	Log    = &config.Log
)

func init() {
	// 構成のロード
	if err := Load(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		os.Exit(1)
	}
}

// Load は環境変数と設定ファイルから設定をロードします
func Load() error {
	// 環境変数のデフォルト値の設定
	viper.SetDefault("app.name", "ecommerce-app")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("server.port", "8080")
	// ログ関連のデフォルト設定
	viper.SetDefault("log.level", "info")  // デフォルトはinfo
	viper.SetDefault("log.format", "json") // デフォルトはJSON

	// 開発環境ではデバッグログを有効化
	if viper.GetString("app.environment") == "development" {
		viper.SetDefault("log.level", "debug")
	}

	// 本番環境では警告以上のみを記録
	if viper.GetString("app.environment") == "production" {
		viper.SetDefault("log.level", "warn")
	}

	// 環境変数の読み込み
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

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
	if err := viper.BindEnv("log.level", "LOG_LEVEL"); err != nil {
		return err
	}
	if err := viper.BindEnv("log.format", "LOG_FORMAT"); err != nil {
		return err
	}

	// 設定ファイルの読み込み（存在する場合）
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 設定ファイルの読み込み（存在しなくてもエラーとしない）
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	// viper.Unmarshalを使って設定を一括で読み込む
	if err := viper.Unmarshal(&config); err != nil {
		return err
	}

	return nil
}
