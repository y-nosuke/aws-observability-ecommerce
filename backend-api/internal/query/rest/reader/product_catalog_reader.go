package reader

import (
	"context"
	"fmt"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/mysql"
	"github.com/stephenafamo/bob/dialect/mysql/dialect"
	"github.com/stephenafamo/bob/dialect/mysql/sm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/database"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"
)

// ProductCatalogReader は商品カタログデータの読み取りを担当
type ProductCatalogReader struct{}

// NewProductCatalogReader は新しいProductCatalogReaderを作成
func NewProductCatalogReader() *ProductCatalogReader {
	return &ProductCatalogReader{}
}

// ProductListParams は商品一覧取得のパラメータ
type ProductListParams struct {
	Page       int
	PageSize   int
	CategoryID *int
	Keyword    *string
}

// ProductsWithTotal は商品一覧と総数を格納する構造体
type ProductsWithTotal struct {
	Products []*models.Product
	Total    int64
}

// FindProductsWithDetails は商品一覧を詳細情報付きで取得
func (r *ProductCatalogReader) FindProductsWithDetails(ctx context.Context, params *ProductListParams) (*ProductsWithTotal, error) {
	return otel.WithSpanValue(ctx, func(spanCtx context.Context, o *otel.Observer) (*ProductsWithTotal, error) {
		db := database.GetDB()

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

		// Bobでクエリを構築（WHERE条件の準備）
		var whereMods []bob.Mod[*dialect.SelectQuery]

		// カテゴリーでフィルタリング
		if params.CategoryID != nil {
			// 安全なintからint32への変換（共通化した関数を使用）
			categoryID, err := utils.SafeIntToInt32(*params.CategoryID)
			if err != nil {
				return nil, fmt.Errorf("category ID conversion error: %w", err)
			}
			whereMods = append(whereMods, models.SelectWhere.Products.CategoryID.EQ(categoryID))
		}

		// キーワード検索
		if params.Keyword != nil && *params.Keyword != "" {
			keyword := "%" + *params.Keyword + "%"
			whereMods = append(whereMods, mysql.WhereOr(
				models.SelectWhere.Products.Name.Like(keyword),
				models.SelectWhere.Products.Description.Like(keyword),
			))
		}

		// 総数取得用のクエリ（WHERE条件のみ）
		countQuery := models.Products.Query(whereMods...)
		total, err := countQuery.Count(spanCtx, db)
		if err != nil {
			return nil, fmt.Errorf("failed to count products: %w", err)
		}

		// 商品一覧取得用のクエリ（WHERE条件 + ORDER BY + LIMIT + OFFSET + プリロード）
		queryMods := make([]bob.Mod[*dialect.SelectQuery], 0, len(whereMods)+5)
		queryMods = append(queryMods, whereMods...)
		queryMods = append(queryMods,
			sm.OrderBy(models.ProductColumns.CreatedAt).Desc(),
			// カテゴリーを事前読み込み（LEFT JOIN）
			models.Preload.Product.Category(),
			// 在庫情報を後続読み込み（別クエリ）
			models.SelectThenLoad.Product.Inventories(),
			sm.Limit(int64(params.PageSize)),
			sm.Offset(int64(offset)),
		)

		// 商品一覧取得
		productsQuery := models.Products.Query(queryMods...)
		products, err := productsQuery.All(spanCtx, db)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch products: %w", err)
		}

		return &ProductsWithTotal{Products: products, Total: total}, nil
	})
}
