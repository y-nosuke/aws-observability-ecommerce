package usecase

import (
	"context"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// GetProductImageUseCase は商品画像取得のユースケース
type GetProductImageUseCase struct {
	imageStorage service.ImageStorage
}

// NewGetProductImageUseCase は新しいGetProductImageUseCaseを作成する
func NewGetProductImageUseCase(
	imageStorage service.ImageStorage,
) *GetProductImageUseCase {
	return &GetProductImageUseCase{
		imageStorage: imageStorage,
	}
}

// Execute は商品画像取得を実行する
func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int, size service.SizeType) (*dto.GetImageResponse, error) {
	return otel.WithSpanValue(ctx, func(spanCtx context.Context, o *otel.Observer) (*dto.GetImageResponse, error) {
		// 画像データを取得
		getImageResult, err := u.imageStorage.GetImageData(spanCtx, productID, size)
		if err != nil {
			return nil, fmt.Errorf("failed to get image data: %w", err)
		}

		// レスポンスを構築
		return dto.NewGetImageResponse(productID, getImageResult.ImageData, getImageResult.ContentType), nil
	})
}
