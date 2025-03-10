#!/bin/bash

# LocalStackのリソースをテストするスクリプト
ENDPOINT_URL="http://localhost:4566"
REGION="us-east-1"

echo "Testing AWS resources in LocalStack..."

# S3バケットのテスト
echo "Testing S3 buckets..."
aws --endpoint-url=$ENDPOINT_URL s3 ls
echo "Creating test file in S3..."
echo "Hello World" > /tmp/test-file.txt
aws --endpoint-url=$ENDPOINT_URL s3 cp /tmp/test-file.txt s3://ecommerce-static-assets/
aws --endpoint-url=$ENDPOINT_URL s3 ls s3://ecommerce-static-assets/

# CloudWatch Logsのテスト
echo "Testing CloudWatch Logs..."
aws --endpoint-url=$ENDPOINT_URL logs describe-log-groups

# ログイベントを送信
LOG_GROUP="/aws/ecommerce/api"
LOG_STREAM="test-stream"

# ログストリームを作成
aws --endpoint-url=$ENDPOINT_URL logs create-log-stream \
  --log-group-name "$LOG_GROUP" \
  --log-stream-name "$LOG_STREAM"

# ログイベントを送信
aws --endpoint-url=$ENDPOINT_URL logs put-log-events \
  --log-group-name "$LOG_GROUP" \
  --log-stream-name "$LOG_STREAM" \
  --log-events timestamp=$(date +%s000),message="Test log event from LocalStack"

# SQSのテスト
echo "Testing SQS..."
aws --endpoint-url=$ENDPOINT_URL sqs list-queues
aws --endpoint-url=$ENDPOINT_URL sqs send-message \
  --queue-url "$ENDPOINT_URL/000000000000/ecommerce-orders" \
  --message-body "Test order message"

# SNSのテスト
echo "Testing SNS..."
aws --endpoint-url=$ENDPOINT_URL sns list-topics

echo "Resource tests completed. Check the output for any errors."
