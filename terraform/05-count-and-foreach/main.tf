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

# resource "aws_instance" "server" {
#   ami           = "ami-053a45fff0a704a47"
#   instance_type = "t2.micro"
#   count         = 3
# }

variable "users" {
  type    = list(string)
  default = ["demo-user", "admin1", "john"]
}

resource "aws_iam_user" "test" {
  for_each = toset(var.users)
  name     = each.key
}
