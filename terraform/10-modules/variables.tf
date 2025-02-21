variable "vpc_cidr_block" {
  default     = "10.16.0.0/16"
  description = "CIDR Block for the VPC"
  type        = string
}

variable "web_subnet" {
  default     = "10.16.1.0/24"
  description = "Web Subnet CIDR Block"
  type        = string
}

variable "main_vpc_name" {

}

variable "my_public_ip" {}

variable "ssh_public_key" {

}
