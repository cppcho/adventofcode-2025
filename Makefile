.PHONY: run build test fmt lint clean help

# Default target
.DEFAULT_GOAL := help

# Binary name
BINARY := aoc

# Run a specific day's solution
# Usage: make run DAY=1
run:
ifndef DAY
	@echo "Error: DAY not specified. Usage: make run DAY=1"
	@exit 1
endif
	go run cmd/aoc/main.go $(DAY)

# Build the CLI binary
build:
	go build -o $(BINARY) cmd/aoc/main.go

# Run all tests
test:
	go test ./...

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Clean build artifacts
clean:
	rm -f $(BINARY)

# Show help
help:
	@echo "Available targets:"
	@echo "  make run DAY=<day>  - Run a specific day's solution (e.g., make run DAY=1)"
	@echo "  make build          - Build the CLI binary"
	@echo "  make test           - Run all tests"
	@echo "  make fmt            - Format code with go fmt"
	@echo "  make lint           - Run linter (requires golangci-lint)"
	@echo "  make clean          - Remove build artifacts"
	@echo "  make help           - Show this help message"
