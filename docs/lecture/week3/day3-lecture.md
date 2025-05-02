# 1. Week 3 - Day 3: Lambda関数とS3連携の実装

## 1.1. 目次

- [1. Week 3 - Day 3: Lambda関数とS3連携の実装](#1-week-3---day-3-lambda関数とs3連携の実装)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. 既存のLambda関数プロジェクトを確認](#141-既存のlambda関数プロジェクトを確認)
    - [1.4.2. 画像リサイズ機能の実装](#142-画像リサイズ機能の実装)
    - [1.4.3. LocalStackでのS3バケット作成とIAM設定](#143-localstackでのs3バケット作成とiam設定)
    - [1.4.4. Lambda関数のビルドとデプロイ](#144-lambda関数のビルドとデプロイ)
    - [1.4.5. S3トリガーの設定](#145-s3トリガーの設定)
    - [1.4.6. バックエンドAPIとの連携実装](#146-バックエンドapiとの連携実装)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. サーバーレスアーキテクチャの概要](#161-サーバーレスアーキテクチャの概要)
    - [1.6.2. AWS Lambda関数の基本構造](#162-aws-lambda関数の基本構造)
    - [1.6.3. S3バケットとイベント連携](#163-s3バケットとイベント連携)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. Lambda関数の最適化テクニック](#171-lambda関数の最適化テクニック)
    - [1.7.2. 本番環境での画像処理の考慮点](#172-本番環境での画像処理の考慮点)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: Lambda関数が起動しない](#181-問題1-lambda関数が起動しない)
    - [1.8.2. 問題2: 画像処理が失敗する](#182-問題2-画像処理が失敗する)
    - [1.8.3. 問題3: S3イベントが正しくトリガーされない](#183-問題3-s3イベントが正しくトリガーされない)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- サーバーレスアーキテクチャとAWS Lambdaの基本概念を理解する
- LocalStackを使ってローカル環境でAWS Lambdaを実装・テストする
- 商品画像リサイズのLambda関数を実装する
- S3バケットとの連携およびS3イベントトリガーを設定する
- イベント駆動型の設計パターンを実践的に学ぶ

## 1.3. 【準備】

今日の実装では、既存のLambda関数プロジェクトを拡張して本格的な画像リサイズ機能を実装します。LocalStackを使ったAWS環境エミュレーションで、S3イベントトリガーでLambda関数を起動する仕組みを構築します。

### 1.3.1. チェックリスト

- [ ] Docker Composeが起動しており、LocalStackコンテナが動作している
- [ ] AWS CLIがインストールされ、LocalStack用に設定されている
- [ ] Go言語の開発環境が整っている（Go 1.18以上）
- [ ] 既存のbackend-image-processorプロジェクトにアクセスできる
- [ ] Week 1-2で実装したバックエンドが動作していることを確認

## 1.4. 【手順】

### 1.4.1. 既存のLambda関数プロジェクトを確認

最初に、既存のLambda関数プロジェクトを確認しましょう。現在のプロジェクトには基本的なLambda関数のスケルトンが実装されていますが、実際の画像処理機能はまだ含まれていません。

```bash
# プロジェクトディレクトリに移動
cd backend-image-processor
```

現在の実装では、S3イベントを受け取り、オブジェクトを取得して、メタデータを追加するだけの簡単な処理を行っています。これを拡張して、実際の画像リサイズ機能を実装します。

### 1.4.2. 画像リサイズ機能の実装

次に、`main.go` ファイルを更新して、画像リサイズ機能を実装します。

```go
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
```

```bash
go mod tidy
```

このコードでは、以下の画像処理機能を実装しています：

1. S3からアップロードされた画像を読み込む
2. 3つの異なるサイズ（サムネイル、中、大）にリサイズする
3. リサイズした画像を「resized/サイズ名/」プレフィックスでS3にアップロードする
4. 各画像のメタデータに処理情報を追加する

### 1.4.3. LocalStackでのS3バケット作成とIAM設定

LocalStackを使用してS3バケットを作成し、必要なIAM設定を行います。`infra/scripts/aws/create-s3-bucket.sh`を修正して実行します。

```bash
#!/bin/bash

# LocalStackのエンドポイントを設定
export AWS_ENDPOINT_URL=http://localhost:4566

# バケット名
BUCKET_NAME="product-images"

# バケットが存在するか確認
if awslocal s3api head-bucket --bucket "${BUCKET_NAME}" 2>/dev/null; then
  echo "バケット ${BUCKET_NAME} は既に存在します。"
else
  # バケットの作成
  awslocal s3 mb s3://${BUCKET_NAME}
  echo "バケット ${BUCKET_NAME} を作成しました。"

  # CORSの設定
  cat > /tmp/cors-config.json << EOF
{
  "CORSRules": [
    {
      "AllowedHeaders": ["*"],
      "AllowedMethods": ["GET", "PUT", "POST", "DELETE"],
      "AllowedOrigins": ["*"],
      "ExposeHeaders": ["ETag"]
    }
  ]
}
EOF

  # CORSの適用
  awslocal s3api put-bucket-cors --bucket ${BUCKET_NAME} --cors-configuration '{
    "CORSRules": [
      {
        "AllowedHeaders": ["*"],
        "AllowedMethods": ["GET", "PUT", "POST", "DELETE"],
        "AllowedOrigins": ["*"],
        "ExposeHeaders": ["ETag"]
      }
    ]
  }'
  echo "バケット ${BUCKET_NAME} にCORS設定を適用しました。"

  # パブリックアクセスの設定
  awslocal s3api put-bucket-policy --bucket ${BUCKET_NAME} --policy '{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Sid": "PublicReadGetObject",
        "Effect": "Allow",
        "Principal": "*",
        "Action": "s3:GetObject",
        "Resource": "arn:aws:s3:::product-images/*"
      }
    ]
  }'
  echo "バケット ${BUCKET_NAME} にパブリックアクセスポリシーを適用しました。"
fi

# バケットのリスト表示
echo "現在のS3バケット一覧:"
awslocal s3 ls
```

### 1.4.4. Lambda関数のビルドとデプロイ

Lambda関数をビルドしてLocalStackにデプロイするための `infra/scripts/aws/deploy-lambda.sh` を修正します。

```bash
#!/bin/bash

set -e

# LocalStackのエンドポイントを設定
export AWS_ENDPOINT_URL=http://localhost:4566

# Lambda関数名
FUNCTION_NAME="image-processor"
OUTPUT_DIR="build"
ZIP_FILE="${OUTPUT_DIR}/${FUNCTION_NAME}.zip"

# バケット名
BUCKET_NAME=product-images

# ビルド用ディレクトリの作成
mkdir -p "$PWD/backend-image-processor/${OUTPUT_DIR}"

echo "Lambda関数 ${FUNCTION_NAME} をAmazon Linux 2互換環境でビルドします..."

# DockerでAmazon Linux 2ベースのGoイメージを使ってバイナリをビルド
docker run --rm -v "$PWD/backend-image-processor":/go/src/app -w /go/src/app \
  golang:1.24 \
  /bin/bash -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o ${OUTPUT_DIR}/${FUNCTION_NAME} ./main.go"


# バイナリをZIP化
cd "backend-image-processor/${OUTPUT_DIR}" && zip -j "${FUNCTION_NAME}.zip" "${FUNCTION_NAME}" && cd -

echo "Lambda関数のビルドとZIP化が完了しました: ${ZIP_FILE}"

# Lambda関数が既に存在するか確認
if awslocal lambda get-function --function-name "${FUNCTION_NAME}" 2>/dev/null; then
  echo "Lambda関数 ${FUNCTION_NAME} を更新します..."

  awslocal lambda update-function-code \
    --function-name "${FUNCTION_NAME}" \
    --zip-file "fileb://backend-image-processor/${ZIP_FILE}"

  echo "Lambda関数 ${FUNCTION_NAME} を更新しました。"
else
  echo "Lambda関数 ${FUNCTION_NAME} を新規作成します..."

  # Lambda関数の作成
  awslocal lambda create-function \
    --function-name "${FUNCTION_NAME}" \
    --runtime "go1.x" \
    --handler "${FUNCTION_NAME}" \
    --timeout 30 \
    --memory-size 512 \
    --role "arn:aws:iam::000000000000:role/lambda-${FUNCTION_NAME}-role" \
    --zip-file "fileb://backend-image-processor/${ZIP_FILE}" \
    --environment "Variables={BUCKET_NAME=product-images}"

  echo "Lambda関数 ${FUNCTION_NAME} を作成しました。"
fi

# 実行ロールとS3アクセスポリシーの処理
ROLE_NAME="lambda-${FUNCTION_NAME}-role"
POLICY_NAME="s3-access-policy"

# ロールが存在するか確認
if awslocal iam get-role --role-name "${ROLE_NAME}" 2>/dev/null; then
  echo "IAMロール ${ROLE_NAME} は既に存在しています。"
else
  echo "IAMロール ${ROLE_NAME} を作成します..."
  # ロールを作成
  awslocal iam create-role \
    --role-name "${ROLE_NAME}" \
    --assume-role-policy-document '{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "lambda.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }'
  echo "IAMロール ${ROLE_NAME} を作成しました。"
fi

# S3アクセスポリシーの確認と設定（共通処理）
if awslocal iam get-role-policy --role-name "${ROLE_NAME}" --policy-name "${POLICY_NAME}" 2>/dev/null; then
  echo "S3アクセスポリシー ${POLICY_NAME} は既にアタッチされています。更新します..."
else
  echo "S3アクセスポリシー ${POLICY_NAME} をアタッチします..."
fi

# ポリシーをアタッチまたは更新
awslocal iam put-role-policy \
  --role-name "${ROLE_NAME}" \
  --policy-name "${POLICY_NAME}" \
  --policy-document '{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:GetObject",
        "s3:PutObject",
        "s3:ListBucket"
      ],
      "Resource": [
        "arn:aws:s3:::'"${BUCKET_NAME}"'",
        "arn:aws:s3:::'"${BUCKET_NAME}"'/*"
      ]
    }
  ]
}'

echo "IAMロールとポリシーの設定が完了しました。"

# 関数の設定を表示
echo "Lambda関数の設定:"
awslocal lambda get-function --function-name "${FUNCTION_NAME}"
```

### 1.4.5. S3トリガーの設定

Lambda関数をS3イベントでトリガーするよう設定します。`infra/scripts/aws/setup-s3-trigger.sh` を修正します。

```bash
#!/bin/bash

# 使用方法表示関数
show_usage() {
    echo "使用方法: $0 [オプション]"
    echo "オプション:"
    echo "  -n    テスト画像のアップロードをスキップ"
    echo "  -s    サイレントモード (対話的な確認をスキップ)"
    echo "  -h    このヘルプメッセージを表示"
    exit 1
}

# S3バケット名とLambda関数名
BUCKET_NAME="product-images"
FUNCTION_NAME="image-processor"

# Lambda関数のARNを取得
LAMBDA_ARN=$(awslocal lambda get-function --function-name "${FUNCTION_NAME}" --query 'Configuration.FunctionArn' --output text)

echo "Lambda関数ARN: ${LAMBDA_ARN}"

# Lambda関数にS3のイベント通知権限を付与
echo "Lambda関数に S3 バケットからのイベント通知権限を付与します..."
# 権限がすでに存在するかチェック
PERMISSION_EXISTS=$(awslocal lambda get-policy --function-name "${FUNCTION_NAME}" 2>/dev/null | grep -c "s3-trigger" || echo "0")

# 整数比較のエラーを修正
if [ "${PERMISSION_EXISTS}" = "0" ]; then
  # 権限が存在しない場合のみ追加
  awslocal lambda add-permission \
    --function-name "${FUNCTION_NAME}" \
    --statement-id "s3-trigger" \
    --action "lambda:InvokeFunction" \
    --principal "s3.amazonaws.com" \
    --source-arn "arn:aws:s3:::${BUCKET_NAME}" \
    --source-account 000000000000
  echo "Lambda権限を追加しました。"
else
  echo "Lambda権限はすでに設定されています。スキップします。"
fi

# S3バケットにLambdaトリガーを設定
echo "S3バケット ${BUCKET_NAME} にLambdaトリガーを設定します..."

# トリガーの重複を避けるため、一度通知設定を削除してから再設定する
# 通知設定を空にする（削除）
awslocal s3api put-bucket-notification-configuration \
  --bucket "${BUCKET_NAME}" \
  --notification-configuration '{}'

# 少し待機して設定が反映されるのを待つ
sleep 1

# 通知設定を追加
awslocal s3api put-bucket-notification-configuration \
  --bucket "${BUCKET_NAME}" \
  --notification-configuration '{
    "LambdaFunctionConfigurations": [
      {
        "LambdaFunctionArn": "'"${LAMBDA_ARN}"'",
        "Events": ["s3:ObjectCreated:*"],
        "Filter": {
          "Key": {
            "FilterRules": [
              {
                "Name": "prefix",
                "Value": "uploads/"
              }
            ]
          }
        }
      }
    ]
  }'

echo "S3トリガーの設定が完了しました。"

# トリガーテスト用の画像アップロード
echo "トリガーテスト用の新しい画像をアップロードします..."

# 引数によるテスト画像アップロードの制御
UPLOAD_TEST=1
INTERACTIVE=1

# コマンドライン引数の確認
while getopts "nsh" opt; do
  case $opt in
    n) # アップロードをスキップ
      UPLOAD_TEST=0
      ;;
    s) # サイレントモード（対話なし）
      INTERACTIVE=0
      ;;
    h) # ヘルプ表示
      show_usage
      ;;
    \?)
      echo "無効なオプション: -$OPTARG" >&2
      show_usage
      ;;
  esac
done

# 対話モードで確認が必要な場合
if [ $INTERACTIVE -eq 1 ] && [ $UPLOAD_TEST -eq 1 ]; then
    read -p "テスト画像をアップロードしますか？ (y/n) [y]: " ANSWER
    if [[ "$ANSWER" =~ ^[Nn]$ ]]; then
        UPLOAD_TEST=0
    fi
fi

if [ $UPLOAD_TEST -eq 1 ]; then
    # 簡単なテスト画像を生成
    curl -L -o /tmp/trigger-test.jpg https://picsum.photos/800/600
    awslocal s3 cp /tmp/trigger-test.jpg s3://product-images/uploads/trigger-test.jpg

    echo "アップロード完了。数秒後に処理済みフォルダを確認してください。"
    echo "処理済み画像の確認コマンド: awslocal s3 ls s3://product-images/resized/"
else
    echo "テスト画像のアップロードをスキップしました。"
    echo "手動でテストする場合は以下のコマンドを使用してください:"
    echo "awslocal s3 cp your-test-image.jpg s3://product-images/uploads/your-test-image.jpg"
    echo "処理済み画像の確認コマンド: awslocal s3 ls s3://product-images/resized/"
fi
```

### 1.4.6. バックエンドAPIとの連携実装

メインのバックエンドアプリケーションに、S3バケットとの連携機能を追加します。以下のファイルを作成します。

```bash
touch backend-api/internal/api/handlers/product_image.go
```

プロダクトイメージハンドラに以下のコードを追加します：

```go
package handlers

import (
 "bytes"
 "context"
 "fmt"
 "io"
 "net/http"
 "path/filepath"
 "strings"
 "time"

 "github.com/aws/aws-sdk-go-v2/aws"
 "github.com/aws/aws-sdk-go-v2/service/s3"
 "github.com/labstack/echo/v4"

 awsconfig "github.com/y-nosuke/aws-observability-ecommerce/backend-api/internal/aws"
)

// S3設定
const (
 S3BucketName = "product-images"
 S3Endpoint   = "http://localstack:4566" // Docker Compose内でのLocalStack接続先
)

// ProductImageHandler は商品画像を管理するハンドラーです
type ProductImageHandler struct {
 s3Client *s3.Client
}

// NewProductImageHandler は新しいProductImageHandlerを作成します
func NewProductImageHandler(awsConfig *awsconfig.Config) (*ProductImageHandler, error) {

 return &ProductImageHandler{
  s3Client: awsConfig.S3,
 }, nil
}

// UploadProductImage は商品画像をアップロードするハンドラーです
func (h *ProductImageHandler) UploadProductImage(c echo.Context) error {
 // リクエストからproductIDを取得
 productID := c.Param("id")
 if productID == "" {
  return c.JSON(http.StatusBadRequest, map[string]string{
   "error": "product ID is required",
  })
 }

 // フォームからファイルを取得
 file, err := c.FormFile("image")
 if err != nil {
  return c.JSON(http.StatusBadRequest, map[string]string{
   "error": fmt.Sprintf("failed to get uploaded file: %v", err),
  })
 }

 // ファイルを開く
 src, err := file.Open()
 if err != nil {
  return c.JSON(http.StatusInternalServerError, map[string]string{
   "error": fmt.Sprintf("failed to open uploaded file: %v", err),
  })
 }
 defer src.Close()

 // ファイル内容を読み込む
 fileBytes, err := io.ReadAll(src)
 if err != nil {
  return c.JSON(http.StatusInternalServerError, map[string]string{
   "error": fmt.Sprintf("failed to read uploaded file: %v", err),
  })
 }

 // ファイル名から拡張子を取得
 fileExt := strings.ToLower(filepath.Ext(file.Filename))
 if fileExt != ".jpg" && fileExt != ".jpeg" && fileExt != ".png" {
  return c.JSON(http.StatusBadRequest, map[string]string{
   "error": "only JPG and PNG images are supported",
  })
 }

 // Content-Typeを設定
 contentType := "image/jpeg"
 if fileExt == ".png" {
  contentType = "image/png"
 }

 // S3へのアップロード先キーを生成
 timestamp := time.Now().Unix()
 s3Key := fmt.Sprintf("uploads/%s-%d%s", productID, timestamp, fileExt)

 // S3にアップロード
 _, err = h.s3Client.PutObject(context.Background(), &s3.PutObjectInput{
  Bucket:      aws.String(S3BucketName),
  Key:         aws.String(s3Key),
  Body:        bytes.NewReader(fileBytes),
  ContentType: aws.String(contentType),
 })
 if err != nil {
  return c.JSON(http.StatusInternalServerError, map[string]string{
   "error": fmt.Sprintf("failed to upload file to S3: %v", err),
  })
 }

 // ファイル名（拡張子なし）
 fileNameWithoutExt := strings.TrimSuffix(filepath.Base(s3Key), fileExt)

 // レスポンス返却
 // 注：Lambda関数がリサイズ処理をしたあと、それぞれのサイズのURLが生成されるまでに少し時間がかかります
 return c.JSON(http.StatusOK, map[string]interface{}{
  "message":   "File uploaded successfully",
  "productId": productID,
  "filename":  file.Filename,
  "s3Key":     s3Key,
  "urls": map[string]string{
   "original":  fmt.Sprintf("http://localhost:4566/%s/%s", S3BucketName, s3Key),
   "thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/%s_thumbnail%s", S3BucketName, fileNameWithoutExt, fileExt),
   "medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/%s_medium%s", S3BucketName, fileNameWithoutExt, fileExt),
   "large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/%s_large%s", S3BucketName, fileNameWithoutExt, fileExt),
  },
 })
}

// GetProductImage は商品画像のURLを取得するハンドラーです
func (h *ProductImageHandler) GetProductImage(c echo.Context) error {
 // リクエストからproductIDを取得
 productID := c.Param("id")
 if productID == "" {
  return c.JSON(http.StatusBadRequest, map[string]string{
   "error": "product ID is required",
  })
 }

 // S3バケットからオブジェクトリストを取得
 resp, err := h.s3Client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
  Bucket: aws.String(S3BucketName),
  Prefix: aws.String(fmt.Sprintf("uploads/%s", productID)),
 })
 if err != nil {
  return c.JSON(http.StatusInternalServerError, map[string]string{
   "error": fmt.Sprintf("failed to list objects in S3: %v", err),
  })
 }

 // 画像が見つからない場合
 if len(resp.Contents) == 0 {
  return c.JSON(http.StatusNotFound, map[string]string{
   "error": "no images found for this product",
  })
 }

 // 最新の画像を取得
 var latestKey string
 var latestTime time.Time
 for _, obj := range resp.Contents {
  if obj.LastModified.After(latestTime) {
   latestTime = *obj.LastModified
   latestKey = *obj.Key
  }
 }

 // ファイル拡張子を取得
 fileExt := strings.ToLower(filepath.Ext(latestKey))

 // ファイル名（拡張子なし）
 fileNameWithoutExt := strings.TrimSuffix(filepath.Base(latestKey), fileExt)

 // URLを生成して返却
 return c.JSON(http.StatusOK, map[string]interface{}{
  "productId": productID,
  "s3Key":     latestKey,
  "urls": map[string]string{
   "original":  fmt.Sprintf("http://localhost:4566/%s/%s", S3BucketName, latestKey),
   "thumbnail": fmt.Sprintf("http://localhost:4566/%s/resized/thumbnail/%s_thumbnail%s", S3BucketName, fileNameWithoutExt, fileExt),
   "medium":    fmt.Sprintf("http://localhost:4566/%s/resized/medium/%s_medium%s", S3BucketName, fileNameWithoutExt, fileExt),
   "large":     fmt.Sprintf("http://localhost:4566/%s/resized/large/%s_large%s", S3BucketName, fileNameWithoutExt, fileExt),
  },
 })
}
```

次に、これらのAPIエンドポイントをOpenAPI仕様に追加し、コード生成を行います。

まず、`openapi.yaml`に新しいエンドポイント定義を追加します：

```yaml
  /products/{id}/image:
    post:
      summary: 商品画像のアップロード
      description: 指定された商品IDに画像をアップロードします（管理者のみ）
      operationId: uploadProductImage
      tags:
        - product
        - admin
      parameters:
        - $ref: "#/components/parameters/ProductIdParam"
      requestBody:
        content:
          multipart/form-data:
            schema:
              type: object
              properties:
                image:
                  type: string
                  format: binary
                  description: アップロードする画像ファイル
              required:
                - image
      responses:
        "200":
          description: 画像が正常にアップロードされました
          content:
            application/json:
              schema:
                type: object
                properties:
                  imageUrl:
                    type: string
                    description: アップロードされた画像のURL
                  thumbnailUrl:
                    type: string
                    description: サムネイル画像のURL
                  mediumUrl:
                    type: string
                    description: 中サイズ画像のURL
                  largeUrl:
                    type: string
                    description: 大サイズ画像のURL
        "400":
          description: 無効なリクエスト
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
        "401":
          description: 認証エラー
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
        "403":
          description: 権限エラー
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
        "500":
          description: 内部サーバーエラー
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
        default:
          description: 予期しないエラー
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"

   get:
      summary: 商品画像の取得
      description: 指定された商品IDの画像を取得します
      operationId: getProductImage
      tags:
        - product
      parameters:
        - $ref: "#/components/parameters/ProductIdParam"
        - name: size
          in: query
          description: 画像サイズ（thumbnail, medium, large）
          schema:
            type: string
            enum: [thumbnail, medium, large]
            default: medium
      responses:
        "200":
          description: 画像が正常に取得されました
          content:
            image/*:
              schema:
                type: string
                format: binary
        "404":
          description: 画像が見つかりません
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
        "500":
          description: 内部サーバーエラー
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
        default:
          description: 予期しないエラー
          content:
            application/json:
              schema:
                $ref: "#/components/schemas/ErrorResponse"
```

次に、OpenAPI仕様からコードを生成します：

```bash
task generate:openapi
```

生成されたコードには、新しく追加したAPIエンドポイントのインターフェース定義が含まれています。これをハンドラーで実装します。ハンドラー実装では、`ServerInterfaceWrapper`がリクエストを受け取り、適切なハンドラーにディスパッチします。

## 1.5. 【確認ポイント】

このDayの作業が正しく完了したことを確認するためのチェックリストです：

- [ ] Lambda関数が画像リサイズ機能を実装している
- [ ] Lambda関数がLocalStackに正常にデプロイされている
- [ ] S3バケットが作成され、CORS設定が完了している
- [ ] S3イベント通知が設定され、Lambda関数がトリガーされる
- [ ] 画像アップロード後、3つのサイズ（thumbnail, medium, large）の画像が生成される
- [ ] バックエンドAPIからS3バケットに画像をアップロードできる
- [ ] アップロードした画像のURLを取得できる

テスト用スクリプトを作成して、機能を検証しましょう：

```bash
# test-upload.shスクリプトを作成
cat <<EOF > scripts/test-upload.sh
#!/bin/bash
set -e

# LocalStack用のAWS CLIエンドポイント設定
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=us-east-1
export AWS_ENDPOINT_URL=http://localhost:4566

# バケット名
BUCKET_NAME=product-images

# テスト用の画像をダウンロード
echo "Downloading test image..."
curl -s -o test-image.jpg https://picsum.photos/800/600

# uploads/プレフィックスを使ってアップロード
aws --endpoint-url=\${AWS_ENDPOINT_URL} s3 cp test-image.jpg s3://\${BUCKET_NAME}/uploads/test-product-1.jpg

echo "Test image uploaded to S3. Lambda function should be triggered."
echo "Waiting for Lambda function to process the image..."
sleep 5

# リサイズされた画像があるか確認
echo "Checking for resized images:"
aws --endpoint-url=\${AWS_ENDPOINT_URL} s3 ls s3://\${BUCKET_NAME}/resized/thumbnail/ || echo "Thumbnail not found"
aws --endpoint-url=\${AWS_ENDPOINT_URL} s3 ls s3://\${BUCKET_NAME}/resized/medium/ || echo "Medium not found"
aws --endpoint-url=\${AWS_ENDPOINT_URL} s3 ls s3://\${BUCKET_NAME}/resized/large/ || echo "Large not found"

echo "Test completed."
EOF

# スクリプトに実行権限を付与して実行
chmod +x scripts/test-upload.sh
./scripts/test-upload.sh
```

## 1.6. 【詳細解説】

### 1.6.1. サーバーレスアーキテクチャの概要

サーバーレスアーキテクチャは、サーバーのプロビジョニングや管理を開発者が行う必要がない実行モデルです。インフラストラクチャの管理からアプリケーション開発者を解放し、コードの開発に集中できるようにするアプローチです。

**サーバーレスの主な特徴**:

1. **イベント駆動**: 関数はイベント（例：HTTPリクエスト、ファイルのアップロード、スケジュール）によってトリガーされる
2. **自動スケーリング**: トラフィックに応じて自動的にスケールアップ・ダウンする
3. **使用量ベースの課金**: 実際に使用した分のみ課金される（コンピューティング時間、メモリ使用量など）
4. **ステートレス**: 関数は基本的にステートレスで、必要な状態は外部サービス（DBなど）に保存する
5. **短命なプロセス**: 関数は短時間実行され、処理終了後にリソースが解放される

**サーバーレスの利点**:

- インフラ管理の手間を削減
- スケーラビリティの向上
- コスト効率の改善（アイドル時間に対する課金がない）
- 開発スピードの向上
- メンテナンスの負担軽減

**サーバーレスの課題**:

- コールドスタート問題（初回起動時の遅延）
- 長時間実行の制限
- ベンダーロックイン
- デバッグの複雑さ
- ローカル開発環境の構築の難しさ

今回実装したLambda関数は、S3バケットに画像がアップロードされるというイベントに反応して自動的に起動し、画像処理を行うという典型的なサーバーレスパターンを実装しています。

### 1.6.2. AWS Lambda関数の基本構造

AWS Lambda関数は、AWSクラウド内で実行されるコードの単位です。トリガーに応答して実行され、必要なコンピューティングリソースを自動的に管理します。

**AWS Lambda関数の主要コンポーネント**:

1. **ハンドラー関数**: Lambda関数のエントリーポイント。イベントとコンテキストを受け取り、処理を行う
2. **イベントオブジェクト**: 関数を呼び出したサービスからの入力データ
3. **コンテキストオブジェクト**: 関数の実行環境に関する情報（タイムアウト、リソース制限など）
4. **環境変数**: 関数の設定や認証情報などを保存
5. **レイヤー**: 共通ライブラリやSDKなどを含む再利用可能なコンポーネント

**Lambda関数のライフサイクル**:

1. **初期化フェーズ**: 関数が初めて呼び出されるか、既存のインスタンスが利用可能でない場合に発生
   - 関数コードのダウンロードと実行環境の準備
   - ランタイムの初期化
   - 関数外のコード（グローバル変数やスコープ外のコード）の実行

2. **呼び出しフェーズ**: 実際の関数実行
   - ハンドラー関数の実行
   - イベント処理
   - レスポンス返却または例外発生

3. **シャットダウンフェーズ**: 関数が一定時間呼び出されない場合
   - リソースの解放
   - 実行環境の終了

**Goでの実装特性**:

Goは静的にコンパイルされるため、Lambda環境では特に高速に起動します。また、並行処理の機能が充実しているため、複数のタスクを効率的に処理できるという利点があります。今回の実装では、Goの提供ランタイム（provided.al2）を使用し、カスタムランタイムとして`bootstrap`という名前の実行ファイルを作成しています。

### 1.6.3. S3バケットとイベント連携

S3（Simple Storage Service）は、AWSのオブジェクトストレージサービスです。高い耐久性、可用性、スケーラビリティを提供し、様々なユースケースに対応できます。

**S3の主な特徴**:

1. **オブジェクトベース**: ファイルやメタデータをオブジェクトとして保存
2. **バケット**: オブジェクトを格納するコンテナ。グローバルに一意の名前が必要
3. **キー**: バケット内でオブジェクトを識別するユニークな名前
4. **バージョニング**: オブジェクトの複数バージョンを保持可能
5. **ライフサイクル管理**: オブジェクトの自動移行や削除ルールを設定可能
6. **イベント通知**: オブジェクトの作成、削除などのイベントを通知

**S3イベント通知**:

S3は様々なイベントタイプをサポートしており、オブジェクトの作成、削除、復元などのイベントを検出して、Lambda関数などのターゲットに通知できます。今回実装したシステムでは、以下のイベントフローを構築しました：

1. ユーザーが商品画像をアップロードする
2. 画像がS3バケットの`uploads/`プレフィックスに保存される
3. S3は`ObjectCreated`イベントを検出し、Lambda関数をトリガーする
4. Lambda関数が画像を処理（リサイズ）し、異なるサイズのバージョンを別のプレフィックスに保存
5. バックエンドAPIは、元の画像と各リサイズバージョンへのURLを提供

この仕組みにより、画像のリサイズ処理をメインのアプリケーションから分離し、サーバーレス関数として実行することで、アプリケーションのスケーラビリティと効率性を向上させています。

## 1.7. 【補足情報】

### 1.7.1. Lambda関数の最適化テクニック

AWS Lambda関数のパフォーマンスを最適化するためのテクニックをいくつか紹介します：

1. **コールドスタートの最小化**:
   - 関数が頻繁に呼び出されるようにする（Provisioned Concurrency、ウォームアップリクエスト）
   - 軽量な言語やランタイムを選択する（Go, Rustなど）
   - 依存関係を最小限に抑える
   - パッケージサイズを小さくする

2. **メモリ割り当ての最適化**:
   - メモリ割り当てを増やすとCPUパワーも増え、処理時間が短縮される
   - ベンチマークテストを行い、コストとパフォーマンスのバランスを見つける

3. **初期化コードの最適化**:
   - ハンドラー関数の外でリソースを初期化する（DB接続、SDKクライアントなど）
   - グローバル変数を適切に活用する

4. **並行処理の活用**:
   - Go言語のgoroutineなどを使用して並行処理を実装
   - 画像処理など、独立した作業は並列に行う

5. **Lambda Layers の活用**:
   - 共通ライブラリやフレームワークをLayersとして共有
   - デプロイパッケージのサイズを削減

今回の実装では、Go言語を使用することでコールドスタートの影響を最小限に抑え、グローバル変数でS3クライアントを初期化するなどの最適化を行っています。

### 1.7.2. 本番環境での画像処理の考慮点

実際の本番環境で画像処理システムを構築する際には、以下の点を考慮すると良いでしょう：

1. **セキュリティ**:
   - ファイルのバリデーション（サイズ制限、形式チェック、マルウェアスキャンなど）
   - 適切なIAMポリシーの設定
   - S3バケットへのアクセス制限
   - リクエスト元の認証・認可

2. **エラー処理とリトライ**:
   - 一時的な障害に対するリトライメカニズム
   - デッドレターキューの設定
   - 処理失敗の通知とログ記録

3. **オブザーバビリティ**:
   - 詳細なロギング
   - メトリクスの収集（処理時間、エラー率など）
   - トレーシングの実装
   - アラートの設定

4. **スケーラビリティ**:
   - 大量のアップロードに対応するための設計
   - ボトルネックの特定と解消
   - コスト管理（リザーブドコンカレンシーの活用など）

5. **コンテンツ配信**:
   - CloudFrontなどのCDNの活用
   - 適切なキャッシュ設定
   - オリジンアクセスアイデンティティによるS3バケット保護

6. **バックアップと障害復旧**:
   - クロスリージョンレプリケーション
   - バージョニングの有効化
   - 定期的なバックアップ

これらの点を考慮することで、より堅牢で安全、かつスケーラブルな画像処理システムを構築することができます。

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: Lambda関数が起動しない

**症状**: S3にファイルをアップロードしても、Lambda関数が起動せず、リサイズ画像が生成されない。

**解決策**:

1. S3イベント通知が正しく設定されているか確認する：

   ```bash
   aws --endpoint-url=http://localhost:4566 s3api get-bucket-notification-configuration --bucket product-images
   ```

2. Lambda関数が正しくデプロイされているか確認する：

   ```bash
   aws --endpoint-url=http://localhost:4566 lambda list-functions
   ```

3. Lambda関数にS3からの呼び出し権限が付与されているか確認する：

   ```bash
   aws --endpoint-url=http://localhost:4566 lambda get-policy --function-name image-resize-function
   ```

4. LocalStackのログを確認して、エラーメッセージがないかチェックする：

   ```bash
   docker-compose logs -f localstack
   ```

5. 正しいプレフィックス（`uploads/`）を使用してファイルをアップロードしていることを確認する

### 1.8.2. 問題2: 画像処理が失敗する

**症状**: Lambda関数は起動するが、画像処理が失敗し、リサイズ画像が生成されない。

**解決策**:

1. Lambda関数のログを確認して、エラーメッセージを特定する：

   ```bash
   aws --endpoint-url=http://localhost:4566 logs describe-log-groups
   aws --endpoint-url=http://localhost:4566 logs describe-log-streams --log-group-name /aws/lambda/image-resize-function
   aws --endpoint-url=http://localhost:4566 logs get-log-events --log-group-name /aws/lambda/image-resize-function --log-stream-name <stream-name>
   ```

2. サポートされていない画像形式ではないことを確認する（現在の実装では、JPEGとPNGのみをサポート）：

   ```bash
   file your-image-file.xxx
   ```

3. 画像が破損していないか確認する：

   ```bash
   identify your-image-file.jpg  # ImageMagickが必要
   ```

4. Go言語の画像処理ライブラリが正しくインストールされているか確認する：

   ```bash
   go list -m golang.org/x/image
   ```

5. Lambda関数に十分なメモリとタイムアウト設定があることを確認する：

   ```bash
   aws --endpoint-url=http://localhost:4566 lambda get-function-configuration --function-name image-resize-function
   ```

### 1.8.3. 問題3: S3イベントが正しくトリガーされない

**症状**: S3にファイルをアップロードしても、イベントが正しくトリガーされない。

**解決策**:

1. S3通知設定の構文を確認する（JSONの形式が正しいか）

2. S3バケットの名前とLambda関数のARNが正しく設定されているか確認する：

   ```bash
   aws --endpoint-url=http://localhost:4566 s3api get-bucket-notification-configuration --bucket product-images
   ```

3. LocalStackのバージョンが最新であることを確認する：

   ```bash
   docker-compose exec localstack localstack --version
   ```

4. S3イベントフィルタールールが期待通りに設定されているか確認する（プレフィックス、サフィックスなど）

5. 手動でLambda関数を呼び出してみて、関数自体が正しく動作するか確認する：

   ```bash
   aws --endpoint-url=http://localhost:4566 lambda invoke --function-name image-resize-function --payload '{"Records":[{"s3":{"bucket":{"name":"product-images"},"object":{"key":"uploads/test.jpg"}}}]}' response.json
   ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **サーバーレスアーキテクチャの基本概念**
   サーバーレスアーキテクチャは、インフラ管理の負担を減らし、イベント駆動型の効率的なアプリケーション開発を可能にします。Lambda関数は短命で、特定のイベントに反応して起動し、実行後に自動的に解放されます。

2. **イベント駆動型設計パターン**
   S3にファイルがアップロードされるというイベントをトリガーにLambda関数を呼び出す設計パターンは、システムを疎結合に保ち、スケーラビリティを向上させます。このパターンはマイクロサービスアーキテクチャとも相性が良く、複雑なワークフローを柔軟に構築できます。

3. **Go言語でのLambda関数実装**
   Go言語はLambda関数の実装に適しており、静的型付け、高速な起動時間、および優れたパフォーマンス特性を提供します。provided.al2ランタイムを使用することで、より高度なカスタマイズが可能になります。

4. **LocalStackによるAWSエミュレーション**
   LocalStackを使用することで、AWSの本番環境を使わずに、ローカル開発環境でLambda、S3などのAWSサービスを効果的にエミュレートできます。これにより、開発コストを抑えつつ、迅速な開発サイクルを実現できます。

5. **バックエンドとサーバーレス関数の連携**
   メインバックエンドとサーバーレス関数の連携は、システム全体のアーキテクチャにおいて重要な設計決定です。処理負荷の高いタスク（画像処理など）をサーバーレスに切り出すことで、メインアプリケーションのスケーラビリティと応答性を向上させることができます。

これらのポイントは、後続のフェーズでより複雑なサーバーレス機能を実装する際の基礎となります。

## 1.10. 【次回の準備】

次回（Day 4）では、高度なバリデーション実装とAPI品質向上に取り組みます。以下の点について事前に確認しておくと良いでしょう：

1. OpenAPI仕様を改めて確認し、バリデーションが必要なエンドポイントを特定する
2. go-playground/validatorパッケージの基本的な使い方を理解する
3. Echo Webフレームワークのバリデーション機能について調査する
4. 適切なHTTPステータスコードとエラーレスポンス形式について考える
5. 必要に応じて、AWS SDKのエラーハンドリングについても復習しておく

また、今日実装したLambda関数とS3連携が正しく動作していることを確認し、問題があれば修正しておきましょう。

## 1.11. 【.envrc サンプル】

以下は本日の実装で使用する.envrcのサンプルです。ご自身の環境に合わせて修正して使用してください。このファイルはgitにコミットしないようにしてください。

```bash
# .envrc サンプル
# LocalStack用の設定
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=us-east-1
export AWS_ENDPOINT_URL=http://localhost:4566

# バケット名
export S3_BUCKET=product-images

# Lambda関数名
export LAMBDA_FUNCTION_NAME=image-resize-function

# ローカル開発用の設定
export GO111MODULE=on
export GOFLAGS=-mod=vendor

# デバッグ設定
export DEBUG=true
```
