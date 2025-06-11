# Goバックエンド基本メトリクス収集 詳細設計書

## 1. 概要

### 1.1. ドキュメント情報

- **ドキュメント名**: Goバックエンド基本メトリクス収集 詳細設計書
- **バージョン**: 1.0.0
- **作成日**: 2025-06-10
- **関連ユーザーストーリー**: US-DEV-OBS-IMPLEMENT-CUSTOM-METRICS
- **対象コンポーネント**: backend-api (Goアプリケーション)

### 1.2. 目的と背景

本設計書は、EコマースアプリケーションのGoバックエンド（Echo）において、REDメトリクス（Rate、Errors、Duration）をOpenTelemetry SDKで収集し、OTel Collector経由でMimirに転送する機能の詳細設計を定義します。

### 1.3. 設計方針

- **段階的実装**: 既存のログミドルウェアを拡張してメトリクス収集を追加
- **非侵襲的**: 既存のビジネスロジックへの影響を最小限に抑制
- **標準準拠**: OpenTelemetry標準に準拠した実装
- **パフォーマンス重視**: メトリクス収集がアプリケーション性能に与える影響を最小化

## 2. 現状分析

### 2.1. 既存実装の状況

#### 2.1.1. OpenTelemetryの実装状況

```go
// 現在: pkg/observability/otel.go (ログのみ)
type OTelManager struct {
    loggerProvider *sdklog.LoggerProvider
    resource       *resource.Resource
}
```

**現状の課題**:

- メトリクス機能が未実装
- メトリクス用のプロバイダーとエクスポーターが存在しない

#### 2.1.2. ミドルウェアの実装状況

```go
// 現在: internal/shared/presentation/rest/middleware/structured_logging.go
func LoggingMiddleware() echo.MiddlewareFunc {
    // HTTPリクエストログは収集済み
    // しかしメトリクス収集は未実装
}
```

### 2.2. インフラストラクチャの準備状況

#### 2.2.1. OTel Collector設定

```yaml
# infra/otel/config.yaml (メトリクスパイプライン設定済み)
service:
  pipelines:
    metrics:
      receivers: [otlp]
      processors: [memory_limiter, attributes/add_env, resource, batch]
      exporters: [debug/stdout, otlphttp/mimir]
```

#### 2.2.2. Mimir設定

- エンドポイント: `http://mimir:9009/otlp`
- Grafana Datasource設定済み

### 2.3. 対象APIエンドポイント

| HTTPメソッド | パス                            | 説明                 | 重要度 |
| ------------ | ------------------------------- | -------------------- | ------ |
| GET          | `/api/health`                   | ヘルスチェック       | 高     |
| GET          | `/api/products`                 | 商品一覧取得         | 高     |
| GET          | `/api/products/{id}`            | 商品詳細取得         | 高     |
| POST         | `/api/products/{id}/image`      | 商品画像アップロード | 中     |
| GET          | `/api/products/{id}/image`      | 商品画像取得         | 中     |
| GET          | `/api/categories`               | カテゴリー一覧取得   | 中     |
| GET          | `/api/categories/{id}/products` | カテゴリー別商品取得 | 中     |

## 3. メトリクス設計

### 3.1. REDメトリクス定義

#### 3.1.1. Rate (リクエスト数)

| 項目               | 値                                               |
| ------------------ | ------------------------------------------------ |
| **メトリクス名**   | `http_requests_total`                            |
| **メトリクス種別** | Counter                                          |
| **説明**           | HTTPリクエストの総数                             |
| **単位**           | requests                                         |
| **ラベル**         | `method`, `route`, `status_code`, `service_name` |

**ラベル詳細**:

- `method`: HTTPメソッド（GET, POST, PUT, DELETE）
- `route`: API ルートパターン（例: `/api/products/{id}`）
- `status_code`: HTTPステータスコード（例: 200, 404, 500）
- `service_name`: サービス名（`aws-observability-ecommerce`）

#### 3.1.2. Errors (エラー率)

| 項目               | 値                                                             |
| ------------------ | -------------------------------------------------------------- |
| **メトリクス名**   | `http_request_errors_total`                                    |
| **メトリクス種別** | Counter                                                        |
| **説明**           | HTTPエラーレスポンス（4xx, 5xx）の総数                         |
| **単位**           | requests                                                       |
| **ラベル**         | `method`, `route`, `status_code`, `error_type`, `service_name` |

**ラベル詳細**:

- `error_type`: エラー種別（`client_error` for 4xx, `server_error` for 5xx）

#### 3.1.3. Duration (レスポンス時間)

| 項目               | 値                                                            |
| ------------------ | ------------------------------------------------------------- |
| **メトリクス名**   | `http_request_duration_seconds`                               |
| **メトリクス種別** | Histogram                                                     |
| **説明**           | HTTPリクエスト処理時間の分布                                  |
| **単位**           | seconds                                                       |
| **ラベル**         | `method`, `route`, `service_name`                             |
| **バケット**       | 0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0 |

### 3.2. 追加メトリクス

#### 3.2.1. Request Size (リクエストサイズ)

| 項目               | 値                                                 |
| ------------------ | -------------------------------------------------- |
| **メトリクス名**   | `http_request_size_bytes`                          |
| **メトリクス種別** | Histogram                                          |
| **説明**           | HTTPリクエストボディサイズの分布                   |
| **単位**           | bytes                                              |
| **ラベル**         | `method`, `route`, `service_name`                  |
| **バケット**       | 64, 256, 1024, 4096, 16384, 65536, 262144, 1048576 |

#### 3.2.2. Response Size (レスポンスサイズ)

| 項目               | 値                                                 |
| ------------------ | -------------------------------------------------- |
| **メトリクス名**   | `http_response_size_bytes`                         |
| **メトリクス種別** | Histogram                                          |
| **説明**           | HTTPレスポンスボディサイズの分布                   |
| **単位**           | bytes                                              |
| **ラベル**         | `method`, `route`, `service_name`                  |
| **バケット**       | 64, 256, 1024, 4096, 16384, 65536, 262144, 1048576 |

## 4. 技術設計

### 4.1. 依存関係の追加

#### 4.1.1. 必要なGoモジュール

```go
// go.mod に追加予定
require (
    // メトリクス関連
    go.opentelemetry.io/otel/metric v1.30.0
    go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.30.0
    go.opentelemetry.io/otel/sdk/metric v1.30.0

    // 可能であればEchoインストルメンテーション
    go.opentelemetry.io/contrib/instrumentation/github.com/labstack/echo/otelecho v0.52.0
)
```

### 4.2. アーキテクチャ設計

#### 4.2.1. コンポーネント図

```
┌─────────────────────────────────────────────────────────────┐
│                    backend-api                              │
├─────────────────────────────────────────────────────────────┤
│  Echo Application                                           │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              Middleware Stack                        │    │
│  │  ┌─────────────────────────────────────────────┐     │    │
│  │  │         Request ID Middleware                │     │    │
│  │  └─────────────────────────────────────────────┘     │    │
│  │  ┌─────────────────────────────────────────────┐     │    │
│  │  │    HTTP Metrics Middleware (NEW)            │     │    │
│  │  │  - Counter: http_requests_total             │     │    │
│  │  │  - Counter: http_request_errors_total       │     │    │
│  │  │  - Histogram: http_request_duration_seconds │     │    │
│  │  │  - Histogram: http_request_size_bytes       │     │    │
│  │  │  - Histogram: http_response_size_bytes      │     │    │
│  │  └─────────────────────────────────────────────┘     │    │
│  │  ┌─────────────────────────────────────────────┐     │    │
│  │  │           Logging Middleware                │     │    │
│  │  └─────────────────────────────────────────────┘     │    │
│  │  ┌─────────────────────────────────────────────┐     │    │
│  │  │       Error Handling Middleware             │     │    │
│  │  └─────────────────────────────────────────────┘     │    │
│  └─────────────────────────────────────────────────────┘    │
├─────────────────────────────────────────────────────────────┤
│  OpenTelemetry Manager                                      │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              OTelManager (EXTENDED)                 │    │
│  │  ┌─────────────────────────────────────────────┐     │    │
│  │  │          LoggerProvider (existing)          │     │    │
│  │  └─────────────────────────────────────────────┘     │    │
│  │  ┌─────────────────────────────────────────────┐     │    │
│  │  │           MeterProvider (NEW)               │     │    │
│  │  │  - OTLP HTTP Exporter                      │     │    │
│  │  │  - Batch Processor                         │     │    │
│  │  └─────────────────────────────────────────────┘     │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
                              │
                              │ OTLP/HTTP
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    OTel Collector                           │
│  ┌─────────────────────────────────────────────────────┐    │
│  │               Metrics Pipeline                      │    │
│  │  Receiver → Processor → Exporter → Mimir           │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                       Mimir                                 │
│  ┌─────────────────────────────────────────────────────┐    │
│  │              Metrics Storage                        │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      Grafana                                │
│  ┌─────────────────────────────────────────────────────┐    │
│  │            Dashboard & Visualization                │    │
│  └─────────────────────────────────────────────────────┘    │
└─────────────────────────────────────────────────────────────┘
```

### 4.3. 実装設計

#### 4.3.1. OTelManager拡張設計

```go
// pkg/observability/otel.go (拡張予定)
type OTelManager struct {
    loggerProvider *sdklog.LoggerProvider
    meterProvider  *sdkmetric.MeterProvider  // NEW
    resource       *resource.Resource
    meter          metric.Meter              // NEW
}

// 新規追加予定のメソッド
func (m *OTelManager) GetMeterProvider() *sdkmetric.MeterProvider
func (m *OTelManager) GetMeter() metric.Meter
```

#### 4.3.2. メトリクス構造体設計

```go
// pkg/observability/metrics.go (新規作成予定)
type HTTPMetrics struct {
    requestsTotal       metric.Int64Counter
    errorsTotal         metric.Int64Counter
    requestDuration     metric.Float64Histogram
    requestSizeBytes    metric.Int64Histogram
    responseSizeBytes   metric.Int64Histogram
}

func NewHTTPMetrics(meter metric.Meter) (*HTTPMetrics, error)
func (m *HTTPMetrics) RecordRequest(method, route string, statusCode int, duration time.Duration, requestSize, responseSize int64)
```

#### 4.3.3. HTTPメトリクスミドルウェア設計

```go
// internal/shared/presentation/rest/middleware/http_metrics.go (新規作成予定)
func HTTPMetricsMiddleware(metrics *observability.HTTPMetrics) echo.MiddlewareFunc {
    return func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            start := time.Now()

            // リクエストサイズ取得
            requestSize := getRequestSize(c.Request())

            // レスポンス監視用ラッパー
            resWrapper := &metricsResponseWriter{ResponseWriter: c.Response().Writer}
            c.Response().Writer = resWrapper

            // ハンドラー実行
            err := next(c)

            // メトリクス記録
            duration := time.Since(start)
            route := getRoutePattern(c)

            metrics.RecordRequest(
                c.Request().Method,
                route,
                c.Response().Status,
                duration,
                requestSize,
                resWrapper.size,
            )

            return err
        }
    }
}
```

### 4.4. 設定拡張設計

#### 4.4.1. 設定構造体拡張

```go
// internal/shared/infrastructure/config/observability_config.go (拡張予定)
type OTelConfig struct {
    ServiceName           string            `mapstructure:"service_name"`
    ServiceVersion        string            `mapstructure:"service_version"`
    ServiceNamespace      string            `mapstructure:"service_namespace"`
    DeploymentEnvironment string            `mapstructure:"deployment_environment"`
    Collector             CollectorConfig   `mapstructure:"collector"`
    Logging               OTelLoggingConfig `mapstructure:"logging"`
    Metrics               OTelMetricsConfig `mapstructure:"metrics"` // NEW
}

// 新規追加予定
type OTelMetricsConfig struct {
    BatchTimeout       time.Duration `mapstructure:"batch_timeout"`
    MaxQueueSize       int           `mapstructure:"max_queue_size"`
    MaxExportBatchSize int           `mapstructure:"max_export_batch_size"`
    ExportTimeout      time.Duration `mapstructure:"export_timeout"`
    HistogramBoundaries []float64     `mapstructure:"histogram_boundaries"`
}
```

## 5. 実装計画

### 5.1. 実装フェーズ

#### 5.1.1. フェーズ1: 基盤実装

1. **依存関係追加**
   - `go.mod`にメトリクス関連ライブラリを追加
   - `go mod tidy`実行

2. **OTelManager拡張**
   - メトリクス用のプロバイダーとエクスポーター追加
   - 設定構造体の拡張

3. **HTTPメトリクス構造体実装**
   - `pkg/observability/metrics.go`新規作成
   - 基本的なREDメトリクス定義

#### 5.1.2. フェーズ2: ミドルウェア実装

1. **HTTPメトリクスミドルウェア実装**
   - `internal/shared/presentation/rest/middleware/http_metrics.go`新規作成
   - ルートパターン取得ロジック実装

2. **ルーター統合**
   - `router.go`でミドルウェア追加
   - 既存ミドルウェアとの順序調整

#### 5.1.3. フェーズ3: テスト・調整

1. **単体テスト実装**
   - メトリクス構造体のテスト
   - ミドルウェアのテスト

2. **統合テスト実行**
   - OTel Collectorとの連携確認
   - Mimirでのメトリクス確認

3. **パフォーマンステスト**
   - メトリクス収集によるオーバーヘッド測定
   - 最適化実施

### 5.2. 実装順序

```
1. go.mod依存関係追加
   ↓
2. observability_config.go拡張
   ↓
3. otel.go拡張（MeterProvider追加）
   ↓
4. metrics.go新規作成
   ↓
5. http_metrics.go新規作成（ミドルウェア）
   ↓
6. router.go修正（ミドルウェア追加）
   ↓
7. config.yaml設定追加
   ↓
8. テスト実装・実行
```

## 6. 品質要件

### 6.1. パフォーマンス要件

| 項目                             | 要件   | 測定方法                                 |
| -------------------------------- | ------ | ---------------------------------------- |
| **レスポンス時間オーバーヘッド** | < 5%   | ベンチマークテストでメトリクス有無の比較 |
| **メモリ使用量増加**             | < 10MB | プロファイリングツールで測定             |
| **CPU使用率増加**                | < 2%   | 負荷テスト時のCPU使用率測定              |

### 6.2. 信頼性要件

| 項目                     | 要件                                |
| ------------------------ | ----------------------------------- |
| **メトリクス収集失敗時** | アプリケーションの正常動作を継続    |
| **OTel Collector障害時** | メトリクス送信失敗でもAPI処理を継続 |
| **メトリクス精度**       | 99.9%以上の正確性                   |

### 6.3. 運用要件

| 項目         | 要件                                 |
| ------------ | ------------------------------------ |
| **設定変更** | 環境変数での動的設定変更対応         |
| **ログ出力** | メトリクス関連エラーの適切なログ出力 |
| **監視**     | メトリクス送信状況の監視機能         |

## 7. 設定仕様

### 7.1. 環境変数

| 環境変数名                           | デフォルト値 | 説明                         |
| ------------------------------------ | ------------ | ---------------------------- |
| `OTEL_METRICS_ENABLED`               | `true`       | メトリクス収集の有効/無効    |
| `OTEL_METRICS_BATCH_TIMEOUT`         | `1s`         | メトリクスバッチ送信間隔     |
| `OTEL_METRICS_MAX_QUEUE_SIZE`        | `2048`       | メトリクスキューの最大サイズ |
| `OTEL_METRICS_MAX_EXPORT_BATCH_SIZE` | `512`        | 一度に送信するメトリクス数   |
| `OTEL_METRICS_EXPORT_TIMEOUT`        | `30s`        | メトリクス送信タイムアウト   |

### 7.2. config.yaml設定例

```yaml
observability:
  otel:
    service_name: "aws-observability-ecommerce"
    service_version: "1.0.0"
    service_namespace: "ecommerce"
    deployment_environment: "development"

    collector:
      endpoint: "otel-collector:4318"
      timeout: "10s"
      retry_enabled: true
      retry_max_attempts: 3
      compression: "gzip"

    metrics:
      batch_timeout: "1s"
      max_queue_size: 2048
      max_export_batch_size: 512
      export_timeout: "30s"
      histogram_boundaries: [0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1.0, 2.5, 5.0, 10.0]
```

## 8. Grafanaダッシュボード設計

### 8.1. 基本ダッシュボード構成

#### 8.1.1. Request Rate パネル

```promql
# リクエストレート（1分間のリクエスト数）
rate(http_requests_total[1m])

# エンドポイント別リクエストレート
sum(rate(http_requests_total[1m])) by (route)
```

#### 8.1.2. Error Rate パネル

```promql
# エラーレート（全体）
sum(rate(http_request_errors_total[1m])) / sum(rate(http_requests_total[1m])) * 100

# エンドポイント別エラーレート
sum(rate(http_request_errors_total[1m])) by (route) / sum(rate(http_requests_total[1m])) by (route) * 100
```

#### 8.1.3. Response Duration パネル

```promql
# 95パーセンタイル レスポンス時間
histogram_quantile(0.95, rate(http_request_duration_seconds_bucket[1m]))

# 平均レスポンス時間
rate(http_request_duration_seconds_sum[1m]) / rate(http_request_duration_seconds_count[1m])
```

### 8.2. 推奨アラートルール

| メトリクス                          | 条件            | 重要度   |
| ----------------------------------- | --------------- | -------- |
| **エラーレート**                    | > 5% for 2分間  | Warning  |
| **エラーレート**                    | > 10% for 1分間 | Critical |
| **95パーセンタイル レスポンス時間** | > 2秒 for 3分間 | Warning  |
| **95パーセンタイル レスポンス時間** | > 5秒 for 1分間 | Critical |

## 9. テスト設計

### 9.1. 単体テスト

#### 9.1.1. HTTPMetrics構造体テスト

```go
// pkg/observability/metrics_test.go
func TestHTTPMetrics_RecordRequest(t *testing.T) {
    // メトリクス記録の正確性テスト
}

func TestHTTPMetrics_CounterIncrement(t *testing.T) {
    // カウンターの増分テスト
}

func TestHTTPMetrics_HistogramRecording(t *testing.T) {
    // ヒストグラムの記録テスト
}
```

#### 9.1.2. ミドルウェアテスト

```go
// internal/shared/presentation/rest/middleware/http_metrics_test.go
func TestHTTPMetricsMiddleware(t *testing.T) {
    // ミドルウェアの動作テスト
}

func TestRoutePatternExtraction(t *testing.T) {
    // ルートパターン抽出ロジックテスト
}
```

### 9.2. 統合テスト

#### 9.2.1. OTel Collector連携テスト

```bash
# infra/scripts/otel-collector/test_metrics.sh
# メトリクス送信とCollector受信の確認スクリプト
```

#### 9.2.2. Mimir連携テスト

```bash
# Mimirでのメトリクス保存確認スクリプト
curl -G 'http://mimir:9009/prometheus/api/v1/query' \
  --data-urlencode 'query=http_requests_total'
```

## 10. リスクと対策

### 10.1. 技術的リスク

| リスク                                     | 影響度 | 対策                               |
| ------------------------------------------ | ------ | ---------------------------------- |
| **メトリクス収集によるパフォーマンス劣化** | 中     | サンプリング実装、バッチ処理最適化 |
| **OTel Collector障害**                     | 低     | 非同期送信、バッファリング         |
| **メモリリーク**                           | 高     | 適切なリソース管理、監視実装       |

### 10.2. 運用リスク

| リスク                 | 影響度 | 対策                                 |
| ---------------------- | ------ | ------------------------------------ |
| **メトリクス設定不正** | 低     | バリデーション実装、デフォルト値設定 |
| **ダッシュボード障害** | 低     | 複数の可視化手段確保                 |
| **アラート疲れ**       | 中     | 適切な閾値設定、段階的アラート       |

## 11. 今後の拡張予定

### 11.1. 短期拡張（次期フェーズ）

- **カスタムビジネスメトリクス**: 商品閲覧数、カート追加数等
- **データベースメトリクス**: クエリ実行時間、接続プール状況
- **AWS サービスメトリクス**: S3アクセス時間、SQS処理状況

### 11.2. 長期拡張（将来計画）

- **ユーザージャーニーメトリクス**: エンドツーエンド処理時間
- **ビジネスKPIメトリクス**: 売上、コンバージョン率
- **セキュリティメトリクス**: 不正アクセス検知、レート制限

---

## 付録

### A. 関連ドキュメント

- [OpenTelemetry Go Documentation](https://opentelemetry.io/docs/languages/go/)
- [Echo Middleware Documentation](https://echo.labstack.com/docs/middleware/)
- [Prometheus Metrics Best Practices](https://prometheus.io/docs/practices/naming/)

### B. 用語集

| 用語                 | 説明                                                               |
| -------------------- | ------------------------------------------------------------------ |
| **RED メトリクス**   | Rate（レート）、Errors（エラー）、Duration（処理時間）の監視指標   |
| **OTLP**             | OpenTelemetry Protocol - テレメトリデータ送信の標準プロトコル      |
| **Mimir**            | Prometheus互換の長期メトリクスストレージ                           |
| **カーディナリティ** | メトリクスのラベル組み合わせ数。高すぎるとパフォーマンス問題の原因 |

---

**レビュー承認**:

- [ ] アーキテクト確認
- [ ] シニア開発者確認
- [ ] SRE確認
- [ ] セキュリティ確認
