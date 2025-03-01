terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "4.55.0"
    }
    tls = {
      source  = "hashicorp/tls"
      version = "4.0.4"
    }
  }
}

provider "aws" {
  region = local.region
}

data "aws_caller_identity" "current" {}
data "aws_availability_zones" "available" {}

resource "aws_ecs_cluster" "tracetest-cluster" {
  name = "tracetest"
  tags = local.tags
}

module "tracetest_ecs_service_security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "tracetest_ecs_service_security_group"
  description = "ECS Service security group"
  vpc_id      = module.network.vpc_id

  ingress_with_cidr_blocks = [
    {
      from_port   = 0
      to_port     = 65535
      protocol    = "tcp"
      description = "HTTP access from VPC"
      cidr_blocks = local.vpc_cidr
  }]

  egress_with_cidr_blocks = [
    {
      from_port   = 0
      to_port     = 65535
      protocol    = "-1"
      description = "HTTP access to anywhere"
      cidr_blocks = "0.0.0.0/0"
  }]
}

resource "aws_iam_role" "tracetest_task_execution_role" {
  name = "tracetest_task_execution_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = ["ecs-tasks.amazonaws.com", "ecs.amazonaws.com"]
        }
      },
    ]
  })
  managed_policy_arns = ["arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"]

  tags = local.tags
}

resource "aws_iam_role_policy" "tracetest_task_execution_role_policy" {
  name = "tracetest_task_execution_role_policy"
  role = aws_iam_role.tracetest_task_execution_role.id

  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "logs:PutLogEvents",
          "logs:CreateLogGroup",
          "logs:CreateLogStream",
          "logs:DescribeLogStreams",
          "logs:DescribeLogGroups",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}
