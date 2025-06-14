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
	Level        string `mapstructure:"level"`
	Format       string `mapstructure:"format"`
	EnableOTel   bool   `mapstructure:"enable_otel"`
	MaxLogSizeMB int    `mapstructure:"max_log_size_mb"`
}

// OTelConfig はOpenTelemetry設定を管理する構造体
type OTelConfig struct {
	ServiceName           string            `mapstructure:"service_name"`
	ServiceVersion        string            `mapstructure:"service_version"`
	ServiceNamespace      string            `mapstructure:"service_namespace"`
	DeploymentEnvironment string            `mapstructure:"deployment_environment"`
	Collector             CollectorConfig   `mapstructure:"collector"`
	Logging               OTelLoggingConfig `mapstructure:"logging"`
	Metrics               OTelMetricsConfig `mapstructure:"metrics"`
	Tracing               OTelTracingConfig `mapstructure:"tracing"`
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

// OTelLoggingConfig はOTelログ設定を管理する構造体
type OTelLoggingConfig struct {
	Enabled            bool          `mapstructure:"enabled"`
	BatchTimeout       time.Duration `mapstructure:"batch_timeout"`
	MaxQueueSize       int           `mapstructure:"max_queue_size"`
	MaxExportBatchSize int           `mapstructure:"max_export_batch_size"`
	ExportTimeout      time.Duration `mapstructure:"export_timeout"`
}

// OTelMetricsConfig はOTelメトリクス設定を管理する構造体
type OTelMetricsConfig struct {
	Enabled             bool          `mapstructure:"enabled"`
	BatchTimeout        time.Duration `mapstructure:"batch_timeout"`
	MaxQueueSize        int           `mapstructure:"max_queue_size"`
	MaxExportBatchSize  int           `mapstructure:"max_export_batch_size"`
	ExportTimeout       time.Duration `mapstructure:"export_timeout"`
	HistogramBoundaries []float64     `mapstructure:"histogram_boundaries"`
}

// OTelTracingConfig はOTelトレース設定を管理する構造体
type OTelTracingConfig struct {
	Enabled            bool          `mapstructure:"enabled"`
	SamplingRatio      float64       `mapstructure:"sampling_ratio"`
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
	viper.SetDefault("observability.otel.collector.endpoint", "otel-collector:4318")
	viper.SetDefault("observability.otel.collector.timeout", "10s")
	viper.SetDefault("observability.otel.collector.retry_enabled", true)
	viper.SetDefault("observability.otel.collector.retry_max_attempts", 3)
	viper.SetDefault("observability.otel.collector.retry_initial_interval", "1s")
	viper.SetDefault("observability.otel.collector.retry_max_interval", "30s")
	viper.SetDefault("observability.otel.collector.compression", "gzip")

	// OTel logging defaults
	viper.SetDefault("observability.otel.logging.enabled", true)
	viper.SetDefault("observability.otel.logging.batch_timeout", "1s")
	viper.SetDefault("observability.otel.logging.max_queue_size", 2048)
	viper.SetDefault("observability.otel.logging.max_export_batch_size", 512)
	viper.SetDefault("observability.otel.logging.export_timeout", "30s")

	// OTel metrics defaults
	viper.SetDefault("observability.otel.metrics.enabled", true)
	viper.SetDefault("observability.otel.metrics.batch_timeout", "1s")
	viper.SetDefault("observability.otel.metrics.max_queue_size", 2048)
	viper.SetDefault("observability.otel.metrics.max_export_batch_size", 512)
	viper.SetDefault("observability.otel.metrics.export_timeout", "30s")
	viper.SetDefault("observability.otel.metrics.histogram_boundaries", []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0})

	// OTel tracing defaults
	viper.SetDefault("observability.otel.tracing.enabled", true)
	viper.SetDefault("observability.otel.tracing.sampling_ratio", 1.0)
	viper.SetDefault("observability.otel.tracing.batch_timeout", "1s")
	viper.SetDefault("observability.otel.tracing.max_queue_size", 2048)
	viper.SetDefault("observability.otel.tracing.max_export_batch_size", 512)
	viper.SetDefault("observability.otel.tracing.export_timeout", "30s")
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
	if err := viper.BindEnv("observability.logging.max_log_size_mb", "OBSERVABILITY_LOGGING_MAX_LOG_SIZE_MB"); err != nil {
		return err
	}

	// OTel service 環境変数
	if err := viper.BindEnv("observability.otel.service_name", "OTEL_SERVICE_NAME"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.service_version", "OTEL_SERVICE_VERSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.service_namespace", "OTEL_SERVICE_NAMESPACE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.deployment_environment", "OTEL_DEPLOYMENT_ENVIRONMENT"); err != nil {
		return err
	}

	// OTel collector 環境変数
	if err := viper.BindEnv("observability.otel.collector.endpoint", "OTEL_EXPORTER_OTLP_ENDPOINT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.collector.timeout", "OTEL_EXPORTER_OTLP_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.collector.retry_enabled", "OTEL_EXPORTER_OTLP_RETRY_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.collector.retry_max_attempts", "OTEL_EXPORTER_OTLP_RETRY_MAX_ATTEMPTS"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.collector.retry_initial_interval", "OTEL_EXPORTER_OTLP_RETRY_INITIAL_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.collector.retry_max_interval", "OTEL_EXPORTER_OTLP_RETRY_MAX_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.collector.compression", "OTEL_EXPORTER_OTLP_COMPRESSION"); err != nil {
		return err
	}

	// OTel logging 環境変数
	if err := viper.BindEnv("observability.otel.logging.enabled", "OTEL_LOGS_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.batch_timeout", "OTEL_LOGS_BATCH_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.max_queue_size", "OTEL_LOGS_MAX_QUEUE_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.max_export_batch_size", "OTEL_LOGS_MAX_EXPORT_BATCH_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.export_timeout", "OTEL_LOGS_EXPORT_TIMEOUT"); err != nil {
		return err
	}

	// OTel metrics 環境変数
	if err := viper.BindEnv("observability.otel.metrics.enabled", "OTEL_METRICS_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.batch_timeout", "OTEL_METRICS_BATCH_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.max_queue_size", "OTEL_METRICS_MAX_QUEUE_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.max_export_batch_size", "OTEL_METRICS_MAX_EXPORT_BATCH_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.export_timeout", "OTEL_METRICS_EXPORT_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.histogram_boundaries", "OTEL_METRICS_HISTOGRAM_BOUNDARIES"); err != nil {
		return err
	}

	// OTel tracing 環境変数
	if err := viper.BindEnv("observability.otel.tracing.enabled", "OTEL_TRACES_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.sampling_ratio", "OTEL_TRACES_SAMPLING_RATIO"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.batch_timeout", "OTEL_TRACES_BATCH_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.max_queue_size", "OTEL_TRACES_MAX_QUEUE_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.max_export_batch_size", "OTEL_TRACES_MAX_EXPORT_BATCH_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.export_timeout", "OTEL_TRACES_EXPORT_TIMEOUT"); err != nil {
		return err
	}

	return nil
}
