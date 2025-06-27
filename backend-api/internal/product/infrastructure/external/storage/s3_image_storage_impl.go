package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.opentelemetry.io/otel/attribute"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	configPkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/errors"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/otel"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/utils"
)

// S3ImageStorageImpl はS3ClientWrapperを使用した画像ストレージの実装
type S3ImageStorageImpl struct {
	s3Client *s3.Client
	config   configPkg.S3Config
}

// NewS3ImageStorageImpl は新しいS3ImageStorageImplを作成する
func NewS3ImageStorageImpl(s3Client *s3.Client, config configPkg.S3Config) service.ImageStorage {
	return &S3ImageStorageImpl{
		s3Client: s3Client,
		config:   config,
	}
}

// UploadImage は商品画像をS3にアップロードし、S3キーとURLマップを返却する
func (s *S3ImageStorageImpl) UploadImage(ctx context.Context, productID int, imageData []byte) (key string, urls map[string]string, err error) {
	spanCtx, o := otel.Start(ctx,
		attribute.String("s3.operation", "upload"),
		attribute.String("s3.bucket", s.config.BucketName),
		attribute.String("s3.key", key),
	)
	defer func() {
		o.End(err)
	}()

	// content type 判定（先頭512バイトで判断）
	contentType, ext := utils.Ext(imageData)

	// S3へのアップロード先キーを生成
	key = fmt.Sprintf("uploads/%d/original.%s", productID, ext)

	o.SetAttributes(attribute.String("s3.content_type", contentType))

	input := &s3.PutObjectInput{
		Bucket:      aws.String(s.config.BucketName),
		Key:         aws.String(key),
		Body:        bytes.NewReader(imageData),
		ContentType: aws.String(contentType),
	}

	if _, err = s.s3Client.PutObject(spanCtx, input); err != nil {
		return "", nil, fmt.Errorf("failed to upload object %s: %w", key, err)
	}

	// URLマップを構築
	bucketName := s.config.BucketName
	urls = map[string]string{
		"original":  fmt.Sprintf("http://localhost:4566/%s/%s", bucketName, key),
		"thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/original_thumbnail.%s", bucketName, ext),
		"medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/original_medium.%s", bucketName, ext),
		"large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/original_large.%s", bucketName, ext),
	}

	return key, urls, nil
}

// GetImageData は指定されたサイズの画像データを取得する
func (s *S3ImageStorageImpl) GetImageData(ctx context.Context, productID int, size service.SizeType) (imageData []byte, contentType string, err error) {
	spanCtx, o := otel.Start(ctx,
		attribute.String("s3.operation", "get"),
		attribute.String("s3.bucket", s.config.BucketName),
	)
	defer func() {
		o.End(err)
	}()
	// サイズのバリデーション
	if size != "thumbnail" && size != "medium" && size != "large" && size != "original" {
		return nil, "", errors.Newf("invalid image size: %s", size)
	}
	// S3のキーを構築
	key := fmt.Sprintf("resized/%d/original_%s.png", productID, size)
	o.SetAttributes(attribute.String("s3.key", key))
	input := &s3.GetObjectInput{
		Bucket: aws.String(s.config.BucketName),
		Key:    aws.String(key),
	}
	result, err := s.s3Client.GetObject(spanCtx, input)
	if err != nil {
		return nil, "", fmt.Errorf("failed to get object %s: %w", key, err)
	}
	reader := result.Body
	defer func(reader io.ReadCloser) {
		if closeErr := reader.Close(); closeErr != nil {
			err = errors.Wrapf(closeErr, "original error: %v, failed to close image data", err)
			return
		}
	}(reader)
	// 画像データを読み込む
	if imageData, err = io.ReadAll(reader); err != nil {
		return nil, "", errors.Wrap(err, "failed to read image data")
	}
	// Content-Typeを拡張子から判断
	contentType = utils.ContentType(imageData)
	return imageData, contentType, nil
}
