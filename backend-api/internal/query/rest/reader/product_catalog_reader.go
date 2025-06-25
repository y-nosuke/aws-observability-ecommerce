package reader

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
)

// ProductCatalogReader は商品カタログデータの読み取りを担当
type ProductCatalogReader struct {
	db boil.ContextExecutor
}

// NewProductCatalogReader は新しいProductCatalogReaderを作成
func NewProductCatalogReader(db boil.ContextExecutor) *ProductCatalogReader {
	return &ProductCatalogReader{
		db: db,
	}
}

// ProductListParams は商品一覧取得のパラメータ
type ProductListParams struct {
	Page       int
	PageSize   int
	CategoryID *int
	Keyword    *string
}

// FindProductsWithDetails は商品一覧を詳細情報付きで取得
func (r *ProductCatalogReader) FindProductsWithDetails(ctx context.Context, params *ProductListParams) ([]*models.Product, int64, error) {
	// ページネーション設定
	if params.Page < 1 {
		params.Page = 1
	}
	if params.PageSize < 1 {
		params.PageSize = 20
	}
	if params.PageSize > 100 {
		params.PageSize = 100
	}

	offset := (params.Page - 1) * params.PageSize

	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.OrderBy("created_at DESC"),
		// カテゴリーと在庫情報を事前読み込み
		qm.Load("Category"),
		qm.Load("Inventories"),
	}

	// カテゴリーでフィルタリング
	if params.CategoryID != nil {
		mods = append(mods, qm.Where("category_id = ?", *params.CategoryID))
	}

	// キーワード検索
	if params.Keyword != nil && *params.Keyword != "" {
		mods = append(mods, qm.Where("name LIKE ? OR description LIKE ?", "%"+*params.Keyword+"%", "%"+*params.Keyword+"%"))
	}

	// 総数を取得（WHERE条件のみを適用）
	var countMods []qm.QueryMod
	if params.CategoryID != nil {
		countMods = append(countMods, qm.Where("category_id = ?", *params.CategoryID))
	}
	if params.Keyword != nil && *params.Keyword != "" {
		countMods = append(countMods, qm.Where("name LIKE ? OR description LIKE ?", "%"+*params.Keyword+"%", "%"+*params.Keyword+"%"))
	}

	var total int64
	var products []*models.Product

	// 総数取得
	total, err := models.Products(countMods...).Count(ctx, r.db)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	// ページネーションを追加
	mods = append(mods, qm.Limit(params.PageSize), qm.Offset(offset))

	// 商品一覧取得
	products, err = models.Products(mods...).All(ctx, r.db)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}

	return products, total, nil
}
