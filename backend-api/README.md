# AWS Observability E-commerce Backend API

AWS オブザーバビリティの学習を目的としたE-commerce Backend APIです。DDD+クリーンアーキテクチャを採用し、AWS サービスとの統合を重視した設計になっています。

## 🏗️ アーキテクチャ

- **設計パターン**: DDD (Domain-Driven Design) + Clean Architecture
- **データベース**: MySQL
- **AWS統合**: S3 (商品画像保存)、CloudWatch (メトリクス)、X-Ray (分散トレーシング)
- **API仕様**: OpenAPI 3.0 / Swagger UI
- **開発環境**: LocalStack対応

## 📁 ディレクトリ構成

```text
backend-api/
├── cmd/api/                           # アプリケーションエントリーポイント
├── internal/
│   ├── product/                       # 商品ドメイン
│   │   ├── application/               # ユースケース・DTO
│   │   ├── domain/                    # エンティティ・値オブジェクト
│   │   ├── infrastructure/            # データベース・外部サービス実装
│   │   └── presentation/              # REST API ハンドラー
│   ├── query/                         # 複数ドメインをまたぐ読み取り専用クエリ
│   └── shared/                        # 共通コンポーネント
│       ├── infrastructure/
│       │   ├── config/                # 設定管理
│       │   ├── aws/                   # AWS サービス統合
│       │   └── models/                # SQLBoiler生成モデル
│       └── presentation/              # 共通プレゼンテーション層
├── migrations/                        # データベースマイグレーション
├── config.yaml                        # 設定ファイル
└── openapi.yaml                       # API仕様
```

## ⚙️ 設定システム

### 設定の優先順位

1. **環境変数** (最優先)
2. **config.yaml** (ファイル設定)
3. **デフォルト値** (コード内定義)

### 設定ファイル構造

#### `config.yaml` の例

```yaml
# アプリケーション基本設定
app:
  name: "aws-observability-ecommerce"
  version: "1.0.0"
  environment: "development"  # development, staging, production

# サーバー設定
server:
  port: 8000
  host: "0.0.0.0"
  read_timeout: 30    # seconds
  write_timeout: 30   # seconds

# データベース設定
database:
  host: "localhost"
  port: 3306
  user: "root"
  password: "password"
  name: "ecommerce"
  max_open_conns: 25
  max_idle_conns: 25
  conn_max_lifetime: 5  # minutes

# AWS設定
aws:
  use_localstack: true
  region: "ap-northeast-1"
  endpoint: "http://localstack:4566"  # LocalStack使用時
  access_key: "test"    # LocalStack用ダミー値
  secret_key: "test"    # LocalStack用ダミー値
  token: ""
  s3:
    bucket_name: "product-images"
    presigned_url_ttl: 3600  # seconds (1時間)
    use_path_style: true     # LocalStack対応
```

### 環境変数による設定オーバーライド

以下の環境変数で設定をオーバーライドできます：

#### アプリケーション設定

```bash
export APP_NAME="my-ecommerce-api"
export APP_VERSION="2.0.0"
export APP_ENV="production"
```

#### サーバー設定

```bash
export PORT=8080
export HOST="127.0.0.1"
```

#### データベース設定

```bash
export DB_HOST="production-db.example.com"
export DB_PORT=3306
export DB_USER="app_user"
export DB_PASSWORD="secure_password"
export DB_NAME="ecommerce_prod"
```

#### AWS設定

```bash
export AWS_USE_LOCALSTACK=false
export AWS_REGION="ap-northeast-1"
export AWS_ACCESS_KEY_ID="your_access_key"
export AWS_SECRET_ACCESS_KEY="your_secret_key"
export AWS_SESSION_TOKEN="your_session_token"
export AWS_S3_BUCKET_NAME="prod-product-images"
```

## 🚀 実行方法

### 1. 前提条件

- Go 1.24以上
- MySQL 8.0以上
- Docker (LocalStack使用時)

### 2. 依存関係のインストール

```bash
go mod download
```

### 3. データベースのセットアップ

```bash
# MySQLの起動 (Dockerの場合)
docker run --name mysql-ecommerce \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=ecommerce \
  -p 3306:3306 -d mysql:8.0

# マイグレーションの実行
go run cmd/migrate/main.go up
```

### 4. LocalStackの起動 (開発環境)

```bash
# LocalStackの起動
docker run --name localstack \
  -p 4566:4566 \
  -e SERVICES=s3 \
  -e DEBUG=1 \
  -d localstack/localstack

# S3バケットの作成
aws --endpoint-url=http://localhost:4566 s3 mb s3://product-images
```

### 5. アプリケーションの起動

#### 開発環境 (LocalStack使用)

```bash
# config.yamlの設定を使用
go run cmd/api/main.go
```

#### 本番環境

```bash
# 環境変数で設定をオーバーライド
export APP_ENV=production
export AWS_USE_LOCALSTACK=false
export DB_HOST=production-db.example.com
export AWS_S3_BUCKET_NAME=prod-product-images

go run cmd/api/main.go
```

## 🔧 開発者向け設定

### デバッグ用設定

```yaml
app:
  environment: "development"

server:
  port: 8000

database:
  host: "localhost"

aws:
  use_localstack: true
  endpoint: "http://localhost:4566"
```

### ステージング環境用設定

```yaml
app:
  environment: "staging"

server:
  port: 8000

database:
  host: "staging-db.internal.com"
  max_open_conns: 50

aws:
  use_localstack: false
  region: "ap-northeast-1"
  s3:
    bucket_name: "staging-product-images"
```

### プロダクション環境用設定

```yaml
app:
  environment: "production"

server:
  port: 8080
  read_timeout: 60
  write_timeout: 60

database:
  host: "prod-db.internal.com"
  max_open_conns: 100
  max_idle_conns: 50
  conn_max_lifetime: 10

aws:
  use_localstack: false
  region: "ap-northeast-1"
  s3:
    bucket_name: "prod-product-images"
    presigned_url_ttl: 1800  # 30分
```

## 📊 API ドキュメント

アプリケーション起動後、以下のURLでSwagger UIにアクセスできます：

- **Swagger UI**: <http://localhost:8000/swagger>
- **OpenAPI仕様**: <http://localhost:8000/openapi.yaml>
- **ヘルスチェック**: <http://localhost:8000/api/health>

### 主要エンドポイント

- `GET /api/health` - ヘルスチェック
- `GET /api/products` - 商品一覧取得
- `GET /api/products/{id}` - 商品詳細取得
- `GET /api/categories` - カテゴリー一覧取得
- `POST /api/products/{id}/images` - 商品画像アップロード
- `GET /api/products/{id}/images` - 商品画像取得

## 🐛 トラブルシューティング

### データベース接続エラー

```bash
Failed to initialize database: failed to ping database
```

**解決方法:**

1. MySQLが起動していることを確認
2. データベース設定を確認
3. ネットワーク接続を確認

```bash
# 接続テスト
mysql -h localhost -u root -p ecommerce
```

### AWS接続エラー

```bash
Failed to initialize AWS services
```

**解決方法:**

1. LocalStack使用時: LocalStackが起動していることを確認
2. 本番環境: AWS認証情報を確認
3. AWS設定を確認

```bash
# LocalStackの確認
curl http://localhost:4566/health

# AWSクレデンシャルの確認
aws sts get-caller-identity
```

### 設定ファイルが読み込まれない

**設定読み込み順序の確認:**

1. アプリケーションのカレントディレクトリに `config.yaml` があるか
2. 環境変数が正しく設定されているか
3. ログで設定値を確認

### ポート番号の競合

```bash
Failed to start server: listen tcp :8000: bind: address already in use
```

**解決方法:**

```bash
# ポートを使用しているプロセスを確認
lsof -i :8000

# 設定でポート番号を変更
export PORT=8080
```

## 📝 ログとメトリクス

### ログ出力例

```bash
2025/01/XX 12:00:00 Connected to database: localhost:3306/ecommerce
2025/01/XX 12:00:00 AWS config loaded for LocalStack environment
2025/01/XX 12:00:00 Starting server on 0.0.0.0:8000
```

### ヘルスチェックレスポンス

```json
{
  "status": "ok",
  "timestamp": "2025-01-XX T12:00:00Z",
  "version": "1.0.0",
  "uptime": 123456,
  "services": {
    "api": {
      "name": "aws-observability-ecommerce",
      "status": "up"
    }
  },
  "resources": {
    "system": {
      "memory": {
        "allocated": 1048576,
        "total": 2097152,
        "system": 4194304
      },
      "goroutines": 10
    }
  }
}
```

## 🔄 設定の動的変更

一部の設定は環境変数の変更により、アプリケーション再起動後に反映されます：

```bash
# ログレベルの変更
export LOG_LEVEL=debug

# データベース接続プールの調整
export DB_MAX_OPEN_CONNS=50
export DB_MAX_IDLE_CONNS=25

# AWS S3設定の変更
export AWS_S3_PRESIGNED_URL_TTL=7200
```

## 🧪 テスト

```bash
# 単体テスト
go test ./...

# 統合テスト (DockerCompose環境)
docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit

# API テスト
./scripts/test_api.sh
```

## 📚 関連ドキュメント

- [アーキテクチャ設計書](../docs/design/architecture.md)
- [API仕様書](./openapi.yaml)
- [データベース設計書](./docs/database/schema.md)
- [デプロイメントガイド](./docs/deployment/README.md)

## 🤝 貢献

1. Forkしてブランチを作成
2. 変更を実装
3. テストを実行
4. Pull Requestを作成

## 📄 ライセンス

MIT License
