terraform {
  backend "s3" {
    bucket         = "aws-observability-ecommerce-terraform-state"
    key            = "dev/terraform.tfstate"
    region         = "us-east-1"
    dynamodb_table = "aws-observability-ecommerce-terraform-locks"
    encrypt        = true
  }
}
