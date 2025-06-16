package reader

import (
	"context"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"

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
	// トレーシングスパンを開始
	// トレーシングスパンを開始
	ctx, span := tracer.Start(ctx, "reader.find_products_with_details", trace.WithAttributes(
		attribute.String("app.layer", "reader"),
		attribute.String("app.domain", "product_catalog"),
		attribute.String("app.operation", "find_products_with_details"),
		attribute.Int("app.page", params.Page),
		attribute.Int("app.page_size", params.PageSize),
	))
	defer span.End()

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

	// フィルタ条件をスパンに記録
	if params.CategoryID != nil {
		span.SetAttributes(attribute.Int("app.filter.category_id", *params.CategoryID))
	}
	if params.Keyword != nil && *params.Keyword != "" {
		span.SetAttributes(attribute.String("app.filter.keyword", *params.Keyword))
	}

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

	// 子スパンでカウント処理
	countCtx, countSpan := tracer.Start(ctx, "reader.count_products", trace.WithAttributes(
		attribute.String("app.operation", "count_products"),
	))
	total, err := models.Products(countMods...).Count(countCtx, r.db)
	if err != nil {
		countSpan.RecordError(err)
		countSpan.SetStatus(codes.Error, err.Error())
		countSpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, 0, fmt.Errorf("failed to count products: %w", err)
	}
	countSpan.SetAttributes(attribute.Int64("app.total_count", total))
	countSpan.End()

	// ページネーションを追加
	mods = append(mods, qm.Limit(params.PageSize), qm.Offset(offset))

	// 子スパンで商品取得処理
	queryCtx, querySpan := tracer.Start(ctx, "reader.query_products", trace.WithAttributes(
		attribute.String("app.operation", "query_products"),
		attribute.Int("app.limit", params.PageSize),
		attribute.Int("app.offset", offset),
	))
	products, err := models.Products(mods...).All(queryCtx, r.db)
	if err != nil {
		querySpan.RecordError(err)
		querySpan.SetStatus(codes.Error, err.Error())
		querySpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, 0, fmt.Errorf("failed to fetch products: %w", err)
	}
	querySpan.SetAttributes(attribute.Int("app.products_found", len(products)))
	querySpan.End()

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("app.products_found", len(products)),
		attribute.Int64("app.total_count", total),
		attribute.Int("app.calculated_offset", offset),
		attribute.Bool("app.has_more_pages", int64(offset+params.PageSize) < total),
	)

	return products, total, nil
}
