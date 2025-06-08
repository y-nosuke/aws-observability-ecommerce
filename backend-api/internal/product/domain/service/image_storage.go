package service

import (
	"context"
)

// ImageStorage は商品画像のストレージインターフェース
type ImageStorage interface {
	// UploadImage は商品画像をアップロードし、S3のキーとURLマップを返却する
	UploadImage(ctx context.Context, productID int64, fileExt string, imageData []byte) (string, map[string]string, error)

	// GetImageData は指定されたサイズの画像データを取得する
	GetImageData(ctx context.Context, productID int64, size string) ([]byte, string, error)
}
