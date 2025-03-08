# フロントエンドホスティング用S3バケット
resource "aws_s3_bucket" "frontend" {
  bucket = "${var.project_name}-frontend-${var.environment}"

  force_destroy = true

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

# S3バケットポリシー - CloudFrontからのアクセスを許可
resource "aws_s3_bucket_policy" "frontend" {
  bucket = aws_s3_bucket.frontend.id
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "AllowCloudFrontServicePrincipal"
        Effect = "Allow"
        Principal = {
          Service = "cloudfront.amazonaws.com"
        }
        Action   = "s3:GetObject"
        Resource = "${aws_s3_bucket.frontend.arn}/*"
        Condition = {
          StringEquals = {
            "AWS:SourceArn" = aws_cloudfront_distribution.frontend.arn
          }
        }
      }
    ]
  })
}

# デプロイアーティファクト保存用S3バケット
resource "aws_s3_bucket" "deployment" {
  bucket = "${var.project_name}-deployment-${var.environment}"

  force_destroy = true

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

# S3バケットのサーバーアクセスログ（オプション）
resource "aws_s3_bucket" "s3_logs" {
  count  = var.enable_s3_logs ? 1 : 0
  bucket = "${var.project_name}-s3-logs-${var.environment}"

  force_destroy = true

  tags = {
    Name = "${var.project_name}-s3-logs"
  }
}

# バケット所有権設定 - ObjectWriter (ACLを有効にする)
resource "aws_s3_bucket_ownership_controls" "s3_logs_ownership" {
  count  = var.enable_s3_logs ? 1 : 0
  bucket = aws_s3_bucket.s3_logs[0].id

  rule {
    object_ownership = "ObjectWriter"
  }
}

# バケットACL設定 - ログ配信が可能になるように
resource "aws_s3_bucket_acl" "s3_logs_acl" {
  count  = var.enable_s3_logs ? 1 : 0
  bucket = aws_s3_bucket.s3_logs[0].id
  acl    = "log-delivery-write"

  # 所有権設定に依存
  depends_on = [aws_s3_bucket_ownership_controls.s3_logs_ownership]
}

# フロントエンドバケットにサーバーアクセスログを設定
resource "aws_s3_bucket_logging" "frontend_logging" {
  count  = var.enable_s3_logs ? 1 : 0
  bucket = aws_s3_bucket.frontend.id

  target_bucket = aws_s3_bucket.s3_logs[0].id
  target_prefix = "s3-access-logs/"
}
