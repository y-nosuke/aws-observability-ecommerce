variable "app_name" {
  description = "アプリケーション名"
  type        = string
  default     = "aws-observability-ecommerce"
}

variable "env" {
  description = "環境名（dev, staging, prod等）"
  type        = string
  default     = "dev"
}

variable "log_groups" {
  description = "作成するロググループの設定"
  type = list(object({
    name      = string
    retention = number
    streams   = list(string)
  }))
  default = []
}
