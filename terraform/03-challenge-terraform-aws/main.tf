resource "aws_vpc" "production" {
  cidr_block = "192.168.0.0/24"

  tags = {
    "Name" = "Production VPC"
  }
}

resource "aws_subnet" "webapps" {
  vpc_id            = aws_vpc.production.id
  cidr_block        = "192.168.0.32/27"
  availability_zone = "us-west-1b"

  tags = {
    "Name" = "Web Applications Subnet"
  }
}
