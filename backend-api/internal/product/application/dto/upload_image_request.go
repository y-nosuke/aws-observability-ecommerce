package dto

// UploadImageRequest は商品画像アップロードのリクエストDTO
type UploadImageRequest struct {
	ProductID int
	ImageData []byte
	Filename  string
}

// NewUploadImageRequest は新しいUploadImageRequestを作成する
func NewUploadImageRequest(productID int, imageData []byte, filename string) *UploadImageRequest {
	return &UploadImageRequest{
		ProductID: productID,
		ImageData: imageData,
		Filename:  filename,
	}
}
