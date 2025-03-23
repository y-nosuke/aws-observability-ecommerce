provider "aws" {
  region                      = "ap-northeast-1"
  access_key                  = "localstack"
  secret_key                  = "localstack"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true

  # tflocalを使用する場合、以下の設定は自動的に行われるため、コメントアウトしておきます
  # endpoints {
  #   cloudwatch     = "http://localhost:4566"
  #   logs           = "http://localhost:4566"
  #   s3             = "http://localhost:4566"
  # }
}
