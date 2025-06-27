package otel

import (
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.34.0"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// NewResource は、アプリケーションの基本的な情報を含むOpenTelemetryリソースを作成します。
func NewResource(config config.OTelConfig) (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(config.ServiceName),
			semconv.ServiceVersion(config.ServiceVersion),
			semconv.ServiceNamespace(config.ServiceNamespace),
			semconv.DeploymentEnvironmentName(config.DeploymentEnvironment),
		),
	)
}
