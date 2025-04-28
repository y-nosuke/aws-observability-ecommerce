package category

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository"
)

// Repository はカテゴリーリポジトリのインターフェース
type Repository interface {
	repository.Repository
	FindByID(ctx context.Context, id int) (*models.Category, error)
	FindAll(ctx context.Context) ([]*models.Category, error)
	Create(ctx context.Context, category *models.Category) error
	Update(ctx context.Context, category *models.Category) error
	Delete(ctx context.Context, id int) error
	GetProductCount(ctx context.Context, categoryID int) (int64, error)
}

// CategoryRepository はカテゴリーリポジトリの実装
type CategoryRepository struct {
	repository.RepositoryBase
}

// New は新しいカテゴリーリポジトリを作成します
func New(executor boil.ContextExecutor) Repository {
	return &CategoryRepository{
		RepositoryBase: repository.NewRepositoryBase(executor),
	}
}

// WithTx はトランザクションを設定したカテゴリーリポジトリを返します
func (r *CategoryRepository) WithTx(tx *sql.Tx) repository.Repository {
	return &CategoryRepository{
		RepositoryBase: r.RepositoryBase.WithTx(tx),
	}
}

// FindByID は指定IDのカテゴリーを取得します
func (r *CategoryRepository) FindByID(ctx context.Context, id int) (*models.Category, error) {
	return models.FindCategory(ctx, r.DB(), id)
}

// FindAll はカテゴリー一覧を取得します
func (r *CategoryRepository) FindAll(ctx context.Context) ([]*models.Category, error) {
	// クエリモディファイアの準備
	mods := []qm.QueryMod{
		qm.OrderBy("name ASC"),
	}

	// SQLBoilerを使用してカテゴリーを取得
	categories, err := models.Categories(mods...).All(ctx, r.DB())
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	return categories, nil
}

// Create は新しいカテゴリーを作成します
func (r *CategoryRepository) Create(ctx context.Context, category *models.Category) error {
	return category.Insert(ctx, r.DB(), boil.Infer())
}

// Update はカテゴリー情報を更新します
func (r *CategoryRepository) Update(ctx context.Context, category *models.Category) error {
	_, err := category.Update(ctx, r.DB(), boil.Infer())
	return err
}

// Delete はカテゴリーを削除します
func (r *CategoryRepository) Delete(ctx context.Context, id int) error {
	category, err := r.FindByID(ctx, id)
	if err != nil {
		return err
	}
	if category == nil {
		return fmt.Errorf("category not found: %d", id)
	}
	_, err = category.Delete(ctx, r.DB())
	return err
}

// GetProductCount は指定カテゴリーの商品数を取得します
func (r *CategoryRepository) GetProductCount(ctx context.Context, categoryID int) (int64, error) {
	mods := []qm.QueryMod{
		qm.Where("category_id = ?", categoryID),
	}
	return models.Products(mods...).Count(ctx, r.DB())
}
