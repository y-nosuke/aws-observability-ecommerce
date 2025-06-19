package middleware

import (
	"context"
	"io"
	"mime/multipart"
	"net/http"
	"regexp"
	"strings"

	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/y-nosuke/aws-observability-ecommerce/backend-api/pkg/logger"
)

// FileUploadMiddleware はファイルアップロード情報を自動記録するミドルウェア
func FileUploadMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// ファイルアップロードエンドポイントかどうかを判定
			if !isFileUploadEndpoint(c) {
				return next(c)
			}

			// トレーシングスパンを取得
			span := trace.SpanFromContext(c.Request().Context())
			ctx := c.Request().Context()

			// ファイル情報を抽出・記録
			extractFileUploadInfo(c, span, ctx)

			return next(c)
		}
	}
}

// isFileUploadEndpoint はファイルアップロードエンドポイントかを判定
func isFileUploadEndpoint(c echo.Context) bool {
	// POSTメソッドかつmultipart/form-dataのContent-Type
	if c.Request().Method != "POST" {
		return false
	}

	contentType := c.Request().Header.Get("Content-Type")
	if contentType == "" {
		return false
	}

	// multipart/form-dataの判定
	matched, err := regexp.MatchString(`multipart/form-data`, contentType)
	if err != nil || !matched {
		return false
	}

	// 画像アップロード系のパスパターンを確認
	path := c.Request().URL.Path
	imageUploadPatterns := []string{
		`/products/\d+/image`,
		`/api/products/\d+/image`,
		`/upload`,
		`/image`,
	}

	for _, pattern := range imageUploadPatterns {
		matched, err := regexp.MatchString(pattern, path)
		if err == nil && matched {
			return true
		}
	}

	return false
}

// extractFileUploadInfo はファイルアップロード情報を抽出・記録
func extractFileUploadInfo(c echo.Context, span trace.Span, ctx context.Context) {
	// multipart formの解析
	err := c.Request().ParseMultipartForm(32 << 20) // 32MB limit
	if err != nil {
		span.SetAttributes(
			attribute.Bool("app.file_upload.parse_error", true),
			attribute.String("app.file_upload.error", err.Error()),
		)
		return
	}

	multipartForm := c.Request().MultipartForm
	if multipartForm == nil || multipartForm.File == nil {
		return
	}

	// ファイル情報を記録
	var totalFiles int
	var totalSize int64

	for fieldName, files := range multipartForm.File {
		for _, fileHeader := range files {
			totalFiles++
			totalSize += fileHeader.Size

			// ファイル拡張子を取得
			if ext := getFileExtension(fileHeader.Filename); ext != "" {
				span.SetAttributes(attribute.String("app.file_upload.extension", ext))
			}

			// 最初のファイルの詳細情報をスパンに記録
			if totalFiles == 1 {
				span.SetAttributes(
					attribute.String("app.file_upload.field_name", fieldName),
					attribute.String("app.file_upload.filename", fileHeader.Filename),
					attribute.Int64("app.file_upload.size", fileHeader.Size),
					attribute.String("app.file_upload.content_type", getFileContentType(fileHeader)),
				)

				// ログ記録
				logger.Info(ctx, "ファイルアップロード検出",
					"field_name", fieldName,
					"filename", fileHeader.Filename,
					"file_size_bytes", fileHeader.Size,
					"content_type", getFileContentType(fileHeader),
					"layer", "middleware",
				)
			}
		}
	}

	// 総合情報をスパンに記録
	span.SetAttributes(
		attribute.Bool("app.file_upload.detected", true),
		attribute.Int("app.file_upload.total_files", totalFiles),
		attribute.Int64("app.file_upload.total_size", totalSize),
	)

	if totalFiles > 1 {
		// 複数ファイルの場合は追加情報をログ記録
		logger.Info(ctx, "複数ファイルアップロード検出",
			"total_files", totalFiles,
			"total_size_bytes", totalSize,
			"layer", "middleware",
		)
	}
}

// getFileContentType はファイルヘッダーからContent-Typeを取得
func getFileContentType(fileHeader *multipart.FileHeader) string {
	// ヘッダーから取得を試行
	if contentType := fileHeader.Header.Get("Content-Type"); contentType != "" {
		return contentType
	}

	// ファイル内容から推測（最初の512バイト）
	file, err := fileHeader.Open()
	if err != nil {
		return "application/octet-stream"
	}
	defer file.Close()

	buffer := make([]byte, 512)
	n, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "application/octet-stream"
	}

	return http.DetectContentType(buffer[:n])
}

// getFileExtension はファイル名から拡張子を取得
func getFileExtension(filename string) string {
	if filename == "" {
		return ""
	}

	// 最後の"."以降を拡張子として取得
	parts := strings.Split(filename, ".")
	if len(parts) < 2 {
		return ""
	}

	return strings.ToLower(parts[len(parts)-1])
}
