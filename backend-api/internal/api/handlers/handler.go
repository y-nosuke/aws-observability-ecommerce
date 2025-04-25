package handlers

import (
	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
)

type Handler struct {
	*HealthHandler
	*ProductHandler
}

// NewProductHandler はProductHandlerのインスタンスを生成する
func NewHandler() *Handler {
	return &Handler{
		HealthHandler:  NewHealthHandler(),
		ProductHandler: NewProductHandler(),
	}
}

// RegisterHandlers はEchoルーターにAPIハンドラーを登録する
func RegisterHandlers(g *echo.Group) error {
	handler := NewHandler()

	// ルーターにハンドラーを登録
	openapi.RegisterHandlers(g, handler)

	return nil
}

// 補助関数
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func boolPtr(b bool) *bool {
	return &b
}
