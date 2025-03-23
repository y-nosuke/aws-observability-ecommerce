resource "aws_cloudwatch_log_group" "this" {
  for_each = { for lg in var.log_groups : lg.name => lg }

  name              = "/${var.app_name}/${var.env}/${each.value.name}"
  retention_in_days = each.value.retention

  tags = {
    Environment = var.env
    Application = var.app_name
  }
}

resource "aws_cloudwatch_log_stream" "this" {
  for_each = {
    for pair in flatten([
      for lg in var.log_groups : [
        for stream in lg.streams : {
          log_group = lg.name
          stream    = stream
        }
      ]
    ]) : "${pair.log_group}-${pair.stream}" => pair
  }

  name           = each.value.stream
  log_group_name = "/${var.app_name}/${var.env}/${each.value.log_group}"

  depends_on = [aws_cloudwatch_log_group.this]
}
