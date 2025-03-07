#!/bin/bash
set -e

# パラメータのチェック
if [ -z "$1" ]; then
    echo "使用方法: $0 <S3バケット名>"
    exit 1
fi

S3_BUCKET=$1

# 現在のディレクトリを保存
CURRENT_DIR=$(pwd)

# フロントエンドディレクトリに移動
cd "$CURRENT_DIR/frontend"

# 必要な依存関係がインストールされていることを確認
if [ ! -d "node_modules" ]; then
    echo "依存関係をインストールしています..."
    npm ci
fi

# フロントエンドをビルド
echo "フロントエンドをビルドしています..."
npm run build

# S3バケットへデプロイ
echo "S3バケット ${S3_BUCKET} にデプロイしています..."
aws s3 sync out/ "s3://${S3_BUCKET}/" --delete

echo "フロントエンドのデプロイが完了しました。"

# CloudFrontのキャッシュを無効化（オプション）
if [ ! -z "$2" ]; then
    CLOUDFRONT_ID=$2
    echo "CloudFrontのキャッシュを無効化しています (Distribution ID: ${CLOUDFRONT_ID})..."
    aws cloudfront create-invalidation --distribution-id "${CLOUDFRONT_ID}" --paths "/*"
    echo "CloudFrontのキャッシュ無効化を開始しました。"
fi

# 元のディレクトリに戻る
cd "$CURRENT_DIR"
