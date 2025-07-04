package otel

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"
)

var tracer trace.Tracer

// NewTracer はトレーサーを初期化します
func NewTracer(provider *sdktrace.TracerProvider) trace.Tracer {
	tracer = provider.Tracer(utils.GetModulePath())
	return tracer
}
