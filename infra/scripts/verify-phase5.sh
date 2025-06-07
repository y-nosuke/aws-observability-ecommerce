#!/bin/bash

# =======================================================
# Phase 5: インフラ統合 動作確認スクリプト
# =======================================================

set -e

echo "🚀 Phase 5: インフラ統合 動作確認を開始します..."

# カラー定義
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# ログ関数
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# 1. Docker Composeサービスの起動確認
log_info "1. Docker Composeサービスの起動確認"

echo "📋 サービス起動状況:"
docker-compose ps

echo ""
log_info "重要なサービスのヘルスチェック..."

# MySQL
if docker-compose exec -T mysql mysqladmin ping -h localhost -u root -prootpassword > /dev/null 2>&1; then
    log_success "MySQL is healthy"
else
    log_error "MySQL is not healthy"
    exit 1
fi

# Loki
if curl -s http://localhost:3100/ready > /dev/null; then
    log_success "Loki is healthy"
else
    log_error "Loki is not healthy"
    exit 1
fi

# OTel Collector
if curl -s http://localhost:13133/ > /dev/null; then
    log_success "OTel Collector is healthy"
else
    log_error "OTel Collector is not healthy"
    exit 1
fi

# Grafana
if curl -s http://localhost:3001/api/health > /dev/null; then
    log_success "Grafana is healthy"
else
    log_error "Grafana is not healthy"
    exit 1
fi

echo ""

# 2. バックエンドAPIの確認
log_info "2. バックエンドAPIの動作確認"

if curl -s http://backend-api.localhost/api/health > /dev/null; then
    log_success "Backend API is responding"
    
    # ヘルスチェックの詳細確認
    echo "📊 ヘルスチェックレスポンス:"
    curl -s http://backend-api.localhost/api/health | jq .
else
    log_warning "Backend API is not responding (may still be starting up)"
fi

echo ""

# 3. ログの送信テスト
log_info "3. ログ送信テストの実行"

# テストリクエストを送信してログ生成
echo "🔄 テストリクエストを送信してログを生成中..."

for i in {1..5}; do
    curl -s http://backend-api.localhost/api/health > /dev/null || true
    curl -s http://backend-api.localhost/api/products > /dev/null || true
    sleep 1
done

log_success "テストリクエスト送信完了"

echo ""

# 4. Lokiでのログ確認
log_info "4. Lokiでのログ確認"

echo "📋 Loki ラベル確認:"
curl -s "http://localhost:3100/loki/api/v1/labels" | jq .

echo ""
echo "📋 最新のログエントリ (直近5分):"
current_time=$(date +%s)
five_minutes_ago=$((current_time - 300))

loki_query="{service_name=\"aws-observability-ecommerce\"}"
curl -s "http://localhost:3100/loki/api/v1/query_range?query=${loki_query}&start=${five_minutes_ago}000000000&end=${current_time}000000000&limit=10" | jq '.data.result[] | .values[] | .[1]' | head -5

echo ""

# 5. Grafanaデータソースの確認
log_info "5. Grafanaデータソースの確認"

echo "📊 Grafana データソース一覧:"
curl -s -u admin:admin http://localhost:3001/api/datasources | jq '.[] | {name: .name, type: .type, url: .url}'

echo ""

# 6. OTel Collectorの統計情報
log_info "6. OTel Collectorの統計情報"

echo "📈 OTel Collector メトリクス:"
curl -s http://localhost:8889/metrics | grep "otelcol_receiver\|otelcol_exporter" | head -10

echo ""

# 7. エンドポイント情報の表示
log_info "7. 重要なエンドポイント情報"

echo "🌐 アクセス可能なエンドポイント:"
echo "  • Grafana Dashboard: http://localhost:3001 (admin/admin)"
echo "  • Grafana (via Traefik): http://grafana.localhost"
echo "  • Backend API: http://backend-api.localhost/api/health"
echo "  • Customer Frontend: http://customer.localhost"
echo "  • Admin Frontend: http://admin.localhost"
echo "  • Loki API: http://localhost:3100"
echo "  • OTel Collector Health: http://localhost:13133"
echo "  • OTel Collector Metrics: http://localhost:8889/metrics"

echo ""

# 8. 設定ファイルの確認
log_info "8. 設定ファイルの存在確認"

config_files=(
    "infra/otel/otel-collector.yaml"
    "infra/loki/loki.yaml"
    "infra/grafana/provisioning/datasources/datasources.yaml"
    "infra/grafana/dashboards/logs-overview.json"
)

for file in "${config_files[@]}"; do
    if [ -f "$file" ]; then
        log_success "✅ $file"
    else
        log_error "❌ $file が見つかりません"
    fi
done

echo ""

# 9. 推奨次ステップ
log_info "9. 推奨次ステップ"

echo "📋 Phase 5完了後の確認ポイント:"
echo "  1. Grafanaにアクセスしてダッシュボードを確認"
echo "  2. Lokiデータソースでログクエリをテスト"
echo "  3. バックエンドAPIにリクエストを送信してログ生成を確認"
echo "  4. 構造化ログの形式が設計通りになっているか確認"

echo ""
echo "🎉 Phase 5: インフラ統合の動作確認が完了しました！"
echo ""
echo "⚠️  注意事項:"
echo "  • 初回起動時は全サービスが起動するまで数分かかる場合があります"
echo "  • ログが表示されない場合は、数分待ってから再度確認してください"
echo "  • Phase 4までの実装（構造化ロガー、ミドルウェア等）が正しく実装されている必要があります"
