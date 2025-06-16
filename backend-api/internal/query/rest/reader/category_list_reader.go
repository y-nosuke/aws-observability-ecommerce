package reader

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	models "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
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
	// トレーシングスパンを開始
	ctx, span := tracer.Start(ctx, "reader.find_categories_with_product_count", trace.WithAttributes(
		attribute.String("app.layer", "reader"),
		attribute.String("app.domain", "category"),
		attribute.String("app.operation", "find_categories_with_product_count"),
	))
	defer span.End()

	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.OrderBy("name ASC"),
	}

	// 子スパンでカテゴリー取得処理
	categoriesCtx, categoriesSpan := tracer.Start(ctx, "reader.query_categories", trace.WithAttributes(
		attribute.String("app.operation", "query_categories"),
	))
	categories, err := models.Categories(mods...).All(categoriesCtx, r.db)
	if err != nil {
		categoriesSpan.RecordError(err)
		categoriesSpan.SetStatus(codes.Error, err.Error())
		categoriesSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}
	categoriesSpan.SetAttributes(attribute.Int("app.categories_found", len(categories)))
	categoriesSpan.End()

	// 各カテゴリーの商品数を取得
	result := make([]*CategoryWithCount, 0, len(categories))
	for i, category := range categories {
		// 子スパンで商品数カウント処理（カテゴリー毎）
		countCtx, countSpan := tracer.Start(ctx, "reader.count_products_by_category", trace.WithAttributes(
			attribute.String("app.operation", "count_products_by_category"),
			attribute.Int("app.category_id", category.ID),
			attribute.String("app.category_name", category.Name),
			attribute.Int("app.category_index", i),
		))

		count, err := models.Products(
			qm.Where("category_id = ?", category.ID),
		).Count(countCtx, r.db)

		if err != nil {
			countSpan.RecordError(err)
			countSpan.SetStatus(codes.Error, err.Error())
			countSpan.End()
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			return nil, fmt.Errorf("failed to count products for category %d: %w", category.ID, err)
		}

		countSpan.SetAttributes(attribute.Int64("app.product_count", count))
		countSpan.End()

		result = append(result, &CategoryWithCount{
			Category:     category,
			ProductCount: count,
		})
	}

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("app.categories_processed", len(categories)),
		attribute.Int("app.total_results", len(result)),
	)

	return result, nil
}
