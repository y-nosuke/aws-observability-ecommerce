package aws

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	configPkg "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/shared/infrastructure/config"
)

// S3ClientWrapper はS3クライアントの薄いラッパー
type S3ClientWrapper struct {
	client *s3.Client
	config configPkg.S3Config
}

// NewS3ClientWrapper は新しいS3クライアントラッパーを作成します
func NewS3ClientWrapper(client *s3.Client, config configPkg.S3Config) *S3ClientWrapper {
	return &S3ClientWrapper{
		client: client,
		config: config,
	}
}

// GetBucketName は設定されているバケット名を返します
func (w *S3ClientWrapper) GetBucketName() string {
	return w.config.BucketName
}

// UploadOptions はアップロードオプション
type UploadOptions struct {
	ContentType     string
	CacheControl    string
	ContentEncoding string
	Metadata        map[string]string
}

// UploadObject はオブジェクトをS3にアップロードします
func (w *S3ClientWrapper) UploadObject(ctx context.Context, key string, body io.Reader, options *UploadOptions) error {
	input := &s3.PutObjectInput{
		Bucket: aws.String(w.config.BucketName),
		Key:    aws.String(key),
		Body:   body,
	}

	if options != nil {
		if options.ContentType != "" {
			input.ContentType = aws.String(options.ContentType)
		}
		if options.CacheControl != "" {
			input.CacheControl = aws.String(options.CacheControl)
		}
		if options.ContentEncoding != "" {
			input.ContentEncoding = aws.String(options.ContentEncoding)
		}
		if len(options.Metadata) > 0 {
			input.Metadata = options.Metadata
		}
	}

	_, err := w.client.PutObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to upload object %s: %w", key, err)
	}

	return nil
}

// GetObject はオブジェクトをS3から取得します
func (w *S3ClientWrapper) GetObject(ctx context.Context, key string) (io.ReadCloser, error) {
	input := &s3.GetObjectInput{
		Bucket: aws.String(w.config.BucketName),
		Key:    aws.String(key),
	}

	result, err := w.client.GetObject(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to get object %s: %w", key, err)
	}

	return result.Body, nil
}

// DeleteObject はオブジェクトをS3から削除します
func (w *S3ClientWrapper) DeleteObject(ctx context.Context, key string) error {
	input := &s3.DeleteObjectInput{
		Bucket: aws.String(w.config.BucketName),
		Key:    aws.String(key),
	}

	_, err := w.client.DeleteObject(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to delete object %s: %w", key, err)
	}

	return nil
}

// ListObjects はオブジェクト一覧を取得します
func (w *S3ClientWrapper) ListObjects(ctx context.Context, prefix string) ([]types.Object, error) {
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String(w.config.BucketName),
		Prefix: aws.String(prefix),
	}

	var objects []types.Object
	paginator := s3.NewListObjectsV2Paginator(w.client, input)

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to list objects with prefix %s: %w", prefix, err)
		}

		objects = append(objects, page.Contents...)
	}

	return objects, nil
}

// GeneratePresignedURL は署名付きURLを生成します
func (w *S3ClientWrapper) GeneratePresignedURL(ctx context.Context, key string, operation string) (string, error) {
	presignClient := s3.NewPresignClient(w.client)

	switch operation {
	case "GET":
		request, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(w.config.BucketName),
			Key:    aws.String(key),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(w.config.PresignedURLTTL) * time.Second
		})
		if err != nil {
			return "", fmt.Errorf("failed to generate presigned GET URL for %s: %w", key, err)
		}
		return request.URL, nil

	case "PUT":
		request, err := presignClient.PresignPutObject(ctx, &s3.PutObjectInput{
			Bucket: aws.String(w.config.BucketName),
			Key:    aws.String(key),
		}, func(opts *s3.PresignOptions) {
			opts.Expires = time.Duration(w.config.PresignedURLTTL) * time.Second
		})
		if err != nil {
			return "", fmt.Errorf("failed to generate presigned PUT URL for %s: %w", key, err)
		}
		return request.URL, nil

	default:
		return "", fmt.Errorf("unsupported operation: %s", operation)
	}
}
