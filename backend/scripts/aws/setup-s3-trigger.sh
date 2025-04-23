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
# 権限がすでに存在するかチェック
PERMISSION_EXISTS=$(awslocal lambda get-policy --function-name "${FUNCTION_NAME}" 2>/dev/null | grep -c "s3-trigger" || echo "0")

if [ "$PERMISSION_EXISTS" -eq "0" ]; then
  # 権限が存在しない場合のみ追加
  awslocal lambda add-permission \
    --function-name "${FUNCTION_NAME}" \
    --statement-id "s3-trigger" \
    --action "lambda:InvokeFunction" \
    --principal "s3.amazonaws.com" \
    --source-arn "arn:aws:s3:::${BUCKET_NAME}"
  echo "Lambda権限を追加しました。"
else
  echo "Lambda権限はすでに設定されています。スキップします。"
fi

# S3バケットにLambdaトリガーを設定
echo "S3バケット ${BUCKET_NAME} にLambdaトリガーを設定します..."

# 新しい通知設定の作成
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
  --notification-configuration file:///tmp/s3-notification-config.json

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
    echo "これはトリガーテスト用の画像です" > /tmp/trigger-test.jpg
    awslocal s3 cp /tmp/trigger-test.jpg s3://product-images/trigger-test.jpg

    echo "アップロード完了。数秒後に処理済みフォルダを確認してください。"
    echo "処理済み画像の確認コマンド: awslocal s3 ls s3://product-images/processed/"
else
    echo "テスト画像のアップロードをスキップしました。"
    echo "手動でテストする場合は以下のコマンドを使用してください:"
    echo "awslocal s3 cp your-test-image.jpg s3://product-images/your-test-image.jpg"
    echo "処理済み画像の確認コマンド: awslocal s3 ls s3://product-images/processed/"
fi
