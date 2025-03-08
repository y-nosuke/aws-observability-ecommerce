# Lambda関数のAPI Gateway統合

# CORS設定は、HTTP APIではAPI Gateway自体の設定として処理される
# (すでにapigateway.tfファイルのaws_apigatewayv2_api.mainリソースでCORS設定を行っています)

# ヘルスチェックエンドポイント
resource "aws_apigatewayv2_route" "health" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /health"
  target    = "integrations/${aws_apigatewayv2_integration.catalog_service.id}"
}

# 商品一覧エンドポイント
resource "aws_apigatewayv2_route" "products" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /products"
  target    = "integrations/${aws_apigatewayv2_integration.catalog_service.id}"
}

# 商品詳細エンドポイント
resource "aws_apigatewayv2_route" "product_detail" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /products/{id}"
  target    = "integrations/${aws_apigatewayv2_integration.catalog_service.id}"
}

# カテゴリ一覧エンドポイント
resource "aws_apigatewayv2_route" "categories" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /categories"
  target    = "integrations/${aws_apigatewayv2_integration.catalog_service.id}"
}

# カテゴリ詳細エンドポイント
resource "aws_apigatewayv2_route" "category_detail" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /categories/{id}"
  target    = "integrations/${aws_apigatewayv2_integration.catalog_service.id}"
}

# カテゴリに属する商品一覧エンドポイント
resource "aws_apigatewayv2_route" "category_products" {
  api_id    = aws_apigatewayv2_api.main.id
  route_key = "GET /categories/{id}/products"
  target    = "integrations/${aws_apigatewayv2_integration.catalog_service.id}"
}

# カタログサービスLambda統合
resource "aws_apigatewayv2_integration" "catalog_service" {
  api_id           = aws_apigatewayv2_api.main.id
  integration_type = "AWS_PROXY"

  connection_type      = "INTERNET"
  description          = "Catalog Service Lambda Integration"
  integration_method   = "POST"
  integration_uri      = aws_lambda_function.catalog_service.invoke_arn
  passthrough_behavior = "WHEN_NO_MATCH"

  payload_format_version = "2.0"
  timeout_milliseconds   = 30000
}

# Lambda関数のAPI Gateway呼び出し許可
resource "aws_lambda_permission" "api_gateway" {
  statement_id  = "AllowExecutionFromAPIGateway"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.catalog_service.function_name
  principal     = "apigateway.amazonaws.com"

  # API Gateway ARN: arn:aws:execute-api:{regionId}:{accountId}:{apiId}/*/{httpMethod}/{resource}
  source_arn = "${aws_apigatewayv2_api.main.execution_arn}/*/*"
}
