# Event-Driven Order Processing System

A distributed system built with Go that demonstrates event-driven architecture using Kafka for asynchronous communication between independent microservices.

## Overview

Four microservices handle different domains of an order processing pipeline. Each service is fully independent, with its own database and deployment lifecycle. Services communicate exclusively through Kafka topics — no direct HTTP calls between services.

## Architecture

```
order-service     →  orders.created     →  payment-service, inventory-service
order-service     →  orders.cancelled   →  payment-service, inventory-service
payment-service   →  payments.processed →  inventory-service, notification-service
payment-service   →  payments.failed    →  notification-service
inventory-service →  inventory.updated  →  notification-service
```

### Services

| Service              | Port | Database   | Responsibility                        |
|----------------------|------|------------|---------------------------------------|
| order-service        | 8080 | PostgreSQL | Create and manage orders              |
| payment-service      | 8081 | PostgreSQL | Process payments                      |
| inventory-service    | 8082 | PostgreSQL + Redis | Manage stock levels          |
| notification-service | 8083 | —          | Consume events and send notifications |

### Tech Stack

- **Go** — all services
- **Kafka** — asynchronous messaging (Confluent 7.5.0)
- **PostgreSQL 16** — persistent storage (one container, separate databases)
- **Redis 7** — caching layer for inventory-service
- **Gin** — HTTP framework
- **Docker Compose** — local infrastructure

## Project Structure

Each service follows the same Clean Architecture structure:

```
service/
├── cmd/service/        # Composition root (main.go)
└── internal/
    ├── domain/         # Entities — no external dependencies
    ├── usecase/        # Business logic + interface definitions
    ├── repository/     # PostgreSQL (+ Redis) implementations
    ├── handler/        # Gin HTTP handlers
    └── kafka/          # Kafka producer/consumer
```

Dependency direction: `handler` → `usecase` → `domain` ← `repository`, `kafka`

## Getting Started

### Prerequisites

- Docker and Docker Compose
- Go 1.21+

### Running with Docker

```bash
# Start infrastructure (Kafka, PostgreSQL, Redis)
make docker-up

# Start everything including services
make docker-up-all

# Stop all containers
make docker-down

# Follow logs
make docker-logs
```

### Running Locally

Copy the environment file for each service:

```bash
cp order-service/.env.example order-service/.env
cp payment-service/.env.example payment-service/.env
cp inventory-service/.env.example inventory-service/.env
cp notification-service/.env.example notification-service/.env
```

Start each service:

```bash
make run-order
make run-payment
make run-inventory
make run-notif
```

### Kafka UI

Available at [http://localhost:8090](http://localhost:8090) — inspect topics, messages and consumer groups.

## API Endpoints

### order-service (port 8080)
```
POST /orders        Create a new order
GET  /orders/:id    Get order by ID
```

### payment-service (port 8081)
```
POST /payments      Process a payment
GET  /payments/:id  Get payment by ID
```

### inventory-service (port 8082)
```
POST /inventory        Create inventory item
GET  /inventory/:id    Get item by product ID
PUT  /inventory/:id    Update stock
```

## Development

```bash
# Build all services
make build

# Run tests
make test

# go mod tidy for all services
make tidy
```

Run tests for a specific service:

```bash
cd order-service && go test ./internal/usecase/...
```

> If `go mod tidy` fails with "go.mod file not found", prefix with `GOPATH=/tmp/gopath`:
> ```bash
> cd order-service && GOPATH=/tmp/gopath go mod tidy
> ```
