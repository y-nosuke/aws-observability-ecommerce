package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/mapper"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/query/rest/reader"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
	// Handler トレーサーを開始
	handler := observability.StartHandler(
		ctx.Request().Context(),
		"list_categories",
		ctx.Request().Method,
		ctx.Request().URL.Path,
		http.StatusOK,
		ctx.Request().UserAgent(),
		ctx.RealIP(),
		ctx.Request().ContentLength,
	)
	defer handler.FinishWithHTTPStatus(http.StatusOK)

	handler.LogInfo("Category list request received")

	// カテゴリー一覧取得
	categories, err := h.reader.FindCategoriesWithProductCount(handler.Context())
	if err != nil {
		handler.FinishWithError(err, "Failed to fetch categories", http.StatusInternalServerError)
		errorResponse := h.mapper.PresentInternalServerError("Failed to fetch categories", err)
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	handler.LogInfo("Categories fetched successfully",
		"category_count", len(categories),
	)

	// レスポンス変換
	response := h.mapper.ToCategoryListResponse(categories)

	return ctx.JSON(http.StatusOK, response)
}
