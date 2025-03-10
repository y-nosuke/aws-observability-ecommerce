# Day 2: CI/CDと基本インフラのセットアップ手順

このガイドでは、Day 1で作成したAWS環境の上にCI/CDパイプラインを構築し、基本的なサービスリソースをプロビジョニングします。

## 前提条件

- Day 1の作業が完了しており、基本的なAWSリソース（VPC、サブネット、セキュリティグループなど）が作成済みであること
- GitHubアカウントを持っていること
- GitHubリポジトリ `aws-observability-ecommerce` が作成済みであること

## 料金の目安

この実装で1日稼働させた場合のAWS料金の目安は以下の通りです：

| サービス    | 想定使用量                                          | 概算費用(USD)/日     |
| ----------- | --------------------------------------------------- | -------------------- |
| CloudFront  | データ転送 1GB, リクエスト10,000件                  | $0.10 - $0.15        |
| S3          | ストレージ 1GB, リクエスト5,000件                   | $0.03 - $0.05        |
| API Gateway | リクエスト10,000件                                  | $0.35 - $0.40        |
| Lambda      | 10,000回の呼び出し (128MB, 平均実行時間200ms)       | $0.04 - $0.06        |
| DynamoDB    | 読み取り/書き込み各5,000リクエスト, ストレージ500MB | $0.05 - $0.10        |
| VPC         | NAT Gateway使用時間24時間                           | $0.25 - $0.30        |
| CloudWatch  | 標準モニタリング, 5GB のログ                        | $0.10 - $0.15        |
| **合計**    |                                                     | **$0.92 - $1.21/日** |

注意:

- この料金は目安であり、実際の使用状況によって変動します
- フリーティア対象のサービスや使用量であれば、実質無料になる項目もあります
- 開発/テスト用途の軽い負荷であれば、月に$10-30程度で運用できる可能性があります

## 1. リポジトリ構造の整備

まず、プロジェクトのリポジトリ構造を整備しましょう。

```bash
# プロジェクトのルートディレクトリで以下のコマンドを実行
mkdir -p frontend
mkdir -p backend
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

### カタログサービスのプロジェクト作成

```bash
# バックエンドディレクトリに移動
cd backend

# カタログサービスのディレクトリを作成
mkdir -p catalog-service

# カタログサービスディレクトリに移動
cd catalog-service

# package.jsonの初期化
npm init -y

# 必要な依存関係のインストール
npm install aws-sdk aws-xray-sdk-core

# プロジェクトルートディレクトリに戻る
cd ../..
```

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

### データ投入用環境のセットアップ

```bash
# scriptsディレクトリに移動
cd scripts

# package.jsonの初期化
npm init -y

# AWS SDK v3をインストール（最新のSDKを使用）
npm install @aws-sdk/client-dynamodb @aws-sdk/lib-dynamodb
```

### データ投入スクリプト

`scripts/seed-dynamodb.js` ファイルを作成します。

> ファイルの内容は `seed-dynamodb.js` Artifactを参照してください。

### スクリプトの実行

```bash
# scriptsディレクトリで
node seed-dynamodb.js
```

または、package.jsonにスクリプトエントリを追加して実行することもできます：

```bash
# package.jsonに以下の内容を追加
echo '{
  "scripts": {
    "seed": "node seed-dynamodb.js"
  }
}' > package.json

# スクリプトを実行
npm run seed
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
