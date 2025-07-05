package reader

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"
)

// ProductDetailReader は商品詳細データの読み取りを担当
type ProductDetailReader struct{}

// NewProductDetailReader は新しいProductDetailReaderを作成
func NewProductDetailReader() *ProductDetailReader {
	return &ProductDetailReader{}
}

// FindProductByID は指定されたIDの商品を詳細情報付きで取得
func (r *ProductDetailReader) FindProductByID(ctx context.Context, id int) (*models.Product, error) {
	db := database.GetDB()

	// 安全なintからint32への変換（共通化した関数を使用）
	productID, err := utils.SafeIntToInt32(id)
	if err != nil {
		return nil, fmt.Errorf("product ID conversion error: %w", err)
	}

	// Bobでクエリを構築
	query := models.Products.Query(
		models.SelectWhere.Products.ID.EQ(productID),
		// カテゴリーを事前読み込み（LEFT JOIN）
		models.Preload.Product.Category(),
		// 在庫情報を後続読み込み（別クエリ）
		models.SelectThenLoad.Product.Inventories(),
	)

	// 商品詳細を取得
	product, err := query.One(ctx, db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found: %d", id)
		}
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	return product, nil
}
