package config

import (
	"time"

	"github.com/spf13/viper"
)

// ObservabilityConfig は監視設定を管理する構造体
type ObservabilityConfig struct {
	Logging LoggingConfig `mapstructure:"logging"`
	OTel    OTelConfig    `mapstructure:"otel"`
}

// LoggingConfig はログ設定を管理する構造体
type LoggingConfig struct {
	Level              string `mapstructure:"level"`
	Format             string `mapstructure:"format"`
	EnableOTel         bool   `mapstructure:"enable_otel"`
	EnableTraceContext bool   `mapstructure:"enable_trace_context"`
	MaxLogSizeMB       int    `mapstructure:"max_log_size_mb"`
}

// OTelConfig はOpenTelemetry設定を管理する構造体
type OTelConfig struct {
	ServiceName           string            `mapstructure:"service_name"`
	ServiceVersion        string            `mapstructure:"service_version"`
	ServiceNamespace      string            `mapstructure:"service_namespace"`
	DeploymentEnvironment string            `mapstructure:"deployment_environment"`
	Collector             CollectorConfig   `mapstructure:"collector"`
	Tracing               TracingConfig     `mapstructure:"tracing"`
	Logging               OTelLoggingConfig `mapstructure:"logging"`
}

// CollectorConfig はOTel Collector設定を管理する構造体
type CollectorConfig struct {
	Endpoint             string        `mapstructure:"endpoint"`
	Timeout              time.Duration `mapstructure:"timeout"`
	RetryEnabled         bool          `mapstructure:"retry_enabled"`
	RetryMaxAttempts     int           `mapstructure:"retry_max_attempts"`
	RetryInitialInterval time.Duration `mapstructure:"retry_initial_interval"`
	RetryMaxInterval     time.Duration `mapstructure:"retry_max_interval"`
	Compression          string        `mapstructure:"compression"`
}

// TracingConfig はトレーシング設定を管理する構造体
type TracingConfig struct {
	Enabled              bool    `mapstructure:"enabled"`
	SampleRate           float64 `mapstructure:"sample_rate"`
	MaxAttributesPerSpan int     `mapstructure:"max_attributes_per_span"`
	MaxEventsPerSpan     int     `mapstructure:"max_events_per_span"`
}

// OTelLoggingConfig はOTelログ設定を管理する構造体
type OTelLoggingConfig struct {
	BatchTimeout       time.Duration `mapstructure:"batch_timeout"`
	MaxQueueSize       int           `mapstructure:"max_queue_size"`
	MaxExportBatchSize int           `mapstructure:"max_export_batch_size"`
	ExportTimeout      time.Duration `mapstructure:"export_timeout"`
}

// SetDefaults はObservabilityConfigのデフォルト値を設定します
func (c *ObservabilityConfig) SetDefaults() {
	// Logging defaults
	viper.SetDefault("observability.logging.level", "info")
	viper.SetDefault("observability.logging.format", "json")
	viper.SetDefault("observability.logging.enable_otel", true)
	viper.SetDefault("observability.logging.enable_trace_context", true)
	viper.SetDefault("observability.logging.max_log_size_mb", 100)

	// OTel service defaults
	viper.SetDefault("observability.otel.service_name", "aws-observability-ecommerce")
	viper.SetDefault("observability.otel.service_version", "1.0.0")
	viper.SetDefault("observability.otel.service_namespace", "ecommerce")
	viper.SetDefault("observability.otel.deployment_environment", "development")

	// OTel collector defaults
	viper.SetDefault("observability.otel.collector.endpoint", "http://otel-collector:4318")
	viper.SetDefault("observability.otel.collector.timeout", "10s")
	viper.SetDefault("observability.otel.collector.retry_enabled", true)
	viper.SetDefault("observability.otel.collector.retry_max_attempts", 3)
	viper.SetDefault("observability.otel.collector.retry_initial_interval", "1s")
	viper.SetDefault("observability.otel.collector.retry_max_interval", "30s")
	viper.SetDefault("observability.otel.collector.compression", "gzip")

	// OTel tracing defaults
	viper.SetDefault("observability.otel.tracing.enabled", true)
	viper.SetDefault("observability.otel.tracing.sample_rate", 1.0)
	viper.SetDefault("observability.otel.tracing.max_attributes_per_span", 128)
	viper.SetDefault("observability.otel.tracing.max_events_per_span", 128)

	// OTel logging defaults
	viper.SetDefault("observability.otel.logging.batch_timeout", "1s")
	viper.SetDefault("observability.otel.logging.max_queue_size", 2048)
	viper.SetDefault("observability.otel.logging.max_export_batch_size", 512)
	viper.SetDefault("observability.otel.logging.export_timeout", "30s")
}

// BindEnvironmentVariables は環境変数をバインドします
func (c *ObservabilityConfig) BindEnvironmentVariables() error {
	// Logging 環境変数
	if err := viper.BindEnv("observability.logging.level", "OBSERVABILITY_LOGGING_LEVEL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.logging.format", "OBSERVABILITY_LOGGING_FORMAT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.logging.enable_otel", "OBSERVABILITY_LOGGING_ENABLE_OTEL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.logging.enable_trace_context", "OBSERVABILITY_LOGGING_ENABLE_TRACE_CONTEXT"); err != nil {
		return err
	}

	// OTel service 環境変数
	if err := viper.BindEnv("observability.otel.service_name", "OTEL_SERVICE_NAME"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.service_version", "OTEL_SERVICE_VERSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.deployment_environment", "OTEL_DEPLOYMENT_ENVIRONMENT"); err != nil {
		return err
	}

	// OTel collector 環境変数
	if err := viper.BindEnv("observability.otel.collector.endpoint", "OTEL_EXPORTER_OTLP_ENDPOINT"); err != nil {
		return err
	}

	// OTel tracing 環境変数
	if err := viper.BindEnv("observability.otel.tracing.enabled", "OTEL_TRACES_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.sample_rate", "OTEL_TRACES_SAMPLER_ARG"); err != nil {
		return err
	}

	return nil
}
