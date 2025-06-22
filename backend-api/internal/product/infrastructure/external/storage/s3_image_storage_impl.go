package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/observability"
)

// S3ImageStorageImpl はS3ClientWrapperを使用した画像ストレージの実装
type S3ImageStorageImpl struct {
	s3Wrapper *aws.S3ClientWrapper
}

// NewS3ImageStorageImpl は新しいS3ImageStorageImplを作成する
func NewS3ImageStorageImpl(s3Wrapper *aws.S3ClientWrapper) service.ImageStorage {
	return &S3ImageStorageImpl{
		s3Wrapper: s3Wrapper,
	}
}

// UploadImage は商品画像をS3にアップロードし、S3キーとURLマップを返却する
func (s *S3ImageStorageImpl) UploadImage(ctx context.Context, productID int64, fileExt string, imageData []byte) (string, map[string]string, error) {
	// Repository トレーサーを開始
	repo := observability.StartRepository(ctx, "upload_image")
	defer repo.Finish(false)

	// S3へのアップロード先キーを生成
	key := fmt.Sprintf("uploads/%d/original%s", productID, fileExt)

	// ファイル拡張子に基づいてContent-Typeを設定
	contentType := "image/jpeg"
	switch strings.ToLower(fileExt) {
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".webp":
		contentType = "image/webp"
	}

	repo.LogInfo("Starting S3 image upload",
		"product_id", productID,
		"s3_key", key,
		"content_type", contentType,
		"file_size_bytes", len(imageData),
	)

	// アップロードオプションを設定
	options := &aws.UploadOptions{
		ContentType: contentType,
	}

	err := repo.AddDatabaseStep("s3_upload", "s3_objects", func(stepCtx context.Context) error {
		return s.s3Wrapper.UploadObject(stepCtx, key, bytes.NewReader(imageData), options)
	})

	if err != nil {
		repo.FinishWithError(err, "Failed to upload image to S3", "s3_key", key)
		return "", nil, fmt.Errorf("failed to upload image to S3: %w", err)
	}

	// URLマップを構築
	urls := s.buildImageURLs(key, fileExt)

	repo.LogInfo("S3 image upload completed successfully",
		"product_id", productID,
		"s3_key", key,
		"generated_urls", len(urls),
	)

	repo.Finish(true, "s3_key", key, "urls_generated", len(urls))
	return key, urls, nil
}

// buildImageURLs は画像のURLマップを構築する
func (s *S3ImageStorageImpl) buildImageURLs(s3Key, fileExt string) map[string]string {
	bucketName := s.s3Wrapper.GetBucketName()
	fileNameWithoutExt := strings.TrimSuffix(filepath.Base(s3Key), fileExt)

	// LocalStack環境のURL構築
	return map[string]string{
		"original":  fmt.Sprintf("http://localhost:4566/%s/%s", bucketName, s3Key),
		"thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/%s_thumbnail%s", bucketName, fileNameWithoutExt, fileExt),
		"medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/%s_medium%s", bucketName, fileNameWithoutExt, fileExt),
		"large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/%s_large%s", bucketName, fileNameWithoutExt, fileExt),
	}
}

// GetImageData は指定されたサイズの画像データを取得する
func (s *S3ImageStorageImpl) GetImageData(ctx context.Context, productID int64, size string) ([]byte, string, error) {
	// Repository トレーサーを開始
	repo := observability.StartRepository(ctx, "get_image_data")
	defer repo.Finish(false)

	// サイズのバリデーション
	if size != "thumbnail" && size != "medium" && size != "large" && size != "original" {
		err := fmt.Errorf("invalid image size: %s", size)
		repo.FinishWithError(err, "Invalid image size provided", "size", size)
		return nil, "", err
	}

	// S3のキーを構築
	key := fmt.Sprintf("resized/%d/original_%s.jpg", productID, size)

	repo.LogInfo("Starting S3 image retrieval",
		"product_id", productID,
		"requested_size", size,
		"s3_key", key,
	)

	var imageData []byte
	var reader io.ReadCloser

	err := repo.AddDatabaseStep("s3_get_object", "s3_objects", func(stepCtx context.Context) error {
		var getErr error
		reader, getErr = s.s3Wrapper.GetObject(stepCtx, key)
		return getErr
	})

	if err != nil {
		repo.FinishWithError(err, "Failed to get object from S3", "s3_key", key)
		return nil, "", fmt.Errorf("failed to get object from S3: %w", err)
	}
	defer reader.Close()

	// 画像データを読み込む
	imageData, err = io.ReadAll(reader)
	if err != nil {
		repo.FinishWithError(err, "Failed to read image data", "s3_key", key)
		return nil, "", fmt.Errorf("failed to read image data: %w", err)
	}

	// Content-Typeを拡張子から判断
	contentType := "image/jpeg" // デフォルト
	ext := strings.ToLower(filepath.Ext(key))
	switch ext {
	case ".png":
		contentType = "image/png"
	case ".gif":
		contentType = "image/gif"
	case ".webp":
		contentType = "image/webp"
	}

	repo.LogInfo("S3 image retrieval completed successfully",
		"product_id", productID,
		"s3_key", key,
		"content_type", contentType,
		"image_size_bytes", len(imageData),
	)

	repo.FinishWithRecordCount(true, 1, "content_type", contentType, "image_size", len(imageData))
	return imageData, contentType, nil
}
