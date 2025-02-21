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

module "vpc" {
  source = "../modules/vpc"
}

module "server" {
  source = "../modules/server"

  vpc_id         = module.vpc.main_vpc_id
  web_subnet_id  = module.vpc.web_subnet_id
  script_name    = "${path.root}/entry-script.sh"
  ssh_public_key = var.ssh_public_key
}
