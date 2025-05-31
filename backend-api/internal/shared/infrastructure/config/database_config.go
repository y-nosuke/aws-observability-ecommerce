package config

import (
	"github.com/spf13/viper"
)

// DatabaseConfig はデータベース設定
type DatabaseConfig struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	User            string `mapstructure:"user"`
	Password        string `mapstructure:"password"`
	Name            string `mapstructure:"name"`
	MaxOpenConns    int    `mapstructure:"max_open_conns"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	ConnMaxLifetime int    `mapstructure:"conn_max_lifetime"` // minutes
}

// SetDefaults はデフォルト値を設定します
func (c *DatabaseConfig) SetDefaults() {
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 3306)
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "password")
	viper.SetDefault("database.name", "ecommerce")
	viper.SetDefault("database.max_open_conns", 25)
	viper.SetDefault("database.max_idle_conns", 25)
	viper.SetDefault("database.conn_max_lifetime", 5) // 5 minutes
}

// BindEnvironmentVariables は環境変数をバインドします
func (c *DatabaseConfig) BindEnvironmentVariables() error {
	dbEnvBindings := map[string]string{
		"database.host":     "DB_HOST",
		"database.port":     "DB_PORT",
		"database.user":     "DB_USER",
		"database.password": "DB_PASSWORD",
		"database.name":     "DB_NAME",
	}

	for viperKey, envKey := range dbEnvBindings {
		if err := viper.BindEnv(viperKey, envKey); err != nil {
			return err
		}
	}

	return nil
}
