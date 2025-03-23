variable "app_name" {
  description = "Application name"
  type        = string
  default     = "aws-observability-ecommerce"
}

variable "environment" {
  description = "Environment (dev, staging, prod)"
  type        = string
  default     = "dev"
}
