# go-template

## Requirements

* [docker and docker-compose for local usage / development](https://www.docker.com/get-started)
* [MariaDB](https://mariadb.org/) or [MySQL](https://www.mysql.com/) (optional only for run on host machine) 
* [Redis](https://redis.io/download//) (optional only for run on host machine) 

## Architecture

This project uses **Domain-Driven Design (DDD)** with **Hexagonal Architecture** principles, following Go best practices.

### Structure

```
go-template
├── cmd
│   └── server              # Entry point
├── internal
│   ├── domain/             # Core business logic (no external dependencies)
│   │   └── example/        # Example domain (rename for your needs)
│   │       └── model/     # Entities & value objects
│   │
│   ├── application/        # Use cases & orchestration
│   │   └── example/
│   │       ├── command/   # Request/Response DTOs
│   │       ├── handler/   # Use case implementation
│   │       └── repository/ # Repository interface (CONSUMER-DEFINED)
│   │
│   ├── infrastructure/   # External adapters
│   │   ├── persistence/   # Database implementations
│   │   │   └── repository/
│   │   └── adapters/      # HTTP/gRPC handlers
│   │       ├── http/
│   │       └── grpc/
│   │
│   ├── app/              # Composition root (DI wiring)
│   │   └── config/
│   │
│   └── api/v1/          # OpenAPI generated files
│
├── pkg                   # Shared utilities (logger, database)
├── proto                 # gRPC definitions
├── deployment            # Kubernetes deployment files
├── builds               # Docker-compose files
├── config.yaml          # Configuration file
└── openapi.yaml        # OpenAPI specification
```

### Layer Responsibilities

| Layer | Responsibility |
|-------|---------------|
| `domain` | Pure business logic, entities, value objects (NO external deps) |
| `application` | Use cases, orchestrates domain, defines repository ports |
| `infrastructure` | Implementations (database, HTTP, gRPC adapters) |
| `app` | Dependency injection, wiring, server startup |
| `api` | OpenAPI generated code |

### Dependency Rule

Following Go best practices: **"Accept Interfaces, Return Structs"**

- **Consumer defines interfaces** (application layer)
- **Implementer provides concrete** (infrastructure layer)
- Dependencies point inward only: infrastructure → application → domain

### Adding a New Domain

To add a new domain (e.g., "users"):

1. Create domain model: `internal/domain/users/model/`
2. Create command DTOs: `internal/application/users/command/`
3. Create repository interface: `internal/application/users/repository/` (consumer-defined)
4. Create use case handler: `internal/application/users/handler/`
5. Create repository adapter: `internal/infrastructure/persistence/repository/`
6. Create HTTP handler: `internal/infrastructure/adapters/http/`

## Configuration

Copy the template file:

```
cp config.yaml_sample config.yaml
```

Copy the env template file:

```
cp .env.local .env
```

Edit the corresponding fields as needed.

*Note: The make script will bind .env variables on each run as ENVIRONMENT variables. For production, inject these env variables into the host.*

## Commands

| Command | Description |
|---------|-------------|
| `make up` | Build and start services with docker-compose |
| `make down` | Stop docker-compose services |
| `make run` | Run the app locally |
| `make mysql-up` | Start MySQL with docker-compose |
| `make mysql-down` | Stop MySQL |
| `make redis-up` | Start Redis with docker-compose |
| `make redis-down` | Stop Redis |
| `make mockery` | Generate mock interfaces |
| `make unit` | Run unit tests |
| `make unit-coverage` | Run tests with coverage report |
| `make gen-api` | Generate code from OpenAPI spec |
| `make gen-api-doc` | Generate HTML API documentation |
| `make proto-gen` | Generate gRPC code |

## Quick Start

### Running with Docker

```
make up
```

The service starts at port `80` by default. Change the port in `config.yaml` if needed.

### Running Locally

Requires Go, MySQL, and Redis installed:

```
make run
```

---

Created by Alfonso Rodriguez (beabys@gmail.com)
