# 1. 週1: プロジェクト基盤と基本API実装

このドキュメントでは、AWSオブザーバビリティ学習用eコマースアプリケーションの開発環境セットアップと基本構造の構築について詳細に説明します。Go Echo+sqlboiler+Next.jsのスタックと、LocalStackによるAWSサービスのエミュレーションを活用します。

## 1.1. 目次

- [1. 週1: プロジェクト基盤と基本API実装](#1-週1-プロジェクト基盤と基本api実装)
  - [1.1. 目次](#11-目次)
  - [1.2. はじめに](#12-はじめに)
    - [1.2.1. 今週の学習に必要な前提条件](#121-今週の学習に必要な前提条件)
    - [1.2.2. 今週の目標](#122-今週の目標)
    - [1.2.3. 今週のユースケース](#123-今週のユースケース)
  - [1.3. go-taskを使用したプロジェクト操作](#13-go-taskを使用したプロジェクト操作)
    - [1.3.1. 主なタスク](#131-主なタスク)
    - [1.3.2. タスクの使用例](#132-タスクの使用例)
  - [1.4. AWSとクラウドネイティブアプリケーションの概要](#14-awsとクラウドネイティブアプリケーションの概要)
    - [1.4.1. クラウドネイティブアプリケーションとは](#141-クラウドネイティブアプリケーションとは)
    - [1.4.2. AWS環境とLocalStack](#142-aws環境とlocalstack)
  - [1.5. プロジェクト初期設定](#15-プロジェクト初期設定)
    - [1.5.1. ディレクトリ構造作成](#151-ディレクトリ構造作成)
    - [1.5.2. GitHubリポジトリのセットアップ](#152-githubリポジトリのセットアップ)
      - [1.5.2.1. リポジトリの作成](#1521-リポジトリの作成)
      - [1.5.2.2. リポジトリのクローンと初期プッシュ](#1522-リポジトリのクローンと初期プッシュ)
      - [1.5.2.3. ブランチ保護ルールの設定](#1523-ブランチ保護ルールの設定)
      - [1.5.2.4. ブランチ戦略の設定](#1524-ブランチ戦略の設定)
      - [1.5.2.5. .gitignoreの設定](#1525-gitignoreの設定)
  - [1.6. 開発環境のセットアップ](#16-開発環境のセットアップ)
    - [1.6.1. Go/Echoバックエンド環境のセットアップ](#161-goechoバックエンド環境のセットアップ)
      - [1.6.1.1. バックエンドのプロジェクト構造](#1611-バックエンドのプロジェクト構造)
      - [1.6.1.2. Goプロジェクトの初期化](#1612-goプロジェクトの初期化)
      - [1.6.1.3. プロジェクト構造の作成](#1613-プロジェクト構造の作成)
      - [1.6.1.4. 必要なパッケージのインストール](#1614-必要なパッケージのインストール)
      - [1.6.1.5. direnvの設定](#1615-direnvの設定)
      - [1.6.1.6. Dockerfileの作成](#1616-dockerfileの作成)
      - [1.6.1.7. airの設定ファイル作成](#1617-airの設定ファイル作成)
    - [1.6.2. フロントエンド環境のセットアップ](#162-フロントエンド環境のセットアップ)
      - [1.6.2.1. フロントエンドプロジェクト構造](#1621-フロントエンドプロジェクト構造)
      - [1.6.2.2. 顧客向けフロントエンド](#1622-顧客向けフロントエンド)
      - [1.6.2.3. 管理者向けフロントエンド](#1623-管理者向けフロントエンド)
    - [1.6.3. Docker Compose環境の構築](#163-docker-compose環境の構築)
      - [1.6.3.1. Traefikの導入について](#1631-traefikの導入について)
      - [1.6.3.2. Traefikの基本設定](#1632-traefikの基本設定)
      - [1.6.3.3. ディレクトリ構造とTraefik設定ファイル](#1633-ディレクトリ構造とtraefik設定ファイル)
      - [1.6.3.4. ALBとのハイブリッド構成](#1634-albとのハイブリッド構成)
      - [1.6.3.5. ホスト名の解決と設定](#1635-ホスト名の解決と設定)
  - [1.7. AWS環境の構築](#17-aws環境の構築)
    - [1.7.1. LocalStackのセットアップ](#171-localstackのセットアップ)
    - [1.7.2. AWS CLI設定](#172-aws-cli設定)
    - [1.7.3. Docker Compose環境の起動とLocalStackへの接続テスト](#173-docker-compose環境の起動とlocalstackへの接続テスト)
  - [1.8. バックエンド実装](#18-バックエンド実装)
    - [1.8.1. 設定管理モジュールの実装](#181-設定管理モジュールの実装)
    - [1.8.2. APIハンドラーの実装](#182-apiハンドラーの実装)
      - [1.8.2.1. ルーターの実装](#1821-ルーターの実装)
      - [1.8.2.2. ヘルスチェックハンドラー](#1822-ヘルスチェックハンドラー)
      - [1.8.2.3. 商品一覧ハンドラー](#1823-商品一覧ハンドラー)
    - [1.8.3. メインアプリケーションの実装](#183-メインアプリケーションの実装)
    - [1.8.4. ハンドラーのユニットテスト](#184-ハンドラーのユニットテスト)
  - [1.9. フロントエンド実装](#19-フロントエンド実装)
    - [1.9.1. 商品一覧APIクライアントの実装](#191-商品一覧apiクライアントの実装)
    - [1.9.2. 商品一覧ページの実装](#192-商品一覧ページの実装)
    - [1.9.3. フロントエンドのテスト実装](#193-フロントエンドのテスト実装)
    - [1.9.4. 管理者向けフロントエンドの実装](#194-管理者向けフロントエンドの実装)
  - [1.10. バックエンドとフロントエンドの連携動作確認](#110-バックエンドとフロントエンドの連携動作確認)
    - [1.10.1. バックエンドの起動と動作確認](#1101-バックエンドの起動と動作確認)
    - [1.10.2. フロントエンドの動作確認](#1102-フロントエンドの動作確認)
      - [1.10.2.1. 動作確認ポイント](#11021-動作確認ポイント)
      - [1.10.2.2. 想定される問題と対処法](#11022-想定される問題と対処法)
      - [1.10.2.3. 次のステップ](#11023-次のステップ)
  - [1.11. 理解度チェックと発展課題](#111-理解度チェックと発展課題)
    - [1.11.1. 理解度チェック](#1111-理解度チェック)
    - [1.11.2. 発展課題](#1112-発展課題)
  - [1.12. 次週の予告](#112-次週の予告)
  - [1.13. 参考資料](#113-参考資料)
    - [1.13.1. Go言語とEcho](#1131-go言語とecho)
    - [1.13.2. Next.js とフロントエンド](#1132-nextjs-とフロントエンド)
    - [1.13.3. AWS関連](#1133-aws関連)
    - [1.13.4. Docker と開発環境](#1134-docker-と開発環境)

## 1.2. はじめに

### 1.2.1. 今週の学習に必要な前提条件

この講義では、以下のソフトウェアバージョンを使用します：

- Docker と Docker Compose
- Go (1.24以上)
- Node.js (23以上) と npm/yarn
- Git
- Visual Studio CodeやGoLandなどのIDE/エディタ
- AWS CLI
- go-task

まだインストールしていない場合は、以下のコマンドでインストールしてください：

```bash
# Go (公式サイトからダウンロードしてインストール)
# https://golang.org/dl/

# Node.js (nvm経由でインストールする場合)
nvm install 23.9.0
nvm use 23.9.0

# go-task (macOSの場合)
brew install go-task/tap/go-task

# go-task (Linuxの場合)
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b ~/.local/bin

# go-task (Windowsの場合)
choco install go-task
# または
scoop install task
```

### 1.2.2. 今週の目標

この週では、AWSオブザーバビリティ学習のための基盤となる開発環境の構築と、基本的なAPIエンドポイントの実装を行います。具体的には以下のような目標を達成します：

1. 開発環境のセットアップ
   - GitHubリポジトリの作成と構成
   - Docker Compose環境の構築
   - Go/Echo開発環境の準備
   - Next.js/TypeScript/TailwindCSSフロントエンド環境の準備

2. AWS環境の基本設定
   - LocalStackによるAWSエミュレーション環境の準備
   - AWS CLIの設定

3. 基本APIエンドポイントの実装
   - ヘルスチェックAPI
   - モック商品一覧API

4. テスト実装
   - バックエンドユニットテスト
   - フロントエンドコンポーネントテスト

### 1.2.3. 今週のユースケース

実装を通じて、以下のユースケースを実現することを目指します：

- **OBS-07**: ヘルスチェック（ヘルスチェックエンドポイントが機能し、状態が報告される）
- **CUS-01**: 基本的な商品閲覧（モックAPIからの商品データ取得ができる）
- **DEV-01**: テスト自動化（基本的なユニットテストとコンポーネントテストが実装される）

## 1.3. go-taskを使用したプロジェクト操作

プロジェクトの操作を簡単にするために、go-task（<https://taskfile.dev/）を導入します。go-taskはシンプルなタスクランナーで、YAMLベースの設定ファイルでプロジェクトタスクを管理できます。>

### 1.3.1. 主なタスク

プロジェクトのルートディレクトリに`Taskfile.yml`を作成し、以下のようなタスクを定義します：

```bash
touch Taskfile.yml
```

- `task start`: 環境を起動します
- `task stop`: 環境を停止します
- `task logs`: コンテナのログを表示します
- `task test:backend`: バックエンドのテストを実行します
- `task test:frontend-customer`: フロントエンドのテストを実行します
- `task shell:backend`: バックエンドのシェルを起動します
- `task localstack:check`: LocalStackの状態をチェックします

利用可能なすべてのタスクを確認するには:

```bash
task --list
```

### 1.3.2. タスクの使用例

環境のセットアップと起動:

```bash
task setup
task start
```

テストの実行:

```bash
task test:backend
task test:frontend-customer
```

ログの確認:

```bash
task logs:backend
task logs:frontend-customer
```

## 1.4. AWSとクラウドネイティブアプリケーションの概要

### 1.4.1. クラウドネイティブアプリケーションとは

クラウドネイティブアプリケーションとは、クラウド環境を最大限に活用するために設計されたアプリケーションです。主に以下の特徴を持ちます：

1. **マイクロサービスアーキテクチャ**: 機能ごとに小さなサービスに分割
2. **コンテナ化**: Docker等のコンテナ技術を使用
3. **DevOpsの実践**: CI/CDパイプラインによる継続的なデリバリー
4. **インフラのコード化**: Infrastructure as Code（IaC）の実践
5. **自動スケーリング**: 負荷に応じて自動的にリソースをスケール

このプロジェクトでは、これらの原則に従いながら、AWSのサービスを活用したEコマースアプリケーションを構築していきます。

### 1.4.2. AWS環境とLocalStack

AWSは多様なクラウドサービスを提供していますが、開発環境では実際のAWSサービスを使用するとコストがかかり、開発効率も低下する可能性があります。LocalStackは、開発やテスト時にAWSサービスをローカル環境でエミュレートするツールです。

LocalStackの主なメリットは以下の通りです：

1. **コスト削減**: 実際のAWSを使わないため、課金されない
2. **高速な開発サイクル**: ネットワーク遅延なしで高速に開発
3. **オフライン開発**: インターネット接続なしでも開発可能
4. **環境の独立性**: 開発者ごとに独立した環境を提供

今回のプロジェクトでは、まずLocalStackを開発環境に構築しておきます。第1週ではまだ活用しませんが、次週以降のAWSサービス連携の準備として導入します。

## 1.5. プロジェクト初期設定

### 1.5.1. ディレクトリ構造作成

まず、プロジェクト全体のディレクトリ構造を作成します:

```bash
# プロジェクトのルートディレクトリを作成
mkdir -p aws-observability-ecommerce
cd aws-observability-ecommerce

# バックエンド、フロントエンド、インフラ用のディレクトリを作成
mkdir -p backend frontend-customer frontend-admin infra
mkdir -p infra/{localstack,terraform}
mkdir -p docs
```

このディレクトリ構造は、マイクロサービス的なアプローチを取りつつも、開発の初期段階ではモノレポ（単一リポジトリ）として管理します。これにより、開発初期のコード共有やデプロイが容易になります。

### 1.5.2. GitHubリポジトリのセットアップ

#### 1.5.2.1. リポジトリの作成

1. GitHubにログインし、「New repository」をクリックします。
2. リポジトリ名（例: `aws-observability-ecommerce`）を入力します。
3. 簡単な説明を追加します。例: 「AWSオブザーバビリティ学習用eコマースアプリケーション」
4. 必要に応じて「Private」を選択して非公開リポジトリにします。
5. 「Initialize this repository with a README」にチェックを入れます。
6. 「.gitignore template」で「Go」を選択します。
7. 「Create repository」をクリックしてリポジトリを作成します。

#### 1.5.2.2. リポジトリのクローンと初期プッシュ

作成したリポジトリをローカルにクローンし、既存のプロジェクトファイルをプッシュします：

```bash
git init
git commit --allow-empty -m "first commit"
touch README.md
git add .
git commit -m "Initial commit with project structure"
git remote add origin git@github.com:y-nosuke/aws-observability-ecommerce.git
git push -u origin main
```

#### 1.5.2.3. ブランチ保護ルールの設定

メインブランチを保護し、プルリクエストとコードレビューを強制します：

1. GitHubリポジトリの「Settings」タブをクリックします。
2. 左サイドバーの「Branches」をクリックします。
3. 「Branch protection rules」セクションで「Add rule」をクリックします。
4. 「Branch name pattern」に「main」と入力します。
5. 以下のオプションを有効にします：
   - 「Require a pull request before merging」
   - 「Require approvals」（最低1人のレビュー）
   - 「Require status checks to pass before merging」
   - 「Require branches to be up to date before merging」
6. 「Create」をクリックして保護ルールを保存します。

#### 1.5.2.4. ブランチ戦略の設定

プロジェクトで使用するシンプルなブランチ戦略を設定します：

```bash
# mainブランチから直接featureブランチを作成
git checkout -b feature/initial-setup
```

実際の開発では、機能ごとに `feature/機能名` のようなブランチを作成し、実装完了後にプルリクエストを通じて `main` ブランチにマージする流れになります。

#### 1.5.2.5. .gitignoreの設定

.gitignoreファイルに必要な除外パターンを追加します。すでにGoのテンプレートが適用されていますが、フロントエンドとLocalStack関連の除外を追加します：

## 1.6. 開発環境のセットアップ

### 1.6.1. Go/Echoバックエンド環境のセットアップ

#### 1.6.1.1. バックエンドのプロジェクト構造

バックエンド（Go/Echo）のプロジェクト構造を以下のように設定します：

```text
backend/
├── cmd/                # アプリケーションのエントリーポイントを含むディレクトリ
│   └── api/            # APIサーバーの実装
│       └── main.go     # メインアプリケーション
├── internal/           # 外部からインポートされないパッケージ
│   ├── api/            # APIに関連するコード
│   │   ├── handlers/   # リクエストハンドラー
│   │   ├── middleware/ # ミドルウェア実装
│   │   └── router/     # ルーター定義
│   ├── config/         # 設定管理
│   └── models/         # データモデル
├── pkg/                # 外部からインポート可能なライブラリコード
├── tests/              # テストコード
│   ├── unit/           # ユニットテスト
│   └── integration/    # 統合テスト
├── go.mod              # Goモジュール定義
├── go.sum              # Goモジュール依存関係
└── Dockerfile.dev      # 開発用Docker定義
```

このようなディレクトリ構造は、Goのプロジェクトで一般的に使用される構造であり、関心事の分離と責任の明確化を促進します。

#### 1.6.1.2. Goプロジェクトの初期化

バックエンドのGoプロジェクトを初期化します：

```bash
cd backend
go mod init github.com/y-nosuke/aws-observability-ecommerce
```

#### 1.6.1.3. プロジェクト構造の作成

バックエンドのディレクトリ構造を作成します：

```bash
mkdir -p cmd/api
mkdir -p internal/{api,config,models}
mkdir -p internal/api/{handlers,middleware,router}
mkdir -p pkg
mkdir -p tests/{unit,integration}
```

#### 1.6.1.4. 必要なパッケージのインストール

必要なGo言語のパッケージをインストールします：

```bash
# Echo Webフレームワーク
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware

# 設定管理
go get github.com/spf13/viper

# テスト関連
go get github.com/stretchr/testify
go get go.uber.org/mock

# その他のユーティリティ
go get github.com/google/uuid
```

#### 1.6.1.5. direnvの設定

direnvを使って環境変数を管理します。これはローカル開発環境の設定を簡単に行うためのツールです。

```bash
touch .envrc
```

`.envrc`のサンプル内容：

```bash
# アプリケーション設定
export APP_NAME=aws-observability-ecommerce
export APP_VERSION=1.0.0
export APP_ENV=development
export PORT=8080
```

direnvを許可します（インストール済みの場合）：

```bash
direnv allow
```

#### 1.6.1.6. Dockerfileの作成

Go/Echoアプリケーション用のDockerfileを作成します：

```bash
touch Dockerfile.dev
```

#### 1.6.1.7. airの設定ファイル作成

ホットリロード用の設定ファイルを作成します：

```bash
touch .air.toml
```

この設定により、ファイルが変更されるたびに自動的にgoimportsでフォーマットしてからビルドが実行されます。

### 1.6.2. フロントエンド環境のセットアップ

#### 1.6.2.1. フロントエンドプロジェクト構造

フロントエンド（Next.js）のプロジェクト構造を以下のように設定します：

```text
frontend-customer/ (または frontend-admin/)
├── app/                # Next.js App Routerの実装
│   ├── layout.tsx      # 共通レイアウト
│   ├── page.tsx        # トップページ
│   └── products/       # 商品関連ページ
│       └── page.tsx    # 商品一覧ページ
├── components/         # 再利用可能なUIコンポーネント
│   ├── ui/             # 基本UIコンポーネント
│   └── products/       # 商品関連コンポーネント
├── lib/                # ユーティリティ関数とAPIクライアント
│   └── api/            # APIクライアント実装
├── public/             # 静的ファイル
├── styles/             # グローバルスタイル
├── __tests__/          # テストファイル
│   ├── components/     # コンポーネントテスト
│   └── pages/          # ページテスト
├── next.config.js      # Next.js設定
├── package.json        # npm依存関係
└── Dockerfile.dev      # 開発用Docker定義
```

この構造は、Next.jsのApp Routerを利用したモダンなReactアプリケーションの構造です。

#### 1.6.2.2. 顧客向けフロントエンド

Next.jsプロジェクトを作成します：

```bash
cd ../frontend-customer

# Next.jsプロジェクトの作成
npx create-next-app@latest . --typescript --eslint --tailwind --app
```

インタラクティブなセットアップで以下のように選択します:

```bash
✔ Would you like to use `src/` directory? … No
✔ Would you like to use Turbopack for next dev? … No
✔ Would you like to customize the default import alias? … No
```

direnvを設定します：

```bash
touch .envrc
```

`.envrc`のサンプル内容：

```bash
export NEXT_PUBLIC_API_URL=http://backend:8080/api
```

Dockerfile.devを作成します：

```bash
touch Dockerfile.dev
```

ESLintと関連プラグインをインストールします

```bash
# ESLintと関連プラグインのインストール
npm install --save-dev @typescript-eslint/eslint-plugin @typescript-eslint/parser eslint-config-prettier eslint-plugin-import eslint-plugin-jest eslint-plugin-jsx-a11y eslint-plugin-react eslint-plugin-react-hooks eslint-plugin-testing-library prettier

# テストライブラリのインストール
npm install --save-dev jest @testing-library/react @testing-library/jest-dom jest-environment-jsdom
# jest: JavaScriptテストフレームワーク
# @testing-library/react: Reactコンポーネントをテストするためのライブラリ
# @testing-library/jest-dom: JestでDOMに関するカスタムマッチャーを追加するライブラリ
npm install --save-dev \
  @types/jest \
  @types/react \
  @types/testing-library__react \
  ts-jest \
  typescript \
  @babel/core \
  @babel/preset-env \
  @babel/preset-react \
  @babel/preset-typescript

touch .eslintrc.js .eslintignore .prettierrc.js .prettierignore

# package.jsonのscriptsセクションを更新（手動で行う必要があります）
# 既存のpackage.jsonファイルを開いて、scriptsセクションに以下のコマンドを追加：
# "lint": "eslint --ext .js,.jsx,.ts,.tsx .",
# "lint:fix": "eslint --ext .js,.jsx,.ts,.tsx . --fix",
# "prettier:check": "prettier --check .",
# "prettier:write": "prettier --write .",
# "validate": "npm run lint && npm run prettier:check && npm run test"
```

#### 1.6.2.3. 管理者向けフロントエンド

管理者向けのNext.jsプロジェクトも顧客向けと同様にセットアップします：

```bash
cd ../frontend-admin
npx create-next-app@latest . --typescript --eslint --tailwind --app
```

direnvと必要なファイルの設定も同様に行います。

### 1.6.3. Docker Compose環境の構築

プロジェクトのルートディレクトリに戻り、Docker Compose設定ファイルを作成します：

```bash
cd ..
touch compose.yml
```

Docker Compose環境では、以下のサービスを定義します：

1. **traefik**: リバースプロキシ・ロードバランサー（ホスト名ベースのルーティング用）
2. **backend**: GoアプリケーションのAPIサーバー
3. **frontend-customer**: 顧客向けNext.jsアプリケーション
4. **frontend-admin**: 管理者向けNext.jsアプリケーション
5. **mysql**: データベースサーバー
6. **localstack**: AWSサービスエミュレーター

#### 1.6.3.1. Traefikの導入について

Traefik はモダンな Web リバースプロキシおよびロードバランサーで、Docker と連携して動的にサービスを検出し、リクエストをルーティングする機能を持っています。開発環境でTraefikを使用する主なメリットは次のとおりです：

1. **ホスト名ベースのルーティング**: ポート番号の代わりにホスト名でサービスにアクセスできます
   - `api.localhost`: バックエンドAPIアクセス用
   - `shop.localhost`: 顧客向けフロントエンド用
   - `admin.localhost`: 管理者向けフロントエンド用
   - `db.localhost`: PhpMyAdmin用
   - `traefik.localhost:8080`: Traefikダッシュボード用

2. **動的設定**: Docker のラベルを使って設定を行い、コンテナの起動・停止に応じて自動で設定が変更されます
   - 新しいサービスを追加した場合も、ラベルを付与するだけで自動的にルーティングが設定される
   - コンテナの再起動時も設定が自動的に再適用される

3. **ダッシュボード**: 可視化されたダッシュボードでルーティング状況を確認できます
   - 現在のルーティングルールとそのステータスをリアルタイムで確認
   - サービスの健全性とメトリクスの監視
   - HTTP/TCPエンドポイント、ミドルウェア、サービスの状態が一目でわかる

4. **セキュリティ機能**: 本番環境と同様のセキュリティ設定をテストできます（必要に応じて）
   - HTTPS対応（Let's Encryptとの統合）
   - レート制限
   - 基本認証
   - IPフィルタリング

#### 1.6.3.2. Traefikの基本設定

Traefikを使用するための基本的な設定は以下の通りです：

1. **Traefikサービスの定義**:
   - Traefikコンテナ自体のポート公開設定（80番と8080番）
   - Dockerソケットへのアクセス（動的設定を可能にするため）
   - 設定ファイルのマウント
   - 基本的なコマンドラインオプション

2. **各サービスへのTraefikラベル追加**:
   - `traefik.enable=true`: サービスをTraefikに登録する
   - `traefik.http.routers.[サービス名].rule=Host(\`[ホスト名]\`)`: ホスト名に基づくルーティングルール
   - `traefik.http.services.[サービス名].loadbalancer.server.port=[ポート]`: サービスのポート設定

3. **ネットワーク設定**:
   - 全サービスを同一のネットワークに配置し、コンテナ間通信を可能にする

#### 1.6.3.3. ディレクトリ構造とTraefik設定ファイル

Traefik用の設定ディレクトリと基本設定ファイルを作成します：

```bash
# Traefik設定ディレクトリの作成
mkdir -p infra/traefik

# 基本設定ファイルの作成
touch infra/traefik/traefik.yml
```

`infra/traefik/traefik.yml`には以下のような基本設定を記述します

#### 1.6.3.4. ALBとのハイブリッド構成

後の講義ではAWS ALB (Application Load Balancer)を学習する予定ですが、開発環境ではTraefikとLocalStack上のALBエミュレーションを組み合わせたハイブリッド構成を検討できます。この構成では：

1. **開発の利便性**: Traefikによりホスト名ベースのルーティングや動的設定などの利点を享受
   - 分かりやすいURLでのアクセス
   - ポート番号を覚える必要がない
   - サービス追加時の設定簡略化

2. **学習の効果**: LocalStack上のALBを使ってAWSの設定や動作を学習
   - ALBの設定パターン
   - ターゲットグループの概念
   - ヘルスチェックの設定
   - ルーティングルールの設定

3. **両方の利点**: 「開発の快適さ」と「AWS環境の学習」を両立
   - 開発中はTraefikによる直感的なアクセス
   - 検証時にはLocalStack上のALBを通したアクセスパスをテスト

具体的には、以下のようなアクセスパスを設定できます：

- **開発アクセス**: ユーザー → Traefik → アプリケーション (`api.localhost`, `shop.localhost`など)
- **ALB学習用**: ユーザー → Traefik → LocalStack ALB → アプリケーション

この構成により、開発者は覚えやすいURLでサービスにアクセスできると同時に、ALB設定のベストプラクティスも学べます。

#### 1.6.3.5. ホスト名の解決と設定

基本的に `.localhost` ドメインは自動的に `127.0.0.1` にリゾルブされますが、問題がある場合は `/etc/hosts` ファイルに以下の設定を追加します：

```text
127.0.0.1 api.localhost shop.localhost admin.localhost traefik.localhost
```

これにより、ブラウザや他のアプリケーションから上記のホスト名にアクセスできるようになります。

## 1.7. AWS環境の構築

### 1.7.1. LocalStackのセットアップ

LocalStackの初期化スクリプトを作成します：

```bash
cd infra/localstack
touch init.sh
chmod +x init.sh
```

### 1.7.2. AWS CLI設定

AWS CLIをインストールして、LocalStackに接続するための設定を行います：

```bash
# AWS CLIのインストール (既にインストール済みの場合は不要)
# macOSの場合
brew install awscli

# Linuxの場合
curl "https://awscli.amazonaws.com/awscli-exe-linux-x86_64.zip" -o "awscliv2.zip"
unzip awscliv2.zip
sudo ./aws/install

# Windows (PowerShell)の場合
# https://awscli.amazonaws.com/AWSCLIV2.msi をダウンロード後、インストール

# LocalStack用のプロファイル設定
aws configure --profile localstack
# AWS Access Key ID: test
# AWS Secret Access Key: test
# Default region name: ap-northeast-1
# Default output format: json
```

### 1.7.3. Docker Compose環境の起動とLocalStackへの接続テスト

プロジェクトのコンテナ環境を起動し、LocalStackへの接続テストを行います：

```bash
# プロジェクトルートディレクトリに移動
cd ../../

# Docker Compose環境を起動
task start

# LocalStackの接続テスト
task localstack:check
```

## 1.8. バックエンド実装

### 1.8.1. 設定管理モジュールの実装

設定を管理するモジュールを実装します：

```bash
cd backend
touch internal/config/config.go
```

`internal/config/config.go`に設定管理コードを実装します。この実装は以下の機能を提供します：

- 環境変数からアプリケーション設定の読み込み
- 設定ファイルからの設定読み込み (存在する場合)
- デフォルト値の設定
- 構造化された設定モデルへのマッピング

### 1.8.2. APIハンドラーの実装

#### 1.8.2.1. ルーターの実装

APIルートを設定するルーターを実装します：

```bash
touch internal/api/router/router.go
```

`internal/api/router/router.go`に以下の機能を持つルーター設定を実装します：

- APIグループの定義 (/api プレフィックス)
- ヘルスチェックエンドポイントのルート設定
- 商品関連エンドポイントのルート設定
- 必要に応じたURLパラメータのマッピング

#### 1.8.2.2. ヘルスチェックハンドラー

システムの健全性を確認するヘルスチェックAPIを実装します：

```bash
touch internal/api/handlers/health_handler.go
```

`internal/api/handlers/health_handler.go`に以下の機能を持つハンドラーを実装します：

- システムの基本状態情報の提供
- アプリケーションのバージョンと稼働時間の情報
- メモリ使用量などのリソース情報
- 各サービス（API、DB等）の状態チェック
- 構造化されたJSON応答の生成

#### 1.8.2.3. 商品一覧ハンドラー

商品一覧を取得するAPIを実装します：

```bash
touch internal/api/handlers/product_handler.go
```

`internal/api/handlers/product_handler.go`に以下の機能を持つハンドラーを実装します：

- モック商品データの提供
- ページネーション処理（page, page_sizeパラメータ対応）
- カテゴリーによるフィルタリング機能
- 構造化されたJSON応答の生成

### 1.8.3. メインアプリケーションの実装

アプリケーションのエントリーポイントを実装します：

```bash
touch cmd/api/main.go
```

`cmd/api/main.go`に以下の機能を持つメインアプリケーションを実装します：

- 設定のロード
- Echoフレームワークの初期化とミドルウェア設定
- APIハンドラーの初期化
- ルーターの設定
- グレースフルシャットダウンの実装
- シグナルハンドリング

### 1.8.4. ハンドラーのユニットテスト

APIハンドラーのユニットテストを実装します：

```bash
mkdir -p tests/unit/api/handlers
touch tests/unit/api/handlers/health_handler_test.go
touch tests/unit/api/handlers/product_handler_test.go
```

ヘルスチェックハンドラーとプロダクトハンドラーのユニットテストを実装し、以下の点を検証します：

- ヘルスチェックAPIが正しいステータスとデータを返すことを確認
- 商品一覧APIがページネーションとフィルタリングを正しく処理することを確認
- レスポンスの構造とデータ型が仕様通りであることを確認

テストを実行するコマンド：

```bash
# バックエンドのすべてのテストを実行
task test:backend

# 特定のパッケージのテストを実行
task test:backend:handlers

# カバレッジレポートの生成
task test:backend:coverage
```

## 1.9. フロントエンド実装

### 1.9.1. 商品一覧APIクライアントの実装

APIとの通信を行うクライアントを実装します：

```bash
cd ../frontend-customer
mkdir -p lib/api
touch lib/api/products.ts
```

`lib/api/products.ts`には以下の機能を実装します：

- REST APIの呼び出し（fetch APIの使用）
- エラーハンドリングと再試行ロジック
- レスポンスの型定義とバリデーション

### 1.9.2. 商品一覧ページの実装

商品一覧を表示するページを実装します：

```bash
mkdir -p app/products
touch app/products/page.tsx
```

`app/products/page.tsx`には以下の機能を実装します：

- サーバーサイドでの商品データ取得
- ページネーションコントロールの実装
- カテゴリーフィルターの実装
- 商品カードコンポーネントの表示
- ローディング状態とエラー状態の処理

### 1.9.3. フロントエンドのテスト実装

フロントエンドコンポーネントのテストを実装します：

```bash
mkdir -p __tests__/components
touch __tests__/components/ProductCard.test.tsx
```

テストでは以下の点を検証します：

- コンポーネントが正しくレンダリングされること
- データが正しく表示されること
- インタラクション（クリックなど）が正しく処理されること
- エラー状態が適切に処理されること

コンテナを使ってテストを実行するコマンド：

```bash
# プロジェクトルートディレクトリから実行
task test:frontend-customer

# 特定のテストファイルのみを実行
task test:frontend-customer:component ProductCard
```

### 1.9.4. 管理者向けフロントエンドの実装

管理者向けフロントエンドにも商品一覧ページを実装します：

```bash
cd ../frontend-admin
mkdir -p app/products
touch app/products/page.tsx
```

管理者向け商品一覧ページには、顧客向けページの機能に加えて以下の機能を追加します：

- 商品一覧の管理者ビュー（在庫状況など追加情報を表示）
- クイック編集機能
- 商品登録ボタン

## 1.10. バックエンドとフロントエンドの連携動作確認

### 1.10.1. バックエンドの起動と動作確認

バックエンドが正常に動作することを確認します：

```bash
# go-taskを使って環境を起動
task start

# ログを確認
task logs:backend

# APIエンドポイントの動作確認
curl http://api.localhost/api/health
curl http://api.localhost/api/products
```

### 1.10.2. フロントエンドの動作確認

フロントエンドの動作を確認するには、まず必要なファイルを実装する必要があります。この段階では、以下の実装を行います：

1. **API通信用クライアント** - バックエンドAPIとの通信機能
2. **基本ページ** - トップページと商品一覧ページ

実装したら、以下の手順で動作確認を行います：

```bash
# フロントエンドのログを確認
task logs:frontend-customer
task logs:frontend-admin
```

#### 1.10.2.1. 動作確認ポイント

1. **顧客向けフロントエンド**:
   - **トップページ**: <http://shop.localhost/> にアクセスし、正しく表示されることを確認
   - **商品一覧**: <http://shop.localhost/products> にアクセスして商品一覧が表示されることを確認
   - バックエンドAPIから商品データが正しく取得され、表示されていることを確認
   - カテゴリーフィルターやページネーションが機能することを確認

2. **管理者向けフロントエンド**:
   - **管理ダッシュボード**: <http://admin.localhost/> にアクセスし、正しく表示されることを確認
   - **商品管理ページ**: <http://admin.localhost/products> にアクセスして商品管理画面が表示されることを確認
   - テーブル形式で商品が表示され、操作ボタン（詳細、編集、削除）が機能することを確認

#### 1.10.2.2. 想定される問題と対処法

- **接続エラー**: バックエンドAPIに接続できない場合は、バックエンドサービスが起動しているか確認
- **画像読み込みエラー**: 画像URLが不正な場合は、コンソールに警告が表示されます
- **データ取得エラー**: API呼び出しに失敗した場合は、エラーメッセージが表示されます

#### 1.10.2.3. 次のステップ

第1週では基本的なフロントエンド実装を行いました。第2週では、より高度な機能（認証、詳細ページ、フォーム処理など）を実装していきます。

> **補足**: 実装の詳細コードはプロジェクトリポジトリの`docs/code-samples/week1`ディレクトリに参照用として格納されています。必要に応じて参照してください。

## 1.11. 理解度チェックと発展課題

### 1.11.1. 理解度チェック

以下の項目について理解度を確認してください：

1. Go/Echoフレームワークの基本構造と機能
2. Docker Composeによる開発環境のセットアップ方法
3. GitHubリポジトリの構成とブランチ戦略
4. バックエンドAPIの設計と実装方法
5. Next.jsによるフロントエンド実装の基本
6. テストの実装方法と役割

### 1.11.2. 発展課題

学習を深めるための発展課題に挑戦しましょう：

1. **追加APIエンドポイントの実装**:
   - 商品詳細を表示するAPIエンドポイント（`/api/products/{id}`）を実装
   - カテゴリー一覧を取得するAPIエンドポイント（`/api/products/categories`）を実装

2. **フロントエンドの拡張**:
   - 商品詳細ページの実装
   - カテゴリーによる商品フィルタリング機能の実装

3. **テスト拡充**:
   - E2E（エンドツーエンド）テストの追加
   - フロントエンドのインテグレーションテスト実装

4. **CI/CD統合**:
   - GitHub Actionsによる自動テストの設定
   - 自動デプロイパイプラインの構築

5. **Traefikの導入**:
   - Traefikを使用したサービス公開の実装
   - SSL/TLS証明書の自動発行設定（開発環境用）

## 1.12. 次週の予告

次週は、以下のトピックについて学びます：

1. **オブザーバビリティの基礎実装（ログ）**
   - slogを利用した構造化ログの実装
   - ログレベル管理の設定
   - CloudWatch Logsへのログ転送
   - フロントエンドログの収集と連携

2. **データベース連携**
   - MySQLスキーマの設計と実装
   - golang-migrateによるマイグレーション管理
   - sqlboilerによるORMの設定

3. **OpenAPI仕様と自動生成**
   - OpenAPI仕様の定義
   - ogenによるコード生成
   - API定義ファーストアプローチ

## 1.13. 参考資料

### 1.13.1. Go言語とEcho

- [Go公式ドキュメント](https://golang.org/doc/)
- [Echo Webフレームワーク](https://echo.labstack.com/)
- [Go言語のテスト入門](https://blog.golang.org/using-go-modules)
- [golangci-lint公式ドキュメント](https://golangci-lint.run/)

### 1.13.2. Next.js とフロントエンド

- [Next.js公式ドキュメント](https://nextjs.org/docs)
- [TypeScript公式ドキュメント](https://www.typescriptlang.org/docs/)
- [TailwindCSS公式ドキュメント](https://tailwindcss.com/docs)
- [Jest公式ドキュメント](https://jestjs.io/docs/getting-started)
- [ESLint公式ドキュメント](https://eslint.org/)
- [Prettier公式ドキュメント](https://prettier.io/)

### 1.13.3. AWS関連

- [AWS SDK for Go](https://aws.github.io/aws-sdk-go/docs/)
- [LocalStack公式ドキュメント](https://docs.localstack.cloud/overview/)

### 1.13.4. Docker と開発環境

- [Docker公式ドキュメント](https://docs.docker.com/)
- [Docker Compose公式ドキュメント](https://docs.docker.com/compose/)
- [direnv公式ドキュメント](https://direnv.net/)
- [Traefik公式ドキュメント](https://doc.traefik.io/traefik/)
