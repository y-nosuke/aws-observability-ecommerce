#####################################################
# Variable
#####################################################

variable "log_group_name" {
  type        = string
  description = "CloudWatch Logsのロググループ名"
}

variable "retention_in_days" {
  type        = number
  description = "ログ保持期間（日）"

  validation {
    condition = contains([
      0, 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, 3653
    ], var.retention_in_days)
    error_message = "retention_in_days には 0, 1, 3, 5, 7, 14, 30, 60, 90, 120, 150, 180, 365, 400, 545, 731, 1827, 2192, 2557, 2922, 3288, 3653 のいずれかの値を指定してください。"
  }
}
