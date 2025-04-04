# 1. Week 1 - Day 4: Traefikによるリバースプロキシの設定

## 1.1. 目次

- [1. Week 1 - Day 4: Traefikによるリバースプロキシの設定](#1-week-1---day-4-traefikによるリバースプロキシの設定)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. Traefik用のディレクトリとファイルの作成](#141-traefik用のディレクトリとファイルの作成)
    - [1.4.2. Traefikの基本設定ファイルの作成](#142-traefikの基本設定ファイルの作成)
    - [1.4.3. ダイナミックコンフィグの作成](#143-ダイナミックコンフィグの作成)
      - [1.4.3.1. middlewares.yml](#1431-middlewaresyml)
      - [1.4.3.2. tls.yml](#1432-tlsyml)
    - [1.4.4. Docker Composeへの統合](#144-docker-composeへの統合)
    - [1.4.5. サービスのラベル設定](#145-サービスのラベル設定)
    - [1.4.6. Traefikの起動と動作確認](#146-traefikの起動と動作確認)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. Traefikの基本アーキテクチャ](#161-traefikの基本アーキテクチャ)
    - [1.6.2. 開発環境向け設定ファイルの詳細説明](#162-開発環境向け設定ファイルの詳細説明)
      - [1.6.2.1. traefik.yml（静的設定）](#1621-traefikyml静的設定)
      - [1.6.2.2. middlewares.yml（ミドルウェア設定）](#1622-middlewaresymlミドルウェア設定)
      - [1.6.2.3. tls.yml（開発環境では使用しないがテンプレートとして存在）](#1623-tlsyml開発環境では使用しないがテンプレートとして存在)
    - [1.6.3. Docker Compose ラベルの詳細説明](#163-docker-compose-ラベルの詳細説明)
    - [1.6.4. ホスト名ベースのルーティング](#164-ホスト名ベースのルーティング)
    - [1.6.5. セキュリティヘッダーと設定の重要性](#165-セキュリティヘッダーと設定の重要性)
    - [1.6.6. ミドルウェア連鎖と適用方法](#166-ミドルウェア連鎖と適用方法)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. 開発環境と本番環境の設定の違い](#171-開発環境と本番環境の設定の違い)
      - [1.7.1.1. AWS環境での一般的な構成](#1711-aws環境での一般的な構成)
      - [1.7.1.2. オンプレミスや他のクラウド環境でTraefikを直接公開する場合](#1712-オンプレミスや他のクラウド環境でtraefikを直接公開する場合)
      - [1.7.1.3. 本番環境に向けた一般的な変更点](#1713-本番環境に向けた一般的な変更点)
    - [1.7.2. カスタムミドルウエアの作成方法](#172-カスタムミドルウエアの作成方法)
    - [1.7.3. 本番環境でのデプロイパターン](#173-本番環境でのデプロイパターン)
      - [1.7.3.1. コンテナオーケストレーションプラットフォームでの活用](#1731-コンテナオーケストレーションプラットフォームでの活用)
        - [1.7.3.1.1. ECS（Elastic Container Service）での使用](#17311-ecselastic-container-serviceでの使用)
        - [1.7.3.1.2. Kubernetes環境での使用](#17312-kubernetes環境での使用)
      - [1.7.3.2. HA構成（高可用性）での設定](#1732-ha構成高可用性での設定)
      - [1.7.3.3. Blue/Greenデプロイの実装](#1733-bluegreenデプロイの実装)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: Traefikダッシュボードにアクセスできない](#181-問題1-traefikダッシュボードにアクセスできない)
    - [1.8.2. 問題2: バックエンドサービスへのルーティングが機能しない](#182-問題2-バックエンドサービスへのルーティングが機能しない)
    - [1.8.3. 問題3: HTTPSリダイレクトで無限リダイレクトループが発生する](#183-問題3-httpsリダイレクトで無限リダイレクトループが発生する)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- Traefikを使用したリバースプロキシの基本概念と設定方法を理解する
- ホスト名ベースのルーティングを実装し、複数のサービスを単一のエントリーポイントで提供する
- バックエンドサービスとNext.jsフロントエンドを連携させるための適切な設定
- セキュリティヘッダーの設定とHTTPSリダイレクトの実装
- Docker Compose環境でのTraefik統合パターンを習得する

## 1.3. 【準備】

Traefikによるリバースプロキシを設定する前に、以下の環境が整っていることを確認します。

### 1.3.1. チェックリスト

- [ ] Docker および Docker Compose がインストールされている
- [ ] プロジェクトの基本構造が作成されている
- [ ] プロジェクトルートディレクトリに移動済み
- [ ] Go/Echo バックエンドの基本構造が実装されている
- [ ] Next.js フロントエンドプロジェクトが作成されている
- [ ] `docker-compose.yml` ファイルが作成されている
- [ ] `.env` ファイルが設定されている
- [ ] インターネット接続が利用可能である（Traefikイメージのダウンロード用）

## 1.4. 【手順】

### 1.4.1. Traefik用のディレクトリとファイルの作成

開発環境におけるTraefikの設定ファイルを保存するためのディレクトリ構造を作成します。

```bash
# プロジェクトルートディレクトリにいることを確認
mkdir -p ./infra/traefik/{config,dynamic}
touch ./infra/traefik/config/traefik.yml
touch ./infra/traefik/dynamic/{middlewares.yml,tls.yml}
```

上記のコマンドにより、以下のディレクトリ構造が作成されます：

```text
./infra/
└── traefik/
    ├── config/
    │   └── traefik.yml     # 静的な設定ファイル
    └── dynamic/
        ├── middlewares.yml # ミドルウェア設定
        └── tls.yml         # TLS設定
```

### 1.4.2. Traefikの基本設定ファイルの作成

`traefik.yml`ファイルに基本的な静的設定を記述します。これはTraefikの起動時に読み込まれる設定です。

以下の内容を`./infra/traefik/config/traefik.yml`ファイルに記述します：

```yaml
# 全体的な設定
global:
  checkNewVersion: true
  sendAnonymousUsage: false

# APIとダッシュボードの設定
api:
  dashboard: true
  insecure: true  # 開発環境のみ。本番では必ずfalseに変更すること

# エントリーポイントの設定
entryPoints:
  web:
    address: ":80"
    # 開発環境ではコメントアウト、本番環境ではHTTPSへのリダイレクトを有効化
    # http:
    #   redirections:
    #     entryPoint:
    #       to: websecure
    #       scheme: https

  websecure:
    address: ":443"

# プロバイダー設定
providers:
  # Docker設定
  docker:
    endpoint: "unix:///var/run/docker.sock"
    exposedByDefault: false
    network: traefik-network

  # ファイルプロバイダー設定
  file:
    directory: "/etc/traefik/dynamic"
    watch: true

# 証明書リゾルバー設定（開発環境ではセルフサイン証明書を使用）
certificatesResolvers:
  default:
    # 開発環境なのでLet's Encryptは使わない
    # letsencrypt:
    #   caServer: "https://acme-v02.api.letsencrypt.org/directory"
    #   email: "admin@example.com"
    #   storage: "/etc/traefik/data/acme.json"
    #   httpChallenge:
    #     entryPoint: web

# ログ設定
log:
  level: "INFO"  # DEBUG, INFO, WARN, ERROR, FATAL, PANIC

# アクセスログ設定
accessLog:
  filePath: "/var/log/traefik/access.log"
  bufferingSize: 100

# プロメテウスメトリクス
metrics:
  prometheus: {}
```

### 1.4.3. ダイナミックコンフィグの作成

次に、動的な設定ファイルを作成します。これらはTraefikの実行中に変更できる設定です。

#### 1.4.3.1. middlewares.yml

`docker/traefik/config/middlewares.yml`ファイルを作成し、共通のミドルウェア設定を記述します：

```yaml
http:
  middlewares:
    # セキュリティヘッダーの設定
    secure-headers:
      headers:
        frameDeny: true
        sslRedirect: false  # 開発環境ではfalse
        browserXssFilter: true
        contentTypeNosniff: true
        forceSTSHeader: true
        stsIncludeSubdomains: true
        stsPreload: true
        stsSeconds: 31536000
        customFrameOptionsValue: "SAMEORIGIN"
        customRequestHeaders:
          X-Forwarded-Proto: "https"

    # 開発用CORS設定
    cors:
      headers:
        accessControlAllowMethods:
          - GET
          - POST
          - PUT
          - DELETE
          - OPTIONS
        accessControlAllowHeaders:
          - "*"
        accessControlAllowOriginList:
          - "http://localhost"
          - "http://localhost:3000"
          - "http://frontend.localhost"
          - "http://backend.localhost"
        accessControlMaxAge: 100
        accessControlAllowCredentials: true
        addVaryHeader: true
```

#### 1.4.3.2. tls.yml

以下の内容を`./infra/traefik/dynamic/tls.yml`ファイルに記述します：

```yaml
# TLS設定

tls:
  options:
    default:
      minVersion: "VersionTLS12"
      cipherSuites:
        - "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256"
        - "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
        - "TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305"
        - "TLS_AES_128_GCM_SHA256"
        - "TLS_AES_256_GCM_SHA384"
        - "TLS_CHACHA20_POLY1305_SHA256"
      curvePreferences:
        - "CurveP521"
        - "CurveP384"
      sniStrict: true
```

### 1.4.4. Docker Composeへの統合

Docker ComposeファイルにTraefikサービスを追加します。プロジェクトルートの`compose.yml`ファイルを以下のように編集します：

```yaml
services:
  # Traefikリバースプロキシサービス
  traefik:
    image: traefik:latest
    container_name: traefik
    restart: unless-stopped
    security_opt:
      - no-new-privileges:true
    ports:
      - "80:80" # HTTP
      - "443:443" # HTTPS
      - "8080:8080" # Dashboard
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./infra/traefik/config/traefik.yml:/etc/traefik/traefik.yml:ro
      - ./infra/traefik/dynamic:/etc/traefik/dynamic:ro
      - ./logs/traefik:/var/log/traefik
    networks:
      - ecommerce-network
    deploy:
      resources:
        limits:
          memory: 256M

  # バックエンドAPIサービス
  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    container_name: backend
    restart: unless-stopped
    expose:
      - "8000"
    environment:
      - APP_NAME=aws-observability-ecommerce
      - APP_VERSION=1.0.0
      - APP_ENV=development
      - PORT=8000
    volumes:
      - ./backend:/app # ホットリロード用
    depends_on:
      mysql:
        condition: service_healthy
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8000/api/health"]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 10s
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.backend.rule=Host(`backend.localhost`)"
      - "traefik.http.routers.backend.entrypoints=web"
      - "traefik.http.services.backend.loadbalancer.server.port=8000"
      - "traefik.http.routers.backend.middlewares=secure-headers@file,cors@file"
    networks:
      - ecommerce-network

  # 顧客向けフロントエンドNext.jsサービス
  frontend-customer:
    build:
      context: ./frontend-customer
      dockerfile: Dockerfile
    container_name: frontend-customer
    restart: unless-stopped
    environment:
      - NEXT_PUBLIC_API_URL=http://backend:8000/api
    volumes:
      - ./frontend-customer:/app # ホットリロード用
      - /app/node_modules # node_modulesはコンテナ内のまま
      - /app/.next # .nextディレクトリはコンテナ内のまま
    expose:
      - "3000"
    depends_on:
      - backend
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.customer.rule=Host(`customer.localhost`)"
      - "traefik.http.routers.customer.entrypoints=web"
      - "traefik.http.services.customer.loadbalancer.server.port=3000"
      - "traefik.http.routers.customer.middlewares=secure-headers@file,cors@file"
    networks:
      - ecommerce-network
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/"]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 15s
    deploy:
      resources:
        limits:
          memory: 512M

  # 管理者向けフロントエンドNext.jsサービス
  frontend-admin:
    build:
      context: ./frontend-admin
      dockerfile: Dockerfile
    container_name: frontend-admin
    restart: unless-stopped
    environment:
      - NEXT_PUBLIC_API_URL=http://backend:8000/api
    volumes:
      - ./frontend-admin:/app
      - /app/node_modules
    expose:
      - "3000"
    depends_on:
      - backend
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.admin.rule=Host(`admin.localhost`)"
      - "traefik.http.routers.admin.entrypoints=web"
      - "traefik.http.services.admin.loadbalancer.server.port=3000"
      - "traefik.http.routers.admin.middlewares=secure-headers@file,cors@file"
    networks:
      - ecommerce-network
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:3000/"]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 15s
    deploy:
      resources:
        limits:
          memory: 512M


  # データベースサービスなどがここに追加される
  # ...

networks:
  ecommerce-network:
    driver: bridge
    name: ecommerce-network
```

### 1.4.5. サービスのラベル設定

Traefikは、Dockerコンテナのラベルを使用して設定を行います：

- `traefik.enable=true` - Traefikでこのサービスを管理することを有効化
- `traefik.http.routers.[name].rule=Host(...)` - ホスト名ベースのルーティングルール
- `traefik.http.routers.[name].entrypoints=...` - 使用するエントリーポイント
- `traefik.http.services.[name].loadbalancer.server.port=...` - コンテナ内部のポート
- `traefik.http.routers.[name].middlewares=...` - 適用するミドルウェア

上記のDocker Composeファイルでは、バックエンドとフロントエンドのサービスに対して次のようなラベルを設定しています：

1. **バックエンドサービスのラベル設定**:
   - `traefik.enable=true`: Traefikでこのサービスを管理することを有効化
   - `traefik.http.routers.backend.rule=Host(`backend.localhost`)`: ホスト名ベースのルーティングルール
   - `traefik.http.routers.backend.entrypoints=web`: 使用するエントリーポイント
   - `traefik.http.services.backend.loadbalancer.server.port=8000`: コンテナ内部のポート
   - CORS関連のミドルウェア設定

2. **顧客向けフロントエンドサービスのラベル設定**:
   - `traefik.enable=true`: Traefikでこのサービスを管理することを有効化
   - `traefik.http.routers.frontend.rule=Host(`customer.localhost`)`: ホスト名ベースのルーティングルール
   - `traefik.http.routers.frontend.entrypoints=web`: 使用するエントリーポイント
   - `traefik.http.services.frontend.loadbalancer.server.port=3000`: コンテナ内部のポート

3. **管理者向けフロントエンドサービスのラベル設定**:
   - `traefik.enable=true`: Traefikでこのサービスを管理することを有効化
   - `traefik.http.routers.frontend.rule=Host(`admin.localhost`)`: ホスト名ベースのルーティングルール
   - `traefik.http.routers.frontend.entrypoints=web`: 使用するエントリーポイント
   - `traefik.http.services.frontend.loadbalancer.server.port=3000`: コンテナ内部のポート

### 1.4.6. Traefikの起動と動作確認

すべての設定が完了したら、Docker Composeを使用してサービスを起動します：

```bash
docker compose up -d
```

起動後、以下のエンドポイントにアクセスして動作を確認できます：

1. Traefikダッシュボード: <http://localhost:8080>
2. バックエンドAPI: <http://backend.localhost>
3. 顧客向けフロントエンドアプリ: <http://customer.localhost>
4. 管理者向けフロントエンドアプリ: <http://admin.localhost>

ホスト名を使用するため、ローカルのホストファイルを編集する必要があります：

```bash
# Linux/Mac
sudo vi /etc/hosts

# Windows
# 管理者権限でメモ帳を開き、C:\Windows\System32\drivers\etc\hosts を編集
```

以下の行を追加します：

```text
127.0.0.1 customer.localhost admin.localhost backend.localhost
```

## 1.5. 【確認ポイント】

実装が正しく完了したことを確認するためのチェックリストです：

- [ ] Traefikコンテナが正常に起動しているか (`docker ps` で確認)
- [ ] Traefikダッシュボードにアクセスできる (<http://localhost:8080>)
- [ ] ダッシュボードにすべてのサービスのルートが表示されている
- [ ] バックエンドAPIにアクセスできる (<http://backend.localhost>)
- [ ] 顧客向けフロントエンドアプリにアクセスできる (<http://customer.localhost>)
- [ ] 管理者向けフロントエンドアプリにアクセスできる (<http://admin.localhost>)
- [ ] CORS設定が機能している (フロントエンドからバックエンドへのリクエストが成功する)
- [ ] セキュリティヘッダーが適切に設定されている (ブラウザのDevToolsで確認)
- [ ] Traefikのアクセスログが生成されている (`./logs/traefik/access.log` を確認)

すべてのチェックポイントを満たしていれば、Traefikによるリバースプロキシの設定は正常に完了しています。

## 1.6. 【詳細解説】

### 1.6.1. Traefikの基本アーキテクチャ

Traefikは、モダンなリバースプロキシとロードバランサーで、特にマイクロサービスとDocker環境に適しています。以下はアーキテクチャの主要コンポーネントです：

1. **エントリーポイント (Entrypoints)**
   - トラフィックの入口となるポートとプロトコル
   - 一般的に`web`（HTTP）と`websecure`（HTTPS）の2つが定義される

2. **ルーター (Routers)**
   - リクエストを適切なサービスに記述するルールを定義
   - ホスト名、パス、ヘッダーなどの条件に基づくマッチング

3. **ミドルウェア (Middlewares)**
   - リクエスト/レスポンスを変更・加工するコンポーネント
   - 例：CORS設定、セキュリティヘッダー追加、認証、リダイレクト、ヘッダー操作、レート制限など

4. **サービス (Services)**
   - 実際のバックエンドサービス（コンテナ）
   - 負荷分散、ヘルスチェックなどの機能を提供

5. **プロバイダー (Providers)**
   - Traefikに設定情報を供給するソース
   - Docker、Kubernetes、ファイルなどから動的に設定を読み込む

これらのコンポーネントが連携して、最終的にリクエストは次のように処理されます：

```text
クライアント -> エントリーポイント -> ルーター -> ミドルウェア -> サービス -> バックエンドサーバー
```

![Traefikアーキテクチャ](https://doc.traefik.io/traefik/assets/img/architecture-overview.png)

### 1.6.2. 開発環境向け設定ファイルの詳細説明

#### 1.6.2.1. traefik.yml（静的設定）

`./infra/traefik/config/traefik.yml`の主な設定項目とその目的を説明します。これは開発環境向けの設定です：

```yaml
# 全体的な設定
global:
  checkNewVersion: true  # 新バージョンの確認（開発環境では有用な情報）
  sendAnonymousUsage: false  # 匿名使用統計の送信を無効化（プライバシー保護）

# APIとダッシュボードの設定
api:
  dashboard: true  # ダッシュボード機能の有効化（設定の可視化とデバッグに重要）
  insecure: true  # 認証なしでダッシュボードにアクセス可能（開発環境のみ）

# エントリーポイントの設定
entryPoints:
  web:
    address: ":80"  # HTTPポート
    # 開発環境ではHTTPSリダイレクトはコメントアウト（開発を容易にするため）

  websecure:
    address: ":443"  # HTTPSポート（開発環境では実際には使用しない）

# プロバイダー設定
providers:
  # Docker設定
  docker:
    endpoint: "unix:///var/run/docker.sock"  # Dockerデーモンとの通信方法
    exposedByDefault: false  # 明示的に設定されたコンテナのみ公開（安全性向上）
    network: traefik-network  # 使用するDockerネットワーク（サービス間通信用）

  # ファイルプロバイダー設定
  file:
    directory: "/etc/traefik/dynamic"  # 動的設定ファイルの場所
    watch: true  # ファイル変更の監視（設定変更の即時反映）

# ログ設定
log:
  level: "INFO"  # ログの詳細度（開発中はDEBUGに変更すると詳細情報取得可能）

# アクセスログ設定
accessLog:
  filePath: "/var/log/traefik/access.log"  # アクセスログの保存場所
  bufferingSize: 100  # バッファサイズ（パフォーマンス最適化）

# プロメテウスメトリクス（開発環境でも監視・モニタリングの学習用）
metrics:
  prometheus: {}  # Prometheusと統合するための基本設定
```

この開発環境向け設定は主に以下の目的で行われています：

1. **開発環境の効率化**: ダッシュボード有効化、HTTPSリダイレクト無効化など
2. **サービスディスカバリー**: Docker統合による各サービスの自動検出
3. **開発中の可視性向上**: アクセスログとダッシュボードによる動作確認
4. **設定の柔軟性**: ファイル変更の監視とホットリロードによる迅速な設定変更

#### 1.6.2.2. middlewares.yml（ミドルウェア設定）

`./infra/traefik/dynamic/middlewares.yml`の主な設定項目とその目的を説明します。これは開発環境向けの設定です：

```yaml
http:
  middlewares:
    # セキュリティヘッダーの設定
    secure-headers:
      headers:
        frameDeny: true  # クリックジャッキング攻撃対策（X-Frame-Options: DENY相当）
        sslRedirect: false  # 開発環境でHTTPSリダイレクト無効（本番では必ずtrueに）
        browserXssFilter: true  # XSS保護フィルターの有効化（X-XSS-Protection: 1; mode=block相当）
        contentTypeNosniff: true  # MIMEタイプスニッフィング防止（X-Content-Type-Options: nosniff相当）
        forceSTSHeader: true  # HSTS強制（Strict-Transport-Security）
        stsIncludeSubdomains: true  # HSTSをサブドメインにも適用
        stsPreload: true  # ブラウザのHSTS事前読み込みリストに登録可能
        stsSeconds: 31536000  # HSTSの有効期間（1年）
        customFrameOptionsValue: "SAMEORIGIN"  # 同一オリジンのフレームのみ許可
        customRequestHeaders:
          X-Forwarded-Proto: "https"  # 開発環境でもHTTPS対応するアプリのためのヘッダー

    # 開発用CORS設定（開発時の利便性を重視した緩めの設定）
    cors:
      headers:
        accessControlAllowMethods:  # 許可するHTTPメソッド
          - GET
          - POST
          - PUT
          - DELETE
          - OPTIONS
        accessControlAllowHeaders:  # 許可するHTTPヘッダー
          - "*"  # すべてのヘッダーを許可（開発環境のみ）
        accessControlAllowOriginList:  # 許可するオリジン
          - "http://localhost"
          - "http://localhost:3000"
          - "http://frontend.localhost"
          - "http://backend.localhost"
        accessControlMaxAge: 100  # プリフライトリクエストの結果キャッシュ時間（秒）
        accessControlAllowCredentials: true  # クレデンシャル（Cookie等）の送信許可
        addVaryHeader: true  # Varyヘッダーの追加（キャッシュ制御の最適化）
```

この開発環境向けの設定は主に以下の目的で行われています：

1. **基本的なセキュリティ確保**: 最低限のセキュリティヘッダーを設定
2. **開発の容易さ**: CORSを緩めに設定し、異なるオリジン間の通信を容易に
3. **HTTPS無効**: 開発環境ではHTTPSリダイレクトを無効化（sslRedirect: false）
4. **一貫したヘッダー適用**: すべてのサービスに統一された設定を適用

#### 1.6.2.3. tls.yml（開発環境では使用しないがテンプレートとして存在）

`./infra/traefik/dynamic/tls.yml`の設定は、開発環境では実際には使用されませんが、将来的な本番環境設定のテンプレートとして配置されています。以下はその内容です：

```yaml
tls:
  options:
    default:
      minVersion: "VersionTLS12"  # 最小TLSバージョン（TLS 1.2以下は脆弱性あり）
      sniStrict: true  # SNI（Server Name Indication）の厳格チェック
      cipherSuites:  # 許可する暗号スイート（安全性の高いものだけ許可）
        - TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384
        - TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384
        - TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256
        - TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256
        - TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305
        - TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305
      curvePreferences: [...]  # 楕円曲線暗号の優先順位
      sniStrict: true  # SNI厳格モード
```

**開発環境でこのファイルが存在する理由**:

1. コード完全性のため（設定ファイル一式として管理）
2. 自己署名証明書を使った開発環境HTTPS設定のオプションを提供
3. 本番環境へのスムーズな移行のためのテンプレート

### 1.6.3. Docker Compose ラベルの詳細説明

Docker Compose内のラベル設定には、各サービスのTraefikルーティング設定が含まれています。実際に使用している設定は以下の通りです：

```yaml
labels:
  - "traefik.enable=true"  # Traefikでこのサービスを管理することを有効化
  - "traefik.http.routers.backend.rule=Host(`backend.localhost`)"  # ホスト名ベースのルーティングルール
  - "traefik.http.routers.backend.entrypoints=web"  # 使用するエントリーポイント
  - "traefik.http.services.backend.loadbalancer.server.port=8000"  # コンテナ内部のポート
  - "traefik.http.routers.backend.middlewares=secure-headers@file,cors@file"  # 適用するミドルウェア
```

この設定の重要な点：

1. **`traefik.enable=true`**:
   - Traefikでこのサービスを管理することを明示的に指定
   - `exposedByDefault: false` と組み合わせて、意図したサービスのみを公開

2. **`traefik.http.routers.backend.rule=Host(`backend.localhost`)`**:
   - リクエストのホストヘッダーが `backend.localhost` の場合にこのルートを使用
   - ホスト名ベースのルーティングの基本設定（各サービスに独自のサブドメインを割り当て）

3. **`traefik.http.routers.backend.entrypoints=web`**:
   - `web` エントリーポイント（ポート80）でリクエストを受け付ける
   - 開発環境では基本的にHTTPのみを使用

4. **`traefik.http.services.backend.loadbalancer.server.port=8000`**:
   - コンテナ内部のサービスが動作するポート
   - Traefikがリクエストを転送する先のポート番号

5. **`traefik.http.routers.backend.middlewares=secure-headers@file,cors@file`**:
   - このルートに適用するミドルウェアを指定
   - `@file` 修飾子でファイルプロバイダーの設定を参照
   - 複数のミドルウェアをカンマ区切りで連鎖（左から右の順に適用）

当初設定されていた以下のCORS関連のラベルについて説明します：

```yaml
- "traefik.http.middlewares.backend-cors.headers.accesscontrolallowmethods=GET,POST,PUT,DELETE,OPTIONS"
- "traefik.http.middlewares.backend-cors.headers.accesscontrolalloworiginlist=*"
- "traefik.http.middlewares.backend-cors.headers.accesscontrolmaxage=100"
- "traefik.http.middlewares.backend-cors.headers.addvaryheader=true"
```

これらのラベルの目的と課題：

1. **Docker ラベルでの直接定義**:
   - `backend-cors` という名前のCORSミドルウェアをラベル経由で直接定義
   - `middlewares.yml` と機能が重複し、設定の一元管理ができない

2. **全てのオリジンを許可**:
   - `accesscontrolalloworiginlist=*` はすべてのオリジン（ドメイン）からのアクセスを許可
   - 開発では便利だが、非常に緩い設定であり、本番環境では避けるべき

3. **動作しなかった理由**:
   - ミドルウェアの定義の競合（ファイルとラベルの両方で同様の設定）
   - 設定の優先順位の問題
   - ファイル設定を参照する `cors@file` が既に適用されていた可能性

この問題の解決策として、すべてのCORS設定をファイルベース（`middlewares.yml`）に統一し、ラベルでは `cors@file` として参照するアプローチが採用されました。これにより、設定の一元管理と再利用性が向上しています。

### 1.6.4. ホスト名ベースのルーティング

ホスト名ベースのルーティングは、同じIPアドレスやポートに対して、リクエストのHostヘッダーに基づいて異なるサービスにルーティングする方法です。

ホスト名ベースのルーティングの利点：

1. **単一のIPとポートで複数のサービスを提供可能**
   - 非常に効率的なリソース利用
   - シンプルなネットワーク設定

2. **企業のアプリケーション対応が容易**
   - バックエンドAPIとフロントエンドを分離したドメインで提供可能
   - バックエンドとフロントエンドの独立したデプロイとスケーリング

3. **マイクロサービスアーキテクチャに最適**
   - 各マイクロサービスを独自のサブドメインで公開可能
   - サービスの追加や削除が容易

4. **開発ワークフローの向上**
   - ローカル開発環境でも本番に近い構成でテスト可能
   - チーム開発でのコンフリクトの軽減

本実装では、ローカル開発環境において、以下のホスト名を設定しています：

- `backend.localhost` - Go/EchoバックエンドAPIサービス
- `customer.localhost` - 顧客向けNext.jsフロントエンドアプリケーション
- `admin.localhost` - 管理者向けNext.jsフロントエンドアプリケーション

### 1.6.5. セキュリティヘッダーと設定の重要性

Traefikのセキュリティヘッダー設定は、Webアプリケーションの全体的なセキュリティ向上に重要な役割を果たします。開発環境でも以下のセキュリティヘッダーの基本を理解することが重要です：

1. **Content-Security-Policy (CSP)**
   - クロスサイトスクリプティング（XSS）攻撃を防ぐ
   - 許可されたソースからのリソースのみを読み込むようブラウザに指示
   - 開発環境では緩めの設定を使用可能

2. **X-Frame-Options**
   - クリックジャッキング攻撃を防止
   - 開発環境では通常 `SAMEORIGIN` を設定

3. **X-Content-Type-Options**
   - MIMEタイプスニッフィングを防止
   - 開発環境でも `nosniff` を設定するのが一般的

4. **X-XSS-Protection**
   - ブラウザのXSSフィルターを有効化
   - 開発環境でも有効化しておくべき基本的な保護

5. **Strict-Transport-Security (HSTS)**
   - 開発環境ではHTTPSを使用しないため、実質的に無効
   - 設定自体は将来の本番環境のために用意

開発環境では、主に学習目的とテストの容易さから、一部のセキュリティ設定（特にHTTPSリダイレクト）が緩和されていますが、基本的なセキュリティヘッダーは維持されています。

### 1.6.6. ミドルウェア連鎖と適用方法

Docker Composeのラベル設定で注目すべき点は、複数のミドルウェアを連鎖させる方法です：

```yaml
- "traefik.http.routers.backend.middlewares=secure-headers@file,cors@file"
```

この設定の重要な点：

1. **ミドルウェアの連鎖**:
   - カンマ区切りで複数のミドルウェアを指定
   - 左から右の順に適用される（secure-headers → cors）

2. **`@file` 修飾子**:
   - ファイルプロバイダーで定義されたミドルウェアを参照
   - 動的設定ファイル（`middlewares.yml`）から読み込む

3. **適用順序の重要性**:
   - セキュリティヘッダーが最初に適用され、基本的な防御を設定
   - その後CORSヘッダーが適用され、APIアクセスを可能に

このアプローチにより、共通のミドルウェア設定を複数のサービスで再利用しながら、必要に応じてサービスごとに異なるミドルウェアの組み合わせを適用できます。

## 1.7. 【補足情報】

### 1.7.1. 開発環境と本番環境の設定の違い

開発環境で構築したTraefik設定を本番環境に移行する際には、いくつかの重要な変更が必要です。以下は主な違いと本番環境向けの設定ポイントです：

| 項目                     | 開発環境                              | 本番環境                       |
| ------------------------ | ------------------------------------- | ------------------------------ |
| **ダッシュボード**       | `insecure: true` （直接アクセス可能） | `insecure: false` （認証必要） |
| **SSL/TLS**              | 無効または自己署名証明書              | 正規のSSL証明書                |
| **HTTPSリダイレクト**    | 無効                                  | 有効                           |
| **セキュリティヘッダー** | 一部緩和                              | 厳格な設定                     |
| **CORS設定**             | 広範なオリジン許可                    | 限定的な許可オリジン           |
| **ログレベル**           | 詳細（DEBUG/INFO）                    | 最小限（INFO/WARN）            |
| **アクセス制御**         | 最小限                                | 認証と詳細な権限設定           |
| **リソース制限**         | 緩め                                  | 厳格（メモリ、CPU制限）        |
| **モニタリング**         | 基本機能のみ                          | 高度な監視と警告設定           |

#### 1.7.1.1. AWS環境での一般的な構成

AWS環境では、Traefikを利用する際、以下のような構成が一般的です：

1. **ALBとTraefikの役割分担**
   - AWS Application Load Balancer (ALB) がインターネットに面するフロントエンド
   - ALBがTLS終端を担当（HTTPS接続の処理）
   - TraefikはALBの後ろでコンテナ間のルーティングとミドルウェア処理を担当

   この構成では、以下の変更が必要です：

   ```yaml
   # traefik.yml
   entryPoints:
     web:
       address: ":80"
       # HTTPSリダイレクトは不要（ALBが担当）
   ```

2. **TLS設定**
   - AWS Certificate Manager (ACM) で証明書を管理
   - ALBのセキュリティポリシーでTLS設定を行う
   - Traefik側のTLS設定は最小限または不要

3. **ECSタスク定義での統合**
   - ECSタスク定義でTraefikをサイドカーコンテナとして設定
   - サービスディスカバリーにAWS Cloud Mapを使用

   ```json
   {
     "containerDefinitions": [
       {
         "name": "traefik",
         "image": "traefik:latest",
         "essential": true,
         "portMappings": [
           {
             "containerPort": 80,
             "hostPort": 80
           }
         ],
         "environment": [
           {
             "name": "AWS_REGION",
             "value": "ap-northeast-1"
           }
         ],
         "mountPoints": [
           {
             "sourceVolume": "traefik-config",
             "containerPath": "/etc/traefik"
           }
         ]
       },
       // アプリケーションコンテナ定義...
     ]
   }
   ```

#### 1.7.1.2. オンプレミスや他のクラウド環境でTraefikを直接公開する場合

ALBなどのロードバランサーがない環境では、Traefikを直接インターネットに公開するケースがあります。その場合の重要な設定：

1. **HTTPS設定の有効化**

   ```yaml
   # traefik.yml
   entryPoints:
     web:
       address: ":80"
       http:
         redirections:
           entryPoint:
             to: websecure
             scheme: https
     websecure:
       address: ":443"
   ```

2. **Let's Encryptを使用した自動証明書発行**

   ```yaml
   # traefik.yml
   certificatesResolvers:
     letsencrypt:
       acme:
         email: youremail@example.com
         storage: /etc/traefik/acme/acme.json
         httpChallenge:
           entryPoint: web
   ```

3. **セキュリティ強化設定**

   ```yaml
   # middlewares.yml
   http:
     middlewares:
       secure-headers:
         headers:
           frameDeny: true
           sslRedirect: true  # 本番環境ではtrueに
           browserXssFilter: true
           contentTypeNosniff: true
           forceSTSHeader: true
           stsIncludeSubdomains: true
           stsPreload: true
           stsSeconds: 31536000
           customFrameOptionsValue: "DENY"  # より厳格に
           contentSecurityPolicy: "default-src 'self'; img-src 'self' https: data:;"  # CSPを追加
   ```

4. **CORS設定の強化**

   ```yaml
   # middlewares.yml
   cors:
     headers:
       accessControlAllowMethods:
         - GET
         - POST
         - PUT
         - DELETE
         - OPTIONS
       accessControlAllowHeaders:
         - Authorization
         - Content-Type
         - X-Requested-With  # 明示的に許可するヘッダーのみ指定
       accessControlAllowOriginList:
         - https://your-domain.com
         - https://api.your-domain.com  # 明示的に許可するオリジンのみ指定
       accessControlMaxAge: 3600
       accessControlAllowCredentials: true
       addVaryHeader: true
   ```

5. **ダッシュボード保護**

   ```yaml
   # traefik.yml
   api:
     dashboard: true
     insecure: false  # 本番環境では必ずfalse

   # dynamic/auth.yml
   http:
     middlewares:
       admin-auth:
         basicAuth:
           users:
             - "admin:$apr1$H6uskkkW$IgXLP6ewTrSuBkTrqE8wj/"  # htpasswd生成パスワード
   ```

#### 1.7.1.3. 本番環境に向けた一般的な変更点

どのような本番環境でも共通して必要な変更点：

1. **リソース制限の設定**

   ```yaml
   # docker-compose.yml または同等の設定
   services:
     traefik:
       deploy:
         resources:
           limits:
             cpus: '0.5'
             memory: 256M
           reservations:
             cpus: '0.1'
             memory: 128M
   ```

2. **ヘルスチェックの強化**

   ```yaml
   # docker-compose.yml または同等の設定
   healthcheck:
     test: ["CMD", "traefik", "healthcheck"]
     interval: 10s
     timeout: 5s
     retries: 3
     start_period: 30s
   ```

3. **ログ設定の最適化**

   ```yaml
   # traefik.yml
   log:
     level: "INFO"  # DEBUGよりINFOの方が本番向け
     format: "json"  # 構造化ログの使用

   accessLog:
     filePath: "/var/log/traefik/access.log"
     format: "json"
     bufferingSize: 100
     fields:
       headers:
         defaultMode: "keep"
         names:
           User-Agent: "keep"
           Authorization: "drop"
           Content-Type: "keep"
   ```

4. **メトリクスとモニタリングの強化**

   ```yaml
   # traefik.yml
   metrics:
     prometheus:
       entryPoint: metrics  # 専用エントリポイント
       addServicesLabels: true
       addEntryPointsLabels: true
       buckets:
         - 0.1
         - 0.3
         - 1.0
         - 2.5
         - 5.0

   # エントリーポイント定義に追加
   entryPoints:
     metrics:
       address: ":9100"
   ```

5. **レート制限の設定**

   ```yaml
   # dynamic/middlewares.yml
   http:
     middlewares:
       rate-limit:
         rateLimit:
           average: 100
           burst: 50
   ```

### 1.7.2. カスタムミドルウエアの作成方法

Traefikの強力な機能の一つは、カスタムミドルウエアを追加してリクエスト処理をカスタマイズできる点です。以下は、基本的なミドルウエアの作成手順です：

1. **ミドルウエア定義ファイルを作成**

   新しいミドルウエアを定義する場合、`./infra/traefik/dynamic/custom-middlewares.yml`のようなファイルを作成します：

    ```yaml
    http:
      middlewares:
        # カスタムヘッダー追加ミドルウエア
        add-custom-headers:
          headers:
            customRequestHeaders:
              X-Custom-Header: "custom-value"
              X-Environment: "development"
              X-Request-ID: "{{.RequestID}}"

        # レスポンス修正ミドルウエア
        response-modifier:
          headers:
            customResponseHeaders:
              X-Response-Time: "{{.ResponseTime}}"
              X-Powered-By: "AWS Observability eCommerce"
    ```

2. **ルーターにミドルウエアを適用**

   作成したミドルウエアをサービスに適用するには、Docker Composeファイルのラベルを修正します：

    ```yaml
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.backend.rule=Host(`backend.localhost`)"
      - "traefik.http.routers.backend.entrypoints=web"
      - "traefik.http.services.backend.loadbalancer.server.port=8000"
      - "traefik.http.routers.backend.middlewares=add-custom-headers@file,response-modifier@file"
    ```

3. **ミドルウエアの連鎖**

   複数のミドルウエアを連鎖させる場合は、カンマ区切りで指定します：

   ```yaml
     - "traefik.http.routers.backend.middlewares=secure-headers@file,cors@file,add-custom-headers@file"
   ```

ミドルウエアの適用順序は左から右で、最初に指定したミドルウエアが最初に適用されます。

### 1.7.3. 本番環境でのデプロイパターン

本番環境でTraefikを使用する場合のデプロイパターンには、以下のようなオプションがあります：

#### 1.7.3.1. コンテナオーケストレーションプラットフォームでの活用

##### 1.7.3.1.1. ECS（Elastic Container Service）での使用

- タスク定義で各サービスとTraefikをリンク
- サービスディスカバリーにAWS Cloud Mapを活用
- ECSタスクIAMロールで必要な権限を付与
- TraefikをECS Service Connect対応として設定

```hcl
# Terraformでの定義例
resource "aws_ecs_task_definition" "app" {
  family                   = "app"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = 1024
  memory                   = 2048
  execution_role_arn       = aws_iam_role.ecs_execution_role.arn
  task_role_arn            = aws_iam_role.ecs_task_role.arn

  container_definitions = jsonencode([
    {
      name      = "traefik"
      image     = "traefik:latest"
      essential = true
      portMappings = [
        {
          containerPort = 80
          hostPort      = 80
        }
      ]
      environment = [
        {
          name  = "AWS_REGION"
          value = "ap-northeast-1"
        }
      ]
      mountPoints = [
        {
          sourceVolume  = "traefik-config"
          containerPath = "/etc/traefik"
        }
      ]
    },
    {
      name      = "app"
      image     = "${aws_ecr_repository.app.repository_url}:latest"
      essential = true
      # 残りの設定...
    }
  ])

  volume {
    name = "traefik-config"
    # 設定ボリューム定義...
  }
}
```

##### 1.7.3.1.2. Kubernetes環境での使用

- Kubernetes Ingress Controllerとして活用
- CustomResourceDefinitions（CRDs）による設定
- Kubernetes Service DiscoveryとTraefikの連携

#### 1.7.3.2. HA構成（高可用性）での設定

本番環境では高可用性を確保するために、以下の設定を検討します：

- 複数のTraefikインスタンスを異なるアベイラビリティゾーンにデプロイ
- 共有ストレージや設定データベースの使用
- Let's Encryptを使用する場合のACMEストレージの共有
- セッション維持の設定（ステートフルなアプリケーションの場合）

```yaml
# traefik.yml（HA構成の例）
providers:
  file:
    directory: "/etc/traefik/dynamic"
    watch: true

  # 静的設定ではなく動的なプロバイダーを使用
  consulCatalog:
    prefix: "traefik"
    endpoint:
      address: "consul:8500"
      scheme: "http"
    refreshInterval: "30s"

# Let's Encrypt設定
certificatesResolvers:
  letsencrypt:
    acme:
      email: "your-email@example.com"
      storage: "/etc/traefik/acme/acme.json"  # 共有ストレージにマウント
      httpChallenge:
        entryPoint: web
```

#### 1.7.3.3. Blue/Greenデプロイの実装

新バージョンのリリースにおいて、Blue/Greenデプロイパターンを実装することも可能です：

- 既存環境（Blue）と並行して新環境（Green）をデプロイ
- 新環境の検証が完了したら、Traefikのルーティングを切り替え
- 問題があれば素早く元の環境に戻せる

```yaml
# ルーティング設定の例
http:
  routers:
    api-blue:
      rule: "Host(`api.example.com`) && HeadersRegexp(`X-Version`, `blue`)"
      service: api-blue-service
      middlewares:
        - rate-limit
        - secure-headers

    api-green:
      rule: "Host(`api.example.com`) && HeadersRegexp(`X-Version`, `green`)"
      service: api-green-service
      middlewares:
        - rate-limit
        - secure-headers

    # デフォルトルート（本番切替時に変更）
    api:
      rule: "Host(`api.example.com`)"
      service: api-blue-service  # ここを変更することで切替
      middlewares:
        - rate-limit
        - secure-headers
```

このように、本番環境ではTraefikを単なるリバースプロキシとしてだけでなく、高度なデプロイ戦略やサービスディスカバリー、高可用性を実現するための重要な構成要素として活用できます。

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: Traefikダッシュボードにアクセスできない

**症状**: <http://localhost:8080> にアクセスしてもダッシュボードが表示されない。

**解決策**:

1. Traefikの設定でダッシュボードが有効化されているか確認する：

   ```yaml

    api:
      dashboard: true
      insecure: true  # 開発環境のみ

   ```

2. Docker Composeのポートマッピングが正しく設定されているか確認する：

   ```yaml

    ports:

    - "8080:8080"  # Dashboard

   ```

3. Traefikコンテナが正常に起動しているか確認する：

   ```bash

    docker ps | grep traefik

   ```

4. コンテナのログを確認する：

   ```bash

    docker logs traefik

   ```

### 1.8.2. 問題2: バックエンドサービスへのルーティングが機能しない

**症状**: <http://backend.localhost> にアクセスしてもバックエンドサービスに到達しない。

**解決策**:

1. hostsファイルに正しくエントリが追加されているか確認する：

    ```text
    127.0.0.1 customer.localhost admin.localhost backend.localhost
    ```

2. バックエンドサービスのDockerコンテナが起動しているか確認する：

    ```bash
    docker ps | grep backend
    ```

3. バックエンドサービスのラベル設定が正しいか確認する：

    ```yaml

      labels:

      - "traefik.enable=true"
      - "traefik.http.routers.backend.rule=Host(`backend.localhost`)"
      - "traefik.http.routers.backend.entrypoints=web"
      - "traefik.http.services.backend.loadbalancer.server.port=8000"

    ```

4. Traefikのダッシュボードでルートが正しく設定されているか確認する：
   - <http://localhost:8080/dashboard/> にアクセス
   - HTTP Routersセクションでバックエンドサービスのルートが表示されているか確認

5. ネットワーク設定を確認する：

    ```bash
    docker network inspect traefik-network
    ```

### 1.8.3. 問題3: HTTPSリダイレクトで無限リダイレクトループが発生する

**症状**: HTTPSリダイレクトを有効化してサイトにアクセスすると、ブラウザが「リダイレクトが多すぎます」と表示してアクセスできない。

**解決策**:

1. 開発環境ではHTTPSリダイレクトを無効化する（基本的には自己署名証明書なしの開発環境でHTTPSは使わない）：

    ```yaml

    entryPoints:
      web:
        address: ":80"
        # 以下をコメントアウト
        # http:
        #   redirections:
        #     entryPoint:
        #       to: websecure
        #       scheme: https

    ```

2. 本番環境でSSL証明書が正しく設定されているか確認する：
   - 正規のSSL証明書が設定されているか
   - Let's Encryptなどの自動証明書発行が正しく設定されているか

3. `sslRedirect`設定とHTTPSリダイレクトの設定が重複していないか確認する：

    ```yaml
      # middlewares.yml
        secure-headers:
          headers:
            sslRedirect: false  # 開発環境ではfalse

      # traefik.yml
      entryPoints:
        web:
          address: ":80"
          # 開発環境では無効化
          # http:
          #   redirections:
          #     entryPoint:
          #       to: websecure
          #       scheme: https
    ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **ホスト名ベースのルーティングの仕組み**
   - 同じポート・ネットワークで複数のサービスを提供する方法
   - バックエンド・APIと複数のフロントエンドアプリを連携させる設計パターン

2. **ミドルウェアによるクロスカッティングな設定**
   - セキュリティ設定を一元管理するアプローチ
   - CORSなどの設定をサービスごとではなくゲートウェイレベルで実装

3. **開発環境と本番環境の設定の分離**
   - コメントと条件付き設定による環境切り替え
   - 本番環境に移行する際の考慮点

4. **複数のフロントエンドアプリケーションの共存設計**
   - 顧客向けと管理者向けの別々のアプリケーションを独立して管理
   - 共通のバックエンドAPIとの連携パターン

これらのポイントは、マイクロサービスアーキテクチャと分散システムの構築において重要な基盤となります。

## 1.10. 【次回の準備】

次回（Day 5）では、管理画面の基本実装に取り組みます。以下の点について事前に確認しておくと良いでしょう：

1. **TailwindCSSの基本的な知識**
   - ユーティリティクラスの基本的な使い方
   - レスポンシブデザインの概念
   - 公式Website: <https://tailwindcss.com/docs>

2. **Reactコンポーネントの基本**
   - 再利用可能なコンポーネントの設計について復習
   - Propsと型定義の使い方

3. **Traefikの動作確認**
   - 今回設定したTraefikが正しく動作しているか確認
   - ダッシュボードでのルート確認とアクセスログの確認

4. **Docker環境の安定性**
   - 全サービスが正常に起動し、連携できているか確認

上記の準備が整っていれば、次回の管理画面実装とTailwindCSSによるスタイリングにスムーズに進むことができます。

## 1.11. 【.envrc サンプル】

以下は本日の実装で使用する.envrcのサンプルです。ご自身の環境に合わせて修正して使用してください。このファイルはgitにコミットしないようにしてください。

```bash
# .envrc サンプル

# プロジェクト全体設定
export PROJECT_ROOT=$(pwd)
export COMPOSE_FILE=docker-compose.yml
export COMPOSE_PROJECT_NAME=aws-observability-ecommerce

# バックエンド環境変数
export BACKEND_PORT=8000
export BACKEND_HOST="backend.localhost"

# フロントエンド環境変数
export CUSTOMER_FRONTEND_PORT=3000
export CUSTOMER_FRONTEND_HOST="customer.localhost"
export ADMIN_FRONTEND_PORT=3000
export ADMIN_FRONTEND_HOST="admin.localhost"

# DB環境変数
export DB_HOST="localhost"
export DB_PORT=3306
export DB_NAME="ecommerce"
export DB_USER="ecommerce_user"
export DB_PASSWORD="password"

# 開発環境設定
export NODE_ENV="development"
export GO_ENV="development"
```
