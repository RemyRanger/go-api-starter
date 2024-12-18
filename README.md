
# Go-api-starter - Microservices Monorepo Example with Golang

This repository provides an example of a **microservices architecture** implemented in **Golang**, following a **monorepo structure**. It demonstrates the use of **OpenAPI specifications** to generate service interfaces and adheres to the principles of **hexagonal architecture (ports and adapters)** for clean, maintainable, and scalable code.

## Features

- **Monorepo Architecture**: A single repository to manage multiple microservices, shared libraries, and common tools.
- **Hexagonal Architecture**: Ensures separation of concerns by decoupling business logic from external dependencies like databases or APIs.
- **OpenAPI Integration**: Automatically generates service interfaces and client code from OpenAPI specifications, reducing boilerplate and ensuring consistency.
- **Golang Implementation**: Focused on performance, simplicity, and a strong type system.
- **Scalable Design**: Ready for production-grade enhancements such as observability, distributed tracing, and more.

## Installation

### Prerequisites

- Go (v1.23 or later)
- Docker (optional, for running services in containers)
- NPX for OpenAPI Generator CLI (for generating code)

### Steps

1. Run database, otel-collector, grafana with docker compose:
    ```bash
    make docker_compose_up
    ```

2. Install Go modules:
    ```bash
    make install
    ```

3. Generate APIs interfaces from oas:
    ```bash
    make generate_oas
    ```

4. Run services:
    ```bash
    make run_control
    make run_gateway
    ```

5. Build all binaries:
    ```bash
    make build
    ```

6. Update and lint Go dependencies:
    ```bash
    make update_deps
    ```

6. Run test and see coverage in html:
    ```bash
    make test
    make test_with_coverage
    ```

7. Check telemetry from Grafana:  http://127.0.0.1:3000/grafana/dashboards

---

## Design Principles

### Hexagonal Architecture

The project follows the ports and adapters pattern:

- **Ports**: Interfaces that define the behavior required by the application (e.g., repositories, external APIs).
- **Adapters**: Implementations of the ports, such as database drivers or HTTP clients.
- **Core Domain**: Business logic is isolated and does not depend on the adapters.

This ensures testability, flexibility, and the ability to swap out dependencies with minimal impact.

---

### Code Generation with OpenAPI

The OpenAPI Generator is used to:

- Create server interfaces for microservices.
- Generate strongly-typed interfaces for inter-service communication.
- Maintain consistency across services.

OAS files can be customized doc/ directory.

---