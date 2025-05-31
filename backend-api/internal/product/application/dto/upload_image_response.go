package dto

// UploadImageResponse は商品画像アップロードのレスポンスDTO
type UploadImageResponse struct {
	Message   string            `json:"message"`
	ProductID int64             `json:"productId"`
	Filename  string            `json:"filename"`
	S3Key     string            `json:"s3Key"`
	URLs      map[string]string `json:"urls"`
}

// NewUploadImageResponse は新しいUploadImageResponseを作成する
func NewUploadImageResponse(productID int64, filename string, s3Key string, urls map[string]string) *UploadImageResponse {
	return &UploadImageResponse{
		Message:   "File uploaded successfully",
		ProductID: productID,
		Filename:  filename,
		S3Key:     s3Key,
		URLs:      urls,
	}
}
