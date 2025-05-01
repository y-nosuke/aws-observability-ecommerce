# 1. Week 2 - Day 4: サーバーレスアーキテクチャの基本

## 1.1. 目次

- [1. Week 2 - Day 4: サーバーレスアーキテクチャの基本](#1-week-2---day-4-サーバーレスアーキテクチャの基本)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. LocalStackの確認と設定](#141-localstackの確認と設定)
    - [1.4.2. S3バケットの作成と設定](#142-s3バケットの作成と設定)
    - [1.4.3. 基本的なLambda関数の作成](#143-基本的なlambda関数の作成)
    - [1.4.4. Lambda関数のデプロイとテスト](#144-lambda関数のデプロイとテスト)
    - [1.4.5. S3トリガーの設定](#145-s3トリガーの設定)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. サーバーレスアーキテクチャの概念](#161-サーバーレスアーキテクチャの概念)
    - [1.6.2. AWS Lambdaの仕組みと特徴](#162-aws-lambdaの仕組みと特徴)
    - [1.6.3. イベント駆動型設計の基本概念](#163-イベント駆動型設計の基本概念)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. サーバーレスの利点と制限](#171-サーバーレスの利点と制限)
    - [1.7.2. AWS Lambda以外のサーバーレスサービス](#172-aws-lambda以外のサーバーレスサービス)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: LocalStackが正しく起動しない](#181-問題1-localstackが正しく起動しない)
    - [1.8.2. 問題2: Lambda関数が正しく実行されない](#182-問題2-lambda関数が正しく実行されない)
    - [1.8.3. 問題3: S3トリガーが動作しない](#183-問題3-s3トリガーが動作しない)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)

## 1.2. 【要点】

- サーバーレスアーキテクチャの基本概念と従来型アーキテクチャとの違いを理解する
- LocalStackを用いてAWS Lambda環境をローカルで構築する方法を習得する
- AWS Lambda関数の作成、デプロイ、実行の基本的な流れを学ぶ
- S3バケットとLambda関数の連携によるイベント駆動型処理のパターンを理解する
- AWS CLIとawslocalを使ったローカル環境でのAWSサービス操作を習得する

## 1.3. 【準備】

本日の実習を始める前に、以下の環境とツールが正しく設定されていることを確認してください。

### 1.3.1. チェックリスト

- [ ] Docker Composeで作成した開発環境が起動している
- [ ] localstackコンテナが正常に動作している
- [ ] AWS CLI v2がインストールされている
- [ ] Python 3.8以上がインストールされている
- [ ] awslocalがインストールされている（LocalStack用のAWS CLIラッパー）
- [ ] LocalStackのエントリポイントが環境変数に設定されている

AWS CLIがインストールされていない場合は、[AWS CLIの公式ドキュメント](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)を参照してインストールしてください。

awslocalがインストールされていない場合は、以下のコマンドでインストールします：

```bash
pip install awscli-local
```

## 1.4. 【手順】

### 1.4.1. LocalStackの確認と設定

まず、LocalStackが正常に動作していることを確認し、必要な環境変数を設定します。

```bash
# LocalStackコンテナの状態を確認
docker ps | grep localstack

# awslocal経由でサービスリストをテスト
awslocal s3 ls
awslocal lambda list-functions
```

出力例：

```bash
CONTAINER ID   IMAGE                  COMMAND                  CREATED        STATUS          PORTS                                         NAMES
a1b2c3d4e5f6   localstack/localstack  "docker-entrypoint.sh"   2 hours ago    Up 2 hours      0.0.0.0:4566-4599->4566-4599/tcp, 5678/tcp   project-localstack-1
```

エラーが発生せず、コマンドが実行できれば準備完了です。まだS3バケットやLambda関数は作成していないため、リスト表示は空になります。

### 1.4.2. S3バケットの作成と設定

LocalStack上にS3バケットを作成します。このバケットは、画像ファイルの保存とLambda関数のトリガーに使用します。

```bash
# バケットを作成するディレクトリを用意
mkdir -p backend/scripts/aws

# バケット作成スクリプトを作成
touch backend/scripts/aws/create-s3-bucket.sh
chmod +x backend/scripts/aws/create-s3-bucket.sh
```

`backend/scripts/aws/create-s3-bucket.sh`の内容：

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
  awslocal s3api put-bucket-cors --bucket ${BUCKET_NAME} --cors-configuration file:///tmp/cors-config.json
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

スクリプトを実行してバケットを作成します：

```bash
./backend/scripts/aws/create-s3-bucket.sh
```

出力例：

```bash
バケット product-images を作成しました。
バケット product-images にCORS設定を適用しました。
バケット product-images にパブリックアクセスポリシーを適用しました。
現在のS3バケット一覧:
2023-04-05 12:34:56 product-images
```

### 1.4.3. 基本的なLambda関数の作成

画像リサイズを行う基本的なLambda関数を作成します。Go言語で実装し、ビルドして配置するデプロイスクリプトも作成します。

AWS Lambda用の依存関係をGo言語のプロジェクトに追加します：

```bash
# Lambdaの依存関係を追加
cd backend
go get github.com/aws/aws-lambda-go/lambda
go get github.com/aws/aws-lambda-go/events
go get github.com/aws/aws-sdk-go/aws
go get github.com/aws/aws-sdk-go/service/s3
cd ..
```

まず、ディレクトリとファイルを作成します：

```bash
# Lambda関数のディレクトリを作成
mkdir -p backend/cmd/lambda/image-processor

# Lambda関数のコードを作成
touch backend/cmd/lambda/image-processor/main.go
```

`backend/cmd/lambda/image-processor/main.go`の内容：

```go
package main

import (
 "bytes"
 "context"
 "encoding/json"
 "fmt"
 "log"
 "strings"

 "github.com/aws/aws-lambda-go/events"
 "github.com/aws/aws-lambda-go/lambda"
 "github.com/aws/aws-sdk-go/aws"
 "github.com/aws/aws-sdk-go/aws/session"
 "github.com/aws/aws-sdk-go/service/s3"
)

// S3Event is the structure of S3 event notifications
type S3Event struct {
 Records []struct {
  S3 struct {
   Bucket struct {
    Name string `json:"name"`
   } `json:"bucket"`
   Object struct {
    Key string `json:"key"`
   } `json:"object"`
  } `json:"s3"`
 } `json:"Records"`
}

func handler(ctx context.Context, s3Event events.S3Event) (string, error) {
 // LocalStackのエンドポイントを設定
 awsEndpoint := "http://localstack:4566"

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

  // AWSセッションの作成
  sess, err := session.NewSession(&aws.Config{
   Endpoint:         aws.String(awsEndpoint),
   Region:           aws.String("us-east-1"),
   S3ForcePathStyle: aws.Bool(true),
  })
  if err != nil {
   log.Printf("セッション作成エラー: %v", err)
   return "", err
  }

  // S3クライアントの作成
  s3Client := s3.New(sess)

  // オリジナル画像の取得
  getResult, err := s3Client.GetObject(&s3.GetObjectInput{
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
  buf.ReadFrom(getResult.Body)
  processedData := buf.Bytes()

  // 処理済み画像の新しいパス
  newKey := fmt.Sprintf("processed/%s", key)

  // 処理済み画像をアップロード
  _, err = s3Client.PutObject(&s3.PutObjectInput{
   Bucket:      aws.String(bucket),
   Key:         aws.String(newKey),
   Body:        bytes.NewReader(processedData),
   ContentType: getResult.ContentType,
   Metadata: map[string]*string{
    "ProcessedBy": aws.String("ImageProcessorLambda"),
    "OriginalKey": aws.String(key),
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
```

次に、Lambda関数をビルドしデプロイするスクリプトを作成します：

```bash
# デプロイスクリプトの作成
touch backend/scripts/aws/deploy-lambda.sh
chmod +x backend/scripts/aws/deploy-lambda.sh
```

`backend/scripts/aws/deploy-lambda.sh`の内容：

```bash
#!/bin/bash

# LocalStackのエンドポイントを設定
export AWS_ENDPOINT_URL=http://localhost:4566

# Lambda関数名
FUNCTION_NAME="image-processor"
HANDLER_DIR="backend/cmd/lambda/${FUNCTION_NAME}"
OUTPUT_DIR="backend/build/lambda"
ZIP_FILE="${OUTPUT_DIR}/${FUNCTION_NAME}.zip"

# ビルド用ディレクトリの作成
mkdir -p "${OUTPUT_DIR}"

echo "Lambda関数 ${FUNCTION_NAME} のビルドを開始します..."

# Go言語でLambda関数をビルド
GOOS=linux GOARCH=amd64 go build -o "${OUTPUT_DIR}/${FUNCTION_NAME}" "${HANDLER_DIR}/main.go"

# ビルドしたバイナリをZIPに圧縮
cd "${OUTPUT_DIR}" && zip -j "${FUNCTION_NAME}.zip" "${FUNCTION_NAME}" && cd -

echo "Lambda関数のビルドとZIP化が完了しました: ${ZIP_FILE}"

# Lambda関数が既に存在するか確認
if awslocal lambda get-function --function-name "${FUNCTION_NAME}" 2>/dev/null; then
  echo "Lambda関数 ${FUNCTION_NAME} を更新します..."

  # 関数の更新
  awslocal lambda update-function-code \
    --function-name "${FUNCTION_NAME}" \
    --zip-file "fileb://${ZIP_FILE}"

  echo "Lambda関数 ${FUNCTION_NAME} を更新しました。"
else
  echo "Lambda関数 ${FUNCTION_NAME} を新規作成します..."

  # 実行ロールの作成（LocalStackでは実際には必要ありませんが、AWS環境に近づけるため実施）
  awslocal iam create-role \
    --role-name "lambda-${FUNCTION_NAME}-role" \
    --assume-role-policy-document '{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":{"Service":"lambda.amazonaws.com"},"Action":"sts:AssumeRole"}]}'

  # S3アクセスポリシーをロールにアタッチ
  awslocal iam attach-role-policy \
    --role-name "lambda-${FUNCTION_NAME}-role" \
    --policy-arn "arn:aws:iam::aws:policy/AmazonS3FullAccess"

  # Lambda関数の作成
  awslocal lambda create-function \
    --function-name "${FUNCTION_NAME}" \
    --runtime "go1.x" \
    --handler "${FUNCTION_NAME}" \
    --role "arn:aws:iam::000000000000:role/lambda-${FUNCTION_NAME}-role" \
    --zip-file "fileb://${ZIP_FILE}" \
    --environment "Variables={BUCKET_NAME=product-images}"

  echo "Lambda関数 ${FUNCTION_NAME} を作成しました。"
fi

# 関数の設定を確認
echo "Lambda関数の設定:"
awslocal lambda get-function --function-name "${FUNCTION_NAME}"
```

### 1.4.4. Lambda関数のデプロイとテスト

作成したスクリプトを実行してLambda関数をビルドし、LocalStackにデプロイします：

```bash
cd backend
./scripts/aws/deploy-lambda.sh
```

出力例：

```bash
Lambda関数 image-processor のビルドを開始します...
Lambda関数のビルドとZIP化が完了しました: backend/build/lambda/image-processor.zip
Lambda関数 image-processor を新規作成します...
Lambda関数 image-processor を作成しました。
Lambda関数の設定:
{
    "Configuration": {
        "FunctionName": "image-processor",
        "FunctionArn": "arn:aws:lambda:us-east-1:000000000000:function:image-processor",
        "Runtime": "go1.x",
        "Role": "arn:aws:iam::000000000000:role/lambda-image-processor-role",
        "Handler": "image-processor",
        "Environment": {
            "Variables": {
                "BUCKET_NAME": "product-images"
            }
        }
    }
}
```

続いて、デプロイしたLambda関数をテスト呼び出しするスクリプトを作成します：

```bash
touch backend/scripts/aws/test-lambda.sh
chmod +x backend/scripts/aws/test-lambda.sh
```

`backend/scripts/aws/test-lambda.sh`の内容：

```bash
#!/bin/bash

# LocalStackのエンドポイントを設定
export AWS_ENDPOINT_URL=http://localhost:4566

# Lambda関数名
FUNCTION_NAME="image-processor"

# テスト用のペイロードを作成
cat > /tmp/lambda-test-payload.json << EOF
{
  "Records": [
    {
      "s3": {
        "bucket": {
          "name": "product-images"
        },
        "object": {
          "key": "test-image.jpg"
        }
      }
    }
  ]
}
EOF

# テスト用の画像をS3にアップロード
echo "テスト用の画像をS3にアップロード中..."
echo "これはテスト画像のダミーコンテンツです" > /tmp/test-image.jpg
awslocal s3 cp /tmp/test-image.jpg s3://product-images/test-image.jpg

# Lambda関数の実行
echo "Lambda関数 ${FUNCTION_NAME} をテスト実行します..."
awslocal lambda invoke \
  --function-name "${FUNCTION_NAME}" \
  --payload file:///tmp/lambda-test-payload.json \
  --cli-binary-format raw-in-base64-out \
  /tmp/lambda-response.json

# 実行結果の表示
echo "Lambda関数の実行結果:"
cat /tmp/lambda-response.json

# 処理済み画像の確認
echo "処理済み画像を確認します:"
awslocal s3 ls s3://product-images/processed/
```

スクリプトを実行してLambda関数をテストします：

```bash
./backend/scripts/aws/test-lambda.sh
```

出力例：

```bash
テスト用の画像をS3にアップロード中...
upload: /tmp/test-image.jpg to s3://product-images/test-image.jpg
Lambda関数 image-processor をテスト実行します...
{
    "StatusCode": 200,
    "ExecutedVersion": "$LATEST"
}
Lambda関数の実行結果:
"画像処理が完了しました"
処理済み画像を確認します:
2023-04-05 12:45:56         46 test-image.jpg
```

### 1.4.5. S3トリガーの設定

S3バケットに新しい画像がアップロードされたときに自動的にLambda関数が実行されるように、S3トリガーを設定します。

```bash
touch backend/scripts/aws/setup-s3-trigger.sh
chmod +x backend/scripts/aws/setup-s3-trigger.sh
```

`backend/scripts/aws/setup-s3-trigger.sh`の内容：

```bash
#!/bin/bash

# LocalStackのエンドポイントを設定
export AWS_ENDPOINT_URL=http://localhost:4566

# S3バケット名とLambda関数名
BUCKET_NAME="product-images"
FUNCTION_NAME="image-processor"

# Lambda関数のARNを取得
LAMBDA_ARN=$(awslocal lambda get-function --function-name "${FUNCTION_NAME}" --query 'Configuration.FunctionArn' --output text)

echo "Lambda関数ARN: ${LAMBDA_ARN}"

# Lambda関数にS3のイベント通知権限を付与
echo "Lambda関数に S3 バケットからのイベント通知権限を付与します..."
awslocal lambda add-permission \
  --function-name "${FUNCTION_NAME}" \
  --statement-id "s3-trigger" \
  --action "lambda:InvokeFunction" \
  --principal "s3.amazonaws.com" \
  --source-arn "arn:aws:s3:::${BUCKET_NAME}"

# S3バケットにLambdaトリガーを設定
echo "S3バケット ${BUCKET_NAME} にLambdaトリガーを設定します..."
cat > /tmp/s3-notification-config.json << EOF
{
  "LambdaFunctionConfigurations": [
    {
      "LambdaFunctionArn": "${LAMBDA_ARN}",
      "Events": ["s3:ObjectCreated:*"],
      "Filter": {
        "Key": {
          "FilterRules": [
            {
              "Name": "prefix",
              "Value": ""
            },
            {
              "Name": "suffix",
              "Value": ".jpg"
            }
          ]
        }
      }
    }
  ]
}
EOF

awslocal s3api put-bucket-notification-configuration \
  --bucket "${BUCKET_NAME}" \
  --notification-configuration file:///tmp/s3-notification-config.json

echo "S3トリガーの設定が完了しました。"

# トリガーテスト用の画像アップロード
echo "トリガーテスト用の新しい画像をアップロードします..."
echo "これはトリガーテスト用の画像です" > /tmp/trigger-test.jpg
awslocal s3 cp /tmp/trigger-test.jpg s3://product-images/trigger-test.jpg

echo "アップロード完了。数秒後に処理済みフォルダを確認してください。"
echo "処理済み画像の確認コマンド: awslocal s3 ls s3://product-images/processed/"
```

スクリプトを実行してS3トリガーを設定し、テストします：

```bash
./backend/scripts/aws/setup-s3-trigger.sh
```

出力例：

```bash
Lambda関数ARN: arn:aws:lambda:us-east-1:000000000000:function:image-processor
Lambda関数に S3 バケットからのイベント通知権限を付与します...
S3バケット product-images にLambdaトリガーを設定します...
S3トリガーの設定が完了しました。
トリガーテスト用の新しい画像をアップロードします...
upload: /tmp/trigger-test.jpg to s3://product-images/trigger-test.jpg
アップロード完了。数秒後に処理済みフォルダを確認してください。
処理済み画像の確認コマンド: awslocal s3 ls s3://product-images/processed/
```

トリガーが正しく動作したことを確認します：

```bash
# 数秒待ってから確認
sleep 5
awslocal s3 ls s3://product-images/processed/
```

出力例：

```bash
2023-04-05 12:50:56         32 trigger-test.jpg
```

## 1.5. 【確認ポイント】

以下の点を確認して、実装が正しく完了していることを確認してください：

- [ ] LocalStackが正常に動作し、AWS CLIコマンドが実行できる
- [ ] S3バケット「product-images」が正常に作成されている
- [ ] Lambda関数「image-processor」が正常にビルドされデプロイされている
- [ ] Lambda関数をテスト実行でき、正しい結果が返ってくる
- [ ] S3トリガーが設定され、新しい画像アップロード時にLambda関数が自動実行される
- [ ] 処理済み画像が「processed/」プレフィックス付きで保存される
- [ ] すべてのスクリプトがエラーなく実行できる

## 1.6. 【詳細解説】

### 1.6.1. サーバーレスアーキテクチャの概念

**サーバーレスアーキテクチャ**とは、サーバーのプロビジョニングや管理を行わずにアプリケーションを構築・実行するためのクラウドコンピューティングの実行モデルです。「サーバーレス」という名前ですが、実際にはサーバーが存在しないわけではなく、開発者がサーバーの管理やスケーリングを意識する必要がないという意味です。

**従来のアーキテクチャとの主な違い**:

1. **インフラストラクチャ管理**:
   - 従来: サーバーのプロビジョニング、パッチ適用、スケーリングなどを手動で管理
   - サーバーレス: プロバイダー（AWSなど）がインフラストラクチャを管理

2. **スケーリング**:
   - 従来: 手動または自動スケーリングルールの設定が必要
   - サーバーレス: 自動的にスケールアップ/ダウン（リクエスト数に応じて）

3. **コスト**:
   - 従来: サーバーの稼働時間に基づく課金（アイドル状態でも課金）
   - サーバーレス: 実行時間とリソース使用量に基づく課金（使用した分だけ支払い）

4. **デプロイモデル**:
   - 従来: アプリケーション全体をデプロイ
   - サーバーレス: 機能単位（Function as a Service）でのデプロイが可能

**サーバーレスが特に適している用途**:

- バッチ処理
- イベント駆動型処理
- スケーリング要件が予測困難なアプリケーション
- マイクロサービスアーキテクチャ
- バックエンドAPI
- 定期的な処理タスク

### 1.6.2. AWS Lambdaの仕組みと特徴

**AWS Lambda**は、AWSが提供するサーバーレスコンピューティングサービスで、イベント発生時にコードを実行し、必要なコンピューティングリソースを自動的に管理します。

**主要コンポーネント**:

1. **Lambda関数**: コードとその依存関係をパッケージ化したもの
2. **イベントソース**: Lambda関数を呼び出すトリガー（例: S3、API Gateway、DynamoDB）
3. **ランタイム環境**: コードが実行される環境（Node.js、Python、Go、Javaなど）
4. **実行ロール**: 関数に割り当てられたIAMロール（他のAWSサービスへのアクセス権を定義）

**Lambdaの実行フロー**:

1. イベントがトリガーされる（例: S3バケットへのファイルアップロード）
2. AWSがLambda実行環境を準備（コールドスタート）または再利用
3. イベントデータがハンドラー関数に渡される
4. 関数が実行され、結果を返す
5. 実行環境はしばらくの間維持され、再利用される可能性がある

**重要な特徴**:

- **ステートレス**: 関数は実行中のみ状態を保持
- **タイムアウト**: デフォルトは3秒、最大15分まで設定可能
- **メモリ割り当て**: 128MB〜10GBまで設定可能
- **コールドスタート**: 初回実行時またはスケールアップ時に発生する遅延
- **同時実行制限**: アカウントごとに同時実行数の制限あり

### 1.6.3. イベント駆動型設計の基本概念

**イベント駆動型アーキテクチャ**は、システムコンポーネント間の通信がイベントの生成と処理を通じて行われる設計パターンです。

**主要な概念**:

1. **イベント**: システム内で発生した重要な変化や行動（例: ファイルのアップロード、ユーザーの登録）
2. **イベント生成者(Producer)**: イベントを発生させるコンポーネント
3. **イベントコンシューマー(Consumer)**: イベントに反応して処理を行うコンポーネント
4. **イベントバス/ブローカー**: イベントの配信を管理するミドルウェア

**設計パターン**:

1. **パブリッシュ/サブスクライブ (Pub/Sub)**:
   - 生成者がイベントをパブリッシュし、関心のあるコンシューマーがそれを購読
   - 生成者とコンシューマーが疎結合になる

2. **イベントソーシング**:
   - システムの状態変化をイベントとして保存
   - 状態はイベントの再生により再構築可能

3. **CQRS (Command Query Responsibility Segregation)**:
   - 書き込み操作（コマンド）と読み取り操作（クエリ）を分離
   - イベントを用いてデータの整合性を維持

**利点**:

- システムコンポーネントの疎結合
- スケーラビリティと回復力の向上
- 非同期処理による応答性の向上
- ビジネスプロセスの明確な表現

**AWS環境でのイベント駆動型サービス**:

- **Amazon S3**: オブジェクト操作でのイベント通知
- **Amazon SNS (Simple Notification Service)**: パブリッシュ/サブスクライブメッセージング
- **Amazon SQS (Simple Queue Service)**: メッセージキュー
- **Amazon EventBridge**: イベントバスとルーティング
- **DynamoDB Streams**: テーブル変更のイベントストリーム

今回の実装では、S3バケットに画像がアップロードされると（イベント）、Lambda関数（コンシューマ）が自動的に起動して処理を行う、シンプルなイベント駆動型パターンを実装しています。

## 1.7. 【補足情報】

### 1.7.1. サーバーレスの利点と制限

**利点**:

1. **運用の負担軽減**:
   - サーバー管理が不要
   - パッチ適用やOSアップデートの心配なし
   - インフラストラクチャの可用性はプロバイダーが保証

2. **コスト最適化**:
   - 使用したリソースに対してのみ課金
   - アイドル時間のコストがない
   - リソースの事前確保が不要

3. **スケーラビリティ**:
   - 自動的にスケールアップ/ダウン
   - トラフィック増加時にも対応可能
   - ピーク時の容量計画が不要

4. **迅速な開発とデプロイ**:
   - 小さな機能単位でのデプロイが可能
   - 継続的なデプロイの実現が容易
   - テスト環境の構築が簡素化

**制限と課題**:

1. **コールドスタート**:
   - 初回実行時の遅延
   - 低頻度の呼び出しで顕著
   - ウォームアップ技術で一部緩和可能

2. **実行時間の制限**:
   - Lambda: 最大15分
   - 長時間実行タスクには不向き

3. **ステートレス性**:
   - 状態の保持が困難
   - 外部ストレージとの連携が必要

4. **ベンダーロックイン**:
   - プロバイダー固有のサービスへの依存
   - 移行コストの増加リスク

5. **デバッグとモニタリングの複雑さ**:
   - 分散システムのデバッグが困難
   - オブザーバビリティの課題

6. **リソース制限**:
   - メモリ、ディスク容量、同時実行数の制限
   - 大規模処理には不向きな場合も

### 1.7.2. AWS Lambda以外のサーバーレスサービス

AWS Lambda以外にも、様々なサーバーレスサービスがあります：

**AWS内のサーバーレスサービス**:

1. **Amazon API Gateway**:
   - RESTful APIの作成と管理
   - Lambdaと組み合わせてサーバーレスAPIを構築
   - 認証、スロットリング、モニタリング機能

2. **AWS Step Functions**:
   - サーバーレスワークフローオーケストレーション
   - 複数のLambda関数を連携させる
   - 状態管理と並列処理

3. **Amazon DynamoDB**:
   - サーバーレスNoSQLデータベース
   - 自動スケーリング
   - オンデマンド容量モード

4. **AWS AppSync**:
   - サーバーレスGraphQL API
   - リアルタイムデータ同期
   - オフライン操作サポート

5. **Amazon EventBridge (CloudWatch Events)**:
   - サーバーレスイベントバス
   - スケジュールベースの実行
   - イベントパターンマッチング

6. **AWS Fargate**:
   - サーバーレスコンテナオーケストレーション
   - ECS/EKSと統合
   - サーバー管理なしでコンテナ実行

**他のクラウドプロバイダーのサーバーレスサービス**:

1. **Google Cloud Functions**:
   - GCPのFaaS
   - HTTPリクエストやクラウドイベントで実行

2. **Azure Functions**:
   - MicrosoftのFaaS
   - トリガーと連携機能が豊富

3. **Firebase Cloud Functions**:
   - モバイル/ウェブ向けサーバーレス
   - Firebaseサービスとの緊密な統合

4. **Cloudflare Workers**:
   - エッジコンピューティングプラットフォーム
   - 世界中のエッジロケーションで実行

**オープンソースのサーバーレスフレームワーク**:

1. **Serverless Framework**:
   - マルチクラウド対応のデプロイツール
   - インフラをコードで定義

2. **OpenFaaS**:
   - Kubernetesでサーバーレス機能を実行
   - DockerとKubernetesの知識が活用可能

3. **Knative**:
   - Kubernetes上のサーバーレスプラットフォーム
   - イベント駆動型アーキテクチャのサポート

これらのサービスを理解し、適切に組み合わせることで、より強力なサーバーレスアプリケーションを構築できます。

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: LocalStackが正しく起動しない

**症状**:

- `awslocal`コマンドが`Could not connect to the endpoint URL`エラーを返す
- Docker Composeログに接続エラーが表示される

**解決策**:

1. LocalStackコンテナの状態を確認する：

    ```bash
    docker ps | grep localstack
    ```

2. LocalStackのログを確認する：

    ```bash
    docker logs $(docker ps -q --filter name=localstack)
    ```

3. LocalStackのポートが正しく公開されているか確認：

    ```bash
    netstat -tuln | grep 4566
    ```

4. コンテナを再起動する：

    ```bash
    docker-compose restart localstack
    ```

5. Docker Composeファイルでのポート設定を確認し、ポートの競合がないようにする

### 1.8.2. 問題2: Lambda関数が正しく実行されない

**症状**:

- Lambda関数の実行でエラーが発生する
- 関数は成功するが、期待した処理が行われない

**解決策**:

1. Lambda関数が正しくビルドされているか確認：

    ```bash
    cd backend/build/lambda
    unzip -l image-processor.zip
    ```

2. Lambda関数の実行権限を確認：

    ```bash
    awslocal lambda get-policy --function-name image-processor
    ```

3. Lambda関数のログを確認（LocalStackでは実際にCloudWatchに出力されないため、ローカルにログファイルを作成）：

    ```bash
    # ログを出力するための環境変数を設定
    export LAMBDA_LOGS=/tmp/lambda-logs.txt
    touch $LAMBDA_LOGS
    tail -f $LAMBDA_LOGS &

    # Lambda関数を実行（LOG_FILE環境変数を追加）
    awslocal lambda update-function-configuration \
      --function-name image-processor \
      --environment "Variables={LOG_FILE=$LAMBDA_LOGS,BUCKET_NAME=product-images}"

    # テスト実行
    ./scripts/aws/test-lambda.sh
    ```

4. Lambda関数のコードを再確認し、エラー処理とロギングを強化する

5. LocalStackとの接続情報が正しいか確認：
   - Lambda内でのAWSエンドポイント設定（`http://localstack:4566`）
   - リージョン設定（通常は`us-east-1`）
   - S3バケット名の一致

### 1.8.3. 問題3: S3トリガーが動作しない

**症状**:

- S3にファイルをアップロードしてもLambda関数が自動的に起動しない

**解決策**:

1. S3バケットの通知設定を確認：

    ```bash
    awslocal s3api get-bucket-notification-configuration --bucket product-images
    ```

2. Lambda関数のアクセス権限を確認：

    ```bash
    awslocal lambda get-policy --function-name image-processor
    ```

3. 権限を手動で追加し直す：

    ```bash
    # 古い権限を削除（上書きされない場合）
    awslocal lambda remove-permission \
      --function-name image-processor \
      --statement-id s3-trigger

    # 新しい権限を追加
    awslocal lambda add-permission \
      --function-name image-processor \
      --statement-id s3-trigger \
      --action lambda:InvokeFunction \
      --principal s3.amazonaws.com \
      --source-arn arn:aws:s3:::product-images
    ```

4. 通知設定を再度適用：

    ```bash
    ./scripts/aws/setup-s3-trigger.sh
    ```

5. Lambda関数とS3バケットの名前が一致していることを確認

6. LocalStackのバージョンが最新であることを確認（古いバージョンではS3トリガーが完全にサポートされていない場合がある）

## 1.9. 【今日の重要なポイント】

本日の実装を通じて学んだ特に重要なポイントは以下の通りです：

1. **サーバーレスアーキテクチャの理解**:
   - サーバー管理からの解放
   - イベント駆動型の処理パターン
   - リソース使用に基づく課金モデル

2. **AWS Lambda関数の実装パターン**:
   - ハンドラー関数の構造
   - イベントデータの処理方法
   - AWSサービスとの連携

3. **S3トリガーによるイベント処理**:
   - オブジェクト作成イベントの検出
   - Lambda関数の自動起動
   - 非同期処理フローの実現

4. **LocalStackを使ったローカル開発環境**:
   - AWSサービスのローカルエミュレーション
   - 開発コストの削減
   - 高速なフィードバックサイクル

これらの概念と実装パターンは、今後のフェーズで実装するより複雑なサーバーレスアプリケーションの基礎となります。今日学んだシンプルなイベント駆動型パターンは、拡張して様々なユースケースに適用できます。

## 1.10. 【次回の準備】

Week 2 - Day 5では、Reactコンポーネントの基礎に焦点を当てます。以下の点について事前に確認しておくとよいでしょう：

1. **Node.jsとnpmの環境**:
   - Node.js v14以上がインストールされていることを確認
   - npm（またはyarn）が正しく動作することを確認

2. **Reactの基本概念**:
   - コンポーネントベースのUI構築
   - JSXの構文
   - propsとstate

3. **TypeScriptの基本構文**:
   - 型定義
   - インターフェース
   - ジェネリック

4. **参考リソース**:
   - [React公式ドキュメント](https://reactjs.org/docs/getting-started.html)
   - [TypeScript公式ドキュメント](https://www.typescriptlang.org/docs/)
   - [React+TypeScriptチートシート](https://github.com/typescript-cheatsheets/react)

5. **環境確認**:
   - フロントエンドのコンテナが起動していることを確認
   - `http://shop.localhost`と`http://admin.localhost`にアクセスできることを確認

6. **コード確認**:
   - Next.jsプロジェクトの基本構造を確認
   - 既存のコンポーネントがあれば概要を把握

次回はこれらの知識を基に、実際にReactコンポーネントを作成し、状態管理やイベントハンドリングなどの基本概念を学びます。
