variable "web_port" {
  description = "Web Port"
  default     = 80
  type        = number
}

variable "aws_region" {
  description = "value"
  type        = string
  default     = "us-east-1"
}

variable "enable_dns" {
  description = "DNS support for the VPC"
  type        = bool
  default     = true
}

variable "azs" {
  description = "AZs in the region"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b", "us-east-1c"]
}

variable "amis" {
  type        = map(string)
  description = "Map from region to AMI ID"
  default = {
    "us-east-1" = "ami-053a45fff0a704a47"
  }
}

variable "my_instance" {
  type    = tuple([string, number, bool])
  default = ["t2.micro", 1, true]
}

variable "egress_dsg" {
  type = object({
    from_port   = number
    to_port     = number
    protocol    = string
    cidr_blocks = list(string)
  })
  default = {
    cidr_blocks = ["100.0.0.0/16", "200.0.0.0/16"]
    from_port   = 0
    to_port     = 65365
    protocol    = "tcp"
  }
}
