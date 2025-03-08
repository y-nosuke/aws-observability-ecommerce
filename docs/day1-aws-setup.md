# Day 1: AWS環境セットアップ手順

このガイドでは、AWSオブザーバビリティ学習用eコマースアプリの環境セットアップを行います。

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
