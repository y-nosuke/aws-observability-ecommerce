package usecase

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"go.opentelemetry.io/otel/attribute"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/tracer"
)

// UploadProductImageUseCase は商品画像アップロードのユースケース
type UploadProductImageUseCase struct {
	imageStorage service.ImageStorage
}

// NewUploadProductImageUseCase は新しいUploadProductImageUseCaseを作成する
func NewUploadProductImageUseCase(imageStorage service.ImageStorage) *UploadProductImageUseCase {
	return &UploadProductImageUseCase{
		imageStorage: imageStorage,
	}
}

// Execute は商品画像アップロードを実行する
func (u *UploadProductImageUseCase) Execute(ctx context.Context, req *dto.UploadImageRequest) (*dto.UploadImageResponse, error) {
	// UseCase トレーサーを開始
	useCaseTracer := tracer.NewUseCaseTracer(ctx, "upload_product_image", "product", req.ProductID)
	defer useCaseTracer.Finish(true)

	// ファイル拡張子の検証
	fileExt := strings.ToLower(filepath.Ext(req.Filename))
	useCaseTracer.SetAttributes(
		attribute.String("app.filename", req.Filename),
		attribute.Int("app.file_size_bytes", len(req.ImageData)),
		attribute.String("app.file_extension", fileExt),
	)

	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
		err := useCaseTracer.ValidationError("only JPG and PNG images are supported", "file_extension", fileExt)
		useCaseTracer.FinishWithError(err, "ファイル拡張子のバリデーションに失敗")
		return nil, err
	}

	// 画像をアップロード
	var s3Key string
	var urls map[string]string
	var err error

	err = useCaseTracer.AddStep("upload_image", func(stepCtx context.Context) error {
		s3Key, urls, err = u.imageStorage.UploadImage(stepCtx, req.ProductID, fileExt, req.ImageData)
		return err
	})

	if err != nil {
		useCaseTracer.FinishWithError(err, "画像アップロードに失敗")
		return nil, fmt.Errorf("failed to upload image: %w", err)
	}

	// 成功時のビジネスイベント記録
	useCaseTracer.BusinessEvent("product_image_uploaded", "product", fmt.Sprint(req.ProductID),
		"s3_key", s3Key,
		"filename", req.Filename,
		"file_size", len(req.ImageData),
	)

	useCaseTracer.LogInfo("画像アップロードが正常に完了",
		"product_id", req.ProductID,
		"s3_key", s3Key,
		"generated_urls", len(urls),
	)

	return dto.NewUploadImageResponse(req.ProductID, req.Filename, s3Key, urls), nil
}
