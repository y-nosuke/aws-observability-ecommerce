# API Gateway エンドポイント出力
output "api_endpoint" {
  description = "API Gateway Endpoint URL"
  value       = aws_apigatewayv2_api.main.api_endpoint
}

# Lambda関数出力
output "catalog_lambda_function_name" {
  description = "Catalog Service Lambda Function Name"
  value       = aws_lambda_function.catalog_service.function_name
}

output "catalog_lambda_function_arn" {
  description = "Catalog Service Lambda Function ARN"
  value       = aws_lambda_function.catalog_service.arn
}

# フロントエンド関連の出力
output "frontend_bucket_name" {
  description = "Frontend S3 Bucket Name"
  value       = aws_s3_bucket.frontend.bucket
}

output "cloudfront_distribution_id" {
  description = "CloudFront Distribution ID"
  value       = aws_cloudfront_distribution.frontend.id
}

output "cloudfront_domain" {
  description = "CloudFront Domain Name"
  value       = aws_cloudfront_distribution.frontend.domain_name
}
