package reader

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
)

// ProductDetailReader は商品詳細データの読み取りを担当
type ProductDetailReader struct {
	db boil.ContextExecutor
}

// NewProductDetailReader は新しいProductDetailReaderを作成
func NewProductDetailReader(db boil.ContextExecutor) *ProductDetailReader {
	return &ProductDetailReader{
		db: db,
	}
}

// FindProductByID は指定されたIDの商品を詳細情報付きで取得
func (r *ProductDetailReader) FindProductByID(ctx context.Context, id int) (*models.Product, error) {
	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.Where("id = ?", id),
		// カテゴリーと在庫情報を事前読み込み
		qm.Load("Category"),
		qm.Load("Inventories"),
	}

	// 商品を取得
	product, err := models.Products(mods...).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("product not found: %d", id)
		}
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	return product, nil
}
