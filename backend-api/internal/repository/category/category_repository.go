package category

import (
	"context"
	"database/sql"

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
	mods := []qm.QueryMod{
		qm.OrderBy("name"),
	}
	return models.Categories(mods...).All(ctx, r.DB())
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
	_, err = category.Delete(ctx, r.DB())
	return err
}
