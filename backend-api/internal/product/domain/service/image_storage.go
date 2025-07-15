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

// UploadResult はアップロード結果を格納する構造体
type UploadResult struct {
	Key  string
	URLs map[string]string
}

// GetImageResult は画像取得結果を格納する構造体
type GetImageResult struct {
	ImageData   []byte
	ContentType string
}

// ImageStorage は商品画像のストレージインターフェース
type ImageStorage interface {
	// UploadImage は商品画像をアップロードし、S3のキーとURLマップを返却する
	UploadImage(ctx context.Context, productID int, imageData []byte) (*UploadResult, error)

	// GetImageData は指定されたサイズの画像データを取得する
	GetImageData(ctx context.Context, productID int, size SizeType) (*GetImageResult, error)
}
