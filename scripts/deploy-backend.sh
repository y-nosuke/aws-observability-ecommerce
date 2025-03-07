#!/bin/bash
set -e

# パラメータのチェック
if [ -z "$1" ] || [ -z "$2" ]; then
    echo "使用方法: $0 <サービス名> <Lambda関数名>"
    exit 1
fi

SERVICE_NAME=$1
FUNCTION_NAME=$2

# 現在のディレクトリを保存
CURRENT_DIR=$(pwd)

# バックエンドサービスディレクトリに移動
cd "$CURRENT_DIR/backend/$SERVICE_NAME"

# 依存関係がインストールされていることを確認
if [ ! -d "node_modules" ]; then
    echo "依存関係をインストールしています..."
    npm ci
fi

# デプロイパッケージの作成
echo "デプロイパッケージを作成しています..."
mkdir -p "$CURRENT_DIR/dist"
zip -r "$CURRENT_DIR/dist/${SERVICE_NAME}.zip" .

# Lambda関数のコードを更新
echo "Lambda関数 ${FUNCTION_NAME} を更新しています..."
aws lambda update-function-code \
    --function-name "${FUNCTION_NAME}" \
    --zip-file "fileb://$CURRENT_DIR/dist/${SERVICE_NAME}.zip"

echo "バックエンドサービス ${SERVICE_NAME} のデプロイが完了しました。"

# 元のディレクトリに戻る
cd "$CURRENT_DIR"
