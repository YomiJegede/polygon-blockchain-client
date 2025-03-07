// using environment variables and multi-stage build for optimization
# Build stage
FROM golang:1.20 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o blockchain-client main.go

# Deployment stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/blockchain-client .
EXPOSE 8080
ENV POLYGON_RPC_URL=https://polygon-rpc.com/
CMD ["./blockchain-client"]
