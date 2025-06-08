package usecase

import (
	"context"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logging"
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
func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int64, size string) (*dto.GetImageResponse, error) {
	completeOp := logging.StartOperation(ctx, "get_product_image",
		"product_id", productID,
		"requested_size", size,
		"layer", "usecase")

	// 画像データを取得
	imageData, contentType, err := u.imageStorage.GetImageData(ctx, productID, size)
	if err != nil {
		logging.WithError(ctx, "画像データの取得に失敗", err,
			"product_id", productID,
			"requested_size", size,
			"layer", "usecase",
			"storage_operation", "get_image_data")

		// 操作失敗を記録
		completeOp(false, "error_type", "storage_failure")
		return nil, fmt.Errorf("failed to get image data: %w", err)
	}

	logging.Info(ctx, "画像データを正常に取得",
		"product_id", productID,
		"content_type", contentType,
		"image_size_bytes", len(imageData),
		"layer", "usecase")

	// 操作成功を記録
	completeOp(true,
		"content_type", contentType,
		"image_size_bytes", len(imageData))

	// レスポンスを構築
	return dto.NewGetImageResponse(productID, imageData, contentType), nil
}
