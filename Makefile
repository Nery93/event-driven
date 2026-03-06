# Infrastructure
docker-up:
    docker compose up -d zookeeper kafka kafka-ui postgres redis

docker-up-all:
    docker compose up -d

docker-down:
    docker compose down

docker-logs:
    docker compose logs -f

# Run services locally
run-order:
    cd order-service && go run cmd/order/main.go

run-payment:
    cd payment-service && go run cmd/payment/main.go

run-inventory:
    cd inventory-service && go run cmd/inventory/main.go

run-notif:
    cd notification-service && go run cmd/notification/main.go

# Build
build:
    cd order-service && go build ./...
    cd payment-service && go build ./...
    cd inventory-service && go build ./...
    cd notification-service && go build ./...

build-order:
    cd order-service && go build ./...

# Test
test:
    cd order-service && go test ./...
    cd payment-service && go test ./...
    cd inventory-service && go test ./...
    cd notification-service && go test ./...

# Tidy
tidy:
    cd order-service && GOPATH=/tmp/gopath go mod tidy
    cd payment-service && GOPATH=/tmp/gopath go mod tidy
    cd inventory-service && GOPATH=/tmp/gopath go mod tidy
    cd notification-service && GOPATH=/tmp/gopath go mod tidy