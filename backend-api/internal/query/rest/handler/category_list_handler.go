package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// CategoryListHandler はカテゴリー一覧APIのハンドラー
type CategoryListHandler struct {
	reader *reader.CategoryListReader
	mapper *mapper.CategoryListMapper
}

// NewCategoryListHandler は新しいCategoryListHandlerを作成
func NewCategoryListHandler() *CategoryListHandler {
	return &CategoryListHandler{
		reader: reader.NewCategoryListReader(),
		mapper: mapper.NewCategoryListMapper(),
	}
}

// ListCategories はカテゴリー一覧を取得する
func (h *CategoryListHandler) ListCategories(ctx echo.Context) error {
	return otel.WithSpan(ctx.Request().Context(), func(spanCtx context.Context, o *otel.Observer) error {
		// カテゴリー一覧取得
		categories, err := h.reader.FindCategoriesWithProductCount(spanCtx)
		if err != nil {
			return fmt.Errorf("failed to find categories: %w", err)
		}

		// レスポンス変換
		response := h.mapper.ToCategoryListResponse(categories)

		return ctx.JSON(http.StatusOK, response)
	})
}
