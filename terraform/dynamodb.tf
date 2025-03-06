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
