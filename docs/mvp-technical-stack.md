# 1. AWSオブザーバビリティ学習用eコマースアプリ - MVP技術要素

## 1.1. 技術スタック概要

### 1.1.1. フロントエンド

- **Next.js**: Reactベースのフレームワーク（初心者向け構成）
- **TypeScript**: 型安全性を確保
- **TailwindCSS**: 直感的なスタイリング
- **AWS S3 + CloudFront**: 静的アセットのホスティングと配信

### 1.1.2. バックエンド

- **Go言語**: 高パフォーマンスのバックエンド実装
  - **Echo**: 軽量で高速なWebフレームワーク
  - **sqlboiler**: MySQLに特化したタイプセーフなORM
  - **slog**: Go 1.21標準ライブラリのロギングフレームワーク
- **AWS Fargate (ECS)**: コンテナ実行環境
- **MySQL (RDS)**: リレーショナルデータベース
- **AWS ALB**: バックエンドへのトラフィックルーティング（MVPと全機能の両方で使用）

### 1.1.3. オブザーバビリティの2つのアプローチ

1. **aws-sdk-go v2 アプローチ**:
   - **X-Ray SDK for Go v2**: AWSネイティブな分散トレーシング
   - **CloudWatch SDK v2**: メトリクス送信
   - **slog + CloudWatch Logs**: 構造化ログの統合

2. **OpenTelemetry アプローチ**:
   - **OpenTelemetry Go SDK**: ベンダー中立な計装
   - **AWS Distro for OpenTelemetry**: AWS環境への最適化
   - **X-Ray Exporter**: OTELからX-Rayへのトレース送信

## 1.2. Next.js初心者のための基本概念

Next.jsは初心者でも始めやすいReactフレームワークです。以下が基本的な概念です：

### 1.2.1. ページベースのルーティング

Next.jsでは、`pages`ディレクトリ内のファイル構造がそのままURLになります。

```text
📁 pages
  ├── index.js         → / (ホームページ)
  ├── products
  │   ├── index.js     → /products (商品一覧ページ)
  │   └── [id].js      → /products/123 (商品詳細ページ)
  └── cart.js          → /cart (カートページ)
```

### 1.2.2. データ取得の方法

Next.jsでは主に3つのデータ取得方法があります：

- **getStaticProps**: ビルド時にデータを取得（商品カタログなど変更頻度が低いデータに最適）
- **getServerSideProps**: リクエスト時にサーバーサイドでデータを取得（ユーザー固有データなど）
- **useEffect + fetch**: クライアント側でデータを取得（インタラクティブなデータ）

### 1.2.3. コンポーネント構造

Next.jsでは、UIを再利用可能なコンポーネントに分割します：

- **ページコンポーネント**: ルーティングに対応する全画面
- **共通コンポーネント**: ヘッダー、フッター、ナビゲーション
- **機能コンポーネント**: 商品カード、カートアイテム、フォーム

## 1.3. バックエンド技術詳細

### 1.3.1. Goとエコシステム

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

### 1.3.2. コンテナ化とデプロイ

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

### 1.3.3. ALBのルーティング機能

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

### 1.3.4. ALBベースの認証

1. **Cognito統合**:
   - ALBのCognito認証機能を活用
   - サービスごとの認証ルール設定

2. **カスタム認証層**:
   - 認証専用のサービスを作成
   - JWTベースの認証トークン

### 1.3.5. ALBでのマイクロサービス最適化

1. **スティッキーセッション**:
   - セッション維持によるパフォーマンス向上
   - アプリケーションレベルのキャッシュ最適化

2. **段階的なトラフィック移行**:
   - ブルー/グリーンデプロイメント
   - 重み付けルーティングによる段階的移行

3. **サービスメッシュとの連携**:
   - ALBと軽量サービスメッシュの組み合わせ
   - 高度な通信制御とトレース

## 1.4. オブザーバビリティ実装詳細

### 1.4.1. AWS SDK v2 アプローチの実装詳細

#### 1.4.1.1. ログ実装 (slog)

- **構造化ログ設計**: JSONフォーマットでの一貫したログ形式
- **ログレベル**: DEBUG, INFO, WARN, ERROR の適切な使い分け
- **コンテキスト対応**: リクエストIDやトレースIDの伝播
- **CloudWatch Logs統合**: ロググループとストリームの設計

#### 1.4.1.2. X-Ray実装

- **セグメントとサブセグメント**: リクエスト処理の階層的追跡
- **Echoミドルウェア**: HTTPリクエストの自動トレース
- **sqlboilerトレース**: データベースクエリのトレース
- **カスタムアノテーション**: ビジネスコンテキストの追加
- **サンプリングルール**: トレース収集の最適化

#### 1.4.1.3. メトリクス実装

- **REDメトリクス**: Request Rate, Error Rate, Duration
- **カスタムメトリクス**: ビジネスメトリクスの送信
- **ディメンション設計**: メトリクスの多次元分析
- **CloudWatch Dashboards**: カスタムダッシュボード設計

### 1.4.2. OpenTelemetry アプローチの実装詳細

#### 1.4.2.1. OTELの基本コンポーネント

- **トレーサー**: 分散トレースの収集
- **メーター**: メトリクスの収集
- **コンテキスト伝播**: サービス間での文脈維持
- **インストルメンテーション**: 自動および手動の計装

#### 1.4.2.2. インテグレーション

- **Echo統合**: オープンソースのOTEL Echo計装
- **データベース統合**: SQLクエリのトレース
- **X-Ray Exporter**: OTELデータをX-Rayフォーマットに変換
- **ADOT Collector**: テレメトリデータの収集と変換

#### 1.4.2.3. 高度な機能

- **バッチ処理**: テレメトリデータの効率的な送信
- **サンプリング**: 負荷とコスト削減のための最適化
- **フィルタリング**: 関連データの選択的収集
- **マルチプル出力**: 複数のバックエンドへの送信

## 1.5. アーキテクチャ設計

### 1.5.1. MVPアーキテクチャ（シンプル）

```text
ユーザー → CloudFront → S3 (Next.js静的資産)
         ↘ ALB → Fargate (Go Echo API) → MySQL (RDS)
```

このシンプルなアーキテクチャでは:

- フロントエンドとバックエンドが明確に分離
- ALBが直接Fargateサービスにトラフィックをルーティング
- APIは一つのFargateサービスとして実装
- オブザーバビリティは包括的だが、シンプルな構成

### 1.5.2. 全機能アーキテクチャ（拡張性）

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

## 1.6. ローカル開発環境

### 1.6.1. Docker Compose設定

- **GoアプリケーションとMySQLのコンテナ化**
- **Hot Reloadによる開発効率の向上**
- **テストデータの自動投入**

### 1.6.2. 開発ツール

- **air**: Go言語のホットリロード
  - ファイル変更検知と自動再起動
  - 開発サイクルの短縮
- **ogen**: API定義からのコード生成
  - スキーマファーストアプローチでの開発
- **golang-migrate**: マイグレーション管理
  - データベーススキーマの進化管理
- **sqlboiler**: コード生成の自動化

### 1.6.3. オブザーバビリティのローカル設定

- **LocalStackによるAWSサービスエミュレーション**
  - CloudWatch, X-Ray, SNS, SQSなどの模倣
  - 実際のAWS環境に近い開発体験
- **ADOT Collectorのローカル実行**
- **Jaeger UIによるトレース可視化**

## 1.7. デプロイパイプライン

### 1.7.1. CI/CD設計

- **GitHub Actionsによる自動化**
- **テスト、ビルド、デプロイのフロー**
- **インフラストラクチャのコード化（Terraform）**
  - AWS環境の宣言的定義
  - 環境間の一貫性
  - 変更履歴のバージョン管理

### 1.7.2. AWS環境設定

- **マルチ環境戦略（開発、ステージング、本番）**
- **自動スケーリングポリシー**
- **セキュリティ設定とIAMロール**

## 1.8. MVPから全機能への移行パス

MVPから全機能への段階的な移行は、以下のステップで行われます：

1. **オブザーバビリティの強化**
   - 合成モニタリングの追加
   - リアルユーザーモニタリングの導入
   - 異常検出とアラートの高度化

2. **機能拡張と最適化**
   - ユーザー認証システムの導入
   - 支払い処理の統合
   - パフォーマンス最適化

3. **マイクロサービス化**
   - モノリスの論理的分割
   - ALBルーティングルールの設定
   - サービスディスカバリの実装
   - サービス間通信の設計

4. **イベント駆動アーキテクチャ**
   - イベントバスの設定
   - 非同期処理の導入
   - データ一貫性パターンの実装

5. **高度なデータ処理**
   - 検索機能の強化
   - 分析データパイプライン
   - レコメンデーションエンジン
