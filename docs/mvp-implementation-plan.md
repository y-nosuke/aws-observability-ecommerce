# 1. AWSオブザーバビリティ学習用eコマースアプリ実装計画

## 1.1. 概要

この実装計画は、Go Echo+sqlboiler+Next.jsで構築するeコマースアプリのMVP開発に関するものです。本プロジェクトの主な目的は、AWSオブザーバビリティパターンを学習することであり、オブザーバビリティの2つのアプローチ（AWS SDK v2とOpenTelemetry）を実装して比較します。

### 1.1.1. 主要技術スタック

**フロントエンド:**

- Next.js（Reactベースのフレームワーク）
- TypeScript
- TailwindCSS
- AWS S3 + CloudFront（静的アセットのホスティング）

**バックエンド:**

- Go言語（Echo Webフレームワーク）
- sqlboiler（MySQLに特化したORM）
- slog（構造化ログ）
- AWS Fargate (ECS)
- MySQL (RDS)
- AWS ALB（Application Load Balancer）

**オブザーバビリティ:**

1. AWS SDK v2アプローチ
   - X-Ray SDK for Go v2
   - CloudWatch SDK v2
   - slog + CloudWatch Logs

2. OpenTelemetryアプローチ
   - OpenTelemetry Go SDK
   - AWS Distro for OpenTelemetry (ADOT)
   - X-Ray Exporter

## 1.2. フェーズ別実装計画

### 1.2.1. フェーズ1: 開発環境のセットアップとプロジェクト骨組み構築（2週間）

1. **開発環境構築**
   - Docker Composeによるローカル開発環境の構築
   - Go開発環境のセットアップ（air、テスト、lint）
   - Next.js開発環境のセットアップ
   - LocalStackによるローカルAWSサービスエミュレーション

2. **プロジェクト基本構造の構築**
   - バックエンドのディレクトリ構造の設計
   - API定義（OpenAPI仕様）
   - データベーススキーマとマイグレーション設計
   - フロントエンドの基本レイアウトとルーティング設計

3. **CI/CD初期設定**
   - GitHub Actions基本設定
   - Terraformによる基本インフラ定義

### 1.2.2. フェーズ2: バックエンド実装（3週間）

1. **データベースモデル実装**
   - MySQLスキーマの実装
   - sqlboilerによるモデル生成
   - データマイグレーションスクリプト

2. **コアサービスの実装**
   - 商品カタログサービス実装（`/api/products/*`）
   - 在庫管理サービス実装（`/api/inventory/*`）
   - 注文処理サービス実装（`/api/orders/*`）
   - 管理者認証サービス実装（`/api/auth/*`）

3. **オブザーバビリティ基盤の初期実装（AWS SDK v2アプローチ）**
   - slogによる構造化ログ実装
   - X-Ray SDK初期統合
   - 基本メトリクス収集設定

### 1.2.3. フェーズ3: フロントエンド実装（3週間）

1. **Next.jsの基本コンポーネント実装**
   - レイアウトコンポーネント（ヘッダー、フッター、ナビゲーション）
   - 共通UIコンポーネント
   - API通信用のカスタムフック

2. **顧客向け画面の実装**
   - ホーム画面（SCR-01）
   - 商品一覧画面（SCR-02）
   - 商品詳細画面（SCR-03）
   - カート画面（SCR-04）
   - 注文手続き画面（SCR-05）
   - 注文完了画面（SCR-06）

3. **管理者向け画面の実装**
   - 管理者ログイン画面（SCR-07）
   - 管理者ダッシュボード（SCR-08）
   - 商品管理画面（SCR-09）
   - 商品登録画面（SCR-10）
   - 在庫管理画面（SCR-11）

### 1.2.4. フェーズ4: AWS環境へのデプロイとオブザーバビリティ強化（2週間）

1. **AWSインフラの構築**
   - Terraformによる本番環境構築
   - AWS Fargate、RDS、CloudFront、S3の設定
   - ALBセットアップとルーティング設定

2. **オブザーバビリティの本格実装**
   - X-Ray詳細設定（セグメント、サブセグメント、アノテーション）
   - CloudWatch Dashboardの構築
   - アラート・通知の設定

3. **初期パフォーマンステストと最適化**
   - ロードテスト
   - ボトルネック特定と修正
   - スケーリング設定

### 1.2.5. フェーズ5: OpenTelemetryアプローチの実装（2週間）

1. **OpenTelemetry基盤の構築**
   - OTELライブラリの統合
   - ADOT Collectorの設定
   - トレーサー実装

2. **OTELコンポーネントの実装**
   - Echoミドルウェア統合
   - sqlboilerとの連携
   - メトリクス収集実装

3. **AWS X-Ray Exporterの設定**
   - OTELからX-Rayへのエクスポート設定
   - サンプリングルールの最適化
   - バックエンド統合

### 1.2.6. フェーズ6: オブザーバビリティ比較と最終調整（2週間）

1. **両アプローチの比較と分析**
   - パフォーマンス比較
   - 運用コスト分析
   - 使いやすさ評価

2. **最終調整と文書化**
   - パフォーマンス最適化
   - ドキュメント作成
   - 学習成果のまとめ

## 1.3. Next.jsの基礎解説

### 1.3.1. Next.jsの概要

Next.jsはReactベースのフレームワークで、サーバーサイドレンダリング（SSR）、静的サイト生成（SSG）、クライアントサイドレンダリング（CSR）をシームレスに組み合わせて使用できます。特に初心者にとって、以下の点が魅力的です：

- **ファイルベースのルーティング**: ディレクトリ構造がそのままURLパスに反映
- **ビルトインの最適化**: 画像最適化、コード分割、パフォーマンス最適化
- **APIルート**: 同じプロジェクト内でAPIエンドポイントを作成可能

### 1.3.2. Next.jsの基本構造

```
my-ecommerce-app/
├── pages/            # ルーティングに直接対応するファイル
│   ├── index.js      # ホームページ (/)
│   ├── products/
│   │   ├── index.js  # 商品一覧ページ (/products)
│   │   └── [id].js   # 商品詳細ページ (/products/123)
│   ├── cart.js       # カートページ (/cart)
│   ├── checkout.js   # 注文手続きページ (/checkout)
│   └── api/          # APIルート
│       └── ...       # (フロントエンド開発用の簡易APIなど)
├── components/       # 再利用可能なコンポーネント
│   ├── Layout.js     # 共通レイアウト
│   ├── ProductCard.js # 商品カードコンポーネント
│   └── ...
├── public/           # 静的ファイル（画像など）
└── styles/           # CSSファイル
```

### 1.3.3. Next.jsのデータ取得パターン

Next.jsでは、3つの主要なデータ取得パターンを使用します：

1. **getStaticProps**: ビルド時にデータを取得（静的生成）

   ```javascript
   // 商品一覧ページ - ビルド時にデータ取得
   export async function getStaticProps() {
     // APIからデータ取得
     const products = await fetchProducts();

     return {
       props: {
         products,
       },
       // 任意の再生成間隔（秒）を指定
       revalidate: 60, // ISR (Incremental Static Regeneration)
     };
   }
   ```

2. **getServerSideProps**: リクエスト時にサーバーでデータを取得

   ```javascript
   // 商品詳細ページ - リクエスト時にデータ取得
   export async function getServerSideProps(context) {
     const { id } = context.params;
     const product = await fetchProductById(id);

     return {
       props: {
         product,
       },
     };
   }
   ```

3. **useEffect＋fetch**: クライアント側でのデータ取得

   ```javascript
   import { useState, useEffect } from 'react';

   function CartPage() {
     const [cartItems, setCartItems] = useState([]);

     useEffect(() => {
       // クライアント側でのみ実行
       const fetchCartItems = async () => {
         const items = await getCartItems();
         setCartItems(items);
       };

       fetchCartItems();
     }, []);

     return (
       <div>
         {/* カート内容表示 */}
       </div>
     );
   }
   ```

### 1.3.4. MVPでのNext.js実装方針

1. **適切なデータ取得パターンの選択**:
   - 商品カタログ: `getStaticProps` + ISR（商品データは頻繁に変わらない）
   - 商品詳細: `getServerSideProps`（在庫状況などのリアルタイムデータ）
   - カート: クライアントサイド状態管理（ローカルストレージ＋useEffect）

2. **コンポーネント設計**:
   - Atomic Designパターンの簡略版を採用
   - 再利用可能なUI部品を作成（ProductCard, CategoryNavなど）
   - Propsによるデータ受け渡しを基本とする

3. **レスポンシブデザイン**:
   - TailwindCSSを活用したモバイルファーストアプローチ
   - メディアクエリによる段階的レイアウト調整

## 1.4. オブザーバビリティ実装詳細

オブザーバビリティとは、システムの内部状態を外部から観測可能にする能力を指します。「3本の柱」と呼ばれる要素で構成されます：

1. **ログ（Logs）**: システムで何が起きたかの記録
2. **メトリクス（Metrics）**: システムの状態を数値で表現したもの
3. **トレース（Traces）**: リクエストがシステム内をどのように流れるかの追跡

### 1.4.1. AWS SDK v2アプローチの詳細

#### 1.4.1.1. ログ実装 (slog + CloudWatch Logs)

1. **構造化ログの設計**:
   - JSONフォーマットでの一貫したログ形式の実装
   - コンテキスト情報（リクエストID、トレースID）の付与

   ```go
   // structuredLogger.go (概念例)
   func NewStructuredLogger() *slog.Logger {
     opts := &slog.HandlerOptions{
       Level: slog.LevelDebug,
     }
     handler := slog.NewJSONHandler(os.Stdout, opts)
     return slog.New(handler)
   }
   ```

2. **Echo統合**:
   - エンドポイントアクセスログの自動記録
   - リクエスト/レスポンス情報のログ記録

3. **CloudWatch Logs設定**:
   - ロググループの階層設計（サービス別）
   - ログストリーム命名規則の確立

#### 1.4.1.2. メトリクス実装 (CloudWatch Metrics)

1. **REDパターンの実装**:
   - **R**equest Rate: リクエスト数/秒
   - **E**rror Rate: エラー率
   - **D**uration: レイテンシー分布

   ```go
   // metricsCollector.go (概念例)
   func recordAPIMetrics(ctx context.Context, path string, statusCode int, duration time.Duration) {
     // CloudWatch Metricsに送信
     // リクエスト数、エラー数、レイテンシーのメトリクスを記録
   }
   ```

2. **カスタムメトリクス**:
   - ビジネスメトリクス（商品閲覧数、カート追加数など）
   - システムメトリクス（DB接続プール、キューサイズなど）

3. **ディメンション設計**:
   - エンドポイント、ステータスコード、HTTPメソッドによる分類
   - サービス、環境などによる分類

#### 1.4.1.3. トレース実装 (X-Ray SDK)

1. **セグメントとサブセグメント**:
   - HTTPリクエストのルートセグメント作成
   - サブセグメントによる処理の詳細化（DB操作、外部API呼び出しなど）

   ```go
   // xrayMiddleware.go (概念例)
   func XRayMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
     return func(c echo.Context) error {
       // X-Ray セグメント開始
       ctx, seg := xray.BeginSegment(c.Request().Context(), "API-Request")
       defer seg.Close(nil)

       // コンテキストをX-Rayセグメント付きに更新
       c.SetRequest(c.Request().WithContext(ctx))

       return next(c)
     }
   }
   ```

2. **Echo統合**:
   - ミドルウェアによる自動トレース
   - リクエスト/レスポンスヘッダーの記録

3. **sqlboilerトレース**:
   - SQLクエリのサブセグメント生成
   - クエリパラメータとパフォーマンス情報の記録

4. **X-Rayダッシュボードとサービスマップ**:
   - サービス間の依存関係可視化
   - エラーとレイテンシーのホットスポット特定

### 1.4.2. OpenTelemetryアプローチの詳細

#### 1.4.2.1. OpenTelemetry基本構成要素

1. **トレーサーと計装**:
   - グローバルトレーサーのセットアップ
   - 自動計装と手動計装の組み合わせ

   ```go
   // otelSetup.go (概念例)
   func InitTracer() *sdktrace.TracerProvider {
     exporter, err := otlptracehttp.New(context.Background())
     if err != nil {
       log.Fatal(err)
     }

     tp := sdktrace.NewTracerProvider(
       sdktrace.WithSampler(sdktrace.AlwaysSample()),
       sdktrace.WithBatcher(exporter),
     )
     otel.SetTracerProvider(tp)

     return tp
   }
   ```

2. **コンテキスト伝播**:
   - W3C Trace Contextプロトコルの実装
   - サービス間でのコンテキスト伝達

3. **メータと測定**:
   - カウンター、ゲージ、ヒストグラムの実装
   - メトリクスの集約と送信

#### 1.4.2.2. ADOT Collector設定

1. **Collectorデプロイメント**:
   - Fargateタスクとしての実行
   - スケーラビリティ設定

2. **パイプライン設定**:
   - トレースとメトリクスの受信設定
   - 処理とフィルタリング
   - AWS X-Rayへのエクスポート設定

   ```yaml
   # adot-collector-config.yaml (概念例)
   receivers:
     otlp:
       protocols:
         grpc:
         http:

   processors:
     batch:

   exporters:
     awsxray:
       region: "us-west-2"
     awsemf:
       region: "us-west-2"

   service:
     pipelines:
       traces:
         receivers: [otlp]
         processors: [batch]
         exporters: [awsxray]
       metrics:
         receivers: [otlp]
         processors: [batch]
         exporters: [awsemf]
   ```

#### 1.4.2.3. Echo統合とデータベース計装

1. **Echoミドルウェア**:
   - リクエストの自動トレースとコンテキスト伝播
   - エラーハンドリングとステータスコード記録

2. **sqlboiler計装**:
   - データベース操作のトレース
   - クエリパラメータとパフォーマンス情報の記録

3. **カスタム属性とイベント**:
   - ビジネスコンテキストの追加
   - 重要なイベントのマーキング

### 1.4.3. オブザーバビリティ可視化

#### 1.4.3.1. CloudWatch Dashboardの構築

1. **APIパフォーマンスダッシュボード**:
   - リクエスト数、エラー率、レイテンシーのグラフ
   - エンドポイント別パフォーマンス比較

2. **インフラストラクチャダッシュボード**:
   - CPU、メモリ使用率
   - RDSパフォーマンス指標

3. **ビジネスメトリクスダッシュボード**:
   - 商品閲覧数、カート追加率
   - 注文完了率

#### 1.4.3.2. X-Rayによる分析

1. **サービスマップ**:
   - サービス間の依存関係可視化
   - エラーとレイテンシーのホットスポット特定

2. **トレース詳細分析**:
   - 遅いリクエストの根本原因分析
   - エラーチェーンの特定

3. **アノテーションとフィルタリング**:
   - 特定のユースケースやエラーケースの識別
   - パターン分析

### 1.4.4. オブザーバビリティ学習効果の最大化

1. **実験と比較**:
   - AWS SDK v2アプローチとOpenTelemetryアプローチの並行実行
   - パフォーマンス、使いやすさ、柔軟性の比較

2. **障害シナリオの模擬**:
   - 意図的なエラーケースの注入
   - レイテンシーの人為的な増加
   - コネクション問題のシミュレーション

3. **段階的なオブザーバビリティの向上**:
   - 基本的なログから開始
   - メトリクスの追加
   - 最後にトレースの実装

4. **学習記録の文書化**:
   - 設定の詳細記録
   - トラブルシューティング事例の記録
   - パフォーマンス改善の効果測定

## 1.5. AWS環境のコスト最適化

予算を考慮した効率的なAWS環境の構築と運用計画：

1. **ヘルスチェックのための稼働時間最適化**:
   - 平日の業務時間のみの稼働（月〜金、8時間/日）
   - 自動起動/停止スクリプトの実装

2. **リソースサイジングの最適化**:
   - Fargateタスク: 0.5 vCPU、1GB RAM
   - RDS: db.t3.micro（開発段階では十分）

3. **CloudWatch と X-Ray のコスト削減**:
   - 最適なサンプリングレートの設定（10-20%）
   - 重要なリクエストのみのトレース収集

## 1.6. 学習および開発マイルストーン

| 週    | 主要タスク                              | 学習フォーカス                      |
| ----- | --------------------------------------- | ----------------------------------- |
| 1-2   | 開発環境構築、プロジェクト骨組み        | Docker、Go Echo、Next.js基礎        |
| 3-5   | バックエンド実装、AWS SDK v2初期実装    | Go ORM、構造化ログ、X-Ray基礎       |
| 6-8   | フロントエンド実装、API統合             | Next.jsデータフェッチ、ルーティング |
| 9-10  | AWS環境デプロイ、オブザーバビリティ強化 | AWS Fargate、CloudWatch、ALB        |
| 11-12 | OpenTelemetry実装                       | OTEL基礎、ADOT Collector            |
| 13-14 | 比較分析、最適化、ドキュメント作成      | オブザーバビリティパターン比較      |

## 1.7. リスクと対策

| リスク                         | 影響             | 対策                                           |
| ------------------------------ | ---------------- | ---------------------------------------------- |
| AWS関連コストの増大            | 予算オーバー     | タイマー機能による自動停止、コストアラート設定 |
| 技術的な複雑さ                 | 開発の遅延       | フェーズ分けによる段階的実装、優先順位付け     |
| 不明確な要件                   | スコープクリープ | MVPの明確な定義、スコープの都度確認            |
| オブザーバビリティ設定の複雑化 | 構成管理の困難   | IaC (Terraform)による構成管理                  |

## 1.8. 次のステップ

1. 開発環境のセットアップ
2. バックエンドの基本APIの実装
3. データベーススキーマ実装とORM生成
4. フロントエンドの基本画面の実装
5. AWS SDK v2アプローチによるオブザーバビリティ実装

MVPの実装後、段階的にOpenTelemetryアプローチを導入し、両アプローチを比較していきます。
