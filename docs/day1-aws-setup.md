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

### `main.tf` - プロバイダーとバージョン設定

```hcl
provider "aws" {
  region = "ap-northeast-1"  # 東京リージョン（任意のリージョンに変更可能）
  
  default_tags {
    tags = {
      Project     = "ecommerce-observability"
      Environment = "dev"
      ManagedBy   = "terraform"
    }
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
  
  required_version = ">= 1.0.0"
}
```

### `variables.tf` - 変数定義

```hcl
variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "ecommerce-observability"
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}
```

### `network.tf` - VPCとネットワーク設定

```hcl
# VPC定義
resource "aws_vpc" "main" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    Name = "ecommerce-vpc"
  }
}

# パブリックサブネット (2つのアベイラビリティーゾーンに)
resource "aws_subnet" "public" {
  count             = 2
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.${count.index + 1}.0/24"
  availability_zone = data.aws_availability_zones.available.names[count.index]
  
  map_public_ip_on_launch = true

  tags = {
    Name = "ecommerce-public-subnet-${count.index + 1}"
  }
}

# プライベートサブネット (2つのアベイラビリティーゾーンに)
resource "aws_subnet" "private" {
  count             = 2
  vpc_id            = aws_vpc.main.id
  cidr_block        = "10.0.${count.index + 10}.0/24"
  availability_zone = data.aws_availability_zones.available.names[count.index]
  
  tags = {
    Name = "ecommerce-private-subnet-${count.index + 1}"
  }
}

# インターネットゲートウェイ
resource "aws_internet_gateway" "main" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "ecommerce-igw"
  }
}

# パブリックルートテーブル
resource "aws_route_table" "public" {
  vpc_id = aws_vpc.main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.main.id
  }

  tags = {
    Name = "ecommerce-public-rt"
  }
}

# パブリックサブネットとルートテーブルの関連付け
resource "aws_route_table_association" "public" {
  count          = 2
  subnet_id      = aws_subnet.public[count.index].id
  route_table_id = aws_route_table.public.id
}

# プライベートルートテーブル
resource "aws_route_table" "private" {
  vpc_id = aws_vpc.main.id

  tags = {
    Name = "ecommerce-private-rt"
  }
}

# プライベートサブネットとルートテーブルの関連付け
resource "aws_route_table_association" "private" {
  count          = 2
  subnet_id      = aws_subnet.private[count.index].id
  route_table_id = aws_route_table.private.id
}

# 利用可能なアベイラビリティーゾーンを取得
data "aws_availability_zones" "available" {
  state = "available"
}
```

### `security.tf` - セキュリティグループ設定

```hcl
# APIゲートウェイ/ALB用セキュリティグループ
resource "aws_security_group" "api" {
  name        = "ecommerce-api-sg"
  description = "Security group for API Gateway/ALB"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port   = 80
    to_port     = 80
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow HTTP from anywhere"
  }
  
  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow HTTPS from anywhere"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name = "ecommerce-api-sg"
  }
}

# アプリケーションサービス用セキュリティグループ
resource "aws_security_group" "app" {
  name        = "ecommerce-app-sg"
  description = "Security group for application services"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 0
    to_port         = 65535
    protocol        = "tcp"
    security_groups = [aws_security_group.api.id]
    description     = "Allow all traffic from API security group"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name = "ecommerce-app-sg"
  }
}

# データベース用セキュリティグループ
resource "aws_security_group" "db" {
  name        = "ecommerce-db-sg"
  description = "Security group for databases"
  vpc_id      = aws_vpc.main.id

  ingress {
    from_port       = 0
    to_port         = 65535
    protocol        = "tcp"
    security_groups = [aws_security_group.app.id]
    description     = "Allow all traffic from application security group"
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
    description = "Allow all outbound traffic"
  }

  tags = {
    Name = "ecommerce-db-sg"
  }
}
```

### `outputs.tf` - 出力値定義

```hcl
output "vpc_id" {
  description = "The ID of the VPC"
  value       = aws_vpc.main.id
}

output "public_subnet_ids" {
  description = "The IDs of the public subnets"
  value       = aws_subnet.public[*].id
}

output "private_subnet_ids" {
  description = "The IDs of the private subnets"
  value       = aws_subnet.private[*].id
}

output "api_security_group_id" {
  description = "The ID of the API security group"
  value       = aws_security_group.api.id
}

output "app_security_group_id" {
  description = "The ID of the application security group"
  value       = aws_security_group.app.id
}

output "db_security_group_id" {
  description = "The ID of the database security group"
  value       = aws_security_group.db.id
}
```

### `iam.tf` - IAMロールとポリシー

```hcl
# Lambda実行用の基本IAMロール
resource "aws_iam_role" "lambda_execution" {
  name = "ecommerce-lambda-execution-role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          Service = "lambda.amazonaws.com"
        }
      }
    ]
  })

  tags = {
    Name = "ecommerce-lambda-execution-role"
  }
}

# Lambda実行に必要な基本ポリシーをアタッチ
resource "aws_iam_role_policy_attachment" "lambda_basic" {
  role       = aws_iam_role.lambda_execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}

# CloudWatch Logsへの書き込み権限
resource "aws_iam_role_policy_attachment" "lambda_logs" {
  role       = aws_iam_role.lambda_execution.name
  policy_arn = "arn:aws:iam::aws:policy/CloudWatchLogsFullAccess"
}

# X-Rayトレース権限
resource "aws_iam_role_policy_attachment" "lambda_xray" {
  role       = aws_iam_role.lambda_execution.name
  policy_arn = "arn:aws:iam::aws:policy/AWSXrayWriteOnlyAccess"
}

# DynamoDBアクセス権限 (後でカスタムポリシーに置き換えるのがベストプラクティス)
resource "aws_iam_role_policy_attachment" "lambda_dynamodb" {
  role       = aws_iam_role.lambda_execution.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
}
```

### `storage.tf` - S3バケット設定

```hcl
# フロントエンドホスティング用S3バケット
resource "aws_s3_bucket" "frontend" {
  bucket = "${var.project_name}-frontend-${var.environment}"

  tags = {
    Name = "${var.project_name}-frontend"
  }
}

# バケットパブリックアクセス設定（CloudFrontから配信するため、パブリックアクセスはブロック）
resource "aws_s3_bucket_public_access_block" "frontend_access" {
  bucket = aws_s3_bucket.frontend.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# バケットのウェブサイト設定
resource "aws_s3_bucket_website_configuration" "frontend_website" {
  bucket = aws_s3_bucket.frontend.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

# デプロイアーティファクト保存用S3バケット
resource "aws_s3_bucket" "deployment" {
  bucket = "${var.project_name}-deployment-${var.environment}"

  tags = {
    Name = "${var.project_name}-deployment"
  }
}

# バケットパブリックアクセス設定（デプロイバケットのパブリックアクセスをブロック）
resource "aws_s3_bucket_public_access_block" "deployment_access" {
  bucket = aws_s3_bucket.deployment.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}
```

### `cloudfront.tf` - CloudFront配信設定

```hcl
# CloudFrontディストリビューション
resource "aws_cloudfront_distribution" "frontend" {
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"
  price_class         = "PriceClass_100" # 北米・欧州・アジアの低価格リージョンのみ使用

  # S3バケットをオリジンとして設定
  origin {
    domain_name = aws_s3_bucket.frontend.bucket_regional_domain_name
    origin_id   = "S3-${aws_s3_bucket.frontend.bucket}"

    s3_origin_config {
      origin_access_identity = aws_cloudfront_origin_access_identity.frontend.cloudfront_access_identity_path
    }
  }

  # デフォルトのキャッシュ動作
  default_cache_behavior {
    allowed_methods  = ["GET", "HEAD", "OPTIONS"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = "S3-${aws_s3_bucket.frontend.bucket}"

    # フォワードヘッダ
    forwarded_values {
      query_string = false
      cookies {
        forward = "none"
      }
    }

    viewer_protocol_policy = "redirect-to-https"
    min_ttl                = 0
    default_ttl            = 3600  # 1時間
    max_ttl                = 86400 # 1日
    compress               = true
  }

  # SPAルーティングのためのカスタムエラーレスポンス（React/Next.jsなどのSPAで必要）
  custom_error_response {
    error_code            = 403
    response_code         = 200
    response_page_path    = "/index.html"
    error_caching_min_ttl = 10
  }

  custom_error_response {
    error_code            = 404
    response_code         = 200
    response_page_path    = "/index.html"
    error_caching_min_ttl = 10
  }

  # 地理的制限なし
  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  # デフォルトの証明書を使用
  viewer_certificate {
    cloudfront_default_certificate = true
  }

  tags = {
    Name = "${var.project_name}-frontend-distribution"
  }
}

# CloudFront Origin Access Identity
resource "aws_cloudfront_origin_access_identity" "frontend" {
  comment = "OAI for ${var.project_name} frontend"
}

# S3バケットポリシー - CloudFrontからのアクセスを許可
resource "aws_s3_bucket_policy" "frontend" {
  bucket = aws_s3_bucket.frontend.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid       = "AllowCloudFrontServicePrincipal"
        Effect    = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action    = "s3:GetObject"
        Resource  = "${aws_s3_bucket.frontend.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = aws_cloudfront_distribution.frontend.arn
          }
        }
      }
    ]
  })
}
```

### `apigateway.tf` - API Gateway設定

```hcl
# HTTP API Gateway
resource "aws_apigatewayv2_api" "main" {
  name          = "${var.project_name}-api"
  protocol_type = "HTTP"
  
  cors_configuration {
    allow_origins = ["https://${aws_cloudfront_distribution.frontend.domain_name}", "http://localhost:3000"]
    allow_methods = ["GET", "POST", "PUT", "DELETE", "OPTIONS"]
    allow_headers = ["Content-Type", "Authorization", "X-Amz-Date", "X-Api-Key", "X-Amz-Security-Token"]
    max_age       = 300
  }

  tags = {
    Name = "${var.project_name}-api"
  }
}

# API Gateway ステージの設定
resource "aws_apigatewayv2_stage" "default" {
  api_id      = aws_apigatewayv2_api.main.id
  name        = "$default"
  auto_deploy = true

  access_log_settings {
    destination_arn = aws_cloudwatch_log_group.api_gateway.arn
    format = jsonencode({
      requestId      = "$context.requestId"
      ip             = "$context.identity.sourceIp"
      requestTime    = "$context.requestTime"
      httpMethod     = "$context.httpMethod"
      routeKey       = "$context.routeKey"
      status         = "$context.status"
      protocol       = "$context.protocol"
      responseLength = "$context.responseLength"
      errorMessage   = "$context.error.message"
      integrationError = "$context.integration.error"
      latency        = "$context.responseLatency"
      integrationLatency = "$context.integration.latency"
    })
  }

  default_route_settings {
    throttling_burst_limit = 100
    throttling_rate_limit  = 50
  }

  tags = {
    Name = "${var.project_name}-default-stage"
  }
}

# ログ設定
resource "aws_cloudwatch_log_group" "api_gateway" {
  name              = "/aws/apigateway/${var.project_name}-api"
  retention_in_days = 7

  tags = {
    Name = "${var.project_name}-api-logs"
  }
}
```

### `dynamodb.tf` - DynamoDBテーブル設定

```hcl
# 商品テーブル
resource "aws_dynamodb_table" "products" {
  name         = "${var.project_name}-products"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }

  attribute {
    name = "category_id"
    type = "S"
  }

  global_secondary_index {
    name               = "CategoryIndex"
    hash_key           = "category_id"
    projection_type    = "ALL"
  }

  tags = {
    Name = "${var.project_name}-products-table"
  }
}

# カテゴリーテーブル
resource "aws_dynamodb_table" "categories" {
  name         = "${var.project_name}-categories"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "id"

  attribute {
    name = "id"
    type = "S"
  }

  tags = {
    Name = "${var.project_name}-categories-table"
  }
}

# 商品閲覧履歴テーブル
resource "aws_dynamodb_table" "product_views" {
  name         = "${var.project_name}-product-views"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "user_id"
  range_key    = "timestamp"

  attribute {
    name = "user_id"
    type = "S"
  }

  attribute {
    name = "timestamp"
    type = "S"
  }

  attribute {
    name = "product_id"
    type = "S"
  }

  global_secondary_index {
    name               = "ProductViewsIndex"
    hash_key           = "product_id"
    range_key          = "timestamp"
    projection_type    = "ALL"
  }

  tags = {
    Name = "${var.project_name}-product-views-table"
  }
}
```

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
