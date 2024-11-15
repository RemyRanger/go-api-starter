# Docker Compose Up
docker_compose_up:
	docker compose up -d

# Docker Compose Up
docker_compose_down:
	docker compose down

# Install all Go dependencies
install:
	go mod download

lint_oas:
	npx -y @redocly/cli@latest lint

# Generate OAS bundled document in /dist
generate_oas:
	npx -y @redocly/cli@latest bundle
	oapi-codegen -version
	go generate -x $$(go list ./... | grep -v '/ports\|/adapters|/core' | tr '\n' ' ')

# Build all binaries
build:
	go build -o cmd/control/sxc_control cmd/control/main.go
	go build -o cmd/gateway/sxc_gateway cmd/gateway/main.go

run_control:
	go run cmd/control/main.go

run_gateway:
	go run cmd/gateway/main.go

update_deps:
	go get -u ./...
	go mod tidy
