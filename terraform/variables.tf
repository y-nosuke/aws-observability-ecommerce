variable "project_name" {
  description = "Name of the project"
  type        = string
  default     = "ecommerce-observability"
}

variable "environment" {
  description = "Deployment environment"
  type        = string
  default     = "dev"
}

variable "enable_cloudfront_logs" {
  description = "Enable CloudFront access logs"
  type        = bool
  default     = false
}

variable "enable_s3_logs" {
  description = "Enable S3 server access logs"
  type        = bool
  default     = false
}
