variable "project_name" {
  description = "Name of the project"
  type        = string
}

variable "environment" {
  description = "Deployment environment (dev, staging, prod)"
  type        = string
}

variable "bucket_name" {
  description = "Name of the S3 bucket"
  type        = string
}

variable "is_public" {
  description = "Whether the bucket should be public"
  type        = bool
  default     = false
}

variable "enable_versioning" {
  description = "Whether to enable bucket versioning"
  type        = bool
  default     = false
}
