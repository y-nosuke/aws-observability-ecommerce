package service

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/db/models"
)

// GetProductByID は指定されたIDの商品を取得します
func (s *ProductService) GetProductByID(ctx context.Context, id int) (*models.Product, error) {
	// リポジトリから商品を取得
	p, err := s.productRepo.FindByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product not found: %d", id)
		}
		return nil, fmt.Errorf("failed to get product: %w", err)
	}

	// Eager Loadingで関連データを取得
	if p.R == nil || p.R.Category == nil {
		// カテゴリーが読み込まれていない場合はカテゴリーも取得
		if err := p.L.LoadCategory(ctx, s.db, true, p, nil); err != nil {
			return nil, fmt.Errorf("failed to load category: %w", err)
		}
	}

	if p.R == nil || p.R.Inventories == nil {
		// 在庫情報が読み込まれていない場合は在庫も取得
		if err := p.L.LoadInventories(ctx, s.db, true, p, nil); err != nil {
			return nil, fmt.Errorf("failed to load inventory: %w", err)
		}
	}

	return p, nil
}
