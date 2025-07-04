package service

import (
	"context"
)

type SizeType string

const (
	Thumbnail = SizeType("thumbnail")
	Medium    = SizeType("medium")
	Large     = SizeType("large")
	Original  = SizeType("original")
)

// ImageStorage は商品画像のストレージインターフェース
type ImageStorage interface {
	// UploadImage は商品画像をアップロードし、S3のキーとURLマップを返却する
	UploadImage(ctx context.Context, productID int, imageData []byte) (string, map[string]string, error)

	// GetImageData は指定されたサイズの画像データを取得する
	GetImageData(ctx context.Context, productID int, size SizeType) ([]byte, string, error)
}
