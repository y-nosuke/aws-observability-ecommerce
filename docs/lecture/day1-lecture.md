# 1. Week 1 - Day 1: Docker Compose 環境の構築

## 1.1. 目次

- [1. Week 1 - Day 1: Docker Compose 環境の構築](#1-week-1---day-1-docker-compose-環境の構築)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. プロジェクト構造の作成](#141-プロジェクト構造の作成)
    - [1.4.2. Docker Compose 設定ファイルの作成](#142-docker-compose-設定ファイルの作成)
    - [1.4.3. Traefik の設定](#143-traefik-の設定)
    - [1.4.4. MySQL の設定](#144-mysql-の設定)
      - [1.4.4.1. 設定ファイル権限の設定方法](#1441-設定ファイル権限の設定方法)
        - [1.4.4.1.1. Linux/macOSの場合](#14411-linuxmacosの場合)
        - [1.4.4.1.2. Windowsの場合（WSL使用時も含む）](#14412-windowsの場合wsl使用時も含む)
    - [1.4.5. LocalStack の設定](#145-localstack-の設定)
    - [1.4.6. Docker Compose 環境の起動と検証](#146-docker-compose-環境の起動と検証)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. Docker Compose の役割と利点](#161-docker-compose-の役割と利点)
    - [1.6.2. Traefik の基本概念と動作原理](#162-traefik-の基本概念と動作原理)
    - [1.6.3. LocalStack によるAWSエミュレーション](#163-localstack-によるawsエミュレーション)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. Docker ネットワークについて](#171-docker-ネットワークについて)
    - [1.7.2. ボリュームによるデータ永続化](#172-ボリュームによるデータ永続化)
    - [1.7.3. LocalStack Desktopの活用](#173-localstack-desktopの活用)
      - [1.7.3.1. LocalStack Desktopのインストール](#1731-localstack-desktopのインストール)
      - [1.7.3.2. LocalStack Desktopの設定と使用方法](#1732-localstack-desktopの設定と使用方法)
      - [1.7.3.3. LocalStack Desktopの主な機能](#1733-localstack-desktopの主な機能)
      - [1.7.3.4. LocalStack Desktopの活用例](#1734-localstack-desktopの活用例)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: ポートの競合](#181-問題1-ポートの競合)
    - [1.8.2. 問題2: Traefikでのホスト名解決の問題](#182-問題2-traefikでのホスト名解決の問題)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- Docker Compose を使用して複数のサービス（MySQL、Traefik、LocalStack）を連携させた開発環境を構築する
- Traefik を使用してホスト名ベースのルーティングを設定し、顧客向けと管理者向けの分離された環境を準備する
- LocalStack を使用してAWSサービスをローカルでエミュレートする準備をする
- サービス間の接続とネットワークを適切に設定する
- 永続データのためのボリュームを設定し、コンテナ再起動後もデータが維持されるようにする

## 1.3. 【準備】

このプロジェクトを始めるにあたり、以下のツールとソフトウェアが必要です。実装を始める前に、すべてのツールが正しくインストールされていることを確認してください。

### 1.3.1. チェックリスト

- [x] Git (バージョン管理)

  ```bash
  git --version
  # git version 2.34.1 以上が望ましい
  ```

- [x] Docker Engine

  ```bash
  docker --version
  # Docker version 20.10.0 以上が望ましい
  ```

- [x] Docker Compose

  ```bash
  docker compose version
  # Docker Compose version v2.10.0 以上が望ましい
  ```

- [x] awslocal (LocalStack用AWS CLIラッパー)

  ```bash
  # まずpipxをインストール
  brew install pipx
  # pipxを使ってawscli-localをインストール
  pipx install awscli-local
  awslocal --version
  # aws-cli/x.x.x Python/x.x.x ...
  ```

- [x] テキストエディタまたはIDE (Visual Studio Code推奨)
- [x] ターミナル (Linuxベースであればどれでも可)
- [x] curl または wget (動作確認用)

  ```bash
  curl --version
  # curl 7.68.0 以上が望ましい
  ```

- [x] direnv (オプション、環境変数管理用)

  ```bash
  direnv --version
  # direnv v2.32.0 以上が望ましい
  ```

## 1.4. 【手順】

### 1.4.1. プロジェクト構造の作成

まず、プロジェクト用の新しいディレクトリを作成し、基本的なファイル構造をセットアップします。

```bash
# プロジェクトのルートディレクトリを作成
mkdir -p aws-observability-ecommerce
cd aws-observability-ecommerce

# Dockerおよび環境関連のディレクトリを作成
mkdir -p infra/{mysql/{initdb.d,conf.d},traefik/dynamic,localstack/init-scripts}

# .gitignoreファイルを作成
touch .gitignore

# Git リポジトリを初期化
git init
```

### 1.4.2. Docker Compose 設定ファイルの作成

プロジェクトのルートディレクトリに `compose.yml` ファイルを作成し、必要なサービスを定義します。

```bash
# Docker Compose 設定ファイルを作成
touch compose.yml
```

`compose.yml` に以下の内容を記述します：

```yaml
services:
  traefik:
    image: traefik:latest
    container_name: traefik
    restart: unless-stopped
    ports:
      - "80:80"
      - "8080:8080" # Traefik ダッシュボード
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./infra/traefik/traefik.yml:/etc/traefik/traefik.yml:ro
      - ./infra/traefik/dynamic:/etc/traefik/dynamic:ro
    networks:
      - ecommerce-network
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.dashboard.rule=Host(`traefik.localhost`)"
      - "traefik.http.routers.dashboard.service=api@internal"
      - "traefik.http.services.dashboard.loadbalancer.server.port=8080"
    deploy:
      resources:
        limits:
          memory: 256M

  mysql:
    image: mysql:latest
    container_name: mysql
    restart: unless-stopped
    environment:
      MYSQL_ROOT_PASSWORD: ${MYSQL_ROOT_PASSWORD:-rootpassword}
      MYSQL_DATABASE: ${MYSQL_DATABASE:-ecommerce}
      MYSQL_USER: ${MYSQL_USER:-ecommerce_user}
      MYSQL_PASSWORD: ${MYSQL_PASSWORD:-ecommerce_password}
    ports:
      - "3306:3306"
    volumes:
      - ./infra/mysql/initdb.d:/docker-entrypoint-initdb.d:ro
      - ./infra/mysql/conf.d:/etc/mysql/conf.d:ro
      - mysql_data:/var/lib/mysql
    healthcheck:
      test:
        [
          "CMD",
          "mysqladmin",
          "ping",
          "-h",
          "localhost",
          "-u",
          "root",
          "-p${MYSQL_ROOT_PASSWORD:-rootpassword}",
        ]
      interval: 5s
      timeout: 5s
      retries: 5
      start_period: 10s
    labels:
      - "traefik.enable=false" # Traefikからの直接アクセスは不要
    networks:
      - ecommerce-network
    deploy:
      resources:
        limits:
          memory: 512M

  phpmyadmin:
    image: phpmyadmin/phpmyadmin
    container_name: phpmyadmin
    restart: always
    expose:
      - "80"
    environment:
      - PMA_HOST=mysql
      - PMA_USER=ecommerce_user
      - PMA_PASSWORD=ecommerce_password
      - UPLOAD_LIMIT=300M
    depends_on:
      - mysql
    labels:
      - "traefik.enable=true"
      - "traefik.http.routers.phpmyadmin.rule=Host(`phpmyadmin.localhost`)"
      - "traefik.http.services.phpmyadmin.loadbalancer.server.port=80"
    networks:
      - ecommerce-network

  localstack:
    image: localstack/localstack:latest
    container_name: localstack
    restart: unless-stopped
    environment:
      - SERVICES=s3,cloudwatch,logs,events,sns,sqs,lambda

      - LAMBDA_EXECUTOR=docker
      - DEFAULT_REGION=ap-northeast-1
      - AWS_DEFAULT_REGION=ap-northeast-1
      - AWS_ACCESS_KEY_ID=localstack
      - AWS_SECRET_ACCESS_KEY=localstack
      - HOSTNAME_EXTERNAL=localstack
      - DOCKER_HOST=unix:///var/run/docker.sock
      - DEBUG=1
    ports:
      - "4566:4566" # LocalStackの主要ポート（すべてのAWSサービスへのアクセスに使用）
      - "4571:4571" # Localstack Gateway
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./infra/localstack/init-scripts:/etc/localstack/init/ready.d:ro
      - localstack_data:/var/lib/localstack
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:4566/_localstack/health"]
      interval: 10s
      timeout: 5s
      retries: 3
      start_period: 20s
    deploy:
      resources:
        limits:
          memory: 1G
    labels:
      - "traefik.enable=false" # Traefikからの直接アクセスは不要
    networks:
      - ecommerce-network

volumes:
  mysql_data:
    driver: local
  localstack_data:
    driver: local

networks:
  ecommerce-network:
    driver: bridge
    name: ecommerce-network
```

### 1.4.3. Traefik の設定

Traefik の設定ファイルを作成し、ホスト名ベースのルーティングを構成します。

```bash
# Traefik 設定ファイルを作成
touch infra/traefik/traefik.yml
touch infra/traefik/dynamic/config.yml
```

`infra/traefik/traefik.yml` に以下の内容を記述します：

```yaml
# Traefik グローバル設定
api:
  dashboard: true
  insecure: true

# エントリーポイント設定
entryPoints:
  web:
    address: ":80"

# Docker Provider 設定
providers:
  docker:
    endpoint: "unix:///var/run/docker.sock"
    exposedByDefault: false
  file:
    directory: "/etc/traefik/dynamic"
    watch: true

# ログ設定
log:
  level: "INFO"

# アクセスログ設定
accessLog: {}
```

`infra/traefik/dynamic/config.yml` に以下の内容を記述します：

```yaml
# 動的設定
http:
  routers:
    # 将来的に追加するバックエンドサービス用のルーター
    backend:
      rule: "Host(`api.localhost`)"
      service: backend-service
      entryPoints:
        - web

    # 将来的に追加する顧客向けフロントエンドサービス用のルーター
    frontend-customer:
      rule: "Host(`shop.localhost`)"
      service: frontend-customer-service
      entryPoints:
        - web

    # 将来的に追加する管理者向けフロントエンドサービス用のルーター
    frontend-admin:
      rule: "Host(`admin.localhost`)"
      service: frontend-admin-service
      entryPoints:
        - web

  # サービス定義（将来追加するサービスのためのプレースホルダー）
  services:
    backend-service:
      loadBalancer:
        servers:
          - url: "http://backend:8080"

    frontend-customer-service:
      loadBalancer:
        servers:
          - url: "http://frontend-customer:3000"

    frontend-admin-service:
      loadBalancer:
        servers:
          - url: "http://frontend-admin:3000"
```

`/etc/hosts` に以下の内容を記述します：

  ```bash
  # ローカルホスト名のマッピングを追加
  127.0.0.1 traefik.localhost api.localhost shop.localhost admin.localhost
  ```

### 1.4.4. MySQL の設定

MySQL の初期化スクリプトと設定ファイルを作成します。

```bash
# MySQL 初期化スクリプトと設定ファイルを作成
touch infra/mysql/initdb.d/01_init.sql
touch infra/mysql/conf.d/my.cnf
```

`infra/mysql/initdb.d/01_init.sql` に以下の内容を記述します：

```sql
-- 基本的な初期化スクリプト
-- 詳細なテーブル定義は Week 2 で実装します

-- データベースが存在しない場合は作成
CREATE DATABASE IF NOT EXISTS `ecommerce`;

-- 権限の設定
GRANT ALL PRIVILEGES ON `ecommerce`.* TO 'ecommerceuser'@'%';
FLUSH PRIVILEGES;

-- ecommerceデータベースを選択
USE `ecommerce`;

-- 基本的なテストテーブル作成
CREATE TABLE IF NOT EXISTS `test` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- テストデータ挿入
INSERT INTO `test` (`name`) VALUES ('This is a test');
```

`infra/mysql/conf.d/my.cnf` に以下の内容を記述します：

```ini
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_0900_ai_ci
default-time-zone='+09:00'

[mysql]
default-character-set=utf8mb4

[client]
default-character-set=utf8mb4
```

#### 1.4.4.1. 設定ファイル権限の設定方法

MySQLの設定ファイルを作成したら、適切な権限を設定する必要があります。WSL環境では、ファイルの権限設定方法がLinux/macOSとは異なります。

##### 1.4.4.1.1. Linux/macOSの場合

```bash
# 設定ファイルの権限を変更（所有者のみ書き込み可能、他は読み取り専用）
chmod 644 infra/mysql/conf.d/my.cnf
```

##### 1.4.4.1.2. Windowsの場合（WSL使用時も含む）

Windows環境またはWSL環境でWindowsファイルシステム上のファイルを操作する場合：

1. エクスプローラーでファイルを右クリック
2. 「プロパティ」を選択
3. 「読み取り専用」にチェックを入れて「OK」をクリック

これは、MySQLがセキュリティ上の理由から誰でも書き込み可能な設定ファイルを無視するため、重要なステップです。権限が適切に設定されていないと、以下のような警告が表示されることがあります：

```text
mysqld: [Warning] World-writable config file '/etc/mysql/conf.d/my.cnf' is ignored.
```

### 1.4.5. LocalStack の設定

LocalStack の初期化スクリプトを作成します。

```bash
# LocalStack 初期化スクリプトを作成
touch infra/localstack/init-scripts/01_init_aws_resources.sh
```

`infra/localstack/init-scripts/01_init_aws_resources.sh` に以下の内容を記述します：

```bash
#!/bin/bash
# LocalStack初期化スクリプト

set -e

echo "LocalStack initializing AWS resources..."

# デフォルトリージョンの設定
REGION=${AWS_DEFAULT_REGION:-ap-northeast-1}
LOCALSTACK_HOST=localhost
ENDPOINT_URL=http://${LOCALSTACK_HOST}:4566

# S3バケット作成
echo "Creating S3 buckets..."
aws --endpoint-url=${ENDPOINT_URL} s3 mb s3://ecommerce-product-images --region ${REGION}
aws --endpoint-url=${ENDPOINT_URL} s3 mb s3://ecommerce-logs --region ${REGION}

# CloudWatch Logsロググループ作成
echo "Creating CloudWatch Logs groups..."
aws --endpoint-url=${ENDPOINT_URL} logs create-log-group --log-group-name /ecommerce/api --region ${REGION}
aws --endpoint-url=${ENDPOINT_URL} logs create-log-group --log-group-name /ecommerce/app --region ${REGION}

# SNSトピック作成
echo "Creating SNS topics..."
aws --endpoint-url=${ENDPOINT_URL} sns create-topic --name ecommerce-notifications --region ${REGION}

# SQSキュー作成
echo "Creating SQS queues..."
aws --endpoint-url=${ENDPOINT_URL} sqs create-queue --queue-name ecommerce-events --region ${REGION}

echo "LocalStack initialization completed!"
```

スクリプトに実行権限を付与します。

```bash
chmod +x infra/localstack/init-scripts/01_init_aws_resources.sh
```

### 1.4.6. Docker Compose 環境の起動と検証

作成した Docker Compose 環境を起動し、正常に動作することを確認します。

```bash
# Docker Compose 環境の起動
docker compose up -d

# サービスの状態確認
docker compose ps
```

## 1.5. 【確認ポイント】

Docker Compose 環境が正しくセットアップされたことを確認するためのチェックリストです：

- [x] すべてのコンテナが起動し、ステータスが「Up」になっている

  ```bash
  $ docker compose ps
  NAME         IMAGE                          COMMAND                  SERVICE      CREATED         STATUS                    PORTS
  localstack   localstack/localstack:latest   "docker-entrypoint.sh"   localstack   9 minutes ago   Up 9 minutes (healthy)   0.0.0.0:4566->4566/tcp, 4510-4559/tcp, 5678/tcp, 0.0.0.0:4571->4571/tcp
  mysql        mysql:latest                   "docker-entrypoint.s…"   mysql        9 minutes ago   Up 9 minutes (healthy)   0.0.0.0:3306->3306/tcp, 33060/tcp
  phpmyadmin   phpmyadmin/phpmyadmin          "/docker-entrypoint.…"   phpmyadmin   9 minutes ago   Up 9 minutes             80/tcp
  traefik      traefik:latest                 "/entrypoint.sh trae…"   traefik      9 minutes ago   Up 9 minutes             0.0.0.0:80->80/tcp, 0.0.0.0:8080->8080/tcp
  ```

- [x] MySQL コンテナに接続できる

  ```bash
  $ docker compose exec mysql mysql -uecommerce_user -pecommerce_password -e "SELECT * FROM ecommerce.test;"
  mysql: [Warning] Using a password on the command line interface can be insecure.
  +----+----------------+---------------------+
  | id | name           | created_at          |
  +----+----------------+---------------------+
  |  1 | This is a test | 2025-03-29 14:30:17 |
  +----+----------------+---------------------+
  # 「This is a test」というデータが表示されればOK
  ```

- [x] Traefik ダッシュボードにアクセスできる

  ```bash
  # ブラウザで以下のURLにアクセス
  # http://traefik.localhost:8080

  # または curl で確認
  $ curl -H "Host: traefik.localhost" http://localhost:8080/api/version
  {"Version":"3.3.4","Codename":"saintnectaire","startDate":"2025-03-29T05:30:08.966459923Z"}
  # バージョン情報が表示されればOK
  ```

- [x] LocalStack が正常に動作し、awslocalを使って操作できる

  ```bash
  # S3バケットのリストを取得
  $ awslocal s3 ls
  2025-03-29 14:30:10 ecommerce-product-images
  2025-03-29 14:30:11 ecommerce-logs
  # ecommerce-product-imagesとecommerce-logsが表示されればOK

  # CloudWatch Logsのロググループ一覧を取得
  $ awslocal logs describe-log-groups
  {
      "logGroups": [
          {
              "logGroupName": "/ecommerce/api",
              "creationTime": 1743226211882,
              "metricFilterCount": 0,
              "arn": "arn:aws:logs:ap-northeast-1:000000000000:log-group:/ecommerce/api:*",
              "storedBytes": 0
          },
          {
              "logGroupName": "/ecommerce/app",
              "creationTime": 1743226212285,
              "metricFilterCount": 0,
              "arn": "arn:aws:logs:ap-northeast-1:000000000000:log-group:/ecommerce/app:*",
              "storedBytes": 0
          }
      ]
  }
  # /ecommerce/apiと/ecommerce/appが表示されればOK

  # SNSトピックの一覧を取得
  $ awslocal sns list-topics
  {
      "Topics": [
          {
              "TopicArn": "arn:aws:sns:ap-northeast-1:000000000000:ecommerce-notifications"
          }
      ]
  }
  # ecommerce-notificationsトピックが表示されればOK

  # SQSキューの一覧を取得
  $ awslocal sqs list-queues
  {
      "QueueUrls": [
          "http://sqs.ap-northeast-1.localhost.localstack.cloud:4566/000000000000/ecommerce-events"
      ]
  }
  # ecommerce-eventsキューが表示されればOK
  ```

## 1.6. 【詳細解説】

### 1.6.1. Docker Compose の役割と利点

Docker Compose は、複数のコンテナを定義し実行するためのツールです。本プロジェクトでは、以下の利点を活かして開発環境を構築しています：

1. **依存関係の明確化**: 各サービス間の依存関係を`docker-compose.yml`ファイル内で明示的に定義できます。今回のプロジェクトでは、フロントエンドとバックエンドがデータベースに依存するという関係を表現しています。

2. **環境の一貫性**: 開発チームの全員が同じ環境で作業できるようになります。「私の環境では動くのに」という問題を避けることができます。

3. **簡単な起動と停止**: `docker compose up` と `docker compose down` コマンドだけで、すべてのサービスをまとめて起動・停止できます。

4. **環境変数の統合**: `.env` ファイルや環境変数を利用して、設定値を柔軟に変更できます。セキュリティ上重要な情報（パスワードなど）を設定ファイルから分離できます。

5. **ボリュームによるデータ永続化**: `volumes` セクションで定義したように、コンテナが削除されてもデータを保持できます。

Docker Compose は、本番環境での使用は想定されていませんが、開発やテスト環境としては非常に有用です。本プロジェクトでは、開発中の効率化とフェーズ6での本番デプロイに向けた準備として使用しています。

### 1.6.2. Traefik の基本概念と動作原理

Traefik は、モダンなHTTPリバースプロキシおよびロードバランサーで、以下の特徴があります：

1. **自動検出**: Docker、Kubernetes、Consulなどと統合して、バックエンドサービスを自動検出します。本プロジェクトでは、Docker Providerを使用して、Dockerコンテナを自動検出しています。

2. **動的再構成**: サービスの追加や削除を検出し、設定を自動的に更新します。再起動が不要です。

3. **ルーティングのカスタマイズ**: 本プロジェクトでは、ホスト名ベースのルーティングを使用しています。例えば、`shop.localhost`へのリクエストは顧客向けフロントエンド、`admin.localhost`へのリクエストは管理者向けフロントエンドにルーティングされます。

4. **ミドルウェア**: リクエスト処理のパイプラインにミドルウェアを追加できます。認証、リダイレクト、レート制限などが可能です。

Traefikの構成は、静的構成（`traefik.yml`）と動的構成（`dynamic/config.yml`や Docker ラベル）に分かれています：

- **静的構成**: Traefikの起動時にのみ読み込まれる設定で、エントリーポイント、プロバイダー、ログなどが含まれます。
- **動的構成**: 実行時に変更可能な設定で、ルーター、サービス、ミドルウェアなどが含まれます。

本プロジェクトでは、将来追加されるサービス（バックエンド、フロントエンド）に対するルーティングを事前に設定しています。これらのサービスが実際に起動していなくても、設定自体はエラーなく読み込まれます。

### 1.6.3. LocalStack によるAWSエミュレーション

LocalStack は、AWSのクラウドサービスをローカル環境でエミュレートするツールです。本プロジェクトでは、オブザーバビリティ機能を学習するために使用します：

1. **コスト削減**: 実際のAWSサービスを使用せずに開発やテストができるため、コストを削減できます。

2. **オフライン開発**: インターネット接続がなくても開発が可能です。

3. **高速な反復**: デプロイや設定変更の反映が高速で、開発サイクルを短縮できます。

4. **統合的なテスト環境**: 複数のAWSサービスを組み合わせたアプリケーションのテストが容易です。

設定したLocalStackでは、以下のAWSサービスが利用可能です（compose.ymlの`SERVICES`パラメータで指定）：

- **S3**: オブジェクトストレージ（商品画像の保存などに使用）
- **CloudWatch**: モニタリングとオブザーバビリティサービス
  - **CloudWatch Logs (logs)**: ログ収集と分析
  - **CloudWatch Events/EventBridge (events)**: イベント処理とスケジューリング
- **Lambda**: サーバーレスコンピューティング
- **SQS**: Simple Queue Service（非同期メッセージキュー）
- **SNS**: Simple Notification Service（パブリッシュ/サブスクライブメッセージング）

初期化スクリプト（`01_init_aws_resources.sh`）では、プロジェクトで必要となる基本的なAWSリソース（S3バケット、CloudWatchロググループ、SNSトピック、SQSキュー）を事前に作成しています。

将来のフェーズで実装するオブザーバビリティ機能（ログ、メトリクス、トレース）は、これらのLocalStackサービスを活用して学習していきます。

## 1.7. 【補足情報】

### 1.7.1. Docker ネットワークについて

Docker Compose 設定では、すべてのサービスが `ecommerce-network` という名前のカスタムブリッジネットワークに接続されています。これにより、以下のメリットがあります：

1. **サービス名による名前解決**: 同じネットワーク内のサービスは、サービス名でお互いを参照できます。例えば、バックエンドサービスからMySQLに接続する場合、ホスト名として `mysql` を使用できます。

2. **ネットワークの分離**: カスタムネットワークを使用することで、プロジェクト外の他のDockerコンテナと分離できます。

3. **セキュリティの向上**: 公開する必要のないポートを外部に公開せず、同じネットワーク内のサービスだけがアクセスできるようにできます。

Docker ネットワークの詳細情報を確認するには、以下のコマンドを使用できます：

```bash
# ネットワーク一覧を表示
docker network ls

# ecommerce-networkの詳細情報を表示
docker network inspect ecommerce-network
```

### 1.7.2. ボリュームによるデータ永続化

Docker コンテナ自体は一時的なものであり、コンテナが削除されると内部のデータも失われます。これを防ぐために、Docker Compose 設定では2つの名前付きボリュームを使用しています：

1. **mysql_data**: MySQL のデータファイルを保存します。
2. **localstack_data**: LocalStack の状態とデータを保存します。

これらのボリュームを使用することで、以下のメリットがあります：

1. **データの永続化**: コンテナを再作成しても、データは失われません。
2. **パフォーマンス**: 名前付きボリュームは、バインドマウントよりも一般的にパフォーマンスが良いです。
3. **バックアップの容易さ**: ボリュームのデータを簡単にバックアップできます。

ボリュームの詳細情報を確認するには、以下のコマンドを使用できます：

```bash
# ボリューム一覧を表示
docker volume ls

# mysql_dataボリュームの詳細情報を表示
docker volume inspect mysql_data
```

ボリュームをバックアップするには、データをコンテナ外に取り出す必要があります：

```bash
# MySQLデータのバックアップ例
docker run --rm -v aws-observability-ecommerce_mysql_data:/source -v $(pwd)/backup:/backup alpine tar -czvf /backup/mysql_data_backup.tar.gz -C /source .
```

### 1.7.3. LocalStack Desktopの活用

LocalStack Desktopは、LocalStackのグラフィカルインターフェースを提供するデスクトップアプリケーションです。Docker Composeで起動したLocalStackインスタンスを視覚的に管理できるため、AWS環境の理解とデバッグが容易になります。

#### 1.7.3.1. LocalStack Desktopのインストール

以下の手順でLocalStack Desktopをインストールします：

1. 公式サイト（<https://docs.localstack.cloud/user-guide/tools/localstack-desktop/>）からLocalStack Desktopをダウンロードします。
   - Windows、macOS、Linux向けのインストーラが提供されています。

2. ダウンロードしたインストーラを実行し、画面の指示に従ってインストールを完了します。
   - macOSの場合は、ダウンロードしたDMGファイルを開き、アプリケーションフォルダにドラッグします。
   - Windowsの場合は、インストーラを実行してウィザードに従います。
   - Linuxの場合は、APTリポジトリを追加するか、AppImageを使用します。

#### 1.7.3.2. LocalStack Desktopの設定と使用方法

1. LocalStack Desktopを起動します。

2. 初回起動時に設定画面が表示されます。必要に応じて設定を行い、「Continue」または「Finish」をクリックします。
   - 基本的にはデフォルト設定で問題ありません。

3. Docker Composeで起動したLocalStackインスタンスを検出するための設定：
   - 「Settings」タブを開きます。
   - 「Docker」セクションで、「Auto-detect local Docker containers」オプションが有効になっていることを確認します。
   - 「Integration」セクションで、「LocalStack Endpoint」が `http://localhost:4566` に設定されていることを確認します。

4. ダッシュボードビューで、Docker Composeで起動したLocalStackインスタンスが表示されていることを確認します。
   - 「Instances」セクションに `ecommerce-localstack` のようなエントリが表示されるはずです。

#### 1.7.3.3. LocalStack Desktopの主な機能

1. **リソースブラウザ**: 作成したAWSリソース（S3バケット、CloudWatchロググループなど）を視覚的に参照できます。
   - 左側のナビゲーションパネルからサービスを選択し、作成したリソースを確認できます。

2. **CloudWatchログの確認**:
   - 「CloudWatch」セクションから「Logs」を選択し、ロググループとログストリームを確認できます。
   - 作成したロググループ（`/ecommerce/api`など）をクリックして、ログイベントを表示します。

3. **S3オブジェクトの管理**:
   - 「S3」セクションから作成したバケットを参照し、オブジェクトのアップロード、ダウンロード、削除ができます。
   - 例えば、`ecommerce-product-images`バケットを選択し、テスト画像をアップロードできます。

4. **Lambda関数のテスト**:
   - 「Lambda」セクションから関数を選択し、テストイベントを作成してデプロイした関数を実行できます。
   - 関数のログや結果を直接確認できます。

5. **リクエストの監視**:
   - 「Activity」タブでは、LocalStackに送信されたAPIリクエストをリアルタイムで確認できます。
   - これは、バックエンドコードのデバッグやAWS SDKの動作理解に役立ちます。

#### 1.7.3.4. LocalStack Desktopの活用例

1. **開発中のリアルタイム監視**:
   - アプリケーションが生成するCloudWatchログをリアルタイムで確認し、デバッグに活用します。

2. **S3バケットの内容確認**:
   - アップロードされた商品画像ファイルがS3バケットに正しく保存されているかを視覚的に確認します。

3. **メッセージングサービスのデバッグ**:
   - SNSトピックやSQSキューに送信されたメッセージを確認し、非同期通信のデバッグを行います。

4. **AWSリソースの手動作成**:
   - GUIを使って追加のAWSリソースを作成し、アプリケーションのテストに活用します。

LocalStack Desktopを使用することで、コマンドラインだけでは難しい視覚的な管理と監視が可能になり、AWS環境の学習とデバッグが効率化されます。

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: ポートの競合

**症状**: Docker Compose の起動時に `Error starting userland proxy: listen tcp 0.0.0.0:80: bind: address already in use` のようなエラーが表示される。

**解決策**:

1. 競合しているポートを使用しているプロセスを特定します：

   ```bash
   # Linuxの場合
   sudo lsof -i :80

   # macOSの場合
   sudo lsof -i :80

   # Windowsの場合
   netstat -aon | findstr :80
   ```

2. 競合しているプロセスを停止するか、Docker Compose の設定で使用するポートを変更します：

   ```yaml
   # compose.ymlの該当部分を変更
   ports:
     - "8080:80"  # ローカルの8080ポートをコンテナの80ポートにマッピング
   ```

3. 変更後、Docker Compose を再起動します：

   ```bash
   docker compose down
   docker compose up -d
   ```

### 1.8.2. 問題2: Traefikでのホスト名解決の問題

**症状**: `shop.localhost` や `admin.localhost` にアクセスしても、正しいサービスにルーティングされない。

**解決策**:

1. `/etc/hosts` ファイルに正しいマッピングが追加されていることを確認します：

   ```bash
   sudo nano /etc/hosts

   # 以下の行を追加または確認
   127.0.0.1 traefik.localhost api.localhost shop.localhost admin.localhost
   ```

2. Traefik の設定が正しく読み込まれていることを確認します：

   ```bash
   # ブラウザでTraefikダッシュボードにアクセス
   # http://traefik.localhost:8080

   # または、Traefikのログを確認
   docker compose logs traefik
   ```

3. Traefik の設定ファイルを修正した場合は、Traefik コンテナを再起動します：

   ```bash
   docker compose restart traefik
   ```

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **Docker Compose による環境の統合**: 複数のサービスを一元管理し、開発環境を簡単に再現できるようになりました。これにより、チーム全体で一貫した開発が可能になります。

2. **Traefik によるルーティング**: ホスト名ベースのルーティングを設定し、顧客向けと管理者向けの分離された環境を準備しました。これにより、フロントエンドを独立して開発・デプロイできます。

3. **LocalStack によるAWSエミュレーション**: AWSサービスをローカルでエミュレートする環境を準備しました。これにより、コストをかけずにオブザーバビリティ機能を学習できます。

4. **持続可能な開発環境**: ボリュームによるデータ永続化とサービス間のネットワーク接続を設定することで、長期的な開発に適した環境を構築しました。

これらのポイントは、次回以降の実装においても基盤となる重要な概念です。特に、Docker環境の理解はプロジェクト全体を通じて必要となります。

## 1.10. 【次回の準備】

次回（Day 2）では、バックエンドの基本構造を実装します。以下の点について事前に確認しておくと良いでしょう：

1. **Go言語の基本**: Go言語の基本的な構文や概念を確認しておくと、スムーズに進められます。

2. **Echoフレームワーク**: Go言語のWebフレームワークであるEchoの基本概念や使い方を確認しておくと良いでしょう。

3. **Docker環境の動作確認**: 今回構築したDocker環境が正常に動作していることを確認しておきましょう。次回もこの環境をベースに開発を進めます。

4. **依存関係管理ツール**: Goのモジュール管理（go.modとgo.sum）について理解しておくと良いでしょう。

次回はこれらの知識をベースに、バックエンドの基本構造を実装していきます。

## 1.11. 【.envrc サンプル】

以下は本日の実装で使用する.envrcのサンプルです。ご自身の環境に合わせて修正して使用してください。このファイルはgitにコミットしないようにしてください。

```bash
# .envrc サンプル
# direnvがインストールされている場合、このディレクトリに入ると自動的に環境変数が設定されます
# このファイルはgitにコミットしないでください

# MySQL設定
export MYSQL_ROOT_PASSWORD=rootpassword
export MYSQL_DATABASE=ecommerce
export MYSQL_USER=ecommerce_user
export MYSQL_PASSWORD=ecommerce_password

# AWS設定（LocalStack用）
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=ap-northeast-1
export AWS_ENDPOINT_URL=http://localhost:4566

# 開発環境設定
export ENVIRONMENT=development
```

.gitignoreファイルに.envrcを追加して、誤ってコミットしないようにしましょう：

```bash
echo ".envrc" >> .gitignore
```
