# polygon-blockchain-client

"# Blockchain Client

This is a simple blockchain client that interacts with Polygon RPC.

## Features
- Fetch the latest block number
- Fetch block details by block number
- API exposed via REST
- Dockerized for easy deployment
- Terraform configuration for AWS ECS Fargate deployment

## Installation
1. Clone repository:
   `git clone git@github.com:YomiJegede/polygon-blockchain-client.git`

2. Install dependencies:
   `go mod tidy`
   
3. Build and run the application:
   `go build .`
   `go run main.go`
4. ## Running with Docker
   `docker build -t jyormie/polygon-blockchain-client .`
   `docker tag polygon-blockchain-client jyormie/polygon-blockchain-client:latest`
   `docker push jyormie/polygon-blockchain-client:latest`

5. ## Running Tests
   `go test ./tests`

6. ## Deploying with Terraform
  1. Configure AWS credentials
  2. Initialize Terraform:
     `cd terraform 
         `terraform init`
         `terraform plan`

7. Enhancements for Production readiness
   1. Monitoting & Logging: Use of a structured logging with Prometheus and Grafana and AWS CloudWatch for better observability and debugging.
   2. Rate-limiting: To prevent abuse and control traffic we can use AWS API Gateway for rate-limiting setting usage plans.
   3. Authentication: Secure your API with API Keys using AWS API Gateway and Amazon Cognito for more complex user management.
   4. CI/CD: Automate the build, test, and deployment process using Azure DevOps or GitLab CI/CD pipelines.
   5. Persistent Storage: Use PostgreSQL for reliable and stateful data storage, leveraging Amazon RDS for managed services or Amazon Aurora for scalability.
   
