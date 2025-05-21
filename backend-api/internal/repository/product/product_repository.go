package product

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository"
)

// Repository は商品リポジトリのインターフェース
type Repository interface {
	repository.Repository
	FindByID(ctx context.Context, id int) (*models.Product, error)
	FindAll(ctx context.Context, limit, offset int, keyword *string) ([]*models.Product, error)
	FindByCategory(ctx context.Context, categoryID int, limit, offset int, keyword *string) ([]*models.Product, error)
	Create(ctx context.Context, product *models.Product) error
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id int) error
	Count(ctx context.Context, keyword *string) (int64, error)
	CountByCategory(ctx context.Context, categoryID int, keyword *string) (int64, error)
}

// ProductRepository は商品リポジトリの実装
type ProductRepository struct {
	repository.RepositoryBase
}

// New は新しい商品リポジトリを作成します
func New(executor boil.ContextExecutor) Repository {
	return &ProductRepository{
		RepositoryBase: repository.NewRepositoryBase(executor),
	}
}

// WithTx はトランザクションを設定した商品リポジトリを返します
func (r *ProductRepository) WithTx(tx *sql.Tx) repository.Repository {
	return &ProductRepository{
		RepositoryBase: r.RepositoryBase.WithTx(tx),
	}
}

// FindByID は指定IDの商品を取得します
func (r *ProductRepository) FindByID(ctx context.Context, id int) (*models.Product, error) {
	return models.FindProduct(ctx, r.DB(), id)
}

// FindAll は商品一覧を取得します（ページネーション対応）
func (r *ProductRepository) FindAll(ctx context.Context, limit, offset int, keyword *string) ([]*models.Product, error) {
	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.Limit(limit),
		qm.Offset(offset),
		qm.OrderBy("created_at DESC"),
		// 商品とカテゴリーを結合して取得
		qm.Load("Category"),
	}

	if keyword != nil && *keyword != "" {
		mods = append(mods, qm.Where("name LIKE ? OR description LIKE ?", "%"+*keyword+"%", "%"+*keyword+"%"))
	}

	// SQLBoilerを使用して商品を取得
	products, err := models.Products(mods...).All(ctx, r.DB())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products: %w", err)
	}

	return products, nil
}

// FindByCategory は指定カテゴリーの商品一覧を取得します（ページネーション対応）
func (r *ProductRepository) FindByCategory(ctx context.Context, categoryID int, limit, offset int, keyword *string) ([]*models.Product, error) {
	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.Where("category_id = ?", categoryID),
		qm.Limit(limit),
		qm.Offset(offset),
		qm.OrderBy("created_at DESC"),
		// 商品とカテゴリーを結合して取得
		qm.Load("Category"),
	}

	if keyword != nil && *keyword != "" {
		mods = append(mods, qm.Where("name LIKE ? OR description LIKE ?", "%"+*keyword+"%", "%"+*keyword+"%"))
	}

	// SQLBoilerを使用して商品を取得
	products, err := models.Products(mods...).All(ctx, r.DB())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch products by category: %w", err)
	}

	return products, nil
}

// Create は新しい商品を作成します
func (r *ProductRepository) Create(ctx context.Context, product *models.Product) error {
	return product.Insert(ctx, r.DB(), boil.Infer())
}

// Update は商品情報を更新します
func (r *ProductRepository) Update(ctx context.Context, product *models.Product) error {
	_, err := product.Update(ctx, r.DB(), boil.Infer())
	return err
}

// Delete は商品を削除します
func (r *ProductRepository) Delete(ctx context.Context, id int) error {
	product, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if product == nil {
		return fmt.Errorf("product not found: %d", id)
	}
	_, err = product.Delete(ctx, r.DB())
	return err
}

// Count は商品の総数を取得します
func (r *ProductRepository) Count(ctx context.Context, keyword *string) (int64, error) {
	mods := []qm.QueryMod{}
	if keyword != nil && *keyword != "" {
		mods = append(mods, qm.Where("name LIKE ? OR description LIKE ?", "%"+*keyword+"%", "%"+*keyword+"%"))
	}
	return models.Products(mods...).Count(ctx, r.DB())
}

// CountByCategory は指定カテゴリーの商品数を取得します
func (r *ProductRepository) CountByCategory(ctx context.Context, categoryID int, keyword *string) (int64, error) {
	mods := []qm.QueryMod{
		qm.Where("category_id = ?", categoryID),
	}
	if keyword != nil && *keyword != "" {
		mods = append(mods, qm.Where("name LIKE ? OR description LIKE ?", "%"+*keyword+"%", "%"+*keyword+"%"))
	}
	return models.Products(mods...).Count(ctx, r.DB())
}
