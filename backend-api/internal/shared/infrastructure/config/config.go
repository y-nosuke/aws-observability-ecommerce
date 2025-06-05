package config

import (
	"errors"

	"github.com/spf13/viper"
)

// ConfigSection は設定セクションのインターフェース
type ConfigSection interface {
	SetDefaults()
	BindEnvironmentVariables() error
}

// Config は全体の設定を統合する構造体
type Config struct {
	App           AppConfig           `mapstructure:"app"`
	Server        ServerConfig        `mapstructure:"server"`
	Database      DatabaseConfig      `mapstructure:"database"`
	AWS           AWSConfig           `mapstructure:"aws"`
	Observability ObservabilityConfig `mapstructure:"observability"`
}

var (
	// 各設定セクションへの直接アクセス用のグローバル変数
	App           AppConfig
	Server        ServerConfig
	Database      DatabaseConfig
	AWS           AWSConfig
	Observability ObservabilityConfig

	// configSections は設定セクションのマップ
	configSections = map[string]ConfigSection{
		"app":           &AppConfig{},
		"server":        &ServerConfig{},
		"database":      &DatabaseConfig{},
		"aws":           &AWSConfig{},
		"observability": &ObservabilityConfig{},
	}
)

// LoadConfig は環境変数と設定ファイルから設定をロードします
func LoadConfig() error {
	// デフォルト値の設定
	for _, section := range configSections {
		section.SetDefaults()
	}

	// 環境変数のバインド
	for _, section := range configSections {
		if err := section.BindEnvironmentVariables(); err != nil {
			return err
		}
	}

	// 設定ファイルの読み込み
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("./internal/shared/infrastructure/config")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return err
		}
		// 設定ファイルが見つからない場合は無視（環境変数のみで動作）
	}

	// 設定の構造体へのマッピング
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return err
	}

	// 各設定セクションをグローバル変数にコピー
	App = cfg.App
	Server = cfg.Server
	Database = cfg.Database
	AWS = cfg.AWS
	Observability = cfg.Observability

	return nil
}
