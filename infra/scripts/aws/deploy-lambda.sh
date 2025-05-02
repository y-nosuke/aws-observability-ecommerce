#!/bin/bash

set -e

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
