package usecase

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/application/dto"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/product/domain/service"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
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
	// トレーシングスパンを開始
	ctx, span := tracer.StartUseCase(ctx, "upload_product_image", "product")
	defer span.End()

	// 追加の属性を設定
	span.SetAttributes(
		attribute.Int64("app.entity_id", req.ProductID),
		attribute.String("app.filename", req.Filename),
		attribute.Int("app.file_size_bytes", len(req.ImageData)),
	)

	completeOp := logger.StartOperation(ctx, "upload_product_image",
		"product_id", req.ProductID,
		"filename", req.Filename,
		"file_size_bytes", len(req.ImageData),
		"layer", "usecase")

	// ファイル拡張子の検証
	fileExt := strings.ToLower(filepath.Ext(req.Filename))
	span.SetAttributes(attribute.String("app.file_extension", fileExt))

	if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
		err := fmt.Errorf("only JPG and PNG images are supported")

		// バリデーションエラーログ
		logger.WithError(ctx, "画像ファイル拡張子がサポート外", err,
			"product_id", req.ProductID,
			"filename", req.Filename,
			"file_extension", fileExt,
			"supported_extensions", "jpg,jpeg,png",
			"layer", "usecase",
			"validation_error", "unsupported_file_extension")

		// スパンにエラー情報を記録
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "validation_error"),
			attribute.String("error.detail", "unsupported_file_extension"),
		)

		completeOp(false, "error_type", "validation_error", "file_extension", fileExt)
		return nil, err
	}

	// バリデーション成功ログ
	logger.Info(ctx, "ファイル拡張子のバリデーション成功",
		"product_id", req.ProductID,
		"file_extension", fileExt,
		"layer", "usecase")

	// バリデーション成功をスパンに記録
	span.SetAttributes(attribute.Bool("app.validation_passed", true))

	// 画像をアップロード
	s3Key, urls, err := u.imageStorage.UploadImage(ctx, req.ProductID, fileExt, req.ImageData)
	if err != nil {
		// アップロードエラーログ
		logger.WithError(ctx, "画像アップロードに失敗", err,
			"product_id", req.ProductID,
			"filename", req.Filename,
			"file_extension", fileExt,
			"layer", "usecase",
			"storage_operation", "upload_image")

		// スパンにエラー情報を記録
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		span.SetAttributes(
			attribute.Bool("error", true),
			attribute.String("error.type", "storage_error"),
		)

		completeOp(false, "error_type", "storage_failure")
		return nil, fmt.Errorf("failed to upload image: %w", err)
	}

	// 成功ログ
	logger.Info(ctx, "画像アップロードが正常に完了",
		"product_id", req.ProductID,
		"filename", req.Filename,
		"s3_key", s3Key,
		"generated_urls", len(urls),
		"layer", "usecase")

	// 成功情報をスパンに記録
	span.SetAttributes(
		attribute.String("app.s3_key", s3Key),
		attribute.Int("app.generated_urls", len(urls)),
		attribute.Bool("app.success", true),
	)

	// 操作成功を記録
	completeOp(true,
		"s3_key", s3Key,
		"generated_urls", len(urls))

	return dto.NewUploadImageResponse(req.ProductID, req.Filename, s3Key, urls), nil
}
