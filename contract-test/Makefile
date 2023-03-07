PROVIDER_BINARY=bin/providerApp
CONSUMER_BINARY=bin/consumerApp

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
up_build: build_provider build_consumer
	@echo "Stopping docker images (if running...)"
	docker-compose down
	@echo "Building (when required) and starting docker images..."
	docker-compose up --build -d
	@echo "Docker images built and started!"

## down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"

## build_consumer: builds the consumer binary as a linux executable
build_consumer:
	@echo "Building consumer binary..."
	cd ./consumer && env GOOS=linux CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ${CONSUMER_BINARY} ./cmd/api
	@echo "Done!"

## build_auth: builds the auth binary as a linux executable
build_provider:
	@echo "Building provider binary..."
	cd ./provider && env GOOS=linux CGO_ENABLED=0 go build -gcflags "all=-N -l" -o ${PROVIDER_BINARY} ./cmd/api
	@echo "Done!"

## localdb: localdb starts
localdb:
	@echo "Creating localdb"
	docker run --name users -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -d postgres
	@echo "Running postgres"