package usecase

import (
	"context"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
)

// UploadProductImageUseCase は商品画像アップロードのユースケース
type UploadProductImageUseCase struct {
	imageStorage service.ImageStorage
}

// NewUploadProductImageUseCase は新しいUploadProductImageUseCaseを作成する
func NewUploadProductImageUseCase(imageStorage service.ImageStorage) *UploadProductImageUseCase {
	return &UploadProductImageUseCase{
		imageStorage: imageStorage,
	}
}

// Execute は商品画像アップロードを実行する
func (u *UploadProductImageUseCase) Execute(ctx context.Context, req *dto.UploadImageRequest) (*dto.UploadImageResponse, error) {
	return otel.WithSpanValue(ctx, func(spanCtx context.Context, o *otel.Observer) (*dto.UploadImageResponse, error) {
		// 画像をアップロード
		uploadResult, err := u.imageStorage.UploadImage(spanCtx, req.ProductID, req.ImageData)
		if err != nil {
			return nil, fmt.Errorf("failed to upload image: %w", err)
		}

		return dto.NewUploadImageResponse(req.ProductID, req.Filename, uploadResult.Key, uploadResult.URLs), nil
	})
}
