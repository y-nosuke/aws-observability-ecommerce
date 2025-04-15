# 1. Week 1 - Day 5: LocalStackとCI/CDの設定

## 1.1. 目次

- [1. Week 1 - Day 5: LocalStackとCI/CDの設定](#1-week-1---day-5-localstackとcicdの設定)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. LocalStackの設定と基本操作](#141-localstackの設定と基本操作)
    - [1.4.2. LocalStack の設定](#142-localstack-の設定)
    - [1.4.3. go-taskによるタスクランナーの設定](#143-go-taskによるタスクランナーの設定)
    - [1.4.4. GitHub Actionsを使った基本CI/CDワークフローの設定](#144-github-actionsを使った基本cicdワークフローの設定)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
    - [1.5.1. LocalStackの確認](#151-localstackの確認)
    - [1.5.2. go-taskとGitHub Actionsの確認](#152-go-taskとgithub-actionsの確認)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. LocalStackとは](#161-localstackとは)
    - [1.6.2. go-taskの利点と使い方](#162-go-taskの利点と使い方)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. AWS CLIプロファイルの活用](#171-aws-cliプロファイルの活用)
    - [1.7.2. LocalStack Desktopの活用](#172-localstack-desktopの活用)
      - [1.7.2.1. LocalStack Desktopのインストール](#1721-localstack-desktopのインストール)
      - [1.7.2.2. LocalStack Desktopの設定と使用方法](#1722-localstack-desktopの設定と使用方法)
      - [1.7.2.3. LocalStack Desktopの主な機能](#1723-localstack-desktopの主な機能)
      - [1.7.2.4. LocalStack Desktopの活用例](#1724-localstack-desktopの活用例)
    - [1.7.3. GitHub Actionsの発展的な使い方](#173-github-actionsの発展的な使い方)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: LocalStackのサービスに接続できない](#181-問題1-localstackのサービスに接続できない)
    - [1.8.2. 問題2: GitHub Actionsワークフローが失敗する](#182-問題2-github-actionsワークフローが失敗する)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- LocalStackを使用してAWSサービスをローカル環境でエミュレートする方法を学ぶ
- S3バケットの基本操作をLocalStack上で実践する
- go-taskを使ったタスク自動化の基盤を構築する
- GitHub Actionsを使った基本的なCI/CDパイプラインを設定する
- ローカル開発環境とCI/CD環境の連携方法を理解する

## 1.3. 【準備】

LocalStackとCI/CD環境をセットアップするために必要なツールとAWS CLIの基本設定を行います。

### 1.3.1. チェックリスト

- [ ] Docker Composeが正常に動作していること
- [ ] awslocal（LocalStack用AWS CLIラッパー）がインストールされていること

  ```bash
  # まずpipxをインストール
  brew install pipx
  # pipxを使ってawscli-localをインストール
  pipx install awscli-local
  awslocal --version
  # aws-cli/x.x.x Python/x.x.x ...
  ```

- [ ] go-taskがインストールされていること（[Installation - Task](https://taskfile.dev/installation/)）
- [ ] GitHubアカウントとリポジトリへのアクセス権があること
- [ ] .envrcファイルかdotenvを使用する準備があること

## 1.4. 【手順】

### 1.4.1. LocalStackの設定と基本操作

LocalStackはAWSサービスをローカル環境でエミュレートするツールです。Docker Composeファイルに追加して、AWSサービスをローカルで使用できるようにします。

```yaml
# docker-compose.ymlに追加
services:
  # 既存のサービス設定...

  localstack:
    container_name: "${PROJECT_NAME:-ecommerce}-localstack"
    image: localstack/localstack:2.2.0
    ports:
      - "4566:4566"            # LocalStackのメインエンドポイント
      - "4510-4559:4510-4559"  # 内部サービス用ポート範囲
    environment:
      - DEBUG=${DEBUG:-0}
      - DOCKER_HOST=unix:///var/run/docker.sock
      - HOSTNAME_EXTERNAL=localstack
      - SERVICES=s3,cloudwatch,logs,lambda,iam,sqs,sns,events
    volumes:
      - "${LOCALSTACK_VOLUME_DIR:-./volume/localstack}:/var/lib/localstack"
      - "/var/run/docker.sock:/var/run/docker.sock"
      - "./infra/localstack/init-scripts:/etc/localstack/init/ready.d:ro"
    networks:
      - app-network
```

LocalStackサービスを起動し、基本操作を確認します。

```bash
# LocalStackの起動
docker-compose up -d localstack

# LocalStackのステータス確認
curl http://localhost:4566/health
```

### 1.4.2. LocalStack の設定

LocalStack の初期化スクリプトを作成します。

```bash
# LocalStack 初期化スクリプトを作成
mkdir -p infra/localstack/init-scripts
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

### 1.4.3. go-taskによるタスクランナーの設定

go-taskはMakefileの代替としてYAML形式でタスクを定義できるツールです。プロジェクトのルートディレクトリに`Taskfile.yml`を作成します。

```yaml
# Taskfile.yml
version: '3'

vars:
  PROJECT_NAME: ecommerce-app

tasks:
  default:
    desc: プロジェクトのヘルプを表示
    cmds:
      - task -l
    silent: true

  setup:
    desc: プロジェクトのセットアップ
    cmds:
      - task: setup:backend
      - task: setup:frontend-customer
      - task: setup:frontend-admin

  setup:backend:
    desc: バックエンドのセットアップ
    dir: backend
    cmds:
      - go mod tidy
      - go mod verify

  setup:frontend-customer:
    desc: 顧客向けフロントエンドのセットアップ
    dir: frontend-customer
    cmds:
      - npm install

  setup:frontend-admin:
    desc: 管理者向けフロントエンドのセットアップ
    dir: frontend-admin
    cmds:
      - npm install

  up:
    desc: 開発環境を起動
    cmds:
      - docker-compose up -d

  down:
    desc: 開発環境を停止
    cmds:
      - docker-compose down

  logs:
    desc: コンテナのログを表示
    cmds:
      - docker-compose logs -f {{.CLI_ARGS}}

  lint:
    desc: コードの静的解析を実行
    cmds:
      - task: lint:backend
      - task: lint:frontend-customer
      - task: lint:frontend-admin

  lint:backend:
    desc: バックエンドのlintを実行
    dir: backend
    cmds:
      - golangci-lint run

  lint:frontend-customer:
    desc: 顧客向けフロントエンドのlintを実行
    dir: frontend-customer
    cmds:
      - npm run lint

  lint:frontend-admin:
    desc: 管理者向けフロントエンドのlintを実行
    dir: frontend-admin
    cmds:
      - npm run lint

  test:
    desc: テストを実行
    cmds:
      - task: test:backend
      - task: test:frontend-customer
      - task: test:frontend-admin

  test:backend:
    desc: バックエンドのテストを実行
    dir: backend
    cmds:
      - go test ./... -v

  test:frontend-customer:
    desc: 顧客向けフロントエンドのテストを実行
    dir: frontend-customer
    cmds:
      - npm test

  test:frontend-admin:
    desc: 管理者向けフロントエンドのテストを実行
    dir: frontend-admin
    cmds:
      - npm test

  localstack:create-bucket:
    desc: S3バケットを作成
    cmds:
      - awslocal s3 mb s3://product-images

  localstack:list-buckets:
    desc: S3バケットを一覧表示
    cmds:
      - awslocal s3 ls
```

go-taskの基本コマンドを試してみましょう:

```bash
# タスク一覧を表示
task -l

# 開発環境を起動
task up

# バックエンドのセットアップ
task setup:backend

# lintを実行
task lint

# テストを実行
task test
```

### 1.4.4. GitHub Actionsを使った基本CI/CDワークフローの設定

GitHubリポジトリに基本的なCI/CDワークフローを設定します。`.github/workflows`ディレクトリを作成し、基本的なワークフローファイルを追加します。

```bash
mkdir -p .github/workflows

touch .github/workflows/{backend-ci.yml,frontend-ci.yml,deploy.yml}
```

バックエンドのCIワークフローを作成:

```yaml
# .github/workflows/backend-ci.yml
name: Backend CI

on:
  push:
    branches: [ main, develop ]
    paths:
      - 'backend/**'
      - '.github/workflows/backend-ci.yml'
  pull_request:
    branches: [ main, develop ]
    paths:
      - 'backend/**'
      - '.github/workflows/backend-ci.yml'

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Install dependencies
        run: |
          cd backend
          go mod download

      - name: Install golangci-lint
        run: |
          curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.53.3

      - name: Lint
        run: |
          cd backend
          golangci-lint run

  test:
    runs-on: ubuntu-latest
    needs: lint
    steps:
      - uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Install dependencies
        run: |
          cd backend
          go mod download

      - name: Test
        run: |
          cd backend
          go test -v ./...
```

フロントエンドのCIワークフローを作成:

```yaml
# .github/workflows/frontend-ci.yml
name: Frontend CI

on:
  push:
    branches: [ main, develop ]
    paths:
      - 'frontend-**/**'
      - '.github/workflows/frontend-ci.yml'
  pull_request:
    branches: [ main, develop ]
    paths:
      - 'frontend-**/**'
      - '.github/workflows/frontend-ci.yml'

jobs:
  lint-and-test:
    runs-on: ubuntu-latest
    strategy:
      matrix:
        app: [frontend-customer, frontend-admin]

    steps:
      - uses: actions/checkout@v3

      - name: Set up Node.js
        uses: actions/setup-node@v3
        with:
          node-version: '18'
          cache: 'npm'
          cache-dependency-path: ${{ matrix.app }}/package-lock.json

      - name: Install dependencies
        run: |
          cd ${{ matrix.app }}
          npm ci

      - name: Lint
        run: |
          cd ${{ matrix.app }}
          npm run lint

      - name: Test
        run: |
          cd ${{ matrix.app }}
          npm test
```

シンプルなデプロイワークフローのサンプル:

```yaml
# .github/workflows/deploy.yml
name: Deploy

on:
  push:
    branches: [ main ]
    paths-ignore:
      - '**.md'

jobs:
  build-and-push:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Login to DockerHub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKER_USERNAME }}
          password: ${{ secrets.DOCKER_PASSWORD }}

      - name: Build and push backend
        uses: docker/build-push-action@v4
        with:
          context: ./backend
          push: true
          tags: y-nosuke/aws-observability-ecommerce-backend:latest

      - name: Build and push frontend-customer
        uses: docker/build-push-action@v4
        with:
          context: ./frontend-customer
          push: true
          tags: y-nosuke/aws-observability-ecommerce-frontend-customer:latest

      - name: Build and push frontend-admin
        uses: docker/build-push-action@v4
        with:
          context: ./frontend-admin
          push: true
          tags: y-nosuke/aws-observability-ecommerce-frontend-admin:latest

  # デプロイステップは実際の環境に合わせて実装
```

## 1.5. 【確認ポイント】

### 1.5.1. LocalStackの確認

- [ ] LocalStackコンテナが正常に起動している（`docker-compose ps`で確認）
- [ ] LocalStackヘルスチェックが正常に応答している（`curl http://localhost:4566/health`で確認）
- [ ] LocalStack が正常に動作し、awslocalを使って操作できる

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

### 1.5.2. go-taskとGitHub Actionsの確認

- [ ] Taskfile.ymlが作成され、主要なタスクが定義されている
- [ ] `task -l`コマンドでタスク一覧が表示される
- [ ] 各環境（バックエンド、顧客向けフロントエンド、管理者向けフロントエンド）のタスクが実行できる
- [ ] GitHub Actionsのワークフローファイルが作成されている
- [ ] CI/CDワークフローの設定が適切である

## 1.6. 【詳細解説】

### 1.6.1. LocalStackとは

LocalStackは、AWSのクラウドサービスをローカル環境でエミュレートするツールです。開発やテストのために、AWS本番環境にデプロイすることなく、AWSのサービスを使用できます。

**主なメリット:**

1. **コスト削減**: 開発環境でのAWS使用料が発生しない
2. **開発速度の向上**: ネットワークレイテンシなしで高速にサービスを利用できる
3. **オフライン開発**: インターネット接続がなくても開発できる
4. **制御とカスタマイズ**: 環境を完全に制御でき、テストシナリオを容易に作成できる

**LocalStackで利用可能な主なAWSサービス:**

- S3（Simple Storage Service）
- Lambda
- SQS（Simple Queue Service）
- SNS（Simple Notification Service）
- DynamoDB
- CloudWatch Logs
- IAM（Identity and Access Management）
- その他多数

**LocalStackの設定方法:**

LocalStackはDockerコンテナとして実行され、環境変数で有効にするサービスを制御します。主な環境変数:

- `SERVICES`: 有効にするサービスのカンマ区切りリスト
- `DEBUG`: デバッグモードの有効化
- `DATA_DIR`: データの永続化ディレクトリ

**AWS CLIとの連携:**

AWS CLIをLocalStackと連携させるには、`awslocal`コマンドを使用します。これはエンドポイントURLとプロファイル設定を自動的に処理してくれるAWS CLIのラッパーです。

```bash
# awslocalのインストール（まだインストールしていない場合）
pipx install awscli-local

# S3バケットの作成
awslocal s3 mb s3://my-bucket

# CloudWatch Logsグループの作成
awslocal logs create-log-group --log-group-name /my/logs
```

awslocalは内部で以下のような設定を自動的に行います：

```bash
# これらの設定は自動的に行われるので、手動で設定する必要はありません
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=ap-northeast-1
export AWS_ENDPOINT_URL=http://localhost:4566
```

**SDKとの連携:**

各言語のAWS SDKでLocalStackを使用するには、エンドポイント設定を変更します。GoのAWS SDK v2の例:

```go
customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
    return aws.Endpoint{
        URL:           "http://localhost:4566",
        SigningRegion: region,
    }, nil
})

cfg, err := config.LoadDefaultConfig(ctx,
    config.WithRegion("ap-northeast-1"),
    config.WithEndpointResolverWithOptions(customResolver),
)
```

### 1.6.2. go-taskの利点と使い方

go-taskは、YAMLファイルを使用してタスクを定義する現代的なタスクランナーで、Makefileの代替として使用できます。

**主なメリット:**

1. **クロスプラットフォーム**: Windows、macOS、Linuxで同じく動作
2. **シンプルな構文**: YAMLを使用しているため理解しやすい
3. **依存関係管理**: タスク間の依存関係を簡単に定義できる
4. **並行実行**: タスクの並行実行をサポート
5. **条件付き実行**: 特定の条件が満たされた場合のみタスクを実行

**基本的な使い方:**

Taskfile.ymlの基本構造:

```yaml
version: '3'

tasks:
  # 基本的なタスク
  hello:
    desc: 挨拶を表示
    cmds:
      - echo "Hello, World!"

  # 依存関係を持つタスク
  greet:
    desc: 名前付きの挨拶を表示
    deps: [hello]
    cmds:
      - echo "Nice to meet you!"

  # ディレクトリを指定するタスク
  build:
    desc: プロジェクトをビルド
    dir: ./backend
    cmds:
      - go build -o bin/app main.go

  # 変数を使うタスク
  say:
    desc: メッセージを表示
    cmds:
      - echo "{{.MESSAGE}}"
    vars:
      MESSAGE: Default message

  # サブタスクに分割したタスク
  deploy:
    desc: デプロイを実行
    cmds:
      - task: deploy:staging
      - task: deploy:production

  deploy:staging:
    desc: ステージング環境にデプロイ
    cmds:
      - echo "Deploying to staging..."

  deploy:production:
    desc: 本番環境にデプロイ
    cmds:
      - echo "Deploying to production..."
```

**主要なコマンド:**

- `task -l`: 定義されているタスクとその説明を一覧表示
- `task <タスク名>`: 特定のタスクを実行
- `task <タスク名> -- <引数>`: 引数付きでタスクを実行
- `task -v <タスク名>`: 詳細モードでタスクを実行（各コマンドの出力を表示）
- `task -d <タスク名>`: ドライランモードでタスクを実行（実際には実行せず、何が実行されるかを表示）

**変数の種類:**

1. **静的変数**: Taskfile内で直接定義

   ```yaml
   vars:
     PROJECT_NAME: myproject
   ```

2. **環境変数**: 環境変数を使用

   ```yaml
   env:
     ENV_VAR: '{{.ENV_VAR}}'
   ```

3. **コマンド出力**: コマンドの出力を変数として使用

   ```yaml
   vars:
     GIT_COMMIT:
       sh: git rev-parse --short HEAD
   ```

4. **タスクごとの変数**: 特定のタスク内でのみ使用する変数

   ```yaml
   tasks:
     build:
       vars:
         BINARY_NAME: app
       cmds:
         - go build -o {{.BINARY_NAME}}
   ```

## 1.7. 【補足情報】

### 1.7.1. AWS CLIプロファイルの活用

AWS CLIはAWSサービスを操作するためのコマンドラインツールです。複数の環境（開発、ステージング、本番など）を扱う場合、プロファイルを活用すると便利です。

**異なる環境用のawslocal設定:**

実際の開発ではLocalStackの環境が複数ある場合があります。そのような場合は、環境変数を切り替えることでawslocalの動作を調整できます。

```bash
# 特定のLocalStack環境のエンドポイントを指定
export LOCALSTACK_ENDPOINT=http://custom-localstack:4566
awslocal s3 ls

# 特定のリージョンを指定
export AWS_DEFAULT_REGION=us-east-1
awslocal s3 ls
```

**複数プロジェクトでの使用例:**

```bash
# プロジェクトAの設定
export LOCALSTACK_ENDPOINT=http://localstack-project-a:4566
export AWS_DEFAULT_REGION=ap-northeast-1
awslocal s3 ls

# プロジェクトBの設定
export LOCALSTACK_ENDPOINT=http://localstack-project-b:4566
export AWS_DEFAULT_REGION=us-west-2
awslocal s3 ls
```

**.envrcを使った設定例:**

```bash
# .envrc
export LOCALSTACK_ENDPOINT=http://localhost:4566
export AWS_DEFAULT_REGION=ap-northeast-1
```

**プロファイルと環境変数の優先順位:**

1. コマンドラインオプション (`--profile`, `--region` など)
2. 環境変数 (`AWS_PROFILE`, `AWS_REGION` など)
3. ~/.aws/configと~/.aws/credentialsファイル

### 1.7.2. LocalStack Desktopの活用

LocalStack Desktopは、LocalStackのグラフィカルインターフェースを提供するデスクトップアプリケーションです。Docker Composeで起動したLocalStackインスタンスを視覚的に管理できるため、AWS環境の理解とデバッグが容易になります。

#### 1.7.2.1. LocalStack Desktopのインストール

以下の手順でLocalStack Desktopをインストールします：

1. 公式サイト（<https://docs.localstack.cloud/user-guide/tools/localstack-desktop/>）からLocalStack Desktopをダウンロードします。
   - Windows、macOS、Linux向けのインストーラが提供されています。

2. ダウンロードしたインストーラを実行し、画面の指示に従ってインストールを完了します。
   - macOSの場合は、ダウンロードしたDMGファイルを開き、アプリケーションフォルダにドラッグします。
   - Windowsの場合は、インストーラを実行してウィザードに従います。
   - Linuxの場合は、APTリポジトリを追加するか、AppImageを使用します。

#### 1.7.2.2. LocalStack Desktopの設定と使用方法

1. LocalStack Desktopを起動します。

2. 初回起動時に設定画面が表示されます。必要に応じて設定を行い、「Continue」または「Finish」をクリックします。
   - 基本的にはデフォルト設定で問題ありません。

3. Docker Composeで起動したLocalStackインスタンスを検出するための設定：
   - 「Settings」タブを開きます。
   - 「Docker」セクションで、「Auto-detect local Docker containers」オプションが有効になっていることを確認します。
   - 「Integration」セクションで、「LocalStack Endpoint」が `http://localhost:4566` に設定されていることを確認します。

4. ダッシュボードビューで、Docker Composeで起動したLocalStackインスタンスが表示されていることを確認します。
   - 「Instances」セクションに `ecommerce-localstack` のようなエントリが表示されるはずです。

#### 1.7.2.3. LocalStack Desktopの主な機能

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

#### 1.7.2.4. LocalStack Desktopの活用例

1. **開発中のリアルタイム監視**:
   - アプリケーションが生成するCloudWatchログをリアルタイムで確認し、デバッグに活用します。

2. **S3バケットの内容確認**:
   - アップロードされた商品画像ファイルがS3バケットに正しく保存されているかを視覚的に確認します。

3. **メッセージングサービスのデバッグ**:
   - SNSトピックやSQSキューに送信されたメッセージを確認し、非同期通信のデバッグを行います。

4. **AWSリソースの手動作成**:
   - GUIを使って追加のAWSリソースを作成し、アプリケーションのテストに活用します。

LocalStack Desktopを使用することで、コマンドラインだけでは難しい視覚的な管理と監視が可能になり、AWS環境の学習とデバッグが効率化されます。

### 1.7.3. GitHub Actionsの発展的な使い方

基本的なCI/CDワークフローの他に、GitHub Actionsではより高度な機能も利用できます。

**1. マトリックスビルド:**

複数のOS、言語バージョン、ブラウザなどの組み合わせでテストを実行できます。

```yaml
jobs:
  test:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ubuntu-latest, windows-latest, macos-latest]
        node-version: [14, 16, 18]
    steps:
      - uses: actions/checkout@v3
      - name: Use Node.js ${{ matrix.node-version }}
        uses: actions/setup-node@v3
        with:
          node-version: ${{ matrix.node-version }}
      - run: npm ci
      - run: npm test
```

**2. 環境変数と秘密情報:**

機密情報を安全に扱うためのシークレットを設定できます。

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    env:
      APP_ENV: production
    steps:
      - uses: actions/checkout@v3
      - name: Deploy to Production
        env:
          API_KEY: ${{ secrets.API_KEY }}
        run: ./deploy.sh $API_KEY
```

**3. アーティファクトの保存と共有:**

ビルド成果物を保存し、ジョブ間で共有できます。

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Build
        run: ./build.sh
      - name: Upload artifact
        uses: actions/upload-artifact@v3
        with:
          name: my-artifact
          path: ./dist

  deploy:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Download artifact
        uses: actions/download-artifact@v3
        with:
          name: my-artifact
          path: ./dist
      - name: Deploy
        run: ./deploy.sh
```

**4. ワークフローの条件付き実行:**

特定の条件でのみワークフローやジョブを実行できます。

```yaml
jobs:
  deploy-staging:
    if: github.ref == 'refs/heads/develop'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Deploy to Staging
        run: ./deploy-staging.sh

  deploy-production:
    if: github.ref == 'refs/heads/main'
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Deploy to Production
        run: ./deploy-production.sh
```

**5. 再利用可能なワークフロー:**

複数のリポジトリで再利用できるワークフローを定義できます。

```yaml
# .github/workflows/reusable-workflow.yml
name: Reusable Workflow
on:
  workflow_call:
    inputs:
      environment:
        required: true
        type: string
    secrets:
      api_key:
        required: true

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - name: Deploy to ${{ inputs.environment }}
        env:
          API_KEY: ${{ secrets.api_key }}
        run: ./deploy.sh ${{ inputs.environment }}
```

別のワークフローからの呼び出し:

```yaml
jobs:
  call-workflow:
    uses: ./.github/workflows/reusable-workflow.yml
    with:
      environment: production
    secrets:
      api_key: ${{ secrets.API_KEY }}
```

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: LocalStackのサービスに接続できない

**症状**: AWS CLIやアプリケーションからLocalStackのサービスにアクセスできない。

**解決策**:

1. LocalStackコンテナが実行中か確認する

   ```bash
   docker-compose ps
   ```

2. LocalStackのヘルスエンドポイントをチェックする

   ```bash
   curl http://localhost:4566/health
   ```

3. エンドポイントURLが正しく設定されているか確認する

   ```bash
   # AWS CLIの場合
   aws --endpoint-url=http://localhost:4566 s3 ls
   ```

4. Docker内からLocalStackにアクセスする場合は、ホスト名としてサービス名を使用する

   ```bash
   # Docker内からの接続では、ホスト名 "localstack" を使用
   export AWS_ENDPOINT_URL=http://localstack:4566
   ```

5. DockerネットワークがLocalStackコンテナとアプリケーションコンテナの間で正しく設定されているか確認する

6. ファイアウォールや他の設定がポート4566へのアクセスをブロックしていないか確認する

### 1.8.2. 問題2: GitHub Actionsワークフローが失敗する

**症状**: CI/CDワークフローが失敗し、デプロイや自動テストが実行されない。

**解決策**:

1. ワークフローの出力ログを確認し、具体的なエラーメッセージを特定する

2. プロジェクトの依存関係が正しくインストールされているか確認する

   ```yaml
   # 依存関係のインストールを詳細化
   - name: Install dependencies
     run: |
       npm ci --verbose
       npm list --depth=0
   ```

3. キャッシュを使用して依存関係のインストールを高速化し、一貫性を確保する

   ```yaml
   - name: Cache dependencies
     uses: actions/cache@v3
     with:
       path: ~/.npm
       key: ${{ runner.os }}-node-${{ hashFiles('**/package-lock.json') }}
       restore-keys: |
         ${{ runner.os }}-node-
   ```

4. 必要なシークレットが正しく設定されているか確認する
   - GitHub リポジトリの Settings > Secrets and variables > Actions に移動
   - 必要なシークレットが定義されているか確認

5. 特定のジョブだけを再実行して問題を切り分ける
   - GitHub Actionsの実行ページで「Re-run jobs」を選択
   - 失敗したジョブのみを再実行

6. ローカル環境でワークフローの手順を再現して、問題が環境固有かどうかを確認する

## 1.9. 【今日の重要なポイント】

- LocalStackはAWSサービスをローカルで利用するための強力なツールであり、開発効率とコスト削減に貢献します
- S3などのAWSサービスをLocalStackで利用する際は、エンドポイントURLの設定が重要です
- go-taskを使うことで、複雑なプロジェクトタスクを整理し、チーム全体で一貫した方法で実行できるようになります
- GitHub Actionsを使ったCI/CDパイプラインは、コードの品質管理と自動デプロイを実現するための基礎です
- ローカル開発環境でのエミュレーションとCI/CD環境との連携方法を理解することで、効率的な開発フローを構築できます
- 適切なAWS CLIプロファイル設定により、複数の環境（ローカル、開発、本番など）を簡単に切り替えられます

## 1.10. 【次回の準備】

Week 2ではデータベースのセットアップとデータモデルの設計を行います。以下の点について前もって確認しておくと良いでしょう：

1. MySQLの基本概念とSQL文の書き方
2. データベーススキーマ設計の基本原則（正規化、リレーションシップなど）
3. ORMとコード生成の基本概念
4. Go言語でのデータベース操作

また、以下のツールのインストールも準備しておくと良いでしょう：

1. MySQLクライアント（CLI/GUI）
2. データベース設計ツール（例: dbdiagram.io、MySQL Workbench）
3. sqlboiler（Go用のORM）のインストール

## 1.11. 【.envrc サンプル】

direnvを使用している場合、以下のような.envrcファイルを作成すると便利です。

```bash
# .envrc
# AWS/LocalStack設定
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_DEFAULT_REGION=ap-northeast-1
export AWS_ENDPOINT_URL=http://localhost:4566

# データベース設定
export DB_HOST=localhost
export DB_PORT=3306
export DB_NAME=ecommerce
export DB_USER=ecommerce_user
export DB_PASSWORD=ecommerce_password

# バックエンド設定
export BACKEND_PORT=8080
export BACKEND_DEBUG=true

# フロントエンド設定
export FRONTEND_CUSTOMER_PORT=3000
export FRONTEND_ADMIN_PORT=3001

# PASSを通すために使用
export PATH="$PWD/node_modules/.bin:$PATH"
```

これにディレクトリに入ると自動的に環境変数が設定されるため、各種コマンドを実行しやすくなります。
