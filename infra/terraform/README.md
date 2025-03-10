# Terraform インフラストラクチャコード

このディレクトリには、AWSオブザーバビリティ学習用eコマースアプリケーションのインフラストラクチャを定義するTerraformコードが含まれています。

## 前提条件

- Terraform 1.0.0以上がインストールされていること
- AWS CLIがインストールされ、適切に設定されていること
- 必要なAWS権限があること

## ディレクトリ構造

```text
terraform/
├── environments/            # 環境固有の設定
│   ├── dev/                 # 開発環境
│   ├── staging/             # ステージング環境
│   └── prod/                # 本番環境
├── modules/                 # 再利用可能なモジュール
│   ├── vpc/                 # VPCモジュール
│   ├── ecr/                 # ECRリポジトリモジュール
│   ├── s3/                  # S3バケットモジュール
│   └── iam/                 # IAMロールとポリシーモジュール
└── versions.tf              # Terraformとプロバイダーのバージョン設定
```

## 使用方法

### 開発環境のデプロイ

```bash
# 開発環境ディレクトリに移動
cd environments/dev

# Terraformの初期化
terraform init

# 実行計画の作成
terraform plan

# インフラストラクチャのデプロイ
terraform apply
```

### リモート状態管理の設定（オプション）

本格的な環境では、チーム間でTerraformの状態を共有するために、S3バケットとDynamoDBテーブルを使用します。

```bash
# S3バケットの作成
aws s3api create-bucket \
  --bucket aws-observability-ecommerce-terraform-state \
  --region us-east-1

# S3バケットのバージョニングを有効化
aws s3api put-bucket-versioning \
  --bucket aws-observability-ecommerce-terraform-state \
  --versioning-configuration Status=Enabled

# S3バケットの暗号化を有効化
aws s3api put-bucket-encryption \
  --bucket aws-observability-ecommerce-terraform-state \
  --server-side-encryption-configuration '{
    "Rules": [
      {
        "ApplyServerSideEncryptionByDefault": {
          "SSEAlgorithm": "AES256"
        }
      }
    ]
  }'

# DynamoDBテーブルの作成（状態ロック用）
aws dynamodb create-table \
  --table-name aws-observability-ecommerce-terraform-locks \
  --attribute-definitions AttributeName=LockID,AttributeType=S \
  --key-schema AttributeName=LockID,KeyType=HASH \
  --billing-mode PAY_PER_REQUEST \
  --region us-east-1
```

`backend.tf`ファイルに以下の内容を追加し、リモート状態管理を有効にします:

```hcl
terraform {
  backend "s3" {
    bucket         = "aws-observability-ecommerce-terraform-state"
    key            = "dev/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "aws-observability-ecommerce-terraform-locks"
    encrypt        = true
  }
}
```

## 注意事項

- 開発環境では、コスト削減のためにAWSリソースを必要最小限に保ちます。
- フェーズ4で本格的なAWS環境のデプロイを行う際は、追加のリソースとモジュールを実装します。
- 本番環境用の設定（高可用性やセキュリティ強化など）は、実際の運用に応じて調整してください。
