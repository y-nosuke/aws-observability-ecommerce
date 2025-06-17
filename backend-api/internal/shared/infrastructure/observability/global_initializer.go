package observability

import (
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/metrics"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
)

// GlobalObservabilityInitializer はグローバルオブザーバビリティの初期化を担当
type GlobalObservabilityInitializer struct {
	providerFactory observability.ProviderFactory
}

// NewGlobalObservabilityInitializer は新しい初期化サービスを作成
func NewGlobalObservabilityInitializer(providerFactory observability.ProviderFactory) *GlobalObservabilityInitializer {
	return &GlobalObservabilityInitializer{
		providerFactory: providerFactory,
	}
}

// Initialize はグローバルオブザーバビリティを初期化します
func (g *GlobalObservabilityInitializer) Initialize(observabilityConfig config.ObservabilityConfig) error {
	// ロガーの初期化
	if loggerProvider, err := g.providerFactory.CreateLoggerProvider(); err != nil {
		return fmt.Errorf("logger provider creation failed: %w", err)
	} else if err := logger.InitWithProvider(loggerProvider, observabilityConfig); err != nil {
		return fmt.Errorf("global logger initialization failed: %w", err)
	}

	// メトリクスの初期化
	if meterProvider, err := g.providerFactory.CreateMeterProvider(); err != nil {
		return fmt.Errorf("meter provider creation failed: %w", err)
	} else if err := metrics.InitWithProvider(meterProvider); err != nil {
		return fmt.Errorf("global metrics initialization failed: %w", err)
	}

	// トレーサーの初期化
	if tracerProvider, err := g.providerFactory.CreateTracerProvider(); err != nil {
		return fmt.Errorf("tracer provider creation failed: %w", err)
	} else if err := tracer.InitWithProvider(tracerProvider); err != nil {
		return fmt.Errorf("global tracer initialization failed: %w", err)
	}

	return nil
}
