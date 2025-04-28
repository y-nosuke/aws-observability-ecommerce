package handlers

import (
	"net/http"
	"strings"

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
	// IDの整合性チェック
	if id <= 0 {
		errorResponse := openapi.ErrorResponse{
			Code:    "invalid_parameter",
			Message: "Invalid product ID",
			Details: &map[string]interface{}{
				"id": "must be a positive integer",
			},
		}
		return ctx.JSON(http.StatusBadRequest, errorResponse)
	}

	// サービスから商品詳細を取得
	product, err := h.productService.GetProductByID(ctx.Request().Context(), int(id))
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if strings.Contains(err.Error(), "product not found") {
			errorResponse := openapi.ErrorResponse{
				Code:    "product_not_found",
				Message: "Product not found",
				Details: &map[string]interface{}{
					"id": id,
				},
			}
			return ctx.JSON(http.StatusNotFound, errorResponse)
		}

		// その他のエラー
		errorResponse := openapi.ErrorResponse{
			Code:    "internal_server_error",
			Message: "Failed to fetch product details",
			Details: &map[string]interface{}{
				"error": err.Error(),
			},
		}
		return ctx.JSON(http.StatusInternalServerError, errorResponse)
	}

	// 在庫状態の取得
	inStock := false
	if product.R != nil && product.R.Inventories != nil && len(product.R.Inventories) > 0 {
		inStock = product.R.Inventories[0].Quantity > 0
	}

	// カテゴリー名の取得
	var categoryName *string
	if product.R != nil && product.R.Category != nil {
		categoryName = &product.R.Category.Name
	}

	// 価格のパース
	price, _ := product.Price.Float64()

	// レスポンスの構築
	response := openapi.Product{
		Id:           int64(product.ID),
		Name:         product.Name,
		Description:  stringPtr(product.Description.String),
		Price:        float32(price),
		ImageUrl:     stringPtr(product.ImageURL.String),
		InStock:      &inStock,
		CategoryId:   int64(product.CategoryID),
		CategoryName: categoryName,
		CreatedAt:    &product.CreatedAt,
		UpdatedAt:    &product.UpdatedAt,
	}

	return ctx.JSON(http.StatusOK, response)
}

// ListProductsByCategory は指定されたカテゴリーの商品一覧を取得する
func (h *ProductHandler) ListProductsByCategory(ctx echo.Context, id int64, params openapi.ListProductsByCategoryParams) error {
	// パラメータの取得
	page := 1
	if params.Page != nil {
		page = *params.Page
	}

	pageSize := 20
	if params.PageSize != nil {
		pageSize = *params.PageSize
	}

	// サービスからカテゴリー別商品一覧を取得
	result, err := h.productService.GetProductsByCategory(ctx.Request().Context(), int(id), page, pageSize)
	if err != nil {
		// エラーの種類に応じて適切なレスポンスを返す
		if strings.Contains(err.Error(), "category not found") {
			errorResponse := openapi.ErrorResponse{
				Code:    "category_not_found",
				Message: "Category not found",
				Details: &map[string]interface{}{
					"id": id,
				},
			}
			return ctx.JSON(http.StatusNotFound, errorResponse)
		}

		// その他のエラー
		errorResponse := openapi.ErrorResponse{
			Code:    "internal_server_error",
			Message: "Failed to fetch products by category",
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
