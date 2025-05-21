package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/repository/product"
)

// ProductService は商品サービスの実装
type ProductService struct {
	db          *sql.DB
	productRepo product.Repository
}

// NewProductService は新しい商品サービスを作成します
func NewProductService(db *sql.DB, productRepo product.Repository) *ProductService {
	return &ProductService{
		db:          db,
		productRepo: productRepo,
	}
}

// CreateProduct は商品を作成します
func (s *ProductService) CreateProduct(ctx context.Context, p *models.Product) error {
	return s.productRepo.Create(ctx, p)
}

// TransferProductCategory は商品のカテゴリーを変更するトランザクション処理の例です
func (s *ProductService) TransferProductCategory(ctx context.Context, productID int, newCategoryID int) error {
	return repository.RunInTransaction(ctx, s.db, func(tx *sql.Tx) error {
		// トランザクション内でリポジトリを使用
		txProductRepo, ok := s.productRepo.WithTx(tx).(product.Repository)
		if !ok {
			return fmt.Errorf("failed to convert to product.Repository")
		}

		// 商品を取得
		p, err := txProductRepo.FindByID(ctx, productID)
		if err != nil {
			return err
		}

		// カテゴリーを更新
		p.CategoryID = newCategoryID

		// 商品を更新
		return txProductRepo.Update(ctx, p)
	})
}

// ProductListResult は商品一覧の取得結果を表す構造体
type ProductListResult struct {
	Items      []*models.Product
	Total      int64
	Page       int
	PageSize   int
	TotalPages int
}

// GetProducts は商品一覧を取得します（ページネーション対応）
func (s *ProductService) GetProducts(ctx context.Context, page, pageSize int, keyword *string) (*ProductListResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100 // 最大制限
	}

	// オフセットを計算
	offset := (page - 1) * pageSize

	// 商品の総数を取得
	total, err := s.productRepo.Count(ctx, keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to count products: %w", err)
	}

	// 総ページ数を計算
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	// 商品一覧を取得
	products, err := s.productRepo.FindAll(ctx, pageSize, offset, keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return &ProductListResult{
		Items:      products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetProductsByCategory はカテゴリー別の商品一覧を取得します
func (s *ProductService) GetProductsByCategory(ctx context.Context, categoryID int, page, pageSize int, keyword *string) (*ProductListResult, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100 // 最大制限
	}

	// オフセットを計算
	offset := (page - 1) * pageSize

	// カテゴリー別商品の総数を取得
	total, err := s.productRepo.CountByCategory(ctx, categoryID, keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to count products by category: %w", err)
	}

	// 総ページ数を計算
	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	// カテゴリー別商品一覧を取得
	products, err := s.productRepo.FindByCategory(ctx, categoryID, pageSize, offset, keyword)
	if err != nil {
		return nil, fmt.Errorf("failed to get products by category: %w", err)
	}

	return &ProductListResult{
		Items:      products,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}
