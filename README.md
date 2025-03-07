# polygon-blockchain-client
# Deployment and Production Considerations

"# Blockchain Client

This is a simple blockchain client that interacts with Polygon RPC.

## Features
- Fetch the latest block number
- Fetch block details by block number
- API exposed via REST
- Dockerized for easy deployment
- Terraform configuration for AWS ECS Fargate deployment

## Installation
1. Clone this repository:
   `git clone git@github.com:YomiJegede/polygon-blockchain-client.git`
2. Install dependencies:
   `go mod tidy`
3. Build and run the application:
   `go build -o blockchain-client main.go ./polygon-blockchain-client`
4. ## Running with Docker
   `docker build -t blockchain-client . docker run -p 8080:8080 -e POLYGON_RPC_URL=https://polygon-rpc.com/ polygon-blockchain-client`
5. ## Running Tests
   `go test ./tests`
6. ## Deploying with Terraform
  1. Configure AWS credentials
  2. Initialize Terraform:
     `cd terraform terraform init`
  3. Plan and Apply:
     `terraform apply`
7. Production Enhancements
- **Logging**: Use structured logging (Zap or Logrus)
- **Rate-limiting**: Prevent abuse using middleware (e.g., `golang.org/x/time/rate`)
- **Authentication**: Use API keys or JWT for secured access
- **CI/CD**: Automate builds and deployments with GitHub Actions or AWS CodePipeline
- **Health Checks**: Implement `/healthz` endpoint to monitor app status
- **Persistent Storage**: Use a database like PostgreSQL for historical data storage
   
