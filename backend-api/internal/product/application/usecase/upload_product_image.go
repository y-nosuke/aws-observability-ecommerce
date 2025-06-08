package usecase

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
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
	// ファイル拡張子の検証
	fileExt := strings.ToLower(filepath.Ext(req.Filename))
	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
		return nil, fmt.Errorf("only JPG and PNG images are supported")
	}

	// 画像をアップロード
	s3Key, urls, err := u.imageStorage.UploadImage(ctx, req.ProductID, fileExt, req.ImageData)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %w", err)
	}

	return dto.NewUploadImageResponse(req.ProductID, req.Filename, s3Key, urls), nil
}
