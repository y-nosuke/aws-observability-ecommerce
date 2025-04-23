#!/bin/bash
# LocalStack初期化スクリプト

set -e

echo "LocalStack initializing AWS resources..."

# デフォルトリージョンの設定
REGION=${AWS_DEFAULT_REGION:-ap-northeast-1}
LOCALSTACK_HOST=localhost
ENDPOINT_URL=http://${LOCALSTACK_HOST}:4566

# S3バケット作成
echo "Creating S3 buckets..."
aws --endpoint-url=${ENDPOINT_URL} s3 mb s3://ecommerce-product-images --region ${REGION}
aws --endpoint-url=${ENDPOINT_URL} s3 mb s3://ecommerce-logs --region ${REGION}

# CloudWatch Logsロググループ作成
echo "Creating CloudWatch Logs groups..."
aws --endpoint-url=${ENDPOINT_URL} logs create-log-group --log-group-name /ecommerce/api --region ${REGION}
aws --endpoint-url=${ENDPOINT_URL} logs create-log-group --log-group-name /ecommerce/app --region ${REGION}

# SNSトピック作成
echo "Creating SNS topics..."
aws --endpoint-url=${ENDPOINT_URL} sns create-topic --name ecommerce-notifications --region ${REGION}

# SQSキュー作成
echo "Creating SQS queues..."
aws --endpoint-url=${ENDPOINT_URL} sqs create-queue --queue-name ecommerce-events --region ${REGION}

echo "LocalStack initialization completed!"
