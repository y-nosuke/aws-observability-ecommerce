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
