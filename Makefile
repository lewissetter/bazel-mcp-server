.PHONY: build test clean install lint fmt help

# Variables
BINARY_NAME=bazel-mcp-server
MAIN_PATH=./cmd/bazel-mcp-server
INSTALL_PATH=$(shell go env GOPATH)/bin

# Default target
.DEFAULT_GOAL := help

## build: Build the binary
build:
	@echo "Building..."
	@go build -o $(BINARY_NAME) $(MAIN_PATH)

## test: Run tests
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...

## coverage: Run tests with coverage report
coverage: test
	@echo "Generating coverage report..."
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## clean: Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@rm -f coverage.out coverage.html
	@go clean

## install: Install the binary to GOPATH/bin
install:
	@echo "Installing to $(INSTALL_PATH)..."
	@go install $(MAIN_PATH)

## fmt: Format code
fmt:
	@echo "Formatting code..."
	@go fmt ./...

## lint: Run linters
lint:
	@echo "Running linters..."
	@golangci-lint run --timeout=5m || echo "golangci-lint not installed. Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"

## vet: Run go vet
vet:
	@echo "Running go vet..."
	@go vet ./...

## mod: Tidy and verify module dependencies
mod:
	@echo "Tidying module dependencies..."
	@go mod tidy
	@go mod verify

## run: Build and run the server
run: build
	@echo "Running server..."
	@./$(BINARY_NAME)

## help: Show this help message
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
