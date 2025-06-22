package usecase

import (
	"context"
	"fmt"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
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
func (u *GetProductImageUseCase) Execute(ctx context.Context, productID int, size string) (*dto.GetImageResponse, error) {
	// UseCase トレーサーを開始
	tracer := observability.StartUseCase(ctx, "get_product_image")
	defer tracer.Finish(true)

	// 画像データを取得
	var imageData []byte
	var contentType string
	var err error

	err = tracer.AddStep("get_image_data", func(stepCtx context.Context) error {
		imageData, contentType, err = u.imageStorage.GetImageData(stepCtx, productID, size)
		return err
	})

	if err != nil {
		tracer.FinishWithError(err, "画像データの取得に失敗", "requested_size", size)
		return nil, fmt.Errorf("failed to get image data: %w", err)
	}

	tracer.LogInfo("画像データを正常に取得",
		"product_id", productID,
		"content_type", contentType,
		"image_size_bytes", len(imageData),
		"requested_size", size,
	)

	// レスポンスを構築
	return dto.NewGetImageResponse(productID, imageData, contentType), nil
}
