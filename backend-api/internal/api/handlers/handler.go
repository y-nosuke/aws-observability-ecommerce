package handlers

import (
	"fmt"

	"github.com/labstack/echo/v4"

	awsconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/aws"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
)

type Handler struct {
	*HealthHandler
	*ProductHandler
	*CategoryHandler
	*ProductImageHandler
}

func NewHandler(awsConfig *awsconfig.Config) (*Handler, error) {
	productImageHandler, err := NewProductImageHandler(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create ProductImageHandler: %w", err)
	}

	return &Handler{
		HealthHandler:       NewHealthHandler(),
		ProductHandler:      NewProductHandler(),
		CategoryHandler:     NewCategoryHandler(),
		ProductImageHandler: productImageHandler,
	}, nil
}

// RegisterHandlers はEchoルーターにAPIハンドラーを登録する
func RegisterHandlers(g *echo.Group, awsConfig *awsconfig.Config) error {
	handler, err := NewHandler(awsConfig)
	if err != nil {
		return fmt.Errorf("failed to create handler: %w", err)

	}

	// ルーターにハンドラーを登録
	openapi.RegisterHandlers(g, handler)

	return nil
}

// 補助関数
func stringPtr(s string) *string {
	return &s
}
