package config

import (
	"github.com/spf13/viper"
)

// ServerConfig はサーバー設定
type ServerConfig struct {
	Port         int    `mapstructure:"port"`
	Host         string `mapstructure:"host"`
	ReadTimeout  int    `mapstructure:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout"`
}

// SetDefaults はデフォルト値を設定します
func (c *ServerConfig) SetDefaults() {
	viper.SetDefault("server.port", 8000)
	viper.SetDefault("server.host", "0.0.0.0")
	viper.SetDefault("server.read_timeout", 30)
	viper.SetDefault("server.write_timeout", 30)
}

// BindEnvironmentVariables は環境変数をバインドします
func (c *ServerConfig) BindEnvironmentVariables() error {
	serverEnvBindings := map[string]string{
		"server.port": "PORT",
		"server.host": "HOST",
	}

	for viperKey, envKey := range serverEnvBindings {
		if err := viper.BindEnv(viperKey, envKey); err != nil {
			return err
		}
	}

	return nil
}
