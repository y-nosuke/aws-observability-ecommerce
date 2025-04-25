package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
)

// ProductHandler はoapi-codegenで生成されたサーバーインターフェースを実装する構造体
type ProductHandler struct {
	// 後でリポジトリなどの依存関係を追加
}

// NewProductHandler はProductHandlerのインスタンスを生成する
func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// ListProducts は商品一覧を取得する
func (h *ProductHandler) ListProducts(ctx echo.Context, params openapi.ListProductsParams) error {
	// 実装はまだダミー
	response := openapi.ProductList{
		Items: []openapi.Product{
			{
				Id:           1,
				Name:         "サンプル商品1",
				Description:  stringPtr("サンプル商品の説明文です。"),
				Price:        1000,
				ImageUrl:     stringPtr("https://example.com/image1.jpg"),
				InStock:      boolPtr(true),
				CategoryId:   1,
				CategoryName: stringPtr("サンプルカテゴリー"),
			},
		},
		Total:      1,
		Page:       1,
		PageSize:   20,
		TotalPages: 1,
	}
	return ctx.JSON(http.StatusOK, response)
}

// GetProductById は指定されたIDの商品を取得する
func (h *ProductHandler) GetProductById(ctx echo.Context, id int64) error {
	// 実装はまだダミー
	if id != 1 {
		errorResponse := openapi.ErrorResponse{
			Code:    "product_not_found",
			Message: "Product not found",
		}
		return ctx.JSON(http.StatusNotFound, errorResponse)
	}

	response := openapi.Product{
		Id:           1,
		Name:         "サンプル商品1",
		Description:  stringPtr("サンプル商品の説明文です。"),
		Price:        1000,
		ImageUrl:     stringPtr("https://example.com/image1.jpg"),
		InStock:      boolPtr(true),
		CategoryId:   1,
		CategoryName: stringPtr("サンプルカテゴリー"),
	}
	return ctx.JSON(http.StatusOK, response)
}

// ListCategories はカテゴリー一覧を取得する
func (h *ProductHandler) ListCategories(ctx echo.Context) error {
	// 実装はまだダミー
	response := openapi.CategoryList{
		Items: []openapi.Category{
			{
				Id:           1,
				Name:         "サンプルカテゴリー",
				Description:  stringPtr("サンプルカテゴリーの説明文です。"),
				ImageUrl:     stringPtr("https://example.com/category1.jpg"),
				ProductCount: intPtr(10),
			},
		},
	}
	return ctx.JSON(http.StatusOK, response)
}

// ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
func (h *ProductHandler) ListProductsByCategory(ctx echo.Context, id int64, params openapi.ListProductsByCategoryParams) error {
	// 実装はまだダミー
	if id != 1 {
		errorResponse := openapi.ErrorResponse{
			Code:    "category_not_found",
			Message: "Category not found",
		}
		return ctx.JSON(http.StatusNotFound, errorResponse)
	}

	response := openapi.ProductList{
		Items: []openapi.Product{
			{
				Id:           1,
				Name:         "サンプル商品1",
				Description:  stringPtr("サンプル商品の説明文です。"),
				Price:        1000,
				ImageUrl:     stringPtr("https://example.com/image1.jpg"),
				InStock:      boolPtr(true),
				CategoryId:   1,
				CategoryName: stringPtr("サンプルカテゴリー"),
			},
		},
		Total:      1,
		Page:       1,
		PageSize:   20,
		TotalPages: 1,
	}
	return ctx.JSON(http.StatusOK, response)
}
