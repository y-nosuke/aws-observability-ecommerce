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
	bucket       string
}

// NewUploadProductImageUseCase は新しいUploadProductImageUseCaseを作成する
func NewUploadProductImageUseCase(imageStorage service.ImageStorage, bucket string) *UploadProductImageUseCase {
	return &UploadProductImageUseCase{
		imageStorage: imageStorage,
		bucket:       bucket,
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
	s3Key, err := u.imageStorage.UploadImage(ctx, req.ProductID, fileExt, req.ImageData)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image: %w", err)
	}

	// ファイル名（拡張子なし）
	fileNameWithoutExt := strings.TrimSuffix(filepath.Base(s3Key), fileExt)

	// レスポンスを構築
	urls := map[string]string{
		"original":  fmt.Sprintf("http://localhost:4566/%s/%s", u.bucket, s3Key),
		"thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/%s_thumbnail%s", u.bucket, fileNameWithoutExt, fileExt),
		"medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/%s_medium%s", u.bucket, fileNameWithoutExt, fileExt),
		"large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/%s_large%s", u.bucket, fileNameWithoutExt, fileExt),
	}

	return dto.NewUploadImageResponse(req.ProductID, req.Filename, s3Key, urls), nil
}
