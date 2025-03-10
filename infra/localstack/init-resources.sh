#!/bin/bash

# LocalStackでAWSリソースを初期化するスクリプト
echo "Initializing AWS resources in LocalStack..."

# 共通変数
ENDPOINT_URL="http://localhost:4566"
REGION="us-east-1"

# S3バケットの作成
echo "Creating S3 buckets..."
aws --endpoint-url=$ENDPOINT_URL s3 mb s3://ecommerce-static-assets
aws --endpoint-url=$ENDPOINT_URL s3api put-bucket-acl --bucket ecommerce-static-assets --acl public-read

# CloudWatch Logsのロググループを作成
echo "Creating CloudWatch Log groups..."
aws --endpoint-url=$ENDPOINT_URL logs create-log-group --log-group-name /aws/ecommerce/api
aws --endpoint-url=$ENDPOINT_URL logs create-log-group --log-group-name /aws/ecommerce/frontend

# X-Ray用のサンプリングルールを作成
echo "Creating X-Ray sampling rules..."
aws --endpoint-url=$ENDPOINT_URL xray create-sampling-rule --cli-input-json '{
  "SamplingRule": {
    "RuleName": "ecommerce-default",
    "ResourceARN": "*",
    "Priority": 10000,
    "FixedRate": 1,
    "ReservoirSize": 5,
    "ServiceName": "*",
    "ServiceType": "*",
    "Host": "*",
    "HTTPMethod": "*",
    "URLPath": "*",
    "Version": 1
  }
}'

# SQSキューの作成（拡張機能用）
echo "Creating SQS queues..."
aws --endpoint-url=$ENDPOINT_URL sqs create-queue --queue-name ecommerce-orders

# SNSトピックの作成（拡張機能用）
echo "Creating SNS topics..."
aws --endpoint-url=$ENDPOINT_URL sns create-topic --name ecommerce-notifications

echo "AWS resources initialization complete!"
