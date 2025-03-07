# terraform/deployment.tf

# フロントエンドデプロイ用S3バケット
resource "aws_s3_bucket" "frontend_deploy" {
  bucket = "ecommerce-frontend-deploy-${random_string.bucket_suffix.result}"
  force_destroy = true
}

# ウェブサイトホスティング設定
resource "aws_s3_bucket_website_configuration" "frontend_website" {
  bucket = aws_s3_bucket.frontend_deploy.id

  index_document {
    suffix = "index.html"
  }

  error_document {
    key = "error.html"
  }
}

# バケットポリシー設定 (CloudFrontからのアクセスを許可)
resource "aws_s3_bucket_policy" "frontend_policy" {
  bucket = aws_s3_bucket.frontend_deploy.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          AWS = aws_cloudfront_origin_access_identity.oai.iam_arn
        }
        Action = "s3:GetObject"
        Resource = "${aws_s3_bucket.frontend_deploy.arn}/*"
      }
    ]
  })
}

# SAM/CloudFormationテンプレート用S3バケット
resource "aws_s3_bucket" "cfn_templates" {
  bucket = "ecommerce-cfn-templates-${random_string.bucket_suffix.result}"
  force_destroy = true
}

# ランダムサフィックス (バケット名の一意性確保)
resource "random_string" "bucket_suffix" {
  length  = 8
  special = false
  upper   = false
}

# SSMパラメータストアにバケット名とCloudFront IDを格納
resource "aws_ssm_parameter" "frontend_bucket" {
  name  = "/ecommerce/frontend/bucket-name"
  type  = "String"
  value = aws_s3_bucket.frontend_deploy.id
}

resource "aws_ssm_parameter" "cloudfront_id" {
  name  = "/ecommerce/frontend/cloudfront-id"
  type  = "String"
  value = aws_cloudfront_distribution.frontend_distribution.id
}
