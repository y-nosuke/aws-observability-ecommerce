package reader

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
	// Repository トレーサーを開始
	repo := observability.StartRepository(ctx, "find_products_with_details")
	defer repo.Finish(false)

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

	repo.LogInfo("Starting product catalog query",
		"page", params.Page,
		"page_size", params.PageSize,
		"offset", offset,
		"category_id", params.CategoryID,
		"keyword", params.Keyword,
	)

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
	err := repo.AddDatabaseStep("count_products", "products", func(stepCtx context.Context) error {
		var countErr error
		total, countErr = models.Products(countMods...).Count(stepCtx, r.db)
		return countErr
	})

	if err != nil {
		repo.FinishWithError(err, "Failed to count products")
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}

	repo.RecordDatabaseOperation("products", "COUNT", int(total))

	// ページネーションを追加
	mods = append(mods, qm.Limit(params.PageSize), qm.Offset(offset))

	// 商品一覧取得
	err = repo.AddDatabaseStep("fetch_products", "products", func(stepCtx context.Context) error {
		var fetchErr error
		products, fetchErr = models.Products(mods...).All(stepCtx, r.db)
		return fetchErr
	})

	if err != nil {
		repo.FinishWithError(err, "Failed to fetch products")
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}

	repo.RecordDatabaseOperation("products", "SELECT", len(products))

	repo.LogInfo("Product catalog query completed successfully",
		"total_count", total,
		"fetched_count", len(products),
		"has_category_filter", params.CategoryID != nil,
		"has_keyword_filter", params.Keyword != nil && *params.Keyword != "",
	)

	repo.FinishWithRecordCount(true, len(products), "total_available", total)
	return products, total, nil
}
