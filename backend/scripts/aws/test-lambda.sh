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
