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

// UploadImage は商品画像をS3にアップロードする
func (s *S3ImageStorageImpl) UploadImage(ctx context.Context, productID int64, fileExt string, imageData []byte) (string, error) {
	// S3へのアップロード先キーを生成
	key := fmt.Sprintf("uploads/%d/original%s", productID, fileExt)

	// アップロードオプションを設定
	options := &aws.UploadOptions{
		ContentType: "image/jpeg",
	}

	// ファイル拡張子に基づいてContent-Typeを調整
	switch strings.ToLower(fileExt) {
	case ".png":
		options.ContentType = "image/png"
	case ".gif":
		options.ContentType = "image/gif"
	case ".webp":
		options.ContentType = "image/webp"
	}

	err := s.s3Wrapper.UploadObject(ctx, key, bytes.NewReader(imageData), options)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to S3: %w", err)
	}

	return key, nil
}

// GetImageData は指定されたサイズの画像データを取得する
func (s *S3ImageStorageImpl) GetImageData(ctx context.Context, productID int64, size string) ([]byte, string, error) {
	// サイズのバリデーション
	if size != "thumbnail" && size != "medium" && size != "large" && size != "original" {
		return nil, "", fmt.Errorf("invalid image size: %s", size)
	}

	// S3のキーを構築
	key := fmt.Sprintf("resized/%d/original_%s.jpg", productID, size)

	// S3からオブジェクトを取得
	reader, err := s.s3Wrapper.GetObject(ctx, key)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get object from S3: %w", err)
	}
	defer func(reader io.ReadCloser) {
		if err = reader.Close(); err != nil {
			fmt.Printf("failed to close reader: %v\n", err)
		}
	}(reader)

	// 画像データを読み込む
	imageData, err := io.ReadAll(reader)
	if err != nil {
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

	return imageData, contentType, nil
}
