#!/bin/bash
# OTel Collectorへのテストデータ送信スクリプト
# 依存: otel-cli (https://github.com/equinix-labs/otel-cli)

set -e

OTEL_COLLECTOR_ENDPOINT="http://localhost:4318"

# テストログ（spanとして送信、log的な属性を付与）
otel-cli span --endpoint $OTEL_COLLECTOR_ENDPOINT --protocol http/protobuf \
  --name "test-log" --service "test-service" --attrs "log.message=test log from otel-cli,env=local"

# テストメトリクス（spanとして送信、metric的な属性を付与）
otel-cli span --endpoint $OTEL_COLLECTOR_ENDPOINT --protocol http/protobuf \
  --name "test-metric" --service "test-service" --attrs "metric.name=test_metric,metric.value=1,env=local"

# テストトレース
otel-cli span --endpoint $OTEL_COLLECTOR_ENDPOINT --protocol http/protobuf \
  --name "test-span" --service "test-service" --kind client --attrs "env=local"

echo "OTelテストデータ送信完了"
