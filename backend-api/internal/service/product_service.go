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
