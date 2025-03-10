package config

import (
	"github.com/spf13/viper"
)

// Config アプリケーション設定の構造体
type Config struct {
	LogLevel string `mapstructure:"log_level"`
	Server   ServerConfig
	Database DatabaseConfig
	AWS      AWSConfig
}

// ServerConfig サーバー関連の設定
type ServerConfig struct {
	Port int `mapstructure:"port"`
}

// DatabaseConfig データベース関連の設定
type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
}

// AWSConfig AWS関連の設定
type AWSConfig struct {
	Region    string `mapstructure:"region"`
	Endpoint  string `mapstructure:"endpoint"`
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
}

// Load 設定ファイルを読み込む
func Load(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.AutomaticEnv()

	// 環境変数の設定
	viper.SetEnvPrefix("APP")
	viper.BindEnv("log_level", "LOG_LEVEL")
	viper.BindEnv("server.port", "SERVER_PORT")
	viper.BindEnv("database.host", "DB_HOST")
	viper.BindEnv("database.port", "DB_PORT")
	viper.BindEnv("database.user", "DB_USER")
	viper.BindEnv("database.password", "DB_PASSWORD")
	viper.BindEnv("database.name", "DB_NAME")
	viper.BindEnv("aws.region", "AWS_REGION")
	viper.BindEnv("aws.endpoint", "AWS_ENDPOINT")
	viper.BindEnv("aws.access_key", "AWS_ACCESS_KEY_ID")
	viper.BindEnv("aws.secret_key", "AWS_SECRET_ACCESS_KEY")

	// デフォルト値の設定
	viper.SetDefault("log_level", "info")
	viper.SetDefault("server.port", 8080)
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.user", "ecommerce")
	viper.SetDefault("database.name", "ecommerce")
	viper.SetDefault("aws.region", "us-east-1")

	if err := viper.ReadInConfig(); err != nil {
		// 設定ファイルがない場合はエラーにしない
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, err
		}
	}

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
