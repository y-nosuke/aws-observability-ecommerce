output "log_group_names" {
  value       = module.cloudwatch_logs.log_group_names
  description = "Names of the created log groups"
}
