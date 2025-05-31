package usecase

import (
	"context"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
)

// GetProductImageUseCase は商品画像取得のユースケース
type GetProductImageUseCase struct {
	imageStorage service.ImageStorage
	bucket       string
}

// NewGetProductImageUseCase は新しいGetProductImageUseCaseを作成する
func NewGetProductImageUseCase(imageStorage service.ImageStorage, bucket string) *GetProductImageUseCase {
	return &GetProductImageUseCase{
		imageStorage: imageStorage,
		bucket:       bucket,
	}
}

// Execute は商品画像取得を実行する
func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int64, size string) (*dto.GetImageResponse, error) {
	// 画像データを取得
	imageData, contentType, err := u.imageStorage.GetImageData(ctx, productID, size)
	if err != nil {
		return nil, fmt.Errorf("failed to get image data: %w", err)
	}

	// レスポンスを構築
	return dto.NewGetImageResponse(productID, imageData, contentType), nil
}
