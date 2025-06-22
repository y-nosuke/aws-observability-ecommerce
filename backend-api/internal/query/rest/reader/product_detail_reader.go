package reader

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
	// Repository トレーサーを開始
	repo := observability.StartRepository(ctx, "find_product_by_id")
	defer repo.Finish(false)

	repo.LogInfo("Starting product detail query",
		"product_id", id,
	)

	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.Where("id = ?", id),
		// カテゴリーと在庫情報を事前読み込み
		qm.Load("Category"),
		qm.Load("Inventories"),
	}

	var product *models.Product

	// 商品詳細を取得
	err := repo.AddDatabaseStep("fetch_product_detail", "products", func(stepCtx context.Context) error {
		var fetchErr error
		product, fetchErr = models.Products(mods...).One(stepCtx, r.db)
		return fetchErr
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			repo.RecordNotFoundError("product", id)
			repo.Finish(true, "result", "not_found") // 見つからないのは正常な結果
			return nil, fmt.Errorf("product not found: %d", id)
		}
		repo.FinishWithError(err, "Failed to fetch product", "product_id", id)
		return nil, fmt.Errorf("failed to fetch product: %w", err)
	}

	repo.RecordDatabaseOperation("products", "SELECT", 1)

	repo.LogInfo("Product detail query completed successfully",
		"product_id", id,
		"product_name", product.Name,
		"has_category", product.R.Category != nil,
		"inventory_count", len(product.R.Inventories),
	)

	repo.FinishWithRecordCount(true, 1, "product_name", product.Name)
	return product, nil
}
