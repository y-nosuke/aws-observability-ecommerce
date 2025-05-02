package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/labstack/echo/v4"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/api/openapi"
	awsconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/aws"
)

// S3設定
const (
	S3BucketName = "product-images"
)

// ProductImageHandler は商品画像を管理するハンドラーです
type ProductImageHandler struct {
	s3Client *s3.Client
}

// NewProductImageHandler は新しいProductImageHandlerを作成します
func NewProductImageHandler(awsConfig *awsconfig.Config) (*ProductImageHandler, error) {

	return &ProductImageHandler{
		s3Client: awsConfig.S3,
	}, nil
}

// UploadProductImage は商品画像をアップロードするハンドラーです
func (h *ProductImageHandler) UploadProductImage(ctx echo.Context, id openapi.ProductIdParam) error {
	// フォームからファイルを取得
	file, err := ctx.FormFile("image")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": fmt.Sprintf("failed to get uploaded file: %v", err),
		})
	}

	// ファイルを開く
	src, err := file.Open()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("failed to open uploaded file: %v", err),
		})
	}
	defer src.Close()

	// ファイル内容を読み込む
	fileBytes, err := io.ReadAll(src)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("failed to read uploaded file: %v", err),
		})
	}

	// ファイル名から拡張子を取得
	fileExt := strings.ToLower(filepath.Ext(file.Filename))
	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": "only JPG and PNG images are supported",
		})
	}

	// Content-Typeを設定
	contentType := "image/jpeg"
	if fileExt == ".png" {
		contentType = "image/png"
	}

	// S3へのアップロード先キーを生成
	timestamp := time.Now().Unix()
	s3Key := fmt.Sprintf("uploads/%d-%d%s", id, timestamp, fileExt)

	// S3にアップロード
	_, err = h.s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(S3BucketName),
		Key:         aws.String(s3Key),
		Body:        bytes.NewReader(fileBytes),
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("failed to upload file to S3: %v", err),
		})
	}

	// ファイル名（拡張子なし）
	fileNameWithoutExt := strings.TrimSuffix(filepath.Base(s3Key), fileExt)

	// レスポンス返却
	// 注：Lambda関数がリサイズ処理をしたあと、それぞれのサイズのURLが生成されるまでに少し時間がかかります
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message":   "File uploaded successfully",
		"productId": id,
		"filename":  file.Filename,
		"s3Key":     s3Key,
		"urls": map[string]string{
			"original":  fmt.Sprintf("http://localhost:4566/%s/%s", S3BucketName, s3Key),
			"thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/%s_thumbnail%s", S3BucketName, fileNameWithoutExt, fileExt),
			"medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/%s_medium%s", S3BucketName, fileNameWithoutExt, fileExt),
			"large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/%s_large%s", S3BucketName, fileNameWithoutExt, fileExt),
		},
	})
}

// GetProductImage は商品画像のURLを取得するハンドラーです
func (h *ProductImageHandler) GetProductImage(ctx echo.Context, id openapi.ProductIdParam, params openapi.GetProductImageParams) error {
	// S3バケットからオブジェクトリストを取得
	resp, err := h.s3Client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(S3BucketName),
		Prefix: aws.String(fmt.Sprintf("uploads/%d", id)),
	})
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": fmt.Sprintf("failed to list objects in S3: %v", err),
		})
	}

	// 画像が見つからない場合
	if len(resp.Contents) == 0 {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "no images found for this product",
		})
	}

	// 最新の画像を取得
	var latestKey string
	var latestTime time.Time
	for _, obj := range resp.Contents {
		if obj.LastModified.After(latestTime) {
			latestTime = *obj.LastModified
			latestKey = *obj.Key
		}
	}

	// ファイル拡張子を取得
	fileExt := strings.ToLower(filepath.Ext(latestKey))

	// ファイル名（拡張子なし）
	fileNameWithoutExt := strings.TrimSuffix(filepath.Base(latestKey), fileExt)

	// URLを生成して返却
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"productId": id,
		"s3Key":     latestKey,
		"urls": map[string]string{
			"original":  fmt.Sprintf("http://localhost:4566/%s/%s", S3BucketName, latestKey),
			"thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/%s_thumbnail%s", S3BucketName, fileNameWithoutExt, fileExt),
			"medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/%s_medium%s", S3BucketName, fileNameWithoutExt, fileExt),
			"large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/%s_large%s", S3BucketName, fileNameWithoutExt, fileExt),
		},
	})
}
