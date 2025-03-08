# CloudFrontディストリビューション
resource "aws_cloudfront_distribution" "frontend" {
  enabled             = true
  is_ipv6_enabled     = true
  default_root_object = "index.html"
  price_class         = "PriceClass_100" # 北米・欧州・アジアの低価格リージョンのみ使用
  comment             = "${var.project_name}-frontend"

  # S3バケットをオリジンとして設定
  origin {
    domain_name              = aws_s3_bucket.frontend.bucket_regional_domain_name
    origin_id                = "S3-${aws_s3_bucket.frontend.bucket}"
    origin_access_control_id = aws_cloudfront_origin_access_control.frontend.id
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

  # アクセスログの設定（オプション）
  dynamic "logging_config" {
    for_each = var.enable_cloudfront_logs ? [1] : []
    content {
      include_cookies = false
      bucket          = aws_s3_bucket.cloudfront_logs[0].bucket_domain_name
      prefix          = "cloudfront/"
    }
  }

  tags = {
    Name = "${var.project_name}-frontend-distribution"
  }
}

# CloudFrontオリジンアクセスコントロール
resource "aws_cloudfront_origin_access_control" "frontend" {
  name                              = "${var.project_name}-frontend-oac"
  description                       = "OAC for ${var.project_name} frontend"
  origin_access_control_origin_type = "s3"
  signing_behavior                  = "always"
  signing_protocol                  = "sigv4"
}

# CloudFrontアクセスログ用のS3バケット（オプション）
resource "aws_s3_bucket" "cloudfront_logs" {
  count  = var.enable_cloudfront_logs ? 1 : 0
  bucket = "${var.project_name}-cloudfront-logs-${var.environment}"

  force_destroy = true

  tags = {
    Name = "${var.project_name}-cloudfront-logs"
  }
}

# バケット所有権設定 - ObjectWriter (ACLを有効にする)
resource "aws_s3_bucket_ownership_controls" "cloudfront_logs_ownership" {
  count  = var.enable_cloudfront_logs ? 1 : 0
  bucket = aws_s3_bucket.cloudfront_logs[0].id

  rule {
    object_ownership = "ObjectWriter"
  }
}

# バケットACL設定 - CloudFrontのログ配信が可能になるように
resource "aws_s3_bucket_acl" "cloudfront_logs_acl" {
  count  = var.enable_cloudfront_logs ? 1 : 0
  bucket = aws_s3_bucket.cloudfront_logs[0].id
  acl    = "private"

  # 所有権設定に依存
  depends_on = [aws_s3_bucket_ownership_controls.cloudfront_logs_ownership]
}
