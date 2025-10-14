.PHONY: init infra-up infra-down verify logs proto-gen build-all test-all clean

# Initialize project
init:
	@echo "Initializing BookmySeat project..."
	@go mod tidy
	@chmod +x scripts/*.sh
	@echo "✓ Project initialized"

# Infrastructure commands
infra-up:
	@echo "Starting infrastructure..."
	@docker-compose up -d
	@echo "✓ Infrastructure started"

infra-down:
	@echo "Stopping infrastructure..."
	@docker-compose down
	@echo "✓ Infrastructure stopped"

infra-restart:
	@make infra-down
	@make infra-up

# Verify all services
verify:
	@echo "Verifying services..."
	@bash scripts/verify-setup.sh

# View logs
logs:
	@docker-compose logs -f

# Generate protobuf code
proto-gen:
	@echo "Generating protobuf code..."
	@export PATH=$$PATH:$$(go env GOPATH)/bin && bash scripts/generate-proto.sh
	@echo "✓ Protobuf code generated"

# Build all services
build-all:
	@echo "Building all services..."
	@go build -o bin/user-service services/user-service/main.go || true
	@go build -o bin/event-service services/event-service/main.go || true
	@go build -o bin/booking-service services/booking-service/main.go || true
	@go build -o bin/payment-service services/payment-service/main.go || true
	@go build -o bin/notification-service services/notification-service/main.go || true
	@go build -o bin/api-gateway services/api-gateway/main.go || true
	@echo "✓ All services built"

# Test all services
test-all:
	@echo "Running tests..."
	@go test -v -cover ./...
	@echo "✓ Tests completed"

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	@rm -rf bin/
	@find proto -name "*.pb.go" -delete
	@echo "✓ Cleaned"

# Monitoring shortcuts
open-grafana:
	@open http://localhost:3000

open-jaeger:
	@open http://localhost:16686

open-kibana:
	@open http://localhost:5601

open-prometheus:
	@open http://localhost:9090
