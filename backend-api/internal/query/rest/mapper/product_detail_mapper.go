package mapper

import (
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/presentation/rest/openapi"
)

// ProductDetailMapper は商品詳細データのマッピングを担当
type ProductDetailMapper struct{}

// NewProductDetailMapper は新しいProductDetailMapperを作成
func NewProductDetailMapper() *ProductDetailMapper {
	return &ProductDetailMapper{}
}

// ToProductResponse はモデルからOpenAPIレスポンスへ変換
func (m *ProductDetailMapper) ToProductResponse(p *models.Product) openapi.Product {
	// 在庫状態の取得
	inStock := false
	var quantity *int
	if p.R != nil && p.R.Inventories != nil && len(p.R.Inventories) > 0 {
		inStock = p.R.Inventories[0].Quantity > 0
		q := p.R.Inventories[0].Quantity
		quantity = &q
	}

	// カテゴリー名の取得
	var categoryName *string
	if p.R != nil && p.R.Category != nil {
		categoryName = &p.R.Category.Name
	}

	// 価格のパース
	price, _ := p.Price.Float64()

	// セール価格は値がある場合のみ設定
	var salePrice *float32
	if !p.SalePrice.IsZero() {
		sp, _ := p.SalePrice.Float64()
		spFloat := float32(sp)
		salePrice = &spFloat
	}

	// null.Stringをポインタに変換
	var description *string
	if p.Description.Valid {
		description = &p.Description.String
	}

	var imageURL *string
	if p.ImageURL.Valid {
		imageURL = &p.ImageURL.String
	}

	isNew := p.IsNew
	isFeatured := p.IsFeatured
	sku := p.Sku

	return openapi.Product{
		Id:            int64(p.ID),
		Name:          p.Name,
		Description:   description,
		Sku:           &sku,
		Price:         float32(price),
		SalePrice:     salePrice,
		ImageUrl:      imageURL,
		InStock:       inStock,
		StockQuantity: quantity,
		CategoryId:    int64(p.CategoryID),
		CategoryName:  categoryName,
		IsNew:         &isNew,
		IsFeatured:    &isFeatured,
		CreatedAt:     &p.CreatedAt,
		UpdatedAt:     &p.UpdatedAt,
	}
}

// PresentError はエラーレスポンスを生成
func (m *ProductDetailMapper) PresentError(code, message string, details any) openapi.ErrorResponse {
	detailsMap := map[string]any{
		"error": details,
	}
	if dm, ok := details.(map[string]any); ok {
		detailsMap = dm
	}

	return openapi.ErrorResponse{
		Code:    code,
		Message: message,
		Details: &detailsMap,
	}
}

// PresentInvalidParameter は無効なパラメーターレスポンスを生成
func (m *ProductDetailMapper) PresentInvalidParameter(message string) openapi.ErrorResponse {
	return m.PresentError("invalid_parameter", message, map[string]any{
		"id": "must be a positive integer",
	})
}

// PresentProductNotFound は商品が見つからない場合のレスポンスを生成
func (m *ProductDetailMapper) PresentProductNotFound(message string, id int) openapi.ErrorResponse {
	return m.PresentError("product_not_found", message, map[string]any{
		"id": id,
	})
}

// PresentInternalServerError は内部サーバーエラーレスポンスを生成
func (m *ProductDetailMapper) PresentInternalServerError(message string, err error) openapi.ErrorResponse {
	return m.PresentError("internal_server_error", message, err.Error())
}
