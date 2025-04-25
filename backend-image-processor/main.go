package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func handler(ctx context.Context, s3Event events.S3Event) (string, error) {
	// LocalStackのエンドポイントを設定
	awsEndpoint := "http://localstack:4566"

	// AWS SDK v2 のカスタム設定を作成
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...any) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               awsEndpoint,
			HostnameImmutable: true,
			SigningRegion:     "us-east-1",
		}, nil
	})

	// AWS SDK v2 の設定を作成
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Printf("設定作成エラー: %v", err)
		return "", err
	}

	// S3クライアントの作成
	s3Client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true // S3 Path Style を有効化
	})

	// 各レコードを処理
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		log.Printf("処理開始: バケット=%s, キー=%s", bucket, key)

		// 処理済み画像のプレフィックスをチェック（無限ループ防止）
		if strings.HasPrefix(key, "processed/") {
			log.Printf("既に処理済みの画像なのでスキップします: %s", key)
			continue
		}

		// オリジナル画像の取得
		getResult, err := s3Client.GetObject(ctx, &s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
		})
		if err != nil {
			log.Printf("画像取得エラー: %v", err)
			return "", err
		}
		defer getResult.Body.Close()

		// 実際の画像処理の代わりに、メタデータを追加するシンプルな処理を行う
		// 注: 実際の画像リサイズ処理には外部ライブラリが必要です
		var buf bytes.Buffer
		if _, err = io.Copy(&buf, getResult.Body); err != nil {
			log.Printf("画像読み込みエラー: %v", err)
			return "", err
		}
		processedData := buf.Bytes()

		// 処理済み画像の新しいパス
		newKey := fmt.Sprintf("processed/%s", key)

		// 処理済み画像をアップロード
		_, err = s3Client.PutObject(ctx, &s3.PutObjectInput{
			Bucket:      aws.String(bucket),
			Key:         aws.String(newKey),
			Body:        bytes.NewReader(processedData),
			ContentType: getResult.ContentType,
			Metadata: map[string]string{
				"ProcessedBy": "ImageProcessorLambda",
				"OriginalKey": key,
			},
		})
		if err != nil {
			log.Printf("処理済み画像アップロードエラー: %v", err)
			return "", err
		}

		log.Printf("画像処理完了: %s -> %s", key, newKey)
	}

	return "画像処理が完了しました", nil
}

func main() {
	lambda.Start(handler)
}
