# 注意: バケットは事前に作成しておく必要があります
terraform {
  backend "s3" {
    bucket  = "jp.physicist00.terraform"
    key     = "ecommerce-observability/terraform.tfstate"
    region  = "ap-northeast-1"
    encrypt = true
    # dynamodb_table = "terraform-state-lock"
  }
}
