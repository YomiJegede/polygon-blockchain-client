FROM golang:1.20 as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o polygon-blockchain-client main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/polygon-blockchain-client .
EXPOSE 8080
CMD ["./polygon-blockchain-client"]
