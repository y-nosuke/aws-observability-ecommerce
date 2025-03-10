#!/bin/bash

# LocalStackのヘルスチェックスクリプト
ENDPOINT_URL="http://localhost:4566"

echo "Checking LocalStack health..."

# ヘルスチェックエンドポイントを呼び出す
HEALTH_RESPONSE=$(curl -s $ENDPOINT_URL/_localstack/health)

if [ $? -ne 0 ]; then
  echo "Error: Failed to connect to LocalStack"
  exit 1
fi

# サービスの状態を確認
echo "Service status:"
echo "$HEALTH_RESPONSE" | jq .

# 必要なサービスが起動しているか確認
REQUIRED_SERVICES=("s3" "cloudwatch" "logs" "xray" "sqs" "sns")
ALL_UP=true

for SERVICE in "${REQUIRED_SERVICES[@]}"; do
  STATUS=$(echo "$HEALTH_RESPONSE" | jq -r --arg service "$SERVICE" '.services[$service].running')

  if [ "$STATUS" != "true" ]; then
    echo "Warning: $SERVICE is not running"
    ALL_UP=false
  else
    echo "OK: $SERVICE is running"
  fi
done

if [ "$ALL_UP" = true ]; then
  echo "All required services are running"
  exit 0
else
  echo "Some services are not running. Check the logs for more information."
  exit 1
fi
