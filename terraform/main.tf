provider "aws" {
  region = "ap-northeast-1"  # 東京リージョン（任意のリージョンに変更可能）

  default_tags {
    tags = {
      Project     = "ecommerce-observability"
      Environment = "dev"
      ManagedBy   = "terraform"
    }
  }
}

terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  required_version = ">= 1.0.0"
}
