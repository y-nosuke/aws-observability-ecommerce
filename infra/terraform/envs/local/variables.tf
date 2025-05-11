#####################################################
# Variable
#####################################################

variable "app_name" {
  description = "アプリケーション名"
  type        = string
  default     = "aws-observability-ecommerce"
}

variable "environment" {
  description = "環境名（local, dev, staging, prod等）"
  type        = string
  default     = "local"
}

variable "log_groups" {
  description = "作成するロググループの設定"
  type = list(object({
    name      = string
    retention = number
  }))
  default = []
}
