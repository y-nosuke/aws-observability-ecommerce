package storage

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

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

// UploadImage は商品画像をS3にアップロードし、S3キーとURLマップを返却する
func (s *S3ImageStorageImpl) UploadImage(ctx context.Context, productID int64, fileExt string, imageData []byte) (string, map[string]string, error) {
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	ctx, span := tracer.Start(ctx, "s3.upload_image", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	// S3へのアップロード先キーを生成
	key := fmt.Sprintf("uploads/%d/original%s", productID, fileExt)
	bucketName := s.s3Wrapper.GetBucketName()

	// スパンに属性を設定
	span.SetAttributes(
		attribute.String("rpc.service", "aws.s3"),
		attribute.String("aws.s3.bucket", bucketName),
		attribute.String("aws.s3.key", key),
		attribute.Int("aws.s3.object_size", len(imageData)),
		attribute.String("aws.s3.operation", "PutObject"),
		attribute.Int64("app.product_id", productID),
		attribute.String("app.file_extension", fileExt),
	)

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

	// Content-Typeをスパンに追加
	span.SetAttributes(attribute.String("http.request.content_type", options.ContentType))

	err := s.s3Wrapper.UploadObject(ctx, key, bytes.NewReader(imageData), options)
	if err != nil {
		// スパンにエラー情報を記録
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "s3_upload_error"),
		)
		return "", nil, fmt.Errorf("failed to upload image to S3: %w", err)
	}

	// URLマップを構築（S3ClientWrapperからbucket名を取得）
	urls := s.buildImageURLs(key, fileExt)

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("aws.s3.generated_urls", len(urls)),
		attribute.Bool("aws.s3.success", true),
	)

	return key, urls, nil
}

// buildImageURLs は画像のURLマップを構築する
func (s *S3ImageStorageImpl) buildImageURLs(s3Key, fileExt string) map[string]string {
	// S3ClientWrapperのconfigからbucket名を取得
	bucketName := s.s3Wrapper.GetBucketName()

	// ファイル名（拡張子なし）
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
	// トレーシングスパンを開始
	tracer := otel.Tracer("aws-observability-ecommerce")
	ctx, span := tracer.Start(ctx, "s3.get_image", trace.WithSpanKind(trace.SpanKindClient))
	defer span.End()

	// サイズのバリデーション
	if size != "thumbnail" && size != "medium" && size != "large" && size != "original" {
		err := fmt.Errorf("invalid image size: %s", size)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "validation_error"),
			attribute.String("error.detail", "invalid_image_size"),
		)
		return nil, "", err
	}

	// S3のキーを構築
	key := fmt.Sprintf("resized/%d/original_%s.jpg", productID, size)
	bucketName := s.s3Wrapper.GetBucketName()

	// スパンに属性を設定
	span.SetAttributes(
		attribute.String("rpc.service", "aws.s3"),
		attribute.String("aws.s3.bucket", bucketName),
		attribute.String("aws.s3.key", key),
		attribute.String("aws.s3.operation", "GetObject"),
		attribute.Int64("app.product_id", productID),
		attribute.String("app.image_size", size),
	)

	// S3からオブジェクトを取得
	reader, err := s.s3Wrapper.GetObject(ctx, key)
	if err != nil {
		// スパンにエラー情報を記録
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "s3_get_error"),
		)
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
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "io_read_error"),
		)
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

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.Int("aws.s3.object_size", len(imageData)),
		attribute.String("http.response.content_type", contentType),
		attribute.Bool("aws.s3.success", true),
	)

	return imageData, contentType, nil
}
