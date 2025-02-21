output "ec2_public_ip" {
  description = "The public IP address of the EC2 instance"
  value       = module.server.instance.public_ip
}

output "vpc_id" {
  description = "VPC ID"
  value       = module.vpc.main_vpc_id
}

output "ami_id" {
  description = "AMI ID"
  value       = module.server.instance.ami
  sensitive   = true
}
