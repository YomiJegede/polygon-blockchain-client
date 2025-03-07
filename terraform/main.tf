provider "aws" {
  region = var.aws_region
}

resource "aws_ecs_cluster" "blockchain" {
  name = "blockchain-cluster"
}

module "network" {
  source = "terraform-aws-modules/vpc/aws"
  name   = "blockchain-vpc"
  cidr   = "10.0.0.0/16"
  
  azs             = ["us-east-1a", "us-east-1b"]
  public_subnets  = ["10.0.1.0/24", "10.0.2.0/24"]
}

resource "aws_security_group" "ecs_sg" {
  vpc_id = module.network.vpc_id

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}
