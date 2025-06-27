package config

import (
	"time"

	"github.com/spf13/viper"
)

// ObservabilityConfig は監視設定を管理する構造体
type ObservabilityConfig struct {
	Logging LoggingConfig `mapstructure:"logging"`
	Metrics MetricsConfig `mapstructure:"metrics"`
	OTel    OTelConfig    `mapstructure:"otel"`
}

// LoggingConfig はログ設定を管理する構造体
type LoggingConfig struct {
	Level        string `mapstructure:"level"`
	Format       string `mapstructure:"format"`
	EnableOTel   bool   `mapstructure:"enable_otel"`
	MaxLogSizeMB int    `mapstructure:"max_log_size_mb"`
}

// MetricsConfig はログ設定を管理する構造体
type MetricsConfig struct {
	RequestTimeHistogramBoundaries []float64 `mapstructure:"request_time_histogram_boundaries"`
	RequestSizeHistogramBoundaries []float64 `mapstructure:"request_size_histogram_boundaries"`
}

// OTelConfig はOpenTelemetry設定を管理する構造体
type OTelConfig struct {
	ServiceName           string            `mapstructure:"service_name"`
	ServiceVersion        string            `mapstructure:"service_version"`
	ServiceNamespace      string            `mapstructure:"service_namespace"`
	DeploymentEnvironment string            `mapstructure:"deployment_environment"`
	Logging               OTelLoggingConfig `mapstructure:"logging"`
	Metrics               OTelMetricsConfig `mapstructure:"metrics"`
	Tracing               OTelTracingConfig `mapstructure:"tracing"`
}

// OTelLoggingConfig はOTelログ設定を管理する構造体
type OTelLoggingConfig struct {
	Enabled              bool          `mapstructure:"enabled"`
	Endpoint             string        `mapstructure:"endpoint"`
	Timeout              time.Duration `mapstructure:"timeout"`
	Compression          string        `mapstructure:"compression"`
	RetryEnabled         bool          `mapstructure:"retry_enabled"`
	RetryInitialInterval time.Duration `mapstructure:"retry_initial_interval"`
	RetryMaxInterval     time.Duration `mapstructure:"retry_max_interval"`
	RetryMaxElapsedTime  time.Duration `mapstructure:"retry_max_elapsed_time"`
	ExportTimeout        time.Duration `mapstructure:"export_timeout"`
	MaxQueueSize         int           `mapstructure:"max_queue_size"`
	MaxExportBatchSize   int           `mapstructure:"max_export_batch_size"`
}

// OTelMetricsConfig はOTelメトリクス設定を管理する構造体
type OTelMetricsConfig struct {
	Enabled              bool          `mapstructure:"enabled"`
	Endpoint             string        `mapstructure:"endpoint"`
	Timeout              time.Duration `mapstructure:"timeout"`
	Compression          string        `mapstructure:"compression"`
	RetryEnabled         bool          `mapstructure:"retry_enabled"`
	RetryInitialInterval time.Duration `mapstructure:"retry_initial_interval"`
	RetryMaxInterval     time.Duration `mapstructure:"retry_max_interval"`
	RetryMaxElapsedTime  time.Duration `mapstructure:"retry_max_elapsed_time"`
	Interval             time.Duration `mapstructure:"interval"`
}

// OTelTracingConfig はOTelトレース設定を管理する構造体
type OTelTracingConfig struct {
	Enabled              bool          `mapstructure:"enabled"`
	Endpoint             string        `mapstructure:"endpoint"`
	Timeout              time.Duration `mapstructure:"timeout"`
	Compression          string        `mapstructure:"compression"`
	RetryEnabled         bool          `mapstructure:"retry_enabled"`
	RetryInitialInterval time.Duration `mapstructure:"retry_initial_interval"`
	RetryMaxInterval     time.Duration `mapstructure:"retry_max_interval"`
	RetryMaxElapsedTime  time.Duration `mapstructure:"retry_max_elapsed_time"`
	BatchTimeout         time.Duration `mapstructure:"batch_timeout"`
	MaxQueueSize         int           `mapstructure:"max_queue_size"`
	MaxExportBatchSize   int           `mapstructure:"max_export_batch_size"`
	SamplingRatio        float64       `mapstructure:"sampling_ratio"`
}

// SetDefaults はObservabilityConfigのデフォルト値を設定します
func (c *ObservabilityConfig) SetDefaults() {
	// Logging defaults
	viper.SetDefault("observability.logging.level", "info")
	viper.SetDefault("observability.logging.format", "json")
	viper.SetDefault("observability.logging.enable_otel", true)
	viper.SetDefault("observability.logging.max_log_size_mb", 100)

	// OTel service defaults
	viper.SetDefault("observability.otel.service_name", "aws-observability-ecommerce")
	viper.SetDefault("observability.otel.service_version", "1.0.0")
	viper.SetDefault("observability.otel.service_namespace", "ecommerce")
	viper.SetDefault("observability.otel.deployment_environment", "development")

	// OTel logging defaults
	viper.SetDefault("observability.otel.logging.enabled", true)
	viper.SetDefault("observability.otel.logging.endpoint", "otel-collector:4317")
	viper.SetDefault("observability.otel.logging.timeout", "10s")
	viper.SetDefault("observability.otel.logging.compression", "gzip")
	viper.SetDefault("observability.otel.logging.retry_enabled", true)
	viper.SetDefault("observability.otel.logging.retry_initial_interval", "1s")
	viper.SetDefault("observability.otel.logging.retry_max_interval", "30s")
	viper.SetDefault("observability.otel.logging.retry_max_elapsed_time", "60s")
	viper.SetDefault("observability.otel.logging.export_timeout", "30s")
	viper.SetDefault("observability.otel.logging.max_queue_size", 2048)
	viper.SetDefault("observability.otel.logging.max_export_batch_size", 512)

	// OTel metrics defaults
	viper.SetDefault("observability.otel.metrics.enabled", true)
	viper.SetDefault("observability.otel.metrics.endpoint", "otel-collector:4317")
	viper.SetDefault("observability.otel.metrics.timeout", "10s")
	viper.SetDefault("observability.otel.metrics.compression", "gzip")
	viper.SetDefault("observability.otel.metrics.retry_enabled", true)
	viper.SetDefault("observability.otel.metrics.retry_initial_interval", "1s")
	viper.SetDefault("observability.otel.metrics.retry_max_interval", "30s")
	viper.SetDefault("observability.otel.metrics.retry_max_elapsed_time", "60s")
	viper.SetDefault("observability.otel.metrics.interval", "1s")
	viper.SetDefault("observability.otel.metrics.request_time_histogram_boundaries", []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0})
	viper.SetDefault("observability.otel.metrics.request_size_histogram_boundaries", []float64{64, 256, 1024, 4096, 16384, 65536, 262144, 1048576})

	// OTel tracing defaults
	viper.SetDefault("observability.otel.tracing.enabled", true)
	viper.SetDefault("observability.otel.tracing.endpoint", "otel-collector:4317")
	viper.SetDefault("observability.otel.tracing.timeout", "10s")
	viper.SetDefault("observability.otel.tracing.compression", "gzip")
	viper.SetDefault("observability.otel.tracing.retry_enabled", true)
	viper.SetDefault("observability.otel.tracing.retry_initial_interval", "1s")
	viper.SetDefault("observability.otel.tracing.retry_max_interval", "30s")
	viper.SetDefault("observability.otel.tracing.retry_max_elapsed_time", "60s")
	viper.SetDefault("observability.otel.tracing.batch_timeout", "1s")
	viper.SetDefault("observability.otel.tracing.max_queue_size", 2048)
	viper.SetDefault("observability.otel.tracing.max_export_batch_size", 512)
	viper.SetDefault("observability.otel.tracing.sampling_ratio", 1.0)
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
	if err := viper.BindEnv("observability.logging.max_log_size_mb", "OBSERVABILITY_LOGGING_MAX_LOG_SIZE_MB"); err != nil {
		return err
	}

	// OTel service 環境変数
	if err := viper.BindEnv("observability.otel.service_name", "OBSERVABILITY_OTEL_SERVICE_NAME"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.service_version", "OBSERVABILITY_OTEL_SERVICE_VERSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.service_namespace", "OBSERVABILITY_OTEL_SERVICE_NAMESPACE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.deployment_environment", "OBSERVABILITY_OTEL_DEPLOYMENT_ENVIRONMENT"); err != nil {
		return err
	}

	// OTel logging 環境変数
	if err := viper.BindEnv("observability.otel.logging.enabled", "OBSERVABILITY_OTEL_LOGGING_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.endpoint", "OBSERVABILITY_OTEL_LOGGING_ENDPOINT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.timeout", "OBSERVABILITY_OTEL_LOGGING_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.compression", "OBSERVABILITY_OTEL_LOGGING_COMPRESSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.retry_enabled", "OBSERVABILITY_OTEL_LOGGING_RETRY_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.retry_initial_interval", "OBSERVABILITY_OTEL_LOGGING_RETRY_INITIAL_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.retry_max_interval", "OBSERVABILITY_OTEL_LOGGING_RETRY_MAX_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.retry_max_elapsed_time", "OBSERVABILITY_OTEL_LOGGING_RETRY_MAX_ELAPSED_TIME"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.export_timeout", "OBSERVABILITY_OTEL_LOGGING_EXPORT_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.max_queue_size", "OBSERVABILITY_OTEL_LOGGING_MAX_QUEUE_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.logging.max_export_batch_size", "OBSERVABILITY_OTEL_LOGGING_MAX_EXPORT_BATCH_SIZE"); err != nil {
		return err
	}

	// OTel metrics 環境変数
	if err := viper.BindEnv("observability.otel.metrics.enabled", "OBSERVABILITY_OTEL_METRICS_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.endpoint", "OBSERVABILITY_OTEL_METRICS_ENDPOINT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.timeout", "OBSERVABILITY_OTEL_METRICS_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.compression", "OBSERVABILITY_OTEL_METRICS_COMPRESSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.retry_enabled", "OBSERVABILITY_OTEL_METRICS_RETRY_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.retry_initial_interval", "OBSERVABILITY_OTEL_METRICS_RETRY_INITIAL_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.retry_max_interval", "OBSERVABILITY_OTEL_METRICS_RETRY_MAX_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.retry_max_elapsed_time", "OBSERVABILITY_OTEL_METRICS_RETRY_MAX_ELAPSED_TIME"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.interval", "OBSERVABILITY_OTEL_METRICS_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.request_time_histogram_boundaries", "OBSERVABILITY_OTEL_METRICS_REQUEST_TIME_HISTOGRAM_BOUNDARIES"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.metrics.request_size_histogram_boundaries", "OBSERVABILITY_OTEL_METRICS_REQUEST_SIZE_HISTOGRAM_BOUNDARIES"); err != nil {
		return err
	}

	// OTel tracing 環境変数
	if err := viper.BindEnv("observability.otel.tracing.enabled", "OBSERVABILITY_OTEL_TRACING_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.endpoint", "OBSERVABILITY_OTEL_TRACING_ENDPOINT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.timeout", "OBSERVABILITY_OTEL_TRACING_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.compression", "OBSERVABILITY_OTEL_TRACING_COMPRESSION"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.retry_enabled", "OBSERVABILITY_OTEL_TRACING_RETRY_ENABLED"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.retry_initial_interval", "OBSERVABILITY_OTEL_TRACING_RETRY_INITIAL_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.retry_max_interval", "OBSERVABILITY_OTEL_TRACING_RETRY_MAX_INTERVAL"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.retry_max_elapsed_time", "OBSERVABILITY_OTEL_TRACING_RETRY_MAX_ELAPSED_TIME"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.batch_timeout", "OBSERVABILITY_OTEL_TRACING_BATCH_TIMEOUT"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.max_queue_size", "OBSERVABILITY_OTEL_TRACING_MAX_QUEUE_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.max_export_batch_size", "OBSERVABILITY_OTEL_TRACING_MAX_EXPORT_BATCH_SIZE"); err != nil {
		return err
	}
	if err := viper.BindEnv("observability.otel.tracing.sampling_ratio", "OBSERVABILITY_OTEL_TRACING_SAMPLING_RATIO"); err != nil {
		return err
	}
	return nil
}
