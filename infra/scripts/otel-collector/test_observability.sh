#!/bin/bash

# Observability Stack Test Script
# Tests Logs (Loki), Metrics (Mimir), and Traces (Tempo) via OTEL Collector

# スクリプトのディレクトリを取得
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TEMPLATE_DIR="$SCRIPT_DIR/templates"

# 設定
OTEL_ENDPOINT="http://localhost:4318"
SERVICE_NAME="observability-test"
SERVICE_VERSION="1.0.0"
ENVIRONMENT="development"

# タイムスタンプとID生成
TIMESTAMP=$(date +%s)000000000
TRACE_ID=$(openssl rand -hex 16 | tr '[:lower:]' '[:upper:]')
SPAN_ID=$(openssl rand -hex 8 | tr '[:lower:]' '[:upper:]')
END_TIME=$(($TIMESTAMP + 1000000000)) # +1 second
RANDOM_VALUE=$((RANDOM % 100))

echo "🚀 Testing Observability Stack..."
echo "Timestamp: $TIMESTAMP"
echo "Trace ID: $TRACE_ID"
echo "Span ID: $SPAN_ID"
echo ""

# JSONテンプレートを変数置換する関数
substitute_template() {
    local template_file="$1"
    local output_file="$2"

    if [ ! -f "$template_file" ]; then
        echo "❌ Template file not found: $template_file"
        return 1
    fi

    # sedコマンドで置換を実行
    sed -e "s/{{SERVICE_NAME}}/$SERVICE_NAME/g" \
        -e "s/{{SERVICE_VERSION}}/$SERVICE_VERSION/g" \
        -e "s/{{ENVIRONMENT}}/$ENVIRONMENT/g" \
        -e "s/{{TIMESTAMP}}/$TIMESTAMP/g" \
        -e "s/{{TRACE_ID}}/$TRACE_ID/g" \
        -e "s/{{SPAN_ID}}/$SPAN_ID/g" \
        -e "s/{{START_TIME}}/$TIMESTAMP/g" \
        -e "s/{{END_TIME}}/$END_TIME/g" \
        -e "s/{{LOGGER_NAME}}/test-logger/g" \
        -e "s/{{SEVERITY}}/INFO/g" \
        -e "s|{{LOG_MESSAGE}}|🔍 Test log message - observability stack verification|g" \
        -e "s/{{METER_NAME}}/test-meter/g" \
        -e "s/{{METRIC_NAME}}/test_gauge/g" \
        -e "s/{{METRIC_DESCRIPTION}}/Test gauge metric for observability verification/g" \
        -e "s/{{METRIC_UNIT}}/1/g" \
        -e "s/{{METRIC_TYPE}}/observability_verification/g" \
        -e "s/{{METRIC_VALUE}}/$RANDOM_VALUE/g" \
        -e "s/{{TRACER_NAME}}/test-tracer/g" \
        -e "s/{{SPAN_NAME}}/observability-verification/g" \
        -e "s/{{OPERATION_TYPE}}/test/g" \
        -e "s/{{TEST_CATEGORY}}/observability/g" \
        "$template_file" > "$output_file"
}

# 一時ファイルディレクトリ作成
TEMP_DIR=$(mktemp -d)
trap "rm -rf $TEMP_DIR" EXIT

# 1. Test Logs (Loki)
echo "📝 Testing Logs -> Loki..."
LOG_PAYLOAD="$TEMP_DIR/logs_payload.json"
substitute_template "$TEMPLATE_DIR/logs.json" "$LOG_PAYLOAD"

if [ $? -eq 0 ] && [ -f "$LOG_PAYLOAD" ]; then
    curl -s -X POST "$OTEL_ENDPOINT/v1/logs" \
        -H "Content-Type: application/json" \
        -d "@$LOG_PAYLOAD"

    if [ $? -eq 0 ]; then
        echo "✅ Logs sent successfully"
    else
        echo "❌ Failed to send logs"
    fi
else
    echo "❌ Failed to prepare logs payload"
fi
echo ""

# 2. Test Metrics (Mimir)
echo "📊 Testing Metrics -> Mimir..."
METRICS_PAYLOAD="$TEMP_DIR/metrics_payload.json"
substitute_template "$TEMPLATE_DIR/metrics.json" "$METRICS_PAYLOAD"

if [ $? -eq 0 ] && [ -f "$METRICS_PAYLOAD" ]; then
    curl -s -X POST "$OTEL_ENDPOINT/v1/metrics" \
        -H "Content-Type: application/json" \
        -d "@$METRICS_PAYLOAD"

    if [ $? -eq 0 ]; then
        echo "✅ Metrics sent successfully (value: $RANDOM_VALUE)"
    else
        echo "❌ Failed to send metrics"
    fi
else
    echo "❌ Failed to prepare metrics payload"
fi
echo ""

# 3. Test Traces (Tempo)
echo "🔍 Testing Traces -> Tempo..."
TRACES_PAYLOAD="$TEMP_DIR/traces_payload.json"
substitute_template "$TEMPLATE_DIR/traces.json" "$TRACES_PAYLOAD"

if [ $? -eq 0 ] && [ -f "$TRACES_PAYLOAD" ]; then
    curl -s -X POST "$OTEL_ENDPOINT/v1/traces" \
        -H "Content-Type: application/json" \
        -d "@$TRACES_PAYLOAD"

    if [ $? -eq 0 ]; then
        echo "✅ Traces sent successfully"
    else
        echo "❌ Failed to send traces"
    fi
else
    echo "❌ Failed to prepare traces payload"
fi
echo ""

echo "🎉 Observability stack test completed!"
echo ""
echo "📋 Next Steps:"
echo "1. Open Grafana: http://grafana.localhost (admin/admin)"
echo "2. Check Explore tab to query:"
echo "   - Logs in Loki: {service_name=\"$SERVICE_NAME\"}"
echo "   - Metrics in Mimir: test_gauge"
echo "   - Traces in Tempo: Search for trace ID $TRACE_ID"
echo ""
echo "🔗 Direct URLs:"
echo "   - Grafana: http://grafana.localhost"
echo "   - Loki: http://localhost:3100"
echo "   - Mimir: http://localhost:8080"
echo "   - Tempo: http://localhost:3200"
echo ""

# デバッグ用：生成されたJSONファイルを表示
if [ "$1" = "--debug" ]; then
    echo "🔍 Debug: Generated JSON payloads:"
    echo ""
    echo "--- Logs Payload ---"
    cat "$LOG_PAYLOAD"
    echo ""
    echo "--- Metrics Payload ---"
    cat "$METRICS_PAYLOAD"
    echo ""
    echo "--- Traces Payload ---"
    cat "$TRACES_PAYLOAD"
    echo ""
fi
