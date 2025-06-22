package dto

// GetImageResponse は画像取得のレスポンスDTO
type GetImageResponse struct {
	ProductID   int    `json:"productId"`
	ImageData   []byte `json:"-"`
	ContentType string `json:"-"`
}

// NewGetImageResponse は新しいGetImageResponseを作成する
func NewGetImageResponse(productID int, imageData []byte, contentType string) *GetImageResponse {
	return &GetImageResponse{
		ProductID:   productID,
		ImageData:   imageData,
		ContentType: contentType,
	}
}
