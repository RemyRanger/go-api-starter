# Go-api-starter

# Run database, otel-collector, grafana with docker compose

```bash
make docker_compose_up
```

Clear running docker-compose. Datas will be cleared.

```bash
make docker_compose_down
```

# Install Go modules

```bash
make install
```

# Generate APIs interfaces from oas

```bash
make generate_oas
```

# Run

```bash
make run_control
make run_gateway
```

# Build all binaries

```bash
make build
```

# Update and lint Go dependencies

```bash
make update_deps
```

# Grafana Dashboard

see : http://127.0.0.1:3000/grafana/dashboards