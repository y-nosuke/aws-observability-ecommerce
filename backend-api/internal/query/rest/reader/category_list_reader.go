package reader

import (
	"context"
	"fmt"

	"github.com/stephenafamo/bob/dialect/mysql/sm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// CategoryListReader はカテゴリー一覧データの読み取りを担当
type CategoryListReader struct{}

// NewCategoryListReader は新しいCategoryListReaderを作成
func NewCategoryListReader() *CategoryListReader {
	return &CategoryListReader{}
}

// CategoryWithCount はカテゴリーと商品数を保持する構造体
type CategoryWithCount struct {
	Category     *models.Category
	ProductCount int64
}

// FindCategoriesWithProductCount はカテゴリー一覧を商品数付きで取得
func (r *CategoryListReader) FindCategoriesWithProductCount(ctx context.Context) ([]*CategoryWithCount, error) {
	return otel.WithSpanValue(ctx, func(spanCtx context.Context, o *otel.Observer) ([]*CategoryWithCount, error) {
		db := database.GetDB()

		// カテゴリー一覧を取得
		query := models.Categories.Query(
			sm.OrderBy(models.CategoryColumns.Name).Asc(),
		)
		categories, err := query.All(spanCtx, db)

		if err != nil {
			return nil, fmt.Errorf("failed to fetch categories: %w", err)
		}

		// 各カテゴリーの商品数を取得
		result := make([]*CategoryWithCount, 0, len(categories))

		for _, category := range categories {
			var count int64

			// 商品数をカウント
			countQuery := models.Products.Query(
				models.SelectWhere.Products.CategoryID.EQ(category.ID),
			)
			count, err := countQuery.Count(spanCtx, db)
			if err != nil {
				return nil, fmt.Errorf("failed to count products for category %d: %w", category.ID, err)
			}

			result = append(result, &CategoryWithCount{
				Category:     category,
				ProductCount: count,
			})
		}

		return result, nil
	})
}
