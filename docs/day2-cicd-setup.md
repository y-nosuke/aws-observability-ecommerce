# Day 2: CI/CDと基本インフラのセットアップ手順

このガイドでは、Day 1で作成したAWS環境の上にCI/CDパイプラインを構築し、基本的なサービスリソースをプロビジョニングします。

## 前提条件

- Day 1の作業が完了しており、基本的なAWSリソース（VPC、サブネット、セキュリティグループなど）が作成済みであること
- GitHubアカウントを持っていること
- GitHubリポジトリ `aws-observability-ecommerce` が作成済みであること

## 1. リポジトリ構造の整備

まず、プロジェクトのリポジトリ構造を整備しましょう。

```bash
# プロジェクトのルートディレクトリで以下のコマンドを実行
mkdir -p frontend
mkdir -p backend/catalog-service
mkdir -p .github/workflows
mkdir -p scripts
```

## 2. AWS認証情報のGitHubシークレットへの登録

GitHub Actionsで使用するAWS認証情報をGitHubシークレットに登録します。

1. GitHubの対象リポジトリに移動
2. リポジトリの「Settings」タブをクリック
3. 左側のメニューから「Secrets and variables」→「Actions」を選択
4. 「New repository secret」ボタンをクリック
5. 以下の3つのシークレットを追加:

   | シークレット名          | 値                                            |
   | ----------------------- | --------------------------------------------- |
   | `AWS_ACCESS_KEY_ID`     | AWS IAMユーザーのアクセスキーID               |
   | `AWS_SECRET_ACCESS_KEY` | AWS IAMユーザーのシークレットアクセスキー     |
   | `AWS_REGION`            | 使用するAWSリージョン（例: `ap-northeast-1`） |

## 3. CI/CDワークフローファイルの作成

### インフラデプロイワークフロー

`.github/workflows/terraform-deploy.yml` ファイルを作成します。このワークフローは、Terraformコードに変更があった場合に自動的に実行されます。

> ファイルの内容は `terraform-deploy.yml` Artifactを参照してください。

### フロントエンドデプロイワークフロー

`.github/workflows/frontend-deploy.yml` ファイルを作成します。このワークフローは、フロントエンドのコードに変更があった場合に自動的に実行されます。

> ファイルの内容は `frontend-deploy.yml` Artifactを参照してください。

### バックエンドデプロイワークフロー

`.github/workflows/backend-deploy.yml` ファイルを作成します。このワークフローは、バックエンドのコードに変更があった場合に自動的に実行されます。

> ファイルの内容は `backend-deploy.yml` Artifactを参照してください。

## 4. S3バケットへのデプロイスクリプトの作成

フロントエンドやバックエンドのデプロイ用スクリプトを作成します。

### フロントエンドデプロイスクリプト

`scripts/deploy-frontend.sh` ファイルを作成します。

> ファイルの内容は `deploy-frontend.sh` Artifactを参照してください。

作成後、実行権限を付与します：

```bash
chmod +x scripts/deploy-frontend.sh
```

### バックエンドデプロイスクリプト

`scripts/deploy-backend.sh` ファイルを作成します。

> ファイルの内容は `deploy-backend.sh` Artifactを参照してください。

作成後、実行権限を付与します：

```bash
chmod +x scripts/deploy-backend.sh
```

## 5. サンプルフロントエンドプロジェクトの準備

CI/CDパイプラインをテストするために、簡単なNext.jsプロジェクトを準備します。

### Next.jsプロジェクトの初期化

```bash
cd frontend
npx create-next-app@latest . --typescript --eslint --tailwind --app --src-dir

# 依存関係のインストール
npm install

# 開発サーバーの起動（確認用）
npm run dev
```

### 基本的なページの作成

`frontend/src/app/page.tsx` ファイルを編集して、簡単なページを作成します。

> ファイルの内容は `page.tsx` Artifactを参照してください。

## 6. サンプルバックエンドサービスの準備

バックエンドサービスとして、Node.jsベースの簡単なLambda関数を作成します。

### カタログサービスのLambda関数

`backend/catalog-service/index.js` ファイルを作成します。

> ファイルの内容は `catalog-service-index.js` Artifactを参照してください。

### パッケージ設定

`backend/catalog-service/package.json` ファイルを作成します。

> ファイルの内容は `catalog-service-package.json` Artifactを参照してください。

## 7. Terraformでのバックエンドリソース定義

バックエンドサービス用のLambda関数とAPI Gateway統合を定義するTerraformファイルを作成します。

### `terraform/lambda.tf` ファイルの作成

> ファイルの内容は `lambda.tf` Artifactを参照してください。

### `terraform/api_routes.tf` ファイルの作成

> ファイルの内容は `api_routes.tf` Artifactを参照してください。

## 8. Terraformコードの適用

Terraformコードを適用して、バックエンドサービスのインフラをプロビジョニングします。

```bash
cd terraform
terraform validate
terraform plan
terraform apply
```

## 9. DynamoDBテーブルへのサンプルデータ投入

作成したDynamoDBテーブルにサンプルデータを投入するスクリプトを作成します。

### データ投入スクリプト

`scripts/seed-dynamodb.js` ファイルを作成します。

> ファイルの内容は `seed-dynamodb.js` Artifactを参照してください。

スクリプトを実行します：

```bash
cd scripts
npm install aws-sdk
node seed-dynamodb.js
```

## 10. 初回デプロイとパイプラインのテスト

リポジトリに変更をプッシュして、CI/CDパイプラインが正常に動作することを確認します。

```bash
git add .
git commit -m "Setup CI/CD pipeline and basic infrastructure"
git push origin main
```

GitHubリポジトリの「Actions」タブで、ワークフローの実行状況を確認します。

## 11. デプロイの確認

### フロントエンド

CloudFrontディストリビューションのドメイン名にアクセスして、フロントエンドが正常にデプロイされていることを確認します。

```bash
# CloudFrontドメイン名を取得
cd terraform
CLOUDFRONT_DOMAIN=$(terraform output -raw cloudfront_domain)
echo "フロントエンドURL: https://${CLOUDFRONT_DOMAIN}"
```

### バックエンドAPI

API Gatewayのエンドポイントにアクセスして、バックエンドAPIが正常に動作していることを確認します。

```bash
# API Gatewayエンドポイントを取得
cd terraform
API_ENDPOINT=$(terraform output -raw api_endpoint)
echo "バックエンドAPIエンドポイント: ${API_ENDPOINT}"

# カテゴリ一覧のAPIを呼び出し
curl "${API_ENDPOINT}/categories"

# 商品一覧のAPIを呼び出し
curl "${API_ENDPOINT}/products"
```

## 次のステップ

Day 2のCI/CDと基本インフラのセットアップが完了しました。Day 3では、バックエンドの基本実装とオブザーバビリティの設定に進みます。以下の準備をしておくとよいでしょう：

1. フロントエンドとバックエンドの連携テスト
2. CloudWatchログの確認方法の理解
3. API GatewayとLambdaのモニタリング設定の確認

これでDay 2の「CI/CDと基本インフラ」のセットアップが完了しました。
