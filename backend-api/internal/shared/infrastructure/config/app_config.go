package config

import (
	"github.com/spf13/viper"
)

// AppConfig はアプリケーション基本設定
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
}

// SetDefaults はデフォルト値を設定します
func (c *AppConfig) SetDefaults() {
	viper.SetDefault("app.name", "aws-observability-ecommerce")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
}

// BindEnvironmentVariables は環境変数をバインドします
func (c *AppConfig) BindEnvironmentVariables() error {
	appEnvBindings := map[string]string{
		"app.name":        "APP_NAME",
		"app.version":     "APP_VERSION",
		"app.environment": "APP_ENV",
	}

	for viperKey, envKey := range appEnvBindings {
		if err := viper.BindEnv(viperKey, envKey); err != nil {
			return err
		}
	}

	return nil
}
