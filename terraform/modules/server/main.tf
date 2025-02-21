resource "aws_default_security_group" "default_sec_group" {
  vpc_id = var.vpc_id

  ingress {
    from_port = 22
    to_port   = 22

    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
    # cidr_blocks = [var.my_public_ip]
  }

  ingress {
    from_port = 80
    to_port   = 80

    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  ingress {
    from_port = 8080
    to_port   = 8080

    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1" # any protocol
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    "Name" = "Default Security Group"
  }
}

resource "aws_key_pair" "test_ssh_key" {
  key_name   = "testing_ssh_key"
  public_key = file(var.ssh_public_key)
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

resource "aws_instance" "my_vm" {
  ami                         = data.aws_ami.latest_amazon_linux2.id # Amazon Linux
  instance_type               = var.server_type
  subnet_id                   = var.web_subnet_id
  vpc_security_group_ids      = [aws_default_security_group.default_sec_group.id]
  associate_public_ip_address = true
  key_name                    = aws_key_pair.test_ssh_key.key_name
  user_data                   = file(var.script_name)

  # connection {
  #  type        = "ssh"
  #  host        = self.public_ip
  #  user        = "ec2-user"
  #  private_key = file("~/.ssh/test_rsa")
  # }

  # provisioner "file" {
  #  source      = "./entry-script.sh"
  #  destination = "/home/ec2-user/entry-script.sh"
  # }

  # provisioner "remote-exec" {
  #  inline = [
  #    "chmod +x /home/ec2-user/entry-script.sh",
  #    "/home/ec2-user/entry-script.sh",
  #    "exit"
  #  ]
  #  on_failure = continue
  # }

  tags = {
    "Name" = "My EC2 Instance - Amazon Linux 2"
  }
}
