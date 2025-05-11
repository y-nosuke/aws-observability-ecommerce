#####################################################
# Output
#####################################################

output "log_group_names" {
  value       = [for log_group in module.cloudwatch_logs : log_group.name]
  description = "作成されたロググループの名前"
}

output "log_group_arns" {
  value       = [for log_group in module.cloudwatch_logs : log_group.arn]
  description = "作成されたロググループのARN"
}
