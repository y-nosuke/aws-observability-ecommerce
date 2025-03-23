output "log_group_arns" {
  description = "作成されたロググループのARN"
  value       = { for k, v in aws_cloudwatch_log_group.this : k => v.arn }
}

output "log_group_names" {
  description = "作成されたロググループの名前"
  value       = { for k, v in aws_cloudwatch_log_group.this : k => v.name }
}
