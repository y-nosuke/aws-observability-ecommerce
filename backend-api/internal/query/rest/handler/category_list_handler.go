package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
)

// CategoryListHandler はカテゴリー一覧APIのハンドラー
type CategoryListHandler struct {
	reader *reader.CategoryListReader
	mapper *mapper.CategoryListMapper
}

// NewCategoryListHandler は新しいCategoryListHandlerを作成
func NewCategoryListHandler(db boil.ContextExecutor) *CategoryListHandler {
	return &CategoryListHandler{
		reader: reader.NewCategoryListReader(db),
		mapper: mapper.NewCategoryListMapper(),
	}
}

// ListCategories はカテゴリー一覧を取得する
func (h *CategoryListHandler) ListCategories(ctx echo.Context) error {
	// カテゴリー一覧を取得
	categories, err := h.reader.FindCategoriesWithProductCount(ctx.Request().Context())
	if err != nil {
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch categories", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンスの構築
	response := h.mapper.ToCategoryListResponse(categories)
	return ctx.JSON(http.StatusOK, response)
}
