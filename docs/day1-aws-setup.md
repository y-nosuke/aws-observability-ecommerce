# Day 1: AWS環境セットアップ手順

このガイドでは、AWSオブザーバビリティ学習用eコマースアプリの環境セットアップを行います。

## 1. AWS アカウント準備

以下の項目を確認・設定してください：

- AWS IAMユーザー作成（管理者権限）
- アクセスキー/シークレットキーの安全な管理設定
- MFAの有効化

## 2. Terraformのインストールと設定

### Terraformのインストール

**MacOS (Homebrew):**
```bash
brew install terraform
```

**Windows (Chocolatey):**
```bash
choco install terraform
```

**Linux:**
```bash
wget https://releases.hashicorp.com/terraform/1.7.5/terraform_1.7.5_linux_amd64.zip
unzip terraform_1.7.5_linux_amd64.zip
sudo mv terraform /usr/local/bin/
```

### バージョン確認
```bash
terraform version
```

## 3. プロジェクト構造の作成

```bash
mkdir aws-observability-ecommerce
cd aws-observability-ecommerce
mkdir terraform
cd terraform
```

## 4. 基本的なTerraformファイルの作成

次のTerraformファイルを作成します。各ファイルのコードは個別のArtifactで提供しています。

- `main.tf` - プロバイダーとバージョン設定
- `variables.tf` - 変数定義
- `network.tf` - VPCとネットワーク設定
- `security.tf` - セキュリティグループ設定
- `outputs.tf` - 出力値定義
- `iam.tf` - IAMロールとポリシー
- `storage.tf` - S3バケット設定
- `cloudfront.tf` - CloudFront配信設定
- `apigateway.tf` - API Gateway設定
- `dynamodb.tf` - DynamoDBテーブル設定

## 5. GitHubリポジトリの作成

1. GitHub上で新しいリポジトリを作成します（例: `aws-observability-ecommerce`）
2. ローカルのプロジェクトをリポジトリに初期化してプッシュします:

```bash
# プロジェクトルートディレクトリで:
git init
git add .
git commit -m "Initial commit with Terraform infrastructure code"
git branch -M main
git remote add origin https://github.com/YOUR_USERNAME/aws-observability-ecommerce.git
git push -u origin main
```

## 6. Terraformの実行と基本リソースのプロビジョニング

```bash
# 初期化
cd terraform
terraform init

# 構文チェック
terraform validate

# 変更内容のプレビュー
terraform plan

# リソースのプロビジョニング
terraform apply
```

`apply`コマンドを実行するとリソース作成の確認プロンプトが表示されますので、確認して「yes」と入力します。

## 7. リソース作成の確認

AWS Management Consoleにログインして、作成したリソースを確認します：

- VPCとサブネット
- セキュリティグループ
- IAMロール
- S3バケット
- CloudFrontディストリビューション
- API Gateway
- DynamoDBテーブル

これでDay 1の「AWS環境セットアップ」が完了しました。次のDay 2では、CI/CDパイプラインの構築と基本サービスリソースのプロビジョニングに進みます。
