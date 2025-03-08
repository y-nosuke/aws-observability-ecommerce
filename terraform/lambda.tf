# Lambda関数のデプロイパッケージ用のS3バケット
resource "aws_s3_bucket" "lambda_bucket" {
  bucket = "${var.project_name}-lambda-${var.environment}"

  tags = {
    Name = "${var.project_name}-lambda-bucket"
  }
}

# バケットパブリックアクセス設定（パブリックアクセスをブロック）
resource "aws_s3_bucket_public_access_block" "lambda_bucket_access" {
  bucket = aws_s3_bucket.lambda_bucket.id

  block_public_acls       = true
  block_public_policy     = true
  ignore_public_acls      = true
  restrict_public_buckets = true
}

# カタログサービス用のCloudWatch Logsロググループ
resource "aws_cloudwatch_log_group" "catalog_service_logs" {
  name              = "/aws/lambda/${var.project_name}-catalog-service-${var.environment}"
  retention_in_days = 7

  tags = {
    Name = "${var.project_name}-catalog-service-logs"
  }
}

# カタログサービス用のLambda関数
resource "aws_lambda_function" "catalog_service" {
  function_name = "${var.project_name}-catalog-service-${var.environment}"
  description   = "Catalog service for e-commerce application"

  role    = aws_iam_role.lambda_execution.arn
  handler = "index.handler"
  runtime = "nodejs18.x"

  # デフォルトのZIPファイル（後で更新される）
  filename = "${path.module}/lambda_dummy.zip"

  memory_size = 256
  timeout     = 10

  environment {
    variables = {
      PROJECT_PREFIX = var.project_name
      ENVIRONMENT    = var.environment
    }
  }

  # X-Ray トレーシングの有効化
  tracing_config {
    mode = "Active"
  }

  # 関数URLの設定（オプション - API Gateway経由でアクセスする場合は不要）
  # reserved_concurrent_executions = 10

  depends_on = [
    aws_cloudwatch_log_group.catalog_service_logs,
    aws_iam_role_policy_attachment.lambda_logs,
    aws_iam_role_policy_attachment.lambda_xray,
    aws_iam_role_policy_attachment.lambda_dynamodb
  ]

  tags = {
    Name = "${var.project_name}-catalog-service"
  }
}

# Lambda用ダミーアーカイブ作成（初回apply用）
resource "local_file" "lambda_dummy" {
  filename = "${path.module}/lambda_dummy.js"
  content  = "exports.handler = async (event) => { return { statusCode: 200, body: JSON.stringify({ message: 'Dummy function' }) }; };"
}

data "archive_file" "lambda_dummy_zip" {
  type        = "zip"
  output_path = "${path.module}/lambda_dummy.zip"
  source_file = local_file.lambda_dummy.filename
  depends_on  = [local_file.lambda_dummy]
}
