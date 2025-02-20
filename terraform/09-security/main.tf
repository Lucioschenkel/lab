terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

# data "aws_secretsmanager_secret_version" "creds" {
#   secret_id = "db_credentials"
# }
#
# locals {
#   db_creds = jsondecode(data.aws_secretsmanager_secret_version.creds.secret_string)
# }

# resource "aws_db_instance" "default" {
#   allocated_storage = 10
#   engine            = "mysql"
#   engine_version    = "5.7"
#   instance_class    = "db.t3.micro"
#   name              = "mydb"
#
#   username             = local.db_creds.username
#   password             = local.db_creds.password
#   parameter_group_name = "default.mysql5.7"
#   skip_final_snapshot  = true
# }

