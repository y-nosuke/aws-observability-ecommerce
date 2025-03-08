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
