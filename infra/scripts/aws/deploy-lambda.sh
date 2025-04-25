#!/bin/bash

set -e

# LocalStackのエンドポイントを設定
export AWS_ENDPOINT_URL=http://localhost:4566

# Lambda関数名
FUNCTION_NAME="image-processor"
HANDLER_DIR="cmd/lambda/${FUNCTION_NAME}"
OUTPUT_DIR="build/lambda"
ZIP_FILE="${OUTPUT_DIR}/${FUNCTION_NAME}.zip"

# ビルド用ディレクトリの作成
mkdir -p "${OUTPUT_DIR}"

echo "Lambda関数 ${FUNCTION_NAME} をAmazon Linux 2互換環境でビルドします..."

# DockerでAmazon Linux 2ベースのGoイメージを使ってバイナリをビルド
docker run --rm -v "$PWD":/go/src/app -w /go/src/app \
  golang:1.24 \
  /bin/bash -c "CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w' -o ${OUTPUT_DIR}/${FUNCTION_NAME} ./${HANDLER_DIR}/main.go"


# バイナリをZIP化
cd "${OUTPUT_DIR}" && zip -j "${FUNCTION_NAME}.zip" "${FUNCTION_NAME}" && cd -

echo "Lambda関数のビルドとZIP化が完了しました: ${ZIP_FILE}"

# Lambda関数が既に存在するか確認
if awslocal lambda get-function --function-name "${FUNCTION_NAME}" 2>/dev/null; then
  echo "Lambda関数 ${FUNCTION_NAME} を更新します..."

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
    --timeout 30 \
    --memory-size 512 \
    --role "arn:aws:iam::000000000000:role/lambda-${FUNCTION_NAME}-role" \
    --zip-file "fileb://${ZIP_FILE}" \
    --environment "Variables={BUCKET_NAME=product-images}"

  echo "Lambda関数 ${FUNCTION_NAME} を作成しました。"
fi

# 関数の設定を表示
echo "Lambda関数の設定:"
awslocal lambda get-function --function-name "${FUNCTION_NAME}"
