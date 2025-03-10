provider "aws" {
  region = var.aws_region
}

module "vpc" {
  source = "../../modules/vpc"

  project_name         = var.project_name
  environment          = var.environment
  vpc_cidr             = var.vpc_cidr
  public_subnet_cidrs  = var.public_subnet_cidrs
  private_subnet_cidrs = var.private_subnet_cidrs
  availability_zones   = var.availability_zones
}

module "ecr_backend" {
  source = "../../modules/ecr"

  project_name    = var.project_name
  environment     = var.environment
  repository_name = "backend"
}

module "ecr_frontend" {
  source = "../../modules/ecr"

  project_name    = var.project_name
  environment     = var.environment
  repository_name = "frontend"
}

module "s3_frontend" {
  source = "../../modules/s3"

  project_name      = var.project_name
  environment       = var.environment
  bucket_name       = "static-assets"
  is_public         = true
  enable_versioning = false
}

module "iam" {
  source = "../../modules/iam"

  project_name            = var.project_name
  environment             = var.environment
  create_s3_access_policy = true
  s3_bucket_arn           = module.s3_frontend.bucket_arn
}

