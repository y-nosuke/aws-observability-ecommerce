# 1. AWSオブザーバビリティ学習用eコマースアプリ - 技術要素

このドキュメントでは、MVPと完成版の技術スタックの違いを一覧化し、各技術要素の特徴と用途を明らかにします。

## 1.1. 目次

- [1. AWSオブザーバビリティ学習用eコマースアプリ - 技術要素](#1-awsオブザーバビリティ学習用eコマースアプリ---技術要素)
  - [1.1. 目次](#11-目次)
  - [1.2. インフラストラクチャ](#12-インフラストラクチャ)
  - [1.3. オブザーバビリティサービス](#13-オブザーバビリティサービス)
  - [1.4. 開発言語・フレームワーク](#14-開発言語フレームワーク)
  - [1.5. オブザーバビリティ実装方法](#15-オブザーバビリティ実装方法)
  - [1.6. 開発環境](#16-開発環境)
  - [1.7. Next.js初心者のための基本概念](#17-nextjs初心者のための基本概念)
    - [1.7.1. ページベースのルーティング](#171-ページベースのルーティング)
    - [1.7.2. データ取得の方法](#172-データ取得の方法)
    - [1.7.3. コンポーネント構造](#173-コンポーネント構造)
  - [1.8. バックエンド技術詳細](#18-バックエンド技術詳細)
    - [1.8.1. Goとエコシステム](#181-goとエコシステム)
    - [1.8.2. コンテナ化とデプロイ](#182-コンテナ化とデプロイ)
    - [1.8.3. ALBのルーティング機能](#183-albのルーティング機能)
    - [1.8.4. ALBベースの認証](#184-albベースの認証)
    - [1.8.5. ALBでのマイクロサービス最適化](#185-albでのマイクロサービス最適化)
  - [1.9. オブザーバビリティ実装詳細](#19-オブザーバビリティ実装詳細)
    - [1.9.1. AWS SDK v2 アプローチの実装詳細](#191-aws-sdk-v2-アプローチの実装詳細)
      - [1.9.1.1. ログ実装 (slog)](#1911-ログ実装-slog)
      - [1.9.1.2. X-Ray実装](#1912-x-ray実装)
      - [1.9.1.3. メトリクス実装](#1913-メトリクス実装)
    - [1.9.2. OpenTelemetry アプローチの実装詳細](#192-opentelemetry-アプローチの実装詳細)
      - [1.9.2.1. OTELの基本コンポーネント](#1921-otelの基本コンポーネント)
      - [1.9.2.2. インテグレーション](#1922-インテグレーション)
      - [1.9.2.3. 高度な機能](#1923-高度な機能)
  - [1.10. アーキテクチャ設計](#110-アーキテクチャ設計)
    - [1.10.1. MVPアーキテクチャ（シンプル）](#1101-mvpアーキテクチャシンプル)
    - [1.10.2. 全機能アーキテクチャ（拡張性）](#1102-全機能アーキテクチャ拡張性)
  - [1.11. ローカル開発環境](#111-ローカル開発環境)
    - [1.11.1. Docker Compose設定](#1111-docker-compose設定)
    - [1.11.2. 開発ツール](#1112-開発ツール)
    - [1.11.3. オブザーバビリティのローカル設定](#1113-オブザーバビリティのローカル設定)
  - [1.12. デプロイパイプライン](#112-デプロイパイプライン)
    - [1.12.1. CI/CD設計](#1121-cicd設計)
    - [1.12.2. AWS環境設定](#1122-aws環境設定)
  - [1.13. MVPと完成版の主な違い](#113-mvpと完成版の主な違い)

## 1.2. インフラストラクチャ

| カテゴリ               | 技術/サービス                   | MVP              | 主な用途・特徴                                |
| ---------------------- | ------------------------------- | ---------------- | --------------------------------------------- |
| **フロントエンド**     | Next.js                         | ✅                | Reactベースのフレームワーク、SSR/SSG対応      |
|                        | S3 + CloudFront                 | ✅                | 静的アセットのホスティングと配信              |
| **API管理**            | Application Load Balancer (ALB) | ✅                | バックエンドへのトラフィックルーティング      |
|                        | API Gateway                     | ❌                | REST APIの一元管理、認証統合                  |
| **コンピューティング** | Fargate (ECS)                   | ✅                | コンテナ実行環境、Goバックエンド実行          |
|                        | Lambda                          | ⚪️ (基本実装のみ) | サーバーレス関数、イベント処理                |
| **データストア**       | RDS (MySQL)                     | ✅                | リレーショナルデータベース、主要データの保存  |
|                        | DynamoDB                        | ❌                | NoSQLデータベース、高スケーラブルなデータ処理 |
| **ストレージ**         | S3                              | ✅                | オブジェクトストレージ、画像・ファイル保存    |
| **メッセージング**     | EventBridge                     | ❌                | イベント駆動型アーキテクチャの中核            |
|                        | SNS                             | ❌                | 通知サービス、イベント配信                    |
|                        | SQS                             | ❌                | メッセージキュー、処理の非同期化              |
| **通知**               | SES                             | ❌                | Eメール送信サービス                           |

## 1.3. オブザーバビリティサービス

| カテゴリ                   | 技術/サービス                               | MVP              | 主な用途・特徴                                     |
| -------------------------- | ------------------------------------------- | ---------------- | -------------------------------------------------- |
| **モニタリング基盤**       | CloudWatch Logs                             | ✅                | アプリケーションログの集約、保存、検索             |
|                            | CloudWatch Metrics                          | ✅                | システムとアプリケーションメトリクスの収集         |
|                            | CloudWatch Alarms                           | ✅                | リソース使用率とエラー率の監視                     |
|                            | CloudWatch Dashboards                       | ✅                | カスタムダッシュボードの作成、指標の可視化         |
|                            | X-Ray                                       | ✅                | 分散トレーシング、サービス間の依存関係可視化       |
|                            | CloudWatch Synthetics                       | ❌                | エンドツーエンドのユーザーフロー監視               |
| **高度な分析**             | DevOps Guru                                 | ❌                | ML駆動の異常検出と問題の自動診断                   |
|                            | Amazon Managed Service for Prometheus (AMP) | ❌                | 高度なメトリクス収集とクエリ                       |
|                            | Amazon Managed Grafana (AMG)                | ❌                | 高度なダッシュボードと可視化                       |
| **標準化**                 | AWS Distro for OpenTelemetry (ADOT)         | ⚪️ (基本実装のみ) | 標準化されたテレメトリデータの収集                 |
| **監査・コンプライアンス** | CloudTrail                                  | ❌                | API呼び出しの監査とコンプライアンス                |
|                            | AWS Config                                  | ❌                | インフラストラクチャの構成変更の追跡               |
| **データ分析**             | OpenSearch Service                          | ❌                | 大規模なログデータの検索と分析                     |
|                            | Athena                                      | ❌                | ログデータへのSQLクエリとアドホック分析            |
| **ネットワーク監視**       | VPC Flow Logs                               | ❌                | ネットワークトラフィックの可視化、セキュリティ分析 |
| **テスト・実験**           | Fault Injection Service                     | ❌                | カオスエンジニアリングとレジリエンステスト         |
| **フロントエンド監視**     | CloudWatch RUM                              | ❌                | 実ユーザー体験とパフォーマンスの測定               |
| **サービスヘルス**         | AWS Health Dashboard                        | ❌                | AWSサービスのヘルスステータス監視                  |

## 1.4. 開発言語・フレームワーク

| カテゴリ           | 技術/サービス    | MVP              | 主な用途・特徴                                   |
| ------------------ | ---------------- | ---------------- | ------------------------------------------------ |
| **フロントエンド** | TypeScript       | ✅                | 静的型付けによる堅牢なコード開発                 |
|                    | Next.js          | ✅                | Reactフレームワーク、ルーティング、SSR           |
|                    | React            | ✅                | UIコンポーネントライブラリ                       |
|                    | TailwindCSS      | ✅                | ユーティリティファーストのCSSフレームワーク      |
| **バックエンド**   | Go (Fargate)     | ✅                | 高性能なバックエンド言語、Echo Webフレームワーク |
|                    | sqlboiler        | ✅                | MySQLに特化したタイプセーフなORM                 |
|                    | OpenAPI/ogen     | ✅                | API定義とコード生成                              |
|                    | slog             | ✅                | Go標準の構造化ログライブラリ                     |
|                    | Node.js (Lambda) | ⚪️ (基本実装のみ) | サーバーレス関数の実装言語                       |

## 1.5. オブザーバビリティ実装方法

| アプローチ                   | 技術/サービス                | MVP              | 主な用途・特徴                  |
| ---------------------------- | ---------------------------- | ---------------- | ------------------------------- |
| **aws-sdk-go v2 アプローチ** | X-Ray SDK for Go v2          | ✅                | AWSネイティブな分散トレーシング |
|                              | CloudWatch SDK v2            | ✅                | メトリクス送信                  |
|                              | slog + CloudWatch Logs       | ✅                | 構造化ログの統合                |
| **OpenTelemetry アプローチ** | OpenTelemetry Go SDK         | ⚪️ (基本実装のみ) | ベンダー中立な計装              |
|                              | AWS Distro for OpenTelemetry | ⚪️ (基本実装のみ) | AWS環境への最適化               |
|                              | X-Ray Exporter               | ⚪️ (基本実装のみ) | OTELからX-Rayへのトレース送信   |

## 1.6. 開発環境

| カテゴリ         | 技術/サービス           | MVP              | 主な用途・特徴                        |
| ---------------- | ----------------------- | ---------------- | ------------------------------------- |
| **コンテナ化**   | Docker                  | ✅                | アプリケーションのコンテナ化          |
|                  | Docker Compose          | ✅                | 複数コンテナの管理と調整              |
| **ローカル開発** | LocalStack              | ✅                | AWSサービスのローカルエミュレーション |
|                  | Traefik                 | ✅                | リバースプロキシ、ルーティング        |
| **デプロイ**     | GitHub Actions          | ✅                | CI/CDパイプライン                     |
|                  | Terraform               | ⚪️ (基本実装のみ) | インフラのコード化                    |
| **テスト**       | Fault Injection Service | ❌                | カオスエンジニアリングテスト          |

## 1.7. Next.js初心者のための基本概念

Next.jsは初心者でも始めやすいReactフレームワークです。以下が基本的な概念です：

### 1.7.1. ページベースのルーティング

Next.jsでは、`pages`ディレクトリ内のファイル構造がそのままURLになります。

```text
📁 pages
  ├── index.js         → / (ホームページ)
  ├── products
  │   ├── index.js     → /products (商品一覧ページ)
  │   └── [id].js      → /products/123 (商品詳細ページ)
  └── cart.js          → /cart (カートページ)
```

### 1.7.2. データ取得の方法

Next.jsでは主に3つのデータ取得方法があります：

- **getStaticProps**: ビルド時にデータを取得（商品カタログなど変更頻度が低いデータに最適）
- **getServerSideProps**: リクエスト時にサーバーサイドでデータを取得（ユーザー固有データなど）
- **useEffect + fetch**: クライアント側でデータを取得（インタラクティブなデータ）

### 1.7.3. コンポーネント構造

Next.jsでは、UIを再利用可能なコンポーネントに分割します：

- **ページコンポーネント**: ルーティングに対応する全画面
- **共通コンポーネント**: ヘッダー、フッター、ナビゲーション
- **機能コンポーネント**: 商品カード、カートアイテム、フォーム

## 1.8. バックエンド技術詳細

### 1.8.1. Goとエコシステム

- **Go言語**: コンパイル型、静的型付け、高パフォーマンス
- **Echo**: 軽量で高速なWebフレームワーク
  - ミドルウェアサポート: ログ、認証、エラーハンドリング
  - 依存性注入の仕組み
  - バリデーション機能
- **sqlboiler**: コード生成型ORM
  - データベーススキーマからタイプセーフなモデルを生成
  - トランザクション管理
  - 複雑なクエリをタイプセーフに構築
- **golang-migrate**: データベースマイグレーション管理
  - バージョン管理されたスキーマ変更
  - ロールバック対応
  - CLI/プログラムからの実行
- **ogen**: OpenAPI仕様からのコード生成
  - API定義からのタイプセーフなクライアント/サーバーコード生成
  - リクエスト/レスポンスの検証
  - エンドポイントの自動実装
- **air**: ホットリローディング開発ツール
  - ファイル変更検知による自動再ビルド
  - 開発効率の向上
  - カスタマイズ可能な設定

### 1.8.2. コンテナ化とデプロイ

- **Docker**: Goアプリケーションのコンテナ化
  - マルチステージビルドによる軽量イメージ
  - distrolessベースイメージの利用
- **AWS Fargate**: サーバーレスコンテナ実行環境
  - オートスケーリング設定
  - CloudWatch連携による監視
- **AWS RDS (MySQL)**: マネージドデータベースサービス
  - バックアップと復元
  - パフォーマンスインサイト
- **terraform**: インフラストラクチャのコード化
  - AWS環境の宣言的定義
  - 環境の再現性確保
  - 状態管理とバージョン管理
- **localstack**: ローカル開発用AWSエミュレーター
  - ローカル環境でのAWSサービスエミュレーション
  - テスト環境の簡易構築
  - AWS SDKとの互換性

### 1.8.3. ALBのルーティング機能

1. **パスベースルーティング**:
   - `/users/*` → ユーザー管理サービス
   - `/products/*` → 商品サービス
   - `/orders/*` → 注文サービス
   - `/notifications/*` → 通知サービス

2. **ホストベースルーティング**:
   - `users.api.example.com` → ユーザー管理サービス
   - `products.api.example.com` → 商品サービス
   - `orders.api.example.com` → 注文サービス

3. **ターゲットグループの設計**:
   - サービスごとに個別のターゲットグループを作成
   - Fargateサービスとの自動統合設定
   - ヘルスチェックの個別設定

### 1.8.4. ALBベースの認証

1. **Cognito統合**:
   - ALBのCognito認証機能を活用
   - サービスごとの認証ルール設定

2. **カスタム認証層**:
   - 認証専用のサービスを作成
   - JWTベースの認証トークン

### 1.8.5. ALBでのマイクロサービス最適化

1. **スティッキーセッション**:
   - セッション維持によるパフォーマンス向上
   - アプリケーションレベルのキャッシュ最適化

2. **段階的なトラフィック移行**:
   - ブルー/グリーンデプロイメント
   - 重み付けルーティングによる段階的移行

3. **サービスメッシュとの連携**:
   - ALBと軽量サービスメッシュの組み合わせ
   - 高度な通信制御とトレース

## 1.9. オブザーバビリティ実装詳細

### 1.9.1. AWS SDK v2 アプローチの実装詳細

#### 1.9.1.1. ログ実装 (slog)

- **構造化ログ設計**: JSONフォーマットでの一貫したログ形式
- **ログレベル**: DEBUG, INFO, WARN, ERROR の適切な使い分け
- **コンテキスト対応**: リクエストIDやトレースIDの伝播
- **CloudWatch Logs統合**: ロググループとストリームの設計

#### 1.9.1.2. X-Ray実装

- **セグメントとサブセグメント**: リクエスト処理の階層的追跡
- **Echoミドルウェア**: HTTPリクエストの自動トレース
- **sqlboilerトレース**: データベースクエリのトレース
- **カスタムアノテーション**: ビジネスコンテキストの追加
- **サンプリングルール**: トレース収集の最適化

#### 1.9.1.3. メトリクス実装

- **REDメトリクス**: Request Rate, Error Rate, Duration
- **カスタムメトリクス**: ビジネスメトリクスの送信
- **ディメンション設計**: メトリクスの多次元分析
- **CloudWatch Dashboards**: カスタムダッシュボード設計

### 1.9.2. OpenTelemetry アプローチの実装詳細

#### 1.9.2.1. OTELの基本コンポーネント

- **トレーサー**: 分散トレースの収集
- **メーター**: メトリクスの収集
- **コンテキスト伝播**: サービス間での文脈維持
- **インストルメンテーション**: 自動および手動の計装

#### 1.9.2.2. インテグレーション

- **Echo統合**: オープンソースのOTEL Echo計装
- **データベース統合**: SQLクエリのトレース
- **X-Ray Exporter**: OTELデータをX-Rayフォーマットに変換
- **ADOT Collector**: テレメトリデータの収集と変換

#### 1.9.2.3. 高度な機能

- **バッチ処理**: テレメトリデータの効率的な送信
- **サンプリング**: 負荷とコスト削減のための最適化
- **フィルタリング**: 関連データの選択的収集
- **マルチプル出力**: 複数のバックエンドへの送信

## 1.10. アーキテクチャ設計

### 1.10.1. MVPアーキテクチャ（シンプル）

```text
ユーザー → CloudFront → S3 (Next.js静的資産)
         ↘ ALB → Fargate (Go Echo API) → MySQL (RDS)
```

このシンプルなアーキテクチャでは:

- フロントエンドとバックエンドが明確に分離
- ALBが直接Fargateサービスにトラフィックをルーティング
- APIは一つのFargateサービスとして実装
- オブザーバビリティは包括的だが、シンプルな構成

### 1.10.2. 全機能アーキテクチャ（拡張性）

```text
ユーザー → CloudFront → S3 (Next.js静的資産)
         ↘ ALB → ┌→ ユーザー管理サービス (Fargate) → RDS
                 ├→ 商品サービス (Fargate) → RDS
                 ├→ 注文サービス (Fargate) → RDS
                 └→ 通知サービス (Fargate) → SQS/SNS
```

全機能版では:

- マイクロサービスアーキテクチャへの移行
- ALBのパスベースルーティングによる複数サービスへのトラフィック分散
- 各サービスの独立したスケーリング
- イベント駆動アーキテクチャの導入

## 1.11. ローカル開発環境

### 1.11.1. Docker Compose設定

- **GoアプリケーションとMySQLのコンテナ化**
- **Hot Reloadによる開発効率の向上**
- **テストデータの自動投入**

### 1.11.2. 開発ツール

- **air**: Go言語のホットリロード
  - ファイル変更検知と自動再起動
  - 開発サイクルの短縮
- **ogen**: API定義からのコード生成
  - スキーマファーストアプローチでの開発
- **golang-migrate**: マイグレーション管理
  - データベーススキーマの進化管理
- **sqlboiler**: コード生成の自動化

### 1.11.3. オブザーバビリティのローカル設定

- **LocalStackによるAWSサービスエミュレーション**
  - CloudWatch, X-Ray, SNS, SQSなどの模倣
  - 実際のAWS環境に近い開発体験
- **ADOT Collectorのローカル実行**
- **Jaeger UIによるトレース可視化**

## 1.12. デプロイパイプライン

### 1.12.1. CI/CD設計

- **GitHub Actionsによる自動化**
- **テスト、ビルド、デプロイのフロー**
- **インフラストラクチャのコード化（Terraform）**
  - AWS環境の宣言的定義
  - 環境間の一貫性
  - 変更履歴のバージョン管理

### 1.12.2. AWS環境設定

- **マルチ環境戦略（開発、ステージング、本番）**
- **自動スケーリングポリシー**
- **セキュリティ設定とIAMロール**

## 1.13. MVPと完成版の主な違い

1. **インフラストラクチャの拡張**:
   - MVPではシンプルなALB + Fargate + RDSの構成
   - 完成版ではAPI Gateway、DynamoDB、メッセージングサービス(SNS/SQS/EventBridge)が追加され、よりイベント駆動型のマイクロサービスアーキテクチャに発展

2. **オブザーバビリティの発展**:
   - MVPでは基本的なCloudWatch LogsとX-Rayの統合
   - 完成版ではADOT、RUM、Synthetics、DevOps Guruなど高度なモニタリングツールが追加され、包括的なオブザーバビリティが実現

3. **開発アプローチの進化**:
   - MVPではAWS SDK v2アプローチを中心に実装
   - 完成版ではOpenTelemetryアプローチへの移行と両アプローチの比較学習が可能

4. **サーバーレスの高度な活用**:
   - MVPではLambdaの基本的な使用（画像処理など）
   - 完成版ではイベント駆動型アーキテクチャ、Step Functions、非同期処理など高度なサーバーレスパターンを実装
