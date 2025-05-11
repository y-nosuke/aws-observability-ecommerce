#####################################################
# Locals
#####################################################

locals {
  log_group_prefix = "/${var.app_name}/${var.environment}"
  log_groups = [
    {
      name = "${local.log_group_prefix}/backend-api"
      retention = 30
    },
    {
      name = "${local.log_group_prefix}/frontend-customer"
      retention = 14
    },
    {
      name = "${local.log_group_prefix}/frontend-admin"
      retention = 14
    }
  ]
}

#####################################################
# Module
#####################################################

module "cloudwatch_logs" {
  source = "../../modules/resources/cloudwatch"

  for_each = {
    for log_group in local.log_groups : log_group.name => log_group
  }

  log_group_name = each.value.name
  retention_in_days = each.value.retention
}

#####################################################
# Resource
#####################################################

#####################################################
# Data
#####################################################
