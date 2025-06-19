package reader

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
)

// CategoryListReader はカテゴリー一覧データの読み取りを担当
type CategoryListReader struct {
	db boil.ContextExecutor
}

// NewCategoryListReader は新しいCategoryListReaderを作成
func NewCategoryListReader(db boil.ContextExecutor) *CategoryListReader {
	return &CategoryListReader{
		db: db,
	}
}

// CategoryWithCount はカテゴリーと商品数を保持する構造体
type CategoryWithCount struct {
	Category     *models.Category
	ProductCount int64
}

// FindCategoriesWithProductCount はカテゴリー一覧を商品数付きで取得
func (r *CategoryListReader) FindCategoriesWithProductCount(ctx context.Context) ([]*CategoryWithCount, error) {
	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.OrderBy("name ASC"),
	}

	// カテゴリー一覧を取得
	categories, err := models.Categories(mods...).All(ctx, r.db)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	// 各カテゴリーの商品数を取得
	result := make([]*CategoryWithCount, 0, len(categories))
	for _, category := range categories {
		// 商品数をカウント (otelsqlが自動でトレーシング)
		count, err := models.Products(
			qm.Where("category_id = ?", category.ID),
		).Count(ctx, r.db)
		if err != nil {
			return nil, fmt.Errorf("failed to count products for category %d: %w", category.ID, err)
		}

		result = append(result, &CategoryWithCount{
			Category:     category,
			ProductCount: count,
		})
	}

	return result, nil
}
