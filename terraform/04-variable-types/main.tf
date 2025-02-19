terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.0"
    }
  }
}

provider "aws" {
  region = var.aws_region
}

resource "aws_vpc" "main_vpc" {
  cidr_block         = "10.0.0.0/16"
  enable_dns_support = var.enable_dns

  tags = {
    "Name" = "Production VPC."
  }
}

resource "aws_subnet" "subnet1" {
  vpc_id            = aws_vpc.main_vpc.id
  cidr_block        = "10.0.10.0/24"
  availability_zone = var.azs[0]

  tags = {
    "Name" = "web-subnet"
  }
}

resource "aws_internet_gateway" "my_web_igw" {
  vpc_id = aws_vpc.main_vpc.id
  tags = {
    "Name" = "Web Internet Gateway"
  }
}

resource "aws_default_route_table" "main_vpc_default_rt" {
  default_route_table_id = aws_vpc.main_vpc.default_route_table_id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.my_web_igw.id
  }

  tags = {
    "Name" = "my-default-rt"
  }
}

resource "aws_default_security_group" "default_sec_group" {
  vpc_id = aws_vpc.main_vpc.id

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port   = var.web_port
    to_port     = var.web_port
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = var.egress_dsg["from_port"]
    to_port     = var.egress_dsg["to_port"]
    protocol    = var.egress_dsg["protocol"]
    cidr_blocks = var.egress_dsg["cidr_blocks"]
  }

  tags = {
    "Name" = "Default Security Group"
  }
}

resource "aws_instance" "server" {
  ami           = var.amis[var.aws_region]
  instance_type = var.my_instance[0]
  count         = 1
  # cpu_core_count              = var.my_instance[1]
  associate_public_ip_address = var.my_instance[2]

  tags = {
    "Name" = "Amazon Linux 2"
  }
}
