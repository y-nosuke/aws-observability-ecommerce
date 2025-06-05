#!/bin/bash

# Continuous Metrics Generator for Observability Testing
# Sends metrics to OTEL Collector continuously

# スクリプトのディレクトリを取得
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEMPLATE_DIR="$SCRIPT_DIR/templates"

# 設定
OTEL_ENDPOINT="http://localhost:4318"
SERVICE_NAME="metrics-generator"
ENVIRONMENT="development"
INTERVAL=${1:-10}  # デフォルト10秒間隔

echo "🚀 Starting continuous metrics generation..."
echo "Interval: ${INTERVAL} seconds"
echo "Service: $SERVICE_NAME"
echo "Press Ctrl+C to stop"
echo ""

# JSONテンプレートを変数置換する関数
substitute_metrics_template() {
    local template_file="$1"
    local output_file="$2"
    local metric_name="$3"
    local metric_value="$4"
    local metric_description="$5"
    local timestamp="$6"

    if [ ! -f "$template_file" ]; then
        echo "❌ Template file not found: $template_file"
        return 1
    fi

    # sedコマンドで置換を実行
    sed -e "s/{{SERVICE_NAME}}/$SERVICE_NAME/g" \
        -e "s/{{ENVIRONMENT}}/$ENVIRONMENT/g" \
        -e "s/{{TIMESTAMP}}/$timestamp/g" \
        -e "s/{{METER_NAME}}/continuous-meter/g" \
        -e "s/{{METRIC_NAME}}/$metric_name/g" \
        -e "s/{{METRIC_DESCRIPTION}}/$metric_description/g" \
        -e "s/{{METRIC_UNIT}}/1/g" \
        -e "s/{{METRIC_TYPE}}/continuous_monitoring/g" \
        -e "s/{{METRIC_VALUE}}/$metric_value/g" \
        "$template_file" > "$output_file"
}

# 一時ファイルディレクトリ作成
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR; echo ''; echo '🛑 Continuous metrics generation stopped.'" EXIT

# カウンター初期化
COUNTER=0

while true; do
    TIMESTAMP=$(date +%s)000000000
    COUNTER=$((COUNTER + 1))

    echo -n "📊 Sending metrics batch #$COUNTER - "

    # CPU使用率メトリクス (0-100%)
    CPU_VALUE=$((RANDOM % 100))
    CPU_PAYLOAD="$TEMP_DIR/cpu_metrics.json"
    substitute_metrics_template "$TEMPLATE_DIR/metrics.json" "$CPU_PAYLOAD" \
        "cpu_usage_percent" "$CPU_VALUE" "Simulated CPU usage percentage" "$TIMESTAMP"

    curl -s -X POST "$OTEL_ENDPOINT/v1/metrics" \
        -H "Content-Type: application/json" \
        -d "@$CPU_PAYLOAD" > /dev/null

    # メモリ使用率メトリクス (0-100%)
    MEMORY_VALUE=$((RANDOM % 100))
    MEMORY_PAYLOAD="$TEMP_DIR/memory_metrics.json"
    substitute_metrics_template "$TEMPLATE_DIR/metrics.json" "$MEMORY_PAYLOAD" \
        "memory_usage_percent" "$MEMORY_VALUE" "Simulated memory usage percentage" "$TIMESTAMP"

    curl -s -X POST "$OTEL_ENDPOINT/v1/metrics" \
        -H "Content-Type: application/json" \
        -d "@$MEMORY_PAYLOAD" > /dev/null

    # リクエスト数メトリクス (0-1000)
    REQUEST_VALUE=$((RANDOM % 1000))
    REQUEST_PAYLOAD="$TEMP_DIR/request_metrics.json"
    substitute_metrics_template "$TEMPLATE_DIR/metrics.json" "$REQUEST_PAYLOAD" \
        "http_requests_total" "$REQUEST_VALUE" "Total HTTP requests" "$TIMESTAMP"

    curl -s -X POST "$OTEL_ENDPOINT/v1/metrics" \
        -H "Content-Type: application/json" \
        -d "@$REQUEST_PAYLOAD" > /dev/null

    # レスポンス時間メトリクス (0-2000ms)
    RESPONSE_TIME=$((RANDOM % 2000))
    RESPONSE_PAYLOAD="$TEMP_DIR/response_metrics.json"
    substitute_metrics_template "$TEMPLATE_DIR/metrics.json" "$RESPONSE_PAYLOAD" \
        "http_response_time_ms" "$RESPONSE_TIME" "HTTP response time in milliseconds" "$TIMESTAMP"

    curl -s -X POST "$OTEL_ENDPOINT/v1/metrics" \
        -H "Content-Type: application/json" \
        -d "@$RESPONSE_PAYLOAD" > /dev/null

    echo "CPU: ${CPU_VALUE}%, Memory: ${MEMORY_VALUE}%, Requests: ${REQUEST_VALUE}, Response: ${RESPONSE_TIME}ms ($(date '+%H:%M:%S'))"

    sleep $INTERVAL
done
