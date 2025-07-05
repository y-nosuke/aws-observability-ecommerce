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
	if len(p.R.Inventories) > 0 {
		inStock = p.R.Inventories[0].Quantity > 0
		q := int(p.R.Inventories[0].Quantity)
		quantity = &q
	}

	// カテゴリー名の取得
	var categoryName *string
	if p.R.Category != nil {
		categoryName = &p.R.Category.Name
	}

	// 価格のパース
	price, _ := p.Price.Float64()

	// セール価格は値がある場合のみ設定
	var salePrice *float32
	if p.SalePrice.Valid {
		sp, _ := p.SalePrice.V.Float64()
		spFloat := float32(sp)
		salePrice = &spFloat
	}

	// null.Stringをポインタに変換
	var description *string
	if p.Description.Valid {
		description = &p.Description.V
	}

	var imageURL *string
	if p.ImageURL.Valid {
		imageURL = &p.ImageURL.V
	}

	isNew := p.IsNew
	isFeatured := p.IsFeatured
	sku := p.Sku

	return openapi.Product{
		Id:            int(p.ID),
		Name:          p.Name,
		Description:   description,
		Sku:           &sku,
		Price:         float32(price),
		SalePrice:     salePrice,
		ImageUrl:      imageURL,
		InStock:       inStock,
		StockQuantity: quantity,
		CategoryId:    int(p.CategoryID),
		CategoryName:  categoryName,
		IsNew:         &isNew,
		IsFeatured:    &isFeatured,
		CreatedAt:     &p.CreatedAt,
		UpdatedAt:     &p.UpdatedAt,
	}
}
