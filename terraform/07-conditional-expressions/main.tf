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

data "aws_ami" "latest_amazon_linux2" {
  owners      = ["amazon"]
  most_recent = true
  filter {
    name   = "name" # filter by name
    values = ["amzn2-ami-kernel-*-x86_64-gp2"]
  }

  filter {
    name   = "architecture"
    values = ["x86_64"]
  }
}

resource "aws_instance" "test_server" {
  ami           = data.aws_ami.latest_amazon_linux2.id
  instance_type = "t2.micro"
  count         = var.is_test == true ? 1 : 0
}

resource "aws_instance" "prod_server" {
  ami           = data.aws_ami.latest_amazon_linux2.id
  instance_type = "t2.large"
  count         = var.is_test == false ? 1 : 0
}
