package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/product"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/service"
)

// ProductHandler はoapi-codegenで生成されたサーバーインターフェースを実装する構造体
type ProductHandler struct {
	productService *service.ProductService
}

// NewProductHandler はProductHandlerのインスタンスを生成する
func NewProductHandler() *ProductHandler {
	// リポジトリの初期化
	productRepo := product.New(config.DB)

	// サービスの初期化
	productService := service.NewProductService(config.DB, productRepo)

	return &ProductHandler{
		productService: productService,
	}
}

// ListProducts は商品一覧を取得する
func (h *ProductHandler) ListProducts(ctx echo.Context, params openapi.ListProductsParams) error {
	// パラメータの取得
	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	pageSize := 20
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}

	var result *service.ProductListResult
	var err error

	// カテゴリーIDでフィルタリングするかどうかを判定
	if params.CategoryId != nil {
		categoryID := int(*params.CategoryId)
		result, err = h.productService.GetProductsByCategory(ctx.Request().Context(), categoryID, page, pageSize)
	} else {
		result, err = h.productService.GetProducts(ctx.Request().Context(), page, pageSize)
	}

	if err != nil {
		// エラーハンドリング
		errorResponse := openapi.ErrorResponse{
			Code:    "internal_server_error",
			Message: "Failed to fetch products",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// レスポンスの構築
	items := make([]openapi.Product, 0, len(result.Items))
	for _, p := range result.Items {
		// 在庫状態の取得
		inStock := false
		if p.R != nil && p.R.Inventories != nil && len(p.R.Inventories) > 0 {
			inStock = p.R.Inventories[0].Quantity > 0
		}

		// カテゴリー名の取得
		var categoryName *string
		if p.R != nil && p.R.Category != nil {
			categoryName = &p.R.Category.Name
		}

		price, _ := p.Price.Float64()
		items = append(items, openapi.Product{
			Id:           int64(p.ID),
			Name:         p.Name,
			Description:  stringPtr(p.Description.String),
			Price:        float32(price),
			ImageUrl:     stringPtr(p.ImageURL.String),
			InStock:      &inStock,
			CategoryId:   int64(p.CategoryID),
			CategoryName: categoryName,
			CreatedAt:    &p.CreatedAt,
			UpdatedAt:    &p.UpdatedAt,
		})
	}

	response := openapi.ProductList{
		Items:      items,
		Total:      int(result.Total),
		Page:       result.Page,
		PageSize:   result.PageSize,
		TotalPages: result.TotalPages,
	}

	return ctx.JSON(http.StatusOK, response)
}

// GetProductById は指定されたIDの商品を取得する
func (h *ProductHandler) GetProductById(ctx echo.Context, id int64) error {
	// 実装は次回に行います
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
