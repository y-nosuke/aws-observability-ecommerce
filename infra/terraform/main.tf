module "cloudwatch_logs" {
  source = "./modules/cloudwatch"

  app_name = var.app_name
  env      = var.environment

  log_groups = [
    {
      name      = "backend"
      retention = 30 # 30日間保持
      streams   = ["api-server", "error"]
    },
    {
      name      = "frontend-customer"
      retention = 14 # 14日間保持
      streams   = ["app", "error", "access"]
    },
    {
      name      = "frontend-admin"
      retention = 14 # 14日間保持
      streams   = ["app", "error", "access"]
    }
  ]
}
