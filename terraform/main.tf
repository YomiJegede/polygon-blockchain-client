provider "aws" {
  region = "us-east-1"
}

resource "aws_ecs_cluster" "blockchain" {
  name = "blockchain-cluster"
}

resource "aws_ecs_task_definition" "blockchain_task" {
  family                   = "blockchain-task"
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  cpu                      = "256"
  memory                   = "512"
  execution_role_arn       = aws_iam_role.ecs_task_execution_role.arn
  container_definitions = jsonencode([
    {
      name      = "blockchain-client"
      image     = "your-dockerhub-username/polygon-blockchain-client:latest"
      cpu       = 256
      memory    = 512
      essential = true
      portMappings = [
        {
          containerPort = 8080
          hostPort      = 8080
        }
      ]
    }
  ])
}

resource "aws_ecs_service" "blockchain_service" {
  name            = "blockchain-service"
  cluster         = aws_ecs_cluster.blockchain.id
  task_definition = aws_ecs_task_definition.blockchain_task.arn
  desired_count   = 1
  launch_type     = "FARGATE"
  network_configuration {
    subnets         = ["subnet-xxxxxxxx"]  
    security_groups = ["sg-xxxxxxxx"]
    assign_public_ip = true
  }
}
