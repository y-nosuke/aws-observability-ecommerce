package reader

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	models "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
	// Repository トレーサーを開始
	repo := observability.StartRepository(ctx, "find_categories_with_product_count")
	defer repo.Finish(false)

	repo.LogInfo("Starting category list with product count query")

	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.OrderBy("name ASC"),
	}

	var categories []*models.Category

	// カテゴリー一覧を取得
	err := repo.AddDatabaseStep("fetch_categories", "categories", func(stepCtx context.Context) error {
		var fetchErr error
		categories, fetchErr = models.Categories(mods...).All(stepCtx, r.db)
		return fetchErr
	})

	if err != nil {
		repo.FinishWithError(err, "Failed to fetch categories")
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	repo.RecordDatabaseOperation("categories", "SELECT", len(categories))

	// 各カテゴリーの商品数を取得
	result := make([]*CategoryWithCount, 0, len(categories))

	repo.LogInfo("Starting product count queries for categories",
		"category_count", len(categories),
	)

	for i, category := range categories {
		var count int64

		// 商品数をカウント
		err := repo.AddDatabaseStep(fmt.Sprintf("count_products_category_%d", category.ID), "products", func(stepCtx context.Context) error {
			var countErr error
			count, countErr = models.Products(
				qm.Where("category_id = ?", category.ID),
			).Count(stepCtx, r.db)
			return countErr
		})

		if err != nil {
			repo.FinishWithError(err, "Failed to count products for category",
				"category_id", category.ID,
				"category_name", category.Name,
			)
			return nil, fmt.Errorf("failed to count products for category %d: %w", category.ID, err)
		}

		result = append(result, &CategoryWithCount{
			Category:     category,
			ProductCount: count,
		})

		// 進捗ログ（多数のカテゴリがある場合の可視性向上）
		if (i+1)%10 == 0 || i == len(categories)-1 {
			repo.LogInfo("Category product count progress",
				"completed", i+1,
				"total", len(categories),
			)
		}
	}

	repo.RecordDatabaseOperation("products", "COUNT", len(categories)) // カテゴリ数分のCOUNTクエリ

	repo.LogInfo("Category list with product count query completed successfully",
		"total_categories", len(categories),
		"result_count", len(result),
	)

	repo.FinishWithRecordCount(true, len(result), "categories_processed", len(categories))
	return result, nil
}
