package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// config はアプリケーション設定を表す構造体
type config struct {
	App struct {
		Name        string
		Version     string
		Environment string
	}
	Server struct {
		Port string
	}
}

var Config *config

func init() {
	var err error
	// 構成のロード
	Config, err = LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
		os.Exit(1)
	}
}

// LoadConfig は環境変数と設定ファイルから設定をロードします
func LoadConfig() (*config, error) {
	// デフォルト値の設定
	viper.SetDefault("app.name", "ecommerce-app")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("server.port", "8080")

	// 環境変数の読み込み
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	// 環境変数のバインド
	if err := viper.BindEnv("app.name", "APP_NAME"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("app.version", "APP_VERSION"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("app.environment", "APP_ENV"); err != nil {
		return nil, err
	}
	if err := viper.BindEnv("server.port", "PORT"); err != nil {
		return nil, err
	}

	// 設定ファイルの読み込み（存在する場合）
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")

	// 設定ファイルの読み込み（存在しなくてもエラーとしない）
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	// 設定を構造体にマッピング
	var c config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
