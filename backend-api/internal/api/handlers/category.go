package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/category"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/service"
)

// CategoryHandler はカテゴリーAPIのハンドラ実装
type CategoryHandler struct {
	categoryService *service.CategoryService
}

// NewCategoryHandler は新しいCategoryHandlerを作成します
func NewCategoryHandler() *CategoryHandler {
	// リポジトリの初期化
	categoryRepo := category.New(config.DB)

	// サービスの初期化
	categoryService := service.NewCategoryService(config.DB, categoryRepo)

	return &CategoryHandler{
		categoryService: categoryService,
	}
}

// ListCategories はカテゴリー一覧を取得します
func (h *CategoryHandler) ListCategories(ctx echo.Context) error {
	// サービスからカテゴリー一覧を取得
	categories, err := h.categoryService.GetCategories(ctx.Request().Context())
	if err != nil {
		// エラーハンドリング
		errorResponse := openapi.ErrorResponse{
			Code:    "internal_server_error",
			Message: "Failed to fetch categories",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンスの構築
	items := make([]openapi.Category, 0, len(categories))
	for _, c := range categories {
		// 商品数を取得
		productCount := int(c.ProductCount)
		// 親カテゴリIDの処理
		var parentId *int64
		if c.Category.ParentID.Valid {
			id := int64(c.Category.ParentID.Int)
			parentId = &id
		}

		items = append(items, openapi.Category{
			Id:           int64(c.Category.ID),
			Name:         c.Category.Name,
			Slug:         c.Category.Slug,
			Description:  stringPtr(c.Category.Description.String),
			ImageUrl:     stringPtr(c.Category.ImageURL.String),
			ParentId:     parentId,
			ProductCount: &productCount,
		})
	}

	response := openapi.CategoryList{
		Items: items,
	}

	return ctx.JSON(http.StatusOK, response)
}
