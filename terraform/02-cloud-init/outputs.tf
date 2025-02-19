output "ec2_public_ip" {
  description = "The public IP address of the EC2 instance"
  value       = aws_instance.my_vm.public_ip
}

output "vpc_id" {
  description = "VPC ID"
  value       = aws_vpc.main.id
}

output "ami_id" {
  description = "AMI ID"
  value       = aws_instance.my_vm.ami
  sensitive   = true
}
