package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/category"
)

// CategoryWithCount はカテゴリー情報と商品数を含む構造体
type CategoryWithCount struct {
	Category     *models.Category
	ProductCount int64
}

// CategoryService はカテゴリーサービスの実装
type CategoryService struct {
	db           *sql.DB
	categoryRepo category.Repository
}

// NewCategoryService は新しいカテゴリーサービスを作成します
func NewCategoryService(db *sql.DB, categoryRepo category.Repository) *CategoryService {
	return &CategoryService{
		db:           db,
		categoryRepo: categoryRepo,
	}
}

// GetCategories はカテゴリー一覧を取得します
func (s *CategoryService) GetCategories(ctx context.Context) ([]*CategoryWithCount, error) {
	// カテゴリー一覧を取得
	categories, err := s.categoryRepo.FindAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories: %w", err)
	}

	// 各カテゴリーの商品数を取得して拡張構造体に変換
	result := make([]*CategoryWithCount, 0, len(categories))
	for _, cat := range categories {
		count, err := s.categoryRepo.GetProductCount(ctx, cat.ID)
		if err != nil {
			// エラーが発生しても続行（カウントは0とする）
			count = 0
		}

		// 拡張構造体に変換
		result = append(result, &CategoryWithCount{
			Category:     cat,
			ProductCount: count,
		})
	}

	return result, nil
}

// GetCategoryByID は指定されたIDのカテゴリーを取得します
func (s *CategoryService) GetCategoryByID(ctx context.Context, id int) (*CategoryWithCount, error) {
	// カテゴリーを取得
	category, err := s.categoryRepo.FindByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("category not found: %d", id)
		}
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	// 商品数を取得
	count, err := s.categoryRepo.GetProductCount(ctx, category.ID)
	if err != nil {
		// エラーが発生した場合はカウントを0とする
		count = 0
	}

	// 拡張構造体を返す
	return &CategoryWithCount{
		Category:     category,
		ProductCount: count,
	}, nil
}
