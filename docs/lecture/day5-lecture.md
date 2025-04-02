# 1. Week 1 - Day 5: LocalStackとGitHubリポジトリ設定

## 1.1. 目次

- [1. Week 1 - Day 5: LocalStackとGitHubリポジトリ設定](#1-week-1---day-5-localstackとgithubリポジトリ設定)
  - [1.1. 目次](#11-目次)
  - [1.2. 【要点】](#12-要点)
  - [1.3. 【準備】](#13-準備)
    - [1.3.1. チェックリスト](#131-チェックリスト)
  - [1.4. 【手順】](#14-手順)
    - [1.4.1. LocalStackの設定と基本操作](#141-localstackの設定と基本操作)
    - [1.4.2. S3バケットの作成とテスト](#142-s3バケットの作成とテスト)
    - [1.4.3. GitHubリポジトリのセットアップ](#143-githubリポジトリのセットアップ)
    - [1.4.4. ブランチ戦略とprotected branchの設定](#144-ブランチ戦略とprotected-branchの設定)
    - [1.4.5. go-taskによるタスクランナーの設定](#145-go-taskによるタスクランナーの設定)
    - [1.4.6. GitHub Actionsを使った基本ムCIワークフローの設定](#146-github-actionsを使った基本ムciワークフローの設定)
  - [1.5. 【確認ポイント】](#15-確認ポイント)
    - [1.5.1. LocalStackの確認](#151-localstackの確認)
    - [1.5.2. GitHubリポジトリの確認](#152-githubリポジトリの確認)
    - [1.5.3. go-taskとGitHub Actionsの確認](#153-go-taskとgithub-actionsの確認)
  - [1.6. 【詳細解説】](#16-詳細解説)
    - [1.6.1. LocalStackとは](#161-localstackとは)
    - [1.6.2. go-taskの利点と使い方](#162-go-taskの利点と使い方)
    - [1.6.3. GitHubのブランチ保護と開発ワークフロー](#163-githubのブランチ保護と開発ワークフロー)
  - [1.7. 【補足情報】](#17-補足情報)
    - [1.7.1. AWS CLIプロファイルの活用](#171-aws-cliプロファイルの活用)
    - [1.7.2. GitHub Actionsの発展的な使い方](#172-github-actionsの発展的な使い方)
  - [1.8. 【よくある問題と解決法】](#18-よくある問題と解決法)
    - [1.8.1. 問題1: LocalStackのサービスに接続できない](#181-問題1-localstackのサービスに接続できない)
    - [1.8.2. 問題2: GitHub Actionsワークフローが失敗する](#182-問題2-github-actionsワークフローが失敗する)
  - [1.9. 【今日の重要なポイント】](#19-今日の重要なポイント)
  - [1.10. 【次回の準備】](#110-次回の準備)
  - [1.11. 【.envrc サンプル】](#111-envrc-サンプル)

## 1.2. 【要点】

- LocalStackを使用してAWSサービス（S3、Lambda、CloudWatch）をローカル環境でエミュレーションする方法を習得する
- GitHubリポジトリの設定とブランチ戦略の確立で効率的なチーム開発の基盤を作る
- go-taskを使用したタスクランナーの設定で開発ワークフローを効率化する
- GitHub Actionsを使った基本的なCI設定でコード品質を継続的に確保する
- プロジェクト全体の構造を整理し、フェーズ1の基盤を完成させる

## 1.3. 【準備】

LocalStackとGitHubリポジトリの設定を行うための準備として、必要なツールとアカウントを確認します。本日の実装では、AWSサービスをローカルでエミュレートするためのLocalStackと、コード管理のためのGitHubリポジトリの設定を行います。

### 1.3.1. チェックリスト

- [ ] Dockerが正常に動作している（docker-compose upで環境が起動できる）
- [ ] AWS CLIがインストールされている（バージョン2以降推奨）
- [ ] AWS CLIの基本的な使い方を理解している
- [ ] GitHubアカウントを持っている
- [ ] Gitの基本コマンド（clone, commit, push, pull, branch）を理解している
- [ ] go-taskがインストールされている（または今日インストール予定）
- [ ] オプション: AWS Management Consoleの基本的な操作経験がある

## 1.4. 【手順】

今日は以下の手順で、LocalStackの設定とGitHubリポジトリのセットアップを行います。

### 1.4.1. LocalStackの設定と基本操作

まず、Docker Composeファイルを更新してLocalStackサービスを追加します。

```bash
# compose.ymlがあるプロジェクトのルートディレクトリに移動
cd /path/to/your/project

# infra/localstackディレクトリを作成
mkdir -p infra/localstack
```

次に、LocalStack用の設定ファイルを作成します：

```bash
# LocalStackの初期化スクリプトを作成
touch infra/localstack/init-aws.sh
chmod +x infra/localstack/init-aws.sh
```

`init-aws.sh`に以下の内容を記述します：

```bash
#!/bin/bash
set -e

# S3バケットの作成
awslocal s3 mb s3://product-images

# CloudWatch Logsのロググループを作成
awslocal logs create-log-group --log-group-name /aws/lambda/image-processor

echo "LocalStack初期化が完了しました"
```

次に、compose.ymlファイルを編集してLocalStackサービスを追加します：

```yaml
# LocalStackサービスの設定を追加
services:
  # 既存のサービス設定...

  localstack:
    image: localstack/localstack:latest
    container_name: localstack
    ports:
      - "4566:4566"      # LocalStackのメインポート
      - "4510-4559:4510-4559"  # 追加ポート範囲
    environment:
      - SERVICES=s3,lambda,cloudwatch,logs,events
      - DEBUG=1
      - DOCKER_HOST=unix:///var/run/docker.sock
      - DEFAULT_REGION=ap-northeast-1
    volumes:
      - "./infra/localstack:/docker-entrypoint-initaws.d"
      - "/var/run/docker.sock:/var/run/docker.sock"
```

Docker Composeを再起動してLocalStackを含む環境を起動します：

```bash
docker-compose down
docker-compose up -d
```

LocalStackが正常に起動しているか確認します：

```bash
docker logs localstack
```

AWS CLIをLocalStackに向けて設定します：

```bash
# ~/.aws/credentials または適切な場所に以下を追加
[localstack]
aws_access_key_id = test
aws_secret_access_key = test

# ~/.aws/config または適切な場所に以下を追加
[profile localstack]
region = ap-northeast-1
output = json
endpoint_url = http://localhost:4566
```

### 1.4.2. S3バケットの作成とテスト

LocalStackが起動した後、S3バケットの作成と基本的なオペレーションをテストします。

```bash
# LocalStackのS3バケットリストを表示
aws --endpoint-url=http://localhost:4566 s3 ls

# もしくは設定したプロファイルを使用する場合
aws --profile localstack s3 ls
```

テスト用のファイルをアップロードしてS3の動作確認をします：

```bash
# テスト用の画像ファイルを用意
mkdir -p tmp
touch tmp/test-image.jpg

# S3にアップロード
aws --endpoint-url=http://localhost:4566 s3 cp tmp/test-image.jpg s3://product-images/

# アップロードされたファイルを確認
aws --endpoint-url=http://localhost:4566 s3 ls s3://product-images/
```

### 1.4.3. GitHubリポジトリのセットアップ

次に、GitHubリポジトリをセットアップして、プロジェクトのコードを管理します。

1. まず、GitHubにログインし、新しいリポジトリを作成します。

   - GitHubにログイン
   - 「New repository」ボタンをクリック
   - リポジトリ名を「aws-observability-ecommerce」など適切な名前に設定
   - 説明文を追加：「AWS オブザーバビリティ学習用 eコマースアプリ」
   - PublicもしくはPrivateを選択
   - 「Add a README file」をチェック
   - 「Add .gitignore」をチェックし「Go」テンプレートを選択
   - 「Create repository」をクリック

2. ローカル環境にリポジトリをクローンします。

    ```bash
    # 新しいリポジトリをクローン
    git clone https://github.com/yourusername/aws-observability-ecommerce.git
    cd aws-observability-ecommerce
    ```

3. 既存のプロジェクトファイルをリポジトリにコピーします。

    ```bash
    # 既存のプロジェクトファイルを新しいリポジトリにコピー
    cp -r /path/to/your/existing/project/* .
    ```

4. .gitignoreファイルを確認し、必要に応じて編集します。下記を追加しておくと良いでしょう：

    ```bash
    # .gitignoreファイルを編集
    vim .gitignore
    ```

    ```text
    # 環境変数ファイル
    .env
    .env.local
    .envrc

    # Node.js
    node_modules/
    .next/
    out/

    # 一時ファイル
    tmp/

    # ビルド生成物
    build/
    dist/
    ```

5. 変更をコミットしてプッシュします。

    ```bash
    git add .
    git commit -m "Initial project setup"
    git push origin main
    ```

### 1.4.4. ブランチ戦略とprotected branchの設定

プロジェクトに適切なブランチ戦略を設定します。一般的なGitflowやGitHub Flowをベースにしたシンプルなブランチ戦略を使用します。

1. GitHubのWebインターフェースで、リポジトリの「Settings」 > 「Branches」 > 「Add protection rule」を開きます。

2. 以下の設定を行います：
   - Branch name pattern: `main`
   - 「Require a pull request before merging」をチェック
   - 「Require approvals」をチェックし、必要な承認数を「1」に設定
   - 「Include administrators」をチェック（オプション）
   - 「Save changes」をクリック

3. ローカルで開発用のブランチを作成します。

```bash
# 機能開発用のブランチを作成
git checkout -b feature/initial-localstack-setup
```

### 1.4.5. go-taskによるタスクランナーの設定

go-taskは、Makeやnpm scriptsに似たタスクランナーで、YAML形式でタスクを定義します。まず、go-taskがインストールされていない場合はインストールします。

```bash
# macOS
brew install go-task/tap/go-task

# Linux
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin

# Windows (Scoop)
scoop install task
```

次に、プロジェクトのルートにTaskfile.ymlを作成します：

```bash
touch Taskfile.yml
```

Taskfile.ymlに以下の内容を記述します：

```yaml
version: '3'

vars:
  BACKEND_DIR: backend
  FRONTEND_CUSTOMER_DIR: frontend-customer
  FRONTEND_ADMIN_DIR: frontend-admin

tasks:
  start:
    desc: Start all services with docker-compose
    cmds:
      - docker-compose up -d

  stop:
    desc: Stop all services
    cmds:
      - docker-compose down

  restart:
    desc: Restart all services
    cmds:
      - task: stop
      - task: start

  logs:
    desc: Follow docker logs
    cmds:
      - docker-compose logs -f {{.CLI_ARGS | default ""}}

  backend:lint:
    desc: Run linters on backend code
    dir: '{{.BACKEND_DIR}}'
    cmds:
      - golangci-lint run ./...

  backend:test:
    desc: Run backend tests
    dir: '{{.BACKEND_DIR}}'
    cmds:
      - go test ./... -v

  frontend:install:
    desc: Install frontend dependencies
    cmds:
      - cd {{.FRONTEND_CUSTOMER_DIR}} && npm install
      - cd {{.FRONTEND_ADMIN_DIR}} && npm install

  frontend:lint:
    desc: Run linter on frontend code
    cmds:
      - cd {{.FRONTEND_CUSTOMER_DIR}} && npm run lint
      - cd {{.FRONTEND_ADMIN_DIR}} && npm run lint

  aws:localstack:status:
    desc: Check LocalStack status
    cmds:
      - aws --endpoint-url=http://localhost:4566 --profile localstack s3 ls
      - echo "LocalStack is running"
```

タスクが正しく設定されているかテストします：

```bash
# タスク一覧を表示
task --list

# LocalStackのステータス確認
task aws:localstack:status
```

### 1.4.6. GitHub Actionsを使った基本ムCIワークフローの設定

プロジェクトの品質を確保するための基本的なCIワークフローをGitHub Actionsで設定します。

```bash
# .github/workflowsディレクトリを作成
mkdir -p .github/workflows

# CIワークフローファイルを作成
touch .github/workflows/ci.yml
```

.github/workflows/ci.ymlに以下の内容を記述します：

```yaml
name: CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  backend-test:
    name: Backend Tests
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Go
      uses: actions/setup-go@v4
      with:
        go-version: '1.20'

    - name: Go modules cache
      uses: actions/cache@v3
      with:
        path: ~/go/pkg/mod
        key: ${{ runner.os }}-go-${{ hashFiles('**/go.sum') }}
        restore-keys: ${{ runner.os }}-go-

    - name: Go to backend directory
      run: cd backend

    - name: Test
      run: |
        cd backend
        go test -v ./...

  frontend-test:
    name: Frontend Tests
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3

    - name: Set up Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18'

    - name: NPM cache (customer)
      uses: actions/cache@v3
      with:
        path: frontend-customer/node_modules
        key: ${{ runner.os }}-npm-customer-${{ hashFiles('frontend-customer/package-lock.json') }}
        restore-keys: ${{ runner.os }}-npm-customer-

    - name: NPM cache (admin)
      uses: actions/cache@v3
      with:
        path: frontend-admin/node_modules
        key: ${{ runner.os }}-npm-admin-${{ hashFiles('frontend-admin/package-lock.json') }}
        restore-keys: ${{ runner.os }}-npm-admin-

    - name: Install dependencies (customer)
      run: |
        cd frontend-customer
        npm ci

    - name: Install dependencies (admin)
      run: |
        cd frontend-admin
        npm ci

    - name: Lint (customer)
      run: |
        cd frontend-customer
        npm run lint

    - name: Lint (admin)
      run: |
        cd frontend-admin
        npm run lint
```

設定が完了したら、変更をコミットしてプッシュします：

```bash
git add .
git commit -m "Add LocalStack configuration, go-task setup, and GitHub Actions CI"
git push origin feature/initial-localstack-setup
```

ブランチがプッシュされたら、GitHubのWebインターフェイスでプルリクエストを作成します。ブランチをmainブランチにマージする前に、CIワークフローが正常に実行されることを確認します。

## 1.5. 【確認ポイント】

実装が正しく完了したかどうか、以下のポイントを確認しましょう：

### 1.5.1. LocalStackの確認

- [ ] Docker ComposeでLocalStackサービスが正常に起動している

  ```bash
  docker ps | grep localstack
  ```

- [ ] S3バケットが正しく作成されている

  ```bash
  aws --endpoint-url=http://localhost:4566 s3 ls
  ```

- [ ] テストファイルをS3にアップロードできる

  ```bash
  aws --endpoint-url=http://localhost:4566 s3 cp tmp/test-image.jpg s3://product-images/
  aws --endpoint-url=http://localhost:4566 s3 ls s3://product-images/
  ```

- [ ] CloudWatch Logsのロググループが作成されている

  ```bash
  aws --endpoint-url=http://localhost:4566 logs describe-log-groups
  ```

### 1.5.2. GitHubリポジトリの確認

- [ ] GitHubリポジトリが正しく作成され、プロジェクトファイルがプッシュされている

  ```bash
  git remote -v
  git status
  ```

- [ ] .gitignoreが適切に設定されている

  ```bash
  cat .gitignore
  ```

- [ ] ブランチ保護ルールがGitHub上で設定されている（Settings > Branchesで確認）

- [ ] 開発用のブランチが作成されている

  ```bash
  git branch
  ```

### 1.5.3. go-taskとGitHub Actionsの確認

- [ ] go-taskがインストールされ、タスクが正しく実行できる

  ```bash
  task --version
  task --list
  ```

- [ ] Taskfile.ymlが正しく設定されている

  ```bash
  cat Taskfile.yml
  ```

- [ ] GitHub Actionsのワークフローファイルが正しく作成されている

  ```bash
  cat .github/workflows/ci.yml
  ```

- [ ] プルリクエスト作成後にCIワークフローが正常に実行される（GitHub上で確認）

## 1.6. 【詳細解説】

本日の実装に関連する技術と概念について詳しく解説します。

### 1.6.1. LocalStackとは

LocalStackは、AWSクラウドサービスをローカル環境でエミュレートするツールです。以下の特徴があります：

1. **開発コストの削減**
   - AWSの実環境を使わずにローカルでテストできるため、コストがかかりません
   - CI/CDパイプラインでも利用でき、テストコストを削減できます

2. **単一エンドポイント**
   - デフォルトではポート4566ですべてのAWSサービスをエミュレート
   - 同じエンドポイントで複数のAWSサービスを模擬するため、簡単に統合テストが行えます

3. **大量のAWSサービスをサポート**
   - S3, Lambda, DynamoDB, Kinesis, SQS, SNS, Cloudwatchなど、数十種類のサービスをサポート
   - 次第にサポートされるサービスが増えています

4. **オフライン開発**
   - インターネット接続がなくても開発とテストが可能
   - 移動中やインターネット接続が制限された環境でも利用可能

5. **AWS CLI互換性**
   - AWS CLIのコマンドをそのまま使用でき、`--endpoint-url`オプションで接続先を切り替えるだけ
   - AWS SDKも同様にエンドポイントを切り替えるだけで使用可能

本プロジェクトでは、オブザーバビリティに関連するAWSサービス（CloudWatch Logs, Metrics, X-Ray）やサーバーレス機能（Lambda, S3）をローカルでテストするためにLocalStackを活用します。

### 1.6.2. go-taskの利点と使い方

go-taskはタスクランナーで、以下のような利点があります：

1. **シンプルなYAML形式**
   - Makefileよりも直感的で読みやすいYAML形式でタスクを定義
   - タブとスペースの困ったMakefileの問題がない

2. **クロスプラットフォーム対応**
   - Windows, macOS, Linuxで同じように動作
   - プラットフォーム固有のシェルコマンドを抽象化できる

3. **タスク間の依存関係管理**
   - タスク間の依存関係を定義でき、順序性のある処理を設計可能
   - 依存タスクの実行状態をキャッシュし、必要なときのみ再実行

4. **変数とテンプレート**
   - 変数を定義してタスク間で再利用可能
   - Goテンプレート構文を使用した柔軟なタスク定義

5. **コマンドライン引数のサポート**
   - CLI引数をタスクに渡すことが可能
   - 柔軟なタスクの実行が可能

主な使い方：

```bash
# タスク一覧表示
task --list

# 特定のタスクを実行
task backend:test

# 引数付きでタスクを実行
task logs -- localstack
```

### 1.6.3. GitHubのブランチ保護と開発ワークフロー

GitHubのブランチ保護機能と開発ワークフローについて解説します：

1. **ブランチ保護ルールの目的**
   - mainブランチの保護：直接プッシュを禁止し、プルリクエストを弾する
   - コードレビューの強制：マージ前に指定した数の承認を必要とする
   - テストの強制：ステータスチェック（CIなど）が通ったものだけマージ可能

2. **効果的な開発ワークフロー**
   - 機能別ブランチ（feature/*, bugfix/*など）で開発
   - プルリクエストでコードレビューと自動テスト
   - 承認後にマージ
   - イシュートラッキングとの連携

3. **GitHub FlowとGitflowの比較**

   **GitHub Flow (シンプルなフロー)**
   - mainブランチは常にデプロイ可能な状態
   - 機能開発はブランチを作成して実施
   - プルリクエスト、レビュー、マージのシンプルなフロー
   - CI/CDとの相性が良い

   **Gitflow (より複雑なフロー)**
   - main, develop, feature/*, release/*, hotfix/*など複数の標準ブランチ
   - リリース管理が複雑なプロジェクトに適している
   - 大規模チームや複数バージョンの管理に向いている

本プロジェクトでは、学習の容易さとシンプルさを重視し、GitHub Flowに近いフローを採用しています。

## 1.7. 【補足情報】

実装において役立つ補足情報を紹介します。

### 1.7.1. AWS CLIプロファイルの活用

AWS CLIではプロファイル機能を使って、異なるAWS環境や設定を簡単に切り替えることができます。LocalStackと実際AWS環境の切り替えもプロファイルで簡単にできます。

**~/.aws/credentials**

```text
[default]
aws_access_key_id = AKIAIOSFODNN7EXAMPLE
aws_secret_access_key = wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY

[localstack]
aws_access_key_id = test
aws_secret_access_key = test
```

**~/.aws/config**

```text
[default]
region = ap-northeast-1
output = json

[profile localstack]
region = ap-northeast-1
output = json
endpoint_url = http://localhost:4566
```

これらのプロファイルを使うには、`--profile`オプションを指定します：

```bash
# LocalStackを使う場合
aws --profile localstack s3 ls

# 実際のAWSを使う場合
aws --profile default s3 ls

# またはプロファイルを環境変数で指定することも可能
export AWS_PROFILE=localstack
aws s3 ls  # これでlocalstackプロファイルが使われる
```

### 1.7.2. GitHub Actionsの発展的な使い方

GitHub Actionsは単純なCI/CD以上のことができます。いくつかの発展的な使い方を紹介します：

1. **環境別のデプロイメント**

    ```yaml
    jobs:
      deploy:
        name: Deploy
        runs-on: ubuntu-latest
        environment: ${{ github.ref == 'refs/heads/main' && 'production' || 'staging' }}
        steps:
          # デプロイ手順
    ```

2. **環境変数とシークレット**

    GitHubリポジトリの Settings > Secrets and variables > Actions で環境変数やシークレットを設定し、ワークフロー内で利用できます。

    ```yaml
    steps:
      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ap-northeast-1
    ```

3. **マトリックスビルド**

    複数のバージョンや環境でテストを実行する場合：

    ```yaml
    jobs:
      test:
        strategy:
          matrix:
            go-version: [1.19.x, 1.20.x]
            os: [ubuntu-latest, macos-latest, windows-latest]
        runs-on: ${{ matrix.os }}
        steps:
          - uses: actions/setup-go@v4
            with:
              go-version: ${{ matrix.go-version }}
          # テスト手順
    ```

4. **ワークフローの再利用**

    再利用可能なワークフローを作成し、複数リポジトリで利用できます。

    ```yaml
    # .github/workflows/reusable-workflow.yml
    name: Reusable workflow
    on:
      workflow_call:
        inputs:
          environment:
            type: string
            required: true

    jobs:
      reusable_job:
        runs-on: ubuntu-latest
        environment: ${{ inputs.environment }}
        steps:
          # 手順
    ```

    別のワークフローから呼び出す：

    ```yaml
    jobs:
      call_reusable:
        uses: ./.github/workflows/reusable-workflow.yml
        with:
          environment: production
    ```

## 1.8. 【よくある問題と解決法】

### 1.8.1. 問題1: LocalStackのサービスに接続できない

**症状**: AWS CLIでLocalStackのサービスに接続しようとすると、「connection refused」や「timeout」エラーが発生する。

**解決策**:

1. **LocalStackコンテナの状態確認**

   ```bash
   docker ps | grep localstack
   ```

   コンテナが実行中であることを確認します。

2. **ポートマッピングの確認**

   ```bash
   docker port localstack
   ```

   4566ポートが正しくマップされているか確認します。

3. **エンドポイントURLの確認**

   ```bash
   aws --endpoint-url=http://localhost:4566 s3 ls
   ```

   --endpoint-urlパラメータを正しく指定しているか確認します。

4. **LocalStackコンテナの再起動**

   ```bash
   docker restart localstack
   ```

   必要なサービスが正しく起動するまで待ちます。

5. **ログの確認**

   ```bash
   docker logs localstack
   ```

   サービスの起動エラーや後ログを確認します。

### 1.8.2. 問題2: GitHub Actionsワークフローが失敗する

**症状**: GitHub Actionsのワークフローが失敗して、プルリクエストがマージできない。

**解決策**:

1. **ワークフローのログ確認**
   GitHubのActionsタブから失敗したワークフローをクリックし、詳細なログを確認します。

2. **ワークフローファイルの有効性確認**

   ```bash
   cat .github/workflows/ci.yml
   ```

   YAML文法が正しいか、ジョブの定義が正確か確認します。

3. **テストのローカル実行**

   ```bash
   cd backend
   go test ./... -v
   ```

   ローカルでテストが成功するか確認し、失敗する場合はその原因を修正します。

4. **環境変数やシークレットの確認**
   ワークフローが必要とする環境変数やシークレットが適切に設定されているか確認します。

5. **ディレクトリ構造の確認**
   ワークフローのステップで実行しているディレクトリパスが実際のリポジトリ構造と一致しているか確認します。

## 1.9. 【今日の重要なポイント】

本日の実装で特に重要なポイントは以下の通りです：

1. **LocalStackとAWSサービスのローカルエミュレーション**

   LocalStackは、AWSのサービスをローカルでエミュレートすることで、コストをかけずに開発とテストができます。本プロジェクトのオブザーバビリティ学習において、S3、CloudWatch Logs、Lambdaなどのサービスを利用することが不可欠であり、LocalStackはこれらを手軽に使える環境を提供します。

2. **プロジェクトのバージョン管理とワークフロー**

   GitHubでの標準的な開発ワークフローを確立することで、チーム開発と自己学習の両方で高品質なコードを維持できます。ブランチ保護やPRベースの開発は実務でも一般的な手法で、自己学習でもこのプラクティスを組み込むことで実践的なスキルを進化させることができます。

3. **統合タスクランナーによる開発効率化**

   go-taskのようなタスクランナーを導入することで、繰り返し実行するコマンドを簡略化し、チーム全体で共通の流れを確立できます。様々な環境（Mac、Linux、Windows）でも同じコマンドで作業できるようになり、開発の効率化に大きく貢献します。

4. **自動テストと品質管理の基盤構築**

   GitHub Actionsを使ったCIワークフローは、自動テストやコード品質チェックを自動化し、継続的にPRごとにコードの品質を確保します。オブザーバビリティ学習のプロジェクトでも、この自動化は重要で、初期から適切なテスト力を確保することで、後のフェーズでの実装が保護されます。

5. **フェーズ1の基盤完成**

   本日でプロジェクトの基盤となる開発環境、リポジトリ設定、ワークフロー・自動化、そしてLocalStackによるAWSエミュレーション環境が整いました。これはフェーズ1の基盤道具が揃ったことを意味し、次週からはその上にデータモデル設計やAPI実装を進めていく準備が整いました。

これらのポイントは次回以降の実装でも總勝利に活用されます。

## 1.10. 【次回の準備】

次回（Week 2 - Day 1）では、データモデルと基本的なAPI設計を行います。以下の点について事前に確認しておくと良いでしょう：

1. **MySQLの基本的な知識**
   - 基本的なSQLクエリ（CREATE TABLE, SELECT, INSERTなど）
   - データ型とフィールドおよび制約（PRIMARY KEY, FOREIGN KEYなど）
   - テーブル間のリレーションシップ（1:1、1:N、N:N）

2. **OpenAPI仕様の基本的な知識**
   - YAML形式とJSON形式の理解
   - APIエンドポイント、パラメータ、レスポンスの定義方法
   - Swagger UIの基本的な使い方

3. **Goの基本的な知識**
   - Goの基本構文と型システム
   - モジュールとパッケージの管理
   - 構造体とメソッドの定義

4. **sqlboilerの概要把握**
   - ORMの基本的なコンセプト
   - sqlboilerと他のORMの違い（コード生成アプローチ）
   - モデル生成と使用方法

5. **以下のツールのインストール確認**
   - golang-migrate（データベースマイグレーションツール）
   - ogen（OpenAPIコードジェネレータ）

前回までの目的は環境の整備であったのに対し、次回からは実際のデータ設計とAPI設計に入ります。基本的なER図やデータモデル設計の知識があると理解がスムーズに進みます。

## 1.11. 【.envrc サンプル】

以下は本プロジェクトで使用する.envrcのサンプルです。direnvがインストールされている場合、プロジェクトディレクトリに移動したときに自動的に環境変数がロードされます。このファイルはGitにコミットしないようにしてください。

```bash
# プロジェクト固有の環境変数

# データベース設定
export DB_HOST="localhost"
export DB_PORT="3306"
export DB_USER="app"
export DB_PASSWORD="password"
export DB_NAME="ecommerce"
export DB_SSL_MODE="false"

# AWS設定（LocalStack用）
export AWS_REGION="ap-northeast-1"
export AWS_ACCESS_KEY_ID="test"
export AWS_SECRET_ACCESS_KEY="test"
export AWS_ENDPOINT_URL="http://localhost:4566"

# アプリケーション設定
export APP_ENV="development"
export API_BASE_URL="http://localhost:8080"
export LOG_LEVEL="debug"

# PATHの設定
# ローカルのツールを使用する場合
export PATH="$PWD/tools/bin:$PATH"
```

これらの環境変数は、アプリケーションの設定やLocalStackへの接続など、開発環境を制御するために使用されます。direnvを使用しない場合は、シェルで直接ロードするか、Docker環境内で環境変数として設定してください。
