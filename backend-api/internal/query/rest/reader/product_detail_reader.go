package reader

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
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
	// トレーシングスパンを開始
	// トレーシングスパンを開始
	ctx, span := tracer.Start(ctx, "reader.find_product_by_id", trace.WithAttributes(
		attribute.String("app.layer", "reader"),
		attribute.String("app.domain", "product"),
		attribute.String("app.operation", "find_product_by_id"),
		attribute.Int("app.product_id", id),
	))
	defer span.End()

	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.Where("id = ?", id),
		// カテゴリーと在庫情報を事前読み込み
		qm.Load("Category"),
		qm.Load("Inventories"),
	}

	// 子スパンで商品詳細取得処理
	queryCtx, querySpan := tracer.Start(ctx, "reader.query_product_with_details", trace.WithAttributes(
		attribute.String("app.operation", "query_product_with_details"),
		attribute.Int("app.product_id", id),
		attribute.Bool("app.include_category", true),
		attribute.Bool("app.include_inventories", true),
	))

	product, err := models.Products(mods...).One(queryCtx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			querySpan.SetAttributes(attribute.Bool("app.product_not_found", true))
			querySpan.SetStatus(codes.Error, "product not found")
			querySpan.End()
			span.SetAttributes(attribute.Bool("app.product_not_found", true))
			span.SetStatus(codes.Error, "product not found")
			return nil, fmt.Errorf("product not found: %d", id)
		}
		querySpan.RecordError(err)
		querySpan.SetStatus(codes.Error, err.Error())
		querySpan.End()
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	// 商品詳細情報をスパンに記録
	querySpan.SetAttributes(
		attribute.String("app.product_name", product.Name),
		attribute.String("app.product_sku", product.Sku),
		attribute.Float64("app.product_price", product.Price.InexactFloat64()),
	)

	// カテゴリー情報の記録
	if product.R != nil && product.R.Category != nil {
		querySpan.SetAttributes(
			attribute.Int("app.category_id", product.R.Category.ID),
			attribute.String("app.category_name", product.R.Category.Name),
		)
	}

	// 在庫情報の記録
	if product.R != nil && len(product.R.Inventories) > 0 {
		totalInventory := int64(0)
		for _, inv := range product.R.Inventories {
			totalInventory += int64(inv.Quantity)
		}
		querySpan.SetAttributes(
			attribute.Int("app.inventory_records", len(product.R.Inventories)),
			attribute.Int64("app.total_inventory", totalInventory),
		)
	}

	querySpan.End()

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Bool("app.product_found", true),
		attribute.String("app.product_name", product.Name),
		attribute.String("app.product_sku", product.Sku),
		attribute.Float64("app.product_price", product.Price.InexactFloat64()),
	)

	return product, nil
}
