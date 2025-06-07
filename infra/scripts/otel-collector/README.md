# Observability Stack Test Scripts

このディレクトリには、OpenTelemetry Collector経由でLoki、Mimir、Tempoに対してログ、メトリクス、トレースを送信するテストスクリプトが含まれています。

## ファイル構成

```text
infra/scripts/otel-collector/
├── test_observability.sh      # 単発テスト実行スクリプト
├── continuous_metrics.sh      # 継続的メトリクス生成スクリプト
├── templates/                 # JSONテンプレートファイル
│   ├── logs.json             # ログ送信用JSONテンプレート
│   ├── metrics.json          # メトリクス送信用JSONテンプレート
│   └── traces.json           # トレース送信用JSONテンプレート
└── README.md                 # このファイル
```

## 前提条件

- OpenTelemetry Collectorが `http://localhost:4318` で起動していること
- Grafana、Loki、Mimir、Tempoが正常に動作していること
- `curl` コマンドが利用可能であること

## 使用方法

### 1. 基本動作確認テスト

すべてのコンポーネント（ログ、メトリクス、トレース）の動作を一度に確認：

```bash
cd /home/yoichi/workspace/aws-observability-ecommerce/infra/scripts/otel-collector
./test_observability.sh
```

### 2. デバッグモード実行

生成されたJSONペイロードを表示してデバッグ：

```bash
./test_observability.sh --debug
```

### 3. 継続的メトリクス生成

リアルタイムでメトリクスを送信し続ける（デフォルト10秒間隔）：

```bash
./continuous_metrics.sh
```

間隔を変更する場合（例：5秒間隔）：

```bash
./continuous_metrics.sh 5
```

### 4. ファイル実行権限の設定

スクリプトが実行できない場合：

```bash
chmod +x test_observability.sh
chmod +x continuous_metrics.sh
```

## JSONテンプレートについて

### templates/logs.json

ログ送信用のOTLP形式JSONテンプレート。以下の変数が置換されます：

- `{{SERVICE_NAME}}`: サービス名
- `{{SERVICE_VERSION}}`: サービスバージョン
- `{{TIMESTAMP}}`: Unix nanosecondタイムスタンプ
- `{{LOG_MESSAGE}}`: ログメッセージ
- `{{TRACE_ID}}`: トレースID（ログとトレースの関連付け用）

### templates/metrics.json

メトリクス送信用のOTLP形式JSONテンプレート。以下の変数が置換されます：

- `{{METRIC_NAME}}`: メトリクス名
- `{{METRIC_VALUE}}`: メトリクス値
- `{{METRIC_DESCRIPTION}}`: メトリクスの説明

### templates/traces.json

トレース送信用のOTLP形式JSONテンプレート。以下の変数が置換されます：

- `{{TRACE_ID}}`: トレースID
- `{{SPAN_ID}}`: スパンID
- `{{SPAN_NAME}}`: スパン名
- `{{START_TIME}}`, `{{END_TIME}}`: スパンの開始・終了時刻

## Grafanaでの確認方法

テスト実行後、以下の手順でGrafanaで確認：

1. **Grafanaにアクセス**: <http://grafana.localhost>
   - ユーザー名: `admin`
   - パスワード: `admin`

2. **Exploreタブでデータを確認**:

   **ログ (Loki データソース選択)**:

   ```bash
   {service_name="observability-test"}
   ```

   **メトリクス (Mimir データソース選択)**:

   ```bash
   test_gauge
   cpu_usage_percent
   memory_usage_percent
   http_requests_total
   http_response_time_ms
   ```

   **トレース (Tempo データソース選択)**:
   - Search画面でTraceIDを検索（スクリプト実行時に表示されます）

## トラブルシューティング

### スクリプトが失敗する場合

1. **コンテナの状態確認**:

   ```bash
   docker ps | grep -E "(otel-collector|loki|mimir|tempo|grafana)"
   ```

2. **OTEL Collectorログ確認**:

   ```bash
   docker logs otel-collector
   ```

3. **各サービスのヘルスチェック**:

   ```bash
   curl http://localhost:4318/  # OTEL Collector
   curl http://localhost:3100/ready  # Loki
   curl http://localhost:8080/ready  # Mimir
   curl http://localhost:3200/ready  # Tempo
   ```

### JSONテンプレートが見つからない場合

テンプレートディレクトリが正しく作成されているか確認：

```bash
ls -la /home/yoichi/workspace/aws-observability-ecommerce/infra/scripts/otel-collector/templates/
```

### Grafanaでデータが表示されない場合

1. データソースの設定確認
2. 時間範囲の調整（Last 15 minutes推奨）
3. ログレベルでのフィルタリング確認

## 高度な使用方法

### カスタムJSONテンプレートの作成

新しいテンプレートファイルを作成し、`substitute_template`関数で必要な変数を置換することで、独自のテレメトリーデータを送信できます。

### 異なる環境での実行

スクリプト内の以下の変数を変更することで、異なる環境に対応できます：

- `OTEL_ENDPOINT`: OTEL Collectorのエンドポイント
- `SERVICE_NAME`: サービス名
- `ENVIRONMENT`: 環境名（development/staging/production）

## 参考リンク

- [OpenTelemetry OTLP Specification](https://opentelemetry.io/docs/specs/otlp/)
- [Grafana Loki Documentation](https://grafana.com/docs/loki/)
- [Grafana Mimir Documentation](https://grafana.com/docs/mimir/)
- [Grafana Tempo Documentation](https://grafana.com/docs/tempo/)
