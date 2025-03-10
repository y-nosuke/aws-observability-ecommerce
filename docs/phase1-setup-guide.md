# 1. フェーズ1: 開発環境のセットアップとプロジェクト骨組み構築ガイド

このドキュメントでは、AWSオブザーバビリティ学習用eコマースアプリケーションの開発環境セットアップと基本構造の構築について詳細に説明します。Go Echo+sqlboiler+Next.jsのスタックと、LocalStackによるAWSサービスのエミュレーションを活用します。

## 1.1. 目次

- [1. フェーズ1: 開発環境のセットアップとプロジェクト骨組み構築ガイド](#1-フェーズ1-開発環境のセットアップとプロジェクト骨組み構築ガイド)
  - [1.1. 目次](#11-目次)
  - [1.2. 前提条件](#12-前提条件)
  - [1.3. プロジェクト初期設定](#13-プロジェクト初期設定)
    - [1.3.1. ディレクトリ構造作成](#131-ディレクトリ構造作成)
    - [1.3.2. GitHubリポジトリのセットアップ](#132-githubリポジトリのセットアップ)
      - [1.3.2.1. リポジトリの作成](#1321-リポジトリの作成)
      - [1.3.2.2. リポジトリのクローンと初期プッシュ](#1322-リポジトリのクローンと初期プッシュ)
      - [1.3.2.3. ブランチ保護ルールの設定](#1323-ブランチ保護ルールの設定)
      - [1.3.2.4. イシューテンプレートの作成](#1324-イシューテンプレートの作成)
  - [1.4. Docker Compose環境の構築](#14-docker-compose環境の構築)
    - [1.4.1. 前提条件](#141-前提条件)
    - [1.4.2. 設定ファイルの作成](#142-設定ファイルの作成)
    - [1.4.3. Dockerfileの作成](#143-dockerfileの作成)
    - [1.4.4. 次のステップ](#144-次のステップ)
  - [1.5. Go開発環境のセットアップ](#15-go開発環境のセットアップ)
    - [1.5.1. Go言語のインストール](#151-go言語のインストール)
    - [1.5.2. プロジェクト初期化](#152-プロジェクト初期化)
    - [1.5.3. 必要なライブラリのインストール](#153-必要なライブラリのインストール)
    - [1.5.4. ディレクトリ構造](#154-ディレクトリ構造)
    - [1.5.5. 基本的なアプリケーションコード](#155-基本的なアプリケーションコード)
    - [1.5.6. ホットリロード設定](#156-ホットリロード設定)
    - [1.5.7. データベーススキーマとマイグレーション](#157-データベーススキーマとマイグレーション)
    - [1.5.8. Git設定の追加](#158-git設定の追加)
    - [1.5.9. コンパイルのテスト](#159-コンパイルのテスト)
    - [1.5.10. 次のステップ](#1510-次のステップ)
  - [1.6. Next.js開発環境のセットアップ](#16-nextjs開発環境のセットアップ)
    - [1.6.1. Next.jsの基本概念](#161-nextjsの基本概念)
      - [1.6.1.1. ページルーティング](#1611-ページルーティング)
      - [1.6.1.2. データ取得方法](#1612-データ取得方法)
      - [1.6.1.3. レンダリング戦略](#1613-レンダリング戦略)
    - [1.6.2. 前提条件](#162-前提条件)
    - [1.6.3. Next.jsプロジェクトの作成](#163-nextjsプロジェクトの作成)
    - [1.6.4. 追加パッケージのインストール](#164-追加パッケージのインストール)
    - [1.6.5. ディレクトリ構造の整理](#165-ディレクトリ構造の整理)
    - [1.6.6. 基本レイアウトの作成](#166-基本レイアウトの作成)
    - [1.6.7. 共通コンポーネントの作成](#167-共通コンポーネントの作成)
    - [1.6.8. APIクライアントの設定](#168-apiクライアントの設定)
    - [1.6.9. 商品一覧ページの作成](#169-商品一覧ページの作成)
    - [1.6.10. 管理者ダッシュボードページの作成](#1610-管理者ダッシュボードページの作成)
    - [1.6.11. ESLintとPrettierの設定](#1611-eslintとprettierの設定)
    - [1.6.12. VS Code設定](#1612-vs-code設定)
    - [1.6.13. Next.jsアプリケーションのテスト実行](#1613-nextjsアプリケーションのテスト実行)
    - [1.6.14. 次のステップ](#1614-次のステップ)
  - [1.7. LocalStackのセットアップ](#17-localstackのセットアップ)
    - [1.7.1. LocalStackの概要](#171-localstackの概要)
    - [1.7.2. Docker Composeでの設定](#172-docker-composeでの設定)
      - [1.7.2.1. 環境変数の説明](#1721-環境変数の説明)
      - [1.7.2.2. ボリュームのマウント](#1722-ボリュームのマウント)
    - [1.7.3. 初期化スクリプト](#173-初期化スクリプト)
    - [1.7.4. LocalStackヘルスチェックスクリプトの作成](#174-localstackヘルスチェックスクリプトの作成)
    - [1.7.5. テストスクリプト](#175-テストスクリプト)
    - [1.7.6. バックエンドとの連携](#176-バックエンドとの連携)
      - [1.7.6.1. AWS SDKの設定](#1761-aws-sdkの設定)
      - [1.7.6.2. AWS SDK v2用のCloudWatch Logsクライアント](#1762-aws-sdk-v2用のcloudwatch-logsクライアント)
      - [1.7.6.3. AWS SDK v2用のX-Rayクライアント](#1763-aws-sdk-v2用のx-rayクライアント)
    - [1.7.7. フロントエンドからLocalStackへの接続設定](#177-フロントエンドからlocalstackへの接続設定)
    - [1.7.8. LocalStack Web UIの利用](#178-localstack-web-uiの利用)
    - [1.7.9. トラブルシューティング](#179-トラブルシューティング)
    - [1.7.10. LocalStackでのデバッグテクニック](#1710-localstackでのデバッグテクニック)
    - [1.7.11. AWS SDK v2とLocalStackの統合テスト](#1711-aws-sdk-v2とlocalstackの統合テスト)
    - [1.7.12. 次のステップ](#1712-次のステップ)
  - [1.8. CI/CD初期設定](#18-cicd初期設定)
    - [1.8.1. GitHub Actions基本設定](#181-github-actions基本設定)
    - [1.8.2. Terraformによる基本インフラ定義](#182-terraformによる基本インフラ定義)
    - [1.8.3. Terraformの初期化と実行手順](#183-terraformの初期化と実行手順)
    - [1.8.4. 運用ガイドラインの作成](#184-運用ガイドラインの作成)
    - [1.8.5. AWSオブザーバビリティ準備の文書化](#185-awsオブザーバビリティ準備の文書化)
    - [1.8.6. 動作確認とテスト](#186-動作確認とテスト)
      - [1.8.6.1. 環境の起動](#1861-環境の起動)
      - [1.8.6.2. 基本的な動作確認](#1862-基本的な動作確認)
      - [1.8.6.3. トラブルシューティング](#1863-トラブルシューティング)
    - [1.8.7. 次のステップ](#187-次のステップ)
  - [1.9. 開発ワークフロー](#19-開発ワークフロー)
  - [1.10. 次のステップ](#110-次のステップ)

## 1.2. 前提条件

開発を始める前に、以下のソフトウェアがインストールされていることを確認してください:

- Docker と Docker Compose
- Go (1.24以上)
- Node.js (23以上) と npm/yarn
- Git
- Visual Studio CodeやGoLandなどのIDE/エディタ
- AWS CLI

## 1.3. プロジェクト初期設定

### 1.3.1. ディレクトリ構造作成

まず、プロジェクト全体のディレクトリ構造を作成します:

```bash
# プロジェクトのルートディレクトリを作成
mkdir -p aws-observability-ecommerce
cd aws-observability-ecommerce

# バックエンド、フロントエンド、インフラ用のディレクトリを作成
mkdir -p backend frontend infra
mkdir -p infra/localstack
mkdir -p infra/terraform
mkdir -p docs
```

### 1.3.2. GitHubリポジトリのセットアップ

#### 1.3.2.1. リポジトリの作成

1. GitHubにログインし、「New repository」をクリックします。
2. リポジトリ名（例: `aws-observability-ecommerce`）を入力します。
3. 簡単な説明を追加します。例: 「AWSオブザーバビリティ学習用eコマースアプリケーション」
4. 必要に応じて「Private」を選択して非公開リポジトリにします。
5. 「Initialize this repository with a README」にチェックを入れます。
6. 「.gitignore template」で「Go」を選択します。
7. 「Create repository」をクリックしてリポジトリを作成します。

#### 1.3.2.2. リポジトリのクローンと初期プッシュ

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

#### 1.3.2.3. ブランチ保護ルールの設定

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

#### 1.3.2.4. イシューテンプレートの作成

フィーチャー開発やバグレポートのためのテンプレートを作成します：

```bash
# イシューテンプレート用のディレクトリを作成
mkdir -p .github/ISSUE_TEMPLATE

# 機能リクエスト用テンプレートを作成
touch .github/ISSUE_TEMPLATE/feature_request.md

# バグレポート用テンプレートを作成
touch .github/ISSUE_TEMPLATE/bug_report.md
```

## 1.4. Docker Compose環境の構築

Docker Composeを使用して開発環境を構築することで、プロジェクトの依存関係を簡単に管理し、開発者間で一貫した環境を共有できます。ここでは、バックエンド(Go)、フロントエンド(Next.js)、データベース(MySQL)、およびAWSサービスエミュレーター(LocalStack)を含む完全な開発環境を構築します。

### 1.4.1. 前提条件

以下のソフトウェアがインストールされていることを確認してください：

- Docker と Docker Compose
- Git

### 1.4.2. 設定ファイルの作成

プロジェクトルートに`compose.yml`ファイルを作成します。このファイルには、バックエンドAPI、MySQL、フロントエンド、LocalStackの設定が含まれます。

   ```bash
   touch compose.yml
   ```

### 1.4.3. Dockerfileの作成

バックエンド用とフロントエンド用のDockerfileを作成します:

1. **バックエンド用Dockerfile (`backend/Dockerfile.dev`)**:
   - Go 1.24 Alpineをベースイメージとして使用
   - Airをインストールしてホットリロード機能を有効化

   ```bash
   touch backend/Dockerfile.dev
   ```

2. **フロントエンド用Dockerfile (`frontend/Dockerfile.dev`)**:
   - Node.js 23 Alpineをベースイメージとして使用

   ```bash
   touch frontend/Dockerfile.dev
   ```

### 1.4.4. 次のステップ

Docker Compose環境が正常に動作していることを確認したら、次のステップに進みます:

1. バックエンドAPIの拡張（ハンドラー、サービス、リポジトリの追加）
2. フロントエンドの基本構造の作成（ページとコンポーネント）
3. データベーススキーマのセットアップとマイグレーション
4. オブザーバビリティコンポーネントの実装開始

これで、開発環境の構築が完了し、実際の機能実装に進むことができます。

## 1.5. Go開発環境のセットアップ

Go言語を使用してバックエンドAPIを開発するための環境を整えます。Echo Webフレームワーク、sqlboiler ORM、slogロギングライブラリなど、eコマースアプリケーションの開発に必要なツールとライブラリを設定します。

### 1.5.1. Go言語のインストール

まず、Go言語がインストールされていることを確認します。インストールされていない場合は、[公式サイト](https://golang.org/dl/)からダウンロードしてインストールしてください。Go 1.21以上を推奨します。

```bash
# Goのバージョンを確認
go version
```

### 1.5.2. プロジェクト初期化

```bash
cd backend
go mod init github.com/y-nosuke/aws-observability-ecommerce
```

### 1.5.3. 必要なライブラリのインストール

```bash
# Webフレームワーク
go get github.com/labstack/echo/v4
go get github.com/labstack/echo/v4/middleware

# ロギング
go get golang.org/x/exp/slog

# 設定管理
go get github.com/spf13/viper

# データベースドライバ
go get github.com/go-sql-driver/mysql

# テスト
go get github.com/stretchr/testify
```

### 1.5.4. ディレクトリ構造

以下のようなディレクトリ構造を作成します:

```bash
mkdir -p cmd/api
mkdir -p internal/api/{handlers,middlewares,routes}
mkdir -p internal/config
mkdir -p internal/domain/{models,repositories}
mkdir -p internal/observability/{logging,metrics,tracing}
mkdir -p internal/services
mkdir -p config
mkdir -p migrations
mkdir -p scripts/init-db
mkdir -p tests
```

### 1.5.5. 基本的なアプリケーションコード

以下の基本的なファイルを作成します:

1. **`cmd/api/main.go`**: アプリケーションのエントリーポイント

   ```bash
   touch cmd/api/main.go
   ```

2. **`internal/config/config.go`**: 設定ファイル読み込み

   ```bash
   touch internal/config/config.go
   ```

3. **`internal/observability/logging/logger.go`**: ロギング設定

   ```bash
   touch internal/observability/logging/logger.go
   ```

設定ファイル `config/config.yaml` も作成します:

```bash
touch config/config.yaml
```

### 1.5.6. ホットリロード設定

`.air.toml`ファイルを作成して、ホットリロードを設定します。

```bash
touch .air.toml
```

### 1.5.7. データベーススキーマとマイグレーション

1. **ORM設定**: `sqlboiler.toml`ファイルを作成します。

   ```bash
   touch sqlboiler.toml
   ```

2. **マイグレーションファイル**:
   - `migrations/000001_create_tables.up.sql`
   - `migrations/000001_create_tables.down.sql`

   ```bash
   touch migrations/{000001_create_tables.up.sql,000001_create_tables.down.sql}
   ```

3. **初期データ**:
   - `scripts/init-db/01-init-data.sql`

   ```bash
   touch scripts/init-db/01-init-data.sql
   ```

### 1.5.8. Git設定の追加

`.gitignore` ファイルを作成して、不要なファイルをGitから除外します:

```ignore
# バイナリファイル
bin/
tmp/

# 依存関係
vendor/

# ログファイル
*.log

# 環境設定ファイル
.env
.env.local

# エディタの設定ファイル
.vscode/
.idea/

# sqlboilerの生成ファイル
internal/domain/models/

# Airの一時ファイル
tmp/
```

### 1.5.9. コンパイルのテスト

すべての設定が正しく行われたか確認するために、プロジェクトをコンパイルします:

```bash
go build ./cmd/api

# もしくは

air
```

エラーがなければ、環境設定は成功です。

### 1.5.10. 次のステップ

これで基本的なGo開発環境のセットアップは完了しました。次のステップでは、以下の作業を行います:

1. ハンドラー、サービス、リポジトリの実装
2. データベース接続の実装
3. オブザーバビリティ機能の追加（AWS SDK v2アプローチ）
4. ユニットテストとインテグレーションテストの作成

これらの実装を通じて、MVPの機能を段階的に構築していきます。

## 1.6. Next.js開発環境のセットアップ

### 1.6.1. Next.jsの基本概念

Next.jsを使い始める前に、その基本概念を理解しておくと役立ちます：

#### 1.6.1.1. ページルーティング

Next.jsではファイルベースのルーティングを採用しています。`app`ディレクトリ（App Router）または`pages`ディレクトリ（Pages Router）内のファイルが自動的にルートになります。

例えば:

- `app/page.tsx` → `/` (ホームページ)
- `app/products/page.tsx` → `/products` (商品一覧ページ)
- `app/products/[id]/page.tsx` → `/products/123` (商品詳細ページ)

#### 1.6.1.2. データ取得方法

Next.jsでは主に以下のデータ取得方法があります:

1. **サーバーコンポーネント内でのデータ取得** (App Router)
   - サーバー側で実行され、ページがレンダリングされる前にデータを取得

2. **getStaticProps / getServerSideProps** (Pages Router)
   - ビルド時（静的）またはリクエスト時（サーバーサイド）にデータを取得

3. **useEffect + fetch / SWR / React Query**
   - クライアント側でのデータ取得

#### 1.6.1.3. レンダリング戦略

1. **静的サイト生成 (SSG)**: ビルド時にHTMLを生成
2. **サーバーサイドレンダリング (SSR)**: リクエスト時にサーバーでHTMLを生成
3. **クライアントサイドレンダリング (CSR)**: ブラウザでJavaScriptを使用してレンダリング
4. **インクリメンタル静的再生成 (ISR)**: 静的ページを一定間隔で自動的に再生成

### 1.6.2. 前提条件

Next.jsの開発環境をセットアップする前に、以下のソフトウェアがインストールされていることを確認してください:

- Node.js (18.x以上)
- npm (9.x以上) または yarn (1.22.x以上)

```bash
# バージョン確認
node -v
npm -v # または yarn -v
```

### 1.6.3. Next.jsプロジェクトの作成

```bash
cd frontend
npx create-next-app@latest .
```

インタラクティブなセットアップで以下のように選択します:

```bash
✔ Would you like to use TypeScript? … Yes
✔ Would you like to use ESLint? … Yes
✔ Would you like to use Tailwind CSS? … Yes
✔ Would you like to use `src/` directory? … No
✔ Would you like to use App Router? (recommended) … Yes
✔ Would you like to use Turbopack for next dev? … No
✔ Would you like to customize the default import alias? … No
```

### 1.6.4. 追加パッケージのインストール

```bash
npm install axios swr react-hook-form zod @hookform/resolvers
npm install -D @types/node @types/react @types/react-dom
```

各パッケージの用途:

- `axios`: HTTPリクエスト用クライアント
- `swr`: データ取得のためのReactフック (Stale-While-Revalidate)
- `react-hook-form`: フォーム管理ライブラリ
- `zod`: スキーマ検証ライブラリ
- `@hookform/resolvers`: react-hook-formとzodの連携

### 1.6.5. ディレクトリ構造の整理

```bash
# 顧客向けページ用ディレクトリ
mkdir -p app/\(customer\)/{products,cart,checkout}

# 管理者向けページ用ディレクトリ
mkdir -p app/\(admin\)/{dashboard,products,inventory}

# コンポーネント用ディレクトリ
mkdir -p components/{common,layout,product,cart,admin}

# API通信用ライブラリなどのディレクトリ
mkdir -p lib/{api,hooks,utils}

# タイプ定義用ディレクトリ
mkdir -p types
```

`.env.local`ファイルを作成します:

```bash
touch .env.local
```

```bash
# APIのベースURL
NEXT_PUBLIC_API_URL=http://localhost:8080

# 開発環境フラグ
NEXT_PUBLIC_DEV_ENV=true
```

注意:

- `NEXT_PUBLIC_`プレフィックスを付けると、クライアントサイドでも利用可能になります
- `.env.local`はGitで管理されないため、機密情報を安全に保管できます

### 1.6.6. 基本レイアウトの作成

以下のコンポーネントを作成します:

1. **アプリケーションレイアウト**: `app/layout.tsx`
2. **顧客向けレイアウト**: `app/(customer)/layout.tsx`

   ```bash
   touch app/\(customer\)/layout.tsx
   ```

3. **管理者向けレイアウト**: `app/(admin)/layout.tsx`

   ```bash
   touch app/\(admin\)/layout.tsx
   ```

### 1.6.7. 共通コンポーネントの作成

1. **ヘッダーコンポーネント**: `components/layout/Header.tsx`

   ```bash
   touch components/layout/Header.tsx
   ```

2. **フッターコンポーネント**: `components/layout/Footer.tsx`

   ```bash
   touch components/layout/Footer.tsx
   ```

3. **管理者向けヘッダーとサイドバーコンポーネント**:
   - `components/admin/AdminHeader.tsx`
   - `components/admin/AdminSidebar.tsx`

   ```bash
   touch components/admin/{AdminHeader.tsx,AdminSidebar.tsx}
   ```

### 1.6.8. APIクライアントの設定

APIクライアントとカスタムフックを設定します:

1. **APIクライアント**: `lib/api/client.ts`

   ```bash
   touch lib/api/client.ts
   ```

2. **商品取得用カスタムフック**: `lib/hooks/useProducts.ts`

   ```bash
   touch lib/hooks/useProducts.ts
   ```

3. **トップページ**: `app/page.tsx`

### 1.6.9. 商品一覧ページの作成

```bash
touch app/\(customer\)/products/page.tsx
touch components/product/ProductList.tsx
```

### 1.6.10. 管理者ダッシュボードページの作成

```bash
touch app/\(admin\)/dashboard/page.tsx
```

### 1.6.11. ESLintとPrettierの設定

コード品質とフォーマットの一貫性を保つために、ESLintとPrettierを設定します:

```bash
npm install -D prettier eslint-config-prettier
```

`.prettierrc`ファイルを作成します:

```json
{
  "semi": true,
  "trailingComma": "es5",
  "singleQuote": true,
  "tabWidth": 2,
  "printWidth": 100
}
```

`.eslintrc.json`を更新します:

```json
{
  "extends": [
    "next/core-web-vitals",
    "prettier"
  ],
  "rules": {
    "react/no-unescaped-entities": "off"
  }
}
```

### 1.6.12. VS Code設定

VS Codeでの開発体験を向上させるために、`.vscode/settings.json`ファイルを作成します:

```json
{
  "editor.formatOnSave": true,
  "editor.defaultFormatter": "esbenp.prettier-vscode",
  "editor.codeActionsOnSave": {
    "source.fixAll.eslint": true
  },
  "typescript.tsdk": "frontend/node_modules/typescript/lib",
  "typescript.enablePromptUseWorkspaceTsdk": true
}
```

### 1.6.13. Next.jsアプリケーションのテスト実行

ここまでの設定が完了したら、Next.jsアプリケーションを起動してテストします:

```bash
npm run dev
```

これにより、開発サーバーが起動し、ブラウザで `http://localhost:3000` にアクセスできるようになります。

### 1.6.14. 次のステップ

基本的なNext.js開発環境のセットアップが完了しました。次のステップでは、以下の作業を行います:

1. 残りのページ（商品詳細、カート、チェックアウトなど）の実装
2. フォームの実装と検証
3. 状態管理の実装（カート機能など）
4. APIと連携したデータ表示の実装
5. レスポンシブデザインの最適化

これらの実装を通じて、MVPの機能を段階的に構築していきます。

## 1.7. LocalStackのセットアップ

LocalStackは、AWSサービスをローカル環境でエミュレートするためのツールです。本番環境にデプロイする前に、ローカルでAWSサービスを使った開発やテストを行うことができます。この文書では、eコマースアプリケーションのオブザーバビリティ機能を実装するためのLocalStackの設定方法を説明します。

### 1.7.1. LocalStackの概要

LocalStackは以下のようなAWSサービスをローカルでエミュレートできます：

- **CloudWatch Logs** - ログデータの収集と分析
- **X-Ray** - 分散トレーシング
- **S3** - オブジェクトストレージ
- **SQS** - メッセージキュー
- **SNS** - 通知サービス
- **その他多数のAWSサービス**

これにより、コストをかけずにAWSサービスを使った開発やテストが可能になります。

LocalStackを使用する主な利点：

- **コスト削減**: AWSリソースの使用料金を払わずに開発・テストが可能
- **オフライン開発**: インターネット接続なしで開発可能
- **高速な反復**: デプロイを待たずに即座にテスト可能
- **デバッグの容易さ**: ローカル環境でのデバッグが容易

本プロジェクトでは特に以下のサービスが重要です：

- **CloudWatch Logs**: アプリケーションログのモニタリング
- **X-Ray**: 分散トレーシングとサービスマップ
- **S3**: 静的アセットの保存
- **SQS/SNS**: メッセージングと通知（拡張機能用）

### 1.7.2. Docker Composeでの設定

既に「Docker Compose環境の構築」セクションで基本的なLocalStackの設定を行いましたが、ここではさらに詳細について説明します。

Docker Compose設定ファイル（`docker-compose.yml`）内のLocalStack部分を再確認しましょう：

```yaml
# AWS サービスエミュレーター (LocalStack)
localstack:
  image: localstack/localstack:latest
  ports:
    - "4566:4566"
  environment:
    - SERVICES=s3,cloudwatch,logs,xray,sqs,sns
    - DEBUG=${DEBUG-}
    - DATA_DIR=/tmp/localstack/data
    - DOCKER_HOST=unix:///var/run/docker.sock
    - AWS_DEFAULT_REGION=us-east-1
    - AWS_ACCESS_KEY_ID=test
    - AWS_SECRET_ACCESS_KEY=test
  volumes:
    - ./infra/localstack:/docker-entrypoint-initaws.d
    - localstack-data:/tmp/localstack
    - /var/run/docker.sock:/var/run/docker.sock
  networks:
    - app-network
```

#### 1.7.2.1. 環境変数の説明

- `SERVICES`: 有効にするAWSサービスのリスト
- `DEBUG`: デバッグモードの有効化（未設定時は無効）
- `DATA_DIR`: LocalStackのデータ保存ディレクトリ
- `DOCKER_HOST`: DockerソケットのURL（コンテナ内からDockerコマンドを実行するため）
- `AWS_DEFAULT_REGION`: デフォルトのAWSリージョン
- `AWS_ACCESS_KEY_ID`と`AWS_SECRET_ACCESS_KEY`: ダミーの認証情報（任意の値で動作）

#### 1.7.2.2. ボリュームのマウント

- `/docker-entrypoint-initaws.d`: 起動時に実行されるスクリプトを格納するディレクトリ
- `/tmp/localstack`: データの永続化用のディレクトリ
- `/var/run/docker.sock`: ホストのDockerソケット（コンテナ内からDockerコマンドを実行するため）

### 1.7.3. 初期化スクリプト

LocalStackの起動時に自動的に実行される初期化スクリプトを作成します。`infra/localstack/init-resources.sh`：

```bash
touch infra/localstack/init-resources.sh
chmod +x infra/localstack/init-resources.sh
```

このスクリプトにより、以下のAWSリソースが自動的に作成されます：

1. **S3バケット**： `ecommerce-static-assets` - フロントエンドの静的アセット用
2. **CloudWatch Logsのロググループ**： APIとフロントエンドのログ用
3. **X-Rayのサンプリングルール**： トレースデータの収集設定
4. **SQSキュー**： `ecommerce-orders` - 注文処理用
5. **SNSトピック**： `ecommerce-notifications` - 通知用

### 1.7.4. LocalStackヘルスチェックスクリプトの作成

LocalStackが正常に起動しているかを確認するためのヘルスチェックスクリプトを作成します。これにより、開発環境の問題をすばやく特定できます。

`infra/localstack/health-check.sh`ファイルを作成します：

```bash
touch infra/localstack/health-check.sh
chmod +x infra/localstack/health-check.sh
```

### 1.7.5. テストスクリプト

LocalStackのリソースをテストするスクリプトを作成します。`infra/localstack/test-resources.sh`：

```bash
touch infra/localstack/test-resources.sh
chmod +x infra/localstack/test-resources.sh
```

このスクリプトは以下の操作を行います:

- S3バケットのテスト（リスト取得、ファイルアップロード）
- CloudWatch Logsのテスト（ロググループ一覧取得、ログイベント送信）
- SQSのテスト（キュー一覧取得、メッセージ送信）
- SNSのテスト（トピック一覧取得）

### 1.7.6. バックエンドとの連携

GoバックエンドアプリケーションからLocalStackに接続するための設定を行います。

#### 1.7.6.1. AWS SDKの設定

バックエンドアプリケーションからLocalStackに接続するための設定を行います。`internal/config/aws.go`ファイルを作成します：

```bash
cd backend
touch internal/config/aws.go
```

この設定により:

- LocalStackのエンドポイントURLを指定できます
- テスト用の静的認証情報を設定できます
- リージョン設定を行えます

#### 1.7.6.2. AWS SDK v2用のCloudWatch Logsクライアント

`internal/observability/logging/cloudwatch.go`ファイルを作成します：

```bash
touch internal/observability/logging/cloudwatch.go
```

#### 1.7.6.3. AWS SDK v2用のX-Rayクライアント

`internal/observability/tracing/xray.go`ファイルの基本構造を作成します：

```bash
touch internal/observability/tracing/xray.go
```

### 1.7.7. フロントエンドからLocalStackへの接続設定

フロントエンドアプリケーションからLocalStackに直接接続することはセキュリティ上推奨されませんが、開発環境ではバックエンドAPIを介して間接的に接続します。

フロントエンドからAWSサービスにアクセスする場合は、以下のアプローチを取ります：

1. **バックエンドプロキシ経由でアクセス**：
   - S3へのファイルアップロードなどの操作は、バックエンドAPIを経由して行います
   - APIエンドポイントを作成し、そこからLocalStackに接続します

2. **CloudFrontのエミュレーション**：
   - 静的アセット配信用に、LocalStackのS3をNginxなどでプロキシします
   - 開発環境の`next.config.js`でアセットパスを設定します

例えば、Next.jsの設定で環境変数を使用してS3のエンドポイントを切り替えることができます：

```javascript
// frontend/next.config.js の一部
const nextConfig = {
  env: {
    // 開発環境ではLocalStack、本番環境では実際のAWSエンドポイントを使用
    S3_ENDPOINT: process.env.NODE_ENV === 'development'
      ? 'http://localhost:4566/ecommerce-static-assets'
      : 'https://ecommerce-static-assets.s3.amazonaws.com'
  },
  // その他の設定...
};
```

フロントエンドコードでは、この環境変数を使用してアセットURLを構築します：

```javascript
// 画像URLの例
const imageUrl = `${process.env.S3_ENDPOINT}/${productImagePath}`;
```

ただし、本番環境ではCORSやセキュリティの問題から、フロントエンドから直接AWSサービスにアクセスすることは避け、バックエンドAPIを経由するアーキテクチャを推奨します。

### 1.7.8. LocalStack Web UIの利用

LocalStackにはWebUIが付属しており、作成したリソースを視覚的に確認することができます。

Web UIにアクセスするには：

- ブラウザで `http://localhost:4566/_localstack/setup` にアクセス
- または `http://localhost:8080` でLocalStack Web UIにアクセス（別途設定が必要）

Web UIを利用することで：

- 作成したリソースの視覚的な確認
- リソースの検索と管理
- LocalStackの状態監視
- ローカルに保存されたデータの閲覧
が可能になります。

### 1.7.9. トラブルシューティング

LocalStackの設定で発生しがちな問題と解決策：

1. **コンテナが起動しない**
   - Dockerのログを確認：`docker-compose logs localstack`
   - メモリ不足の場合はDocker設定のリソース割り当てを確認
   - 必要なポート（4566）が他のプロセスで使用されていないか確認

2. **AWS CLIからリソースにアクセスできない**
   - `--endpoint-url` パラメータが正しいか確認
   - AWS認証情報が設定されているか確認（開発環境では任意の値で動作）
   - サービス名が正しく指定されているか確認

3. **初期化スクリプトが実行されない**
   - スクリプトに実行権限があるか確認：`chmod +x infra/localstack/*.sh`
   - 改行コードがLF（Unix形式）になっているか確認
   - スクリプト内のシンタックスエラーを確認

4. **特定のサービスが動作しない**
   - `SERVICES` 環境変数にそのサービスが含まれているか確認
   - サービスが最新バージョンのLocalStackでサポートされているか確認
   - Pro版が必要なサービスではないか確認（一部のサービスはLocalStack Proのみ）

5. **リソース作成が失敗する**
   - LocalStackのログを詳細に確認
   - AWS CLIのコマンドを手動で実行して動作確認
   - LocalStackのバージョンを最新に更新

### 1.7.10. LocalStackでのデバッグテクニック

効果的なデバッグのためのヒント：

1. **詳細なログを有効にする**:

   ```yaml
   environment:
     - DEBUG=1
     - LS_LOG=trace
   ```

2. **リソースの状態を確認する**:

   ```bash
   # S3バケットの一覧
   aws --endpoint-url=http://localhost:4566 s3 ls

   # CloudWatch Logsのグループ一覧
   aws --endpoint-url=http://localhost:4566 logs describe-log-groups

   # SQSキューの一覧
   aws --endpoint-url=http://localhost:4566 sqs list-queues
   ```

3. **LocalStackのヘルスチェック**:

   ```bash
   curl http://localhost:4566/_localstack/health
   ```

### 1.7.11. AWS SDK v2とLocalStackの統合テスト

開発環境でAWS SDK v2がLocalStackと正しく連携できるかテストするためのスクリプトを作成します。

`backend/cmd/test-localstack/main.go`ファイルを作成します：

```bash
mkdir -p backend/cmd/test-localstack
touch backend/cmd/test-localstack/main.go
```

### 1.7.12. 次のステップ

LocalStackの基本的な設定が完了したら、以下の作業を行います：

1. AWS SDK v2を使用したオブザーバビリティ機能の実装
2. バックエンドサービスのCloudWatch Logs統合
3. X-Rayによる分散トレーシングの実装
4. メトリクス収集システムの構築
5. OpenTelemetryとの統合準備

これらの実装を通じて、AWSのオブザーバビリティサービスを深く理解し、効果的に活用する方法を学んでいきます。

## 1.8. CI/CD初期設定

### 1.8.1. GitHub Actions基本設定

GitHub Actionsを使用してCIパイプラインを設定します:

```bash
mkdir -p .github/workflows
touch .github/workflows/{backend-ci.yml,frontend-ci.yml,terraform-validate.yml}
```

1. **バックエンド検証用ワークフロー**: `.github/workflows/backend-ci.yml`
   - リント（golangci-lint）
   - テスト実行
   - ビルド確認

2. **フロントエンド検証用ワークフロー**: `.github/workflows/frontend-ci.yml`
   - リント（ESLint）
   - ビルド確認

3. **Terraformのバリデーションワークフロー**: `.github/workflows/terraform-validate.yml`
   - バリデーション（terraform validate）
   - ビルド確認

### 1.8.2. Terraformによる基本インフラ定義

Infrastructure as Code (IaC)を実践するため、Terraformを設定します:

1. **プロジェクト構造作成**:

   ```bash
   mkdir -p infra/terraform/{environments,modules}
   mkdir -p infra/terraform/environments/{dev,staging,prod}
   mkdir -p infra/terraform/modules/{vpc,ecr,s3,iam}
   ```

2. **共通設定ファイル**: `infra/terraform/versions.tf`

   ```bash
   touch infra/terraform/versions.tf
   ```

3. **VPCモジュール**:
   - `infra/terraform/modules/vpc/main.tf`
   - `infra/terraform/modules/vpc/variables.tf`
   - `infra/terraform/modules/vpc/outputs.tf`

   ```bash
   touch infra/terraform/modules/vpc/{main.tf,variables.tf,outputs.tf}
   ```

4. **ECRモジュール**:
   - `infra/terraform/modules/ecr/main.tf`
   - `infra/terraform/modules/ecr/variables.tf`
   - `infra/terraform/modules/ecr/outputs.tf`

   ```bash
   touch infra/terraform/modules/ecr/{main.tf,variables.tf,outputs.tf}
   ```

5. **S3モジュール**:
   - `infra/terraform/modules/s3/main.tf`
   - `infra/terraform/modules/s3/variables.tf`
   - `infra/terraform/modules/s3/outputs.tf`

   ```bash
   touch infra/terraform/modules/s3/{main.tf,variables.tf,outputs.tf}
   ```

6. **IAMモジュール**:
   - `infra/terraform/modules/iam/main.tf`
   - `infra/terraform/modules/iam/variables.tf`
   - `infra/terraform/modules/iam/outputs.tf`

   ```bash
   touch infra/terraform/modules/iam/{main.tf,variables.tf,outputs.tf}
   ```

7. **開発環境設定**:
   - `infra/terraform/environments/dev/main.tf`
   - `infra/terraform/environments/dev/variables.tf`
   - `infra/terraform/environments/dev/outputs.tf`
   - `infra/terraform/environments/dev/backend.tf`

   ```bash
   touch infra/terraform/environments/dev/{main.tf,variables.tf,outputs.tf,backend.tf}
   ```

### 1.8.3. Terraformの初期化と実行手順

Terraformコードを初めて実行する際の手順書を作成します：

`infra/terraform/README.md`ファイルを作成します：

```bash
touch infra/terraform/README.md
```

### 1.8.4. 運用ガイドラインの作成

プロジェクトのルートディレクトリに運用ガイドラインのドキュメントを作成します：

`CONTRIBUTING.md`ファイルを作成します：

```bash
touch CONTRIBUTING.md
```

### 1.8.5. AWSオブザーバビリティ準備の文書化

オブザーバビリティの実装に向けた準備として、ガイドラインドキュメントを作成します：

`docs/observability-setup.md`ファイルを作成します：

```bash
touch docs/observability-setup.md
```

### 1.8.6. 動作確認とテスト

#### 1.8.6.1. 環境の起動

設定が完了したら、Docker Composeを使って開発環境を起動します:

```bash
# プロジェクトルートディレクトリから
docker-compose up --build
```

このコマンドにより、以下のサービスが起動します:

- バックエンドAPI (Go Echo)
- MySQL データベース
- フロントエンド (Next.js)
- LocalStack (AWSサービスエミュレーター)

#### 1.8.6.2. 基本的な動作確認

以下のエンドポイントにアクセスして、各サービスが正常に動作していることを確認します:

1. **バックエンドAPI**: <http://localhost:8080/health>
   - 正常に動作していれば、以下のようなJSONレスポンスが返ります:

   ```json
   {"status":"ok","version":"0.1.0"}
   ```

2. **フロントエンド**: <http://localhost:3000>
   - Next.jsのウェルカムページが表示されます

3. **LocalStack**: <http://localhost:4566/health>
   - LocalStackのヘルスステータスが返ります

#### 1.8.6.3. トラブルシューティング

よくある問題とその解決策:

1. **ポートの競合**:
   - エラーメッセージ: `Error starting userland proxy: listen tcp 0.0.0.0:8080: bind: address already in use`
   - 解決策: 使用中のポートを変更するか、競合しているプロセスを終了します

2. **MySQLコネクション問題**:
   - エラーメッセージ: `dial tcp: connect: connection refused`
   - 解決策: MySQLサービスが起動しているか確認します。必要に応じて`docker-compose restart mysql`で再起動します

3. **ボリュームマウント問題**:
   - エラーメッセージ: `cannot start service xxx: error while creating mount source path`
   - 解決策: `docker-compose down -v`を実行してボリュームをクリーンアップし、再度起動します

4. **LocalStackリソース初期化問題**:
   - 解決策: LocalStackコンテナのログを確認し、初期化スクリプトに実行権限があることを確認します

### 1.8.7. 次のステップ

これでCI/CD初期設定は完了しました。以下の作業を行ってフェーズ1を完了させます：

1. プロジェクトファイルをGitHubリポジトリにコミットしてプッシュします：

```bash
# 変更を追加してコミット
git add .
git commit -m "[CI/CD] Add initial CI/CD setup and infrastructure code"

# リモートリポジトリにプッシュ
git push origin main
```

1. GitHub Actionsワークフローが正常に実行されることを確認します。

1. 次のフェーズ「バックエンド実装」に向けて準備します。

フェーズ1の完了により、以下の準備が整いました：

- 開発環境のDockerコンテナ設定
- プロジェクトの基本構造
- コード品質を保証するCI設定
- インフラストラクチャのコード化（IaC）
- 各環境（開発、ステージング、本番）の基本設定

これらの基盤を元に、フェーズ2でバックエンドの実装を進めていきます。# CI/CD初期設定

継続的インテグレーション(CI)と継続的デリバリー(CD)は、現代のソフトウェア開発において重要なプラクティスです。このセクションでは、AWSオブザーバビリティ学習用eコマースアプリケーションのためのCI/CD環境の初期設定について説明します。GitHub ActionsとTerraformを使用して、コードの検証からインフラストラクチャのデプロイまでを自動化します。

## 1.9. 開発ワークフロー

環境を設定した後の基本的な開発ワークフローは次のとおりです:

1. コンテナを起動:

   ```bash
   docker-compose up -d
   ```

2. バックエンドのコード変更:
   - `backend`ディレクトリ内のファイルを編集
   - Airによる自動リロードが行われる

3. フロントエンドのコード変更:
   - `frontend`ディレクトリ内のファイルを編集
   - Next.jsの開発サーバーによる自動リロードが行われる

4. AWS関連の設定変更:
   - LocalStackのリソースを更新する場合は、`init-resources.sh`を編集し、コンテナを再起動

5. コンテナを停止:

   ```bash
   docker-compose down
   ```

## 1.10. 次のステップ

フェーズ1が完了したら、以下を確認してください:

1. バックエンドAPIが正常に起動し、ヘルスチェックエンドポイントが機能している
2. フロントエンドアプリケーションが正常に起動し、基本的なUIが表示される
3. LocalStackが正常に動作し、必要なAWSリソースが作成されている
4. GitHub Actions CIが設定され、コードの検証が自動的に行われる
5. Terraformの基本設定が完了し、インフラをコードで管理できる状態になっている

これらが確認できたら、次のフェーズ「バックエンド実装」に進みます。フェーズ2では:

1. データベースモデルの実装と完成
2. コアサービス（商品、在庫、注文処理など）の実装
3. AWS SDK v2を使用したオブザーバビリティ基盤の初期実装

を行います。

フェーズ1で構築した開発環境と基本構造を活用し、実際の機能実装に取り組んでいきましょう。
