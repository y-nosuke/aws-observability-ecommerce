package main

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	awsconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-image-processor/internal/aws"
	"github.com/y-nosuke/aws-observability-ecommerce/backend-image-processor/internal/config"
	"golang.org/x/image/draw"
)

// ImageSize は画像サイズを定義します
type ImageSize struct {
	Width  int
	Height int
}

// リサイズする画像サイズを定義
var (
	ThumbnailSize = ImageSize{Width: 200, Height: 200}
	MediumSize    = ImageSize{Width: 600, Height: 600}
	LargeSize     = ImageSize{Width: 1200, Height: 1200}
)

// handler は S3 イベントを処理します
func handler(ctx context.Context, s3Event events.S3Event) (string, error) {
	// 設定をロード
	if err := config.Load(); err != nil {
		log.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}

	// AWS設定オプションの準備
	awsOptions := awsconfig.Options{
		UseLocalStack: config.AWS.UseLocalStack,
		Region:        config.AWS.Region,
		Endpoint:      config.AWS.Endpoint,
		Credentials: awsconfig.Credentials{
			AccessKey: config.AWS.AccessKey,
			SecretKey: config.AWS.SecretKey,
			Token:     config.AWS.Token,
		},
	}

	awsConfig, err := awsconfig.NewAWSConfig(ctx, awsOptions)
	if err != nil {
		log.Printf("AWS設定エラー: %v", err)
		return "", err
	}

	// 各レコードを処理
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		log.Printf("処理開始: バケット=%s, キー=%s", bucket, key)

		// 処理済み画像のプレフィックスをチェック（無限ループ防止）
		if !strings.HasPrefix(key, "uploads/") ||
			strings.HasPrefix(key, "resized/") {
			log.Printf("処理対象外のオブジェクトなのでスキップします: %s", key)
			continue
		}

		// オリジナル画像の取得
		getResult, err := awsConfig.S3.GetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			log.Printf("画像取得エラー: %v", err)
			return "", err
		}
		defer getResult.Body.Close()

		// 画像フォーマットの判定
		contentType := ""
		if getResult.ContentType != nil {
			contentType = *getResult.ContentType
		}

		var format string
		if strings.Contains(contentType, "jpeg") || strings.Contains(contentType, "jpg") {
			format = "jpeg"
		} else if strings.Contains(contentType, "png") {
			format = "png"
		} else {
			// Content-Typeからフォーマットを判断できない場合はファイル名の拡張子で判断
			ext := strings.ToLower(filepath.Ext(key))
			if ext == ".jpg" || ext == ".jpeg" {
				format = "jpeg"
			} else if ext == ".png" {
				format = "png"
			} else {
				log.Printf("サポートされていない画像フォーマット: %s", contentType)
				return "", fmt.Errorf("サポートされていない画像フォーマット: %s", contentType)
			}
		}

		// 画像データを読み込む
		imgData, err := io.ReadAll(getResult.Body)
		if err != nil {
			log.Printf("画像読み込みエラー: %v", err)
			return "", err
		}

		// 元のファイル名から拡張子を取り除く
		basename := filepath.Base(key)
		extension := filepath.Ext(basename)
		filenameWithoutExt := strings.TrimSuffix(basename, extension)

		// 各サイズにリサイズして保存
		sizes := map[string]ImageSize{
			"thumbnail": ThumbnailSize,
			"medium":    MediumSize,
			"large":     LargeSize,
		}

		for sizeName, size := range sizes {
			// リサイズした画像データを作成
			resizedData, err := resizeImage(bytes.NewReader(imgData), format, size)
			if err != nil {
				log.Printf("%sサイズのリサイズエラー: %v", sizeName, err)
				continue
			}

			// 保存先のキーを生成
			resizedKey := fmt.Sprintf("resized/%s/%s_%s%s", sizeName, filenameWithoutExt, sizeName, extension)

			// Content-Typeの設定
			resizedContentType := "image/jpeg"
			if format == "png" {
				resizedContentType = "image/png"
			}

			// リサイズした画像をアップロード
			_, err = awsConfig.S3.PutObject(ctx, &s3.PutObjectInput{
				Bucket:      aws.String(bucket),
				Key:         aws.String(resizedKey),
				Body:        bytes.NewReader(resizedData),
				ContentType: aws.String(resizedContentType),
				Metadata: map[string]string{
					"ProcessedBy": "ImageProcessorLambda",
					"OriginalKey": key,
					"Size":        fmt.Sprintf("%dx%d", size.Width, size.Height),
				},
			})
			if err != nil {
				log.Printf("リサイズ画像アップロードエラー: %v", err)
				continue
			}

			log.Printf("%sサイズの画像をアップロードしました: %s", sizeName, resizedKey)
		}

		log.Printf("画像処理完了: %s", key)
	}

	return "画像処理が完了しました", nil
}

// resizeImage は画像をリサイズします
func resizeImage(src io.Reader, format string, size ImageSize) ([]byte, error) {
	// 画像をデコード
	var img image.Image
	var err error

	switch format {
	case "jpeg":
		img, err = jpeg.Decode(src)
	case "png":
		img, err = png.Decode(src)
	default:
		return nil, fmt.Errorf("サポートされていない画像フォーマット: %s", format)
	}

	if err != nil {
		return nil, fmt.Errorf("画像のデコードに失敗しました: %w", err)
	}

	// リサイズ用の新しい矩形を作成
	dst := image.NewRGBA(image.Rect(0, 0, size.Width, size.Height))

	// 画像をリサイズ (CatmullRom は高品質なリサイズアルゴリズム)
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	// リサイズした画像をエンコード
	var buf bytes.Buffer
	switch format {
	case "jpeg":
		err = jpeg.Encode(&buf, dst, &jpeg.Options{Quality: 85})
	case "png":
		err = png.Encode(&buf, dst)
	}

	if err != nil {
		return nil, fmt.Errorf("画像のエンコードに失敗しました: %w", err)
	}

	return buf.Bytes(), nil
}

func main() {
	lambda.Start(handler)
}
