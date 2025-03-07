# terraform/iam.tf

# GitHub ActionsがAssumeできるIAMロール
resource "aws_iam_role" "github_actions" {
  name = "GitHubActionsRole"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
      }
    ]
  })
}

# デプロイに必要な権限ポリシー
resource "aws_iam_policy" "deployment_policy" {
  name = "DeploymentPolicy"
  
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = [
          "s3:PutObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:DeleteObject"
        ]
        Resource = [
          aws_s3_bucket.frontend_deploy.arn,
          "${aws_s3_bucket.frontend_deploy.arn}/*",
          aws_s3_bucket.cfn_templates.arn,
          "${aws_s3_bucket.cfn_templates.arn}/*"
        ]
      },
      {
        Effect = "Allow"
        Action = [
          "cloudfront:CreateInvalidation",
          "cloudfront:GetInvalidation",
          "cloudfront:ListInvalidations"
        ]
        Resource = aws_cloudfront_distribution.frontend_distribution.arn
      },
      {
        Effect = "Allow"
        Action = [
          "lambda:UpdateFunctionCode",
          "lambda:GetFunction",
          "lambda:CreateFunction",
          "lambda:DeleteFunction",
          "lambda:UpdateFunctionConfiguration",
          "lambda:TagResource"
        ]
        Resource = "arn:aws:lambda:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:function:ecommerce-*"
      },
      {
        Effect = "Allow"
        Action = [
          "apigateway:PUT",
          "apigateway:POST",
          "apigateway:GET",
          "apigateway:DELETE",
          "apigateway:PATCH"
        ]
        Resource = "arn:aws:apigateway:${data.aws_region.current.name}::*"
      },
      {
        Effect = "Allow"
        Action = [
          "dynamodb:DescribeTable",
          "dynamodb:CreateTable",
          "dynamodb:UpdateTable",
          "dynamodb:DeleteTable"
        ]
        Resource = "arn:aws:dynamodb:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:table/ecommerce-*"
      },
      {
        Effect = "Allow"
        Action = [
          "ssm:GetParameter",
          "ssm:GetParameters"
        ]
        Resource = "arn:aws:ssm:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:parameter/ecommerce/*"
      },
      {
        Effect = "Allow"
        Action = [
          "cloudformation:*Stack*"
        ]
        Resource = "arn:aws:cloudformation:${data.aws_region.current.name}:${data.aws_caller_identity.current.account_id}:stack/ecommerce-*/*"
      }
    ]
  })
}

# ポリシーをロールにアタッチ
resource "aws_iam_role_policy_attachment" "github_actions_policy" {
  role       = aws_iam_role.github_actions.name
  policy_arn = aws_iam_policy.deployment_policy.arn
}

# 現在のAWSアカウント情報を取得
data "aws_caller_identity" "current" {}
data "aws_region" "current" {}
