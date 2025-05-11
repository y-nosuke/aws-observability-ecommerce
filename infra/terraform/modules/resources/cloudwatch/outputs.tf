#####################################################
# Output
#####################################################

output "arn" {
  description = "作成されたロググループのARN"
  value       = aws_cloudwatch_log_group.this.arn
}

output "name" {
  description = "作成されたロググループの名前"
  value       = aws_cloudwatch_log_group.this.name
}
