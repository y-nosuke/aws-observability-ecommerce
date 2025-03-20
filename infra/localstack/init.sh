#!/bin/bash

echo "Initializing LocalStack..."

# 必要なサービスが起動していることを確認
echo "Checking LocalStack services..."
awslocal cloudwatch list-metrics

echo "LocalStack initialization completed."
