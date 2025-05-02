#!/bin/bash

set -e

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
