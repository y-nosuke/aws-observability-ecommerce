output "vpc_id" {
  description = "The ID of the VPC"
  value       = module.vpc.vpc_id
}

output "public_subnet_ids" {
  description = "The IDs of the public subnets"
  value       = module.vpc.public_subnet_ids
}

output "private_subnet_ids" {
  description = "The IDs of the private subnets"
  value       = module.vpc.private_subnet_ids
}

output "backend_repository_url" {
  description = "The URL of the backend ECR repository"
  value       = module.ecr_backend.repository_url
}

output "frontend_repository_url" {
  description = "The URL of the frontend ECR repository"
  value       = module.ecr_frontend.repository_url
}

output "static_assets_bucket" {
  description = "The S3 bucket for static assets"
  value       = module.s3_frontend.bucket_id
}

output "ecs_task_execution_role_arn" {
  description = "The ARN of the ECS task execution role"
  value       = module.iam.ecs_task_execution_role_arn
}

output "ecs_task_role_arn" {
  description = "The ARN of the ECS task role"
  value       = module.iam.ecs_task_role_arn
}
