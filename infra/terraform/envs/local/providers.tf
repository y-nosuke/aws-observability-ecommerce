#####################################################
# Provider
#####################################################

provider "aws" {
  region                      = "ap-northeast-1"
  access_key                  = "localstack"
  secret_key                  = "localstack"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  default_tags {
    tags = {
      Environment = title(var.environment)
      Application = var.app_name
    }
  }
}
