.PHONY: help test run clean

help:
	@echo "Available commands:"
	@echo "  make test            - Run all tests"
	@echo "  make test-verbose    - Run tests with verbose output"
	@echo "  make run             - Run the application on port 8080"
	@echo "  make clean           - Clean artifacts"
	@echo "  make help            - Show this help message"

test:
	@echo "Running tests..."
	go test ./...

test-verbose:
	@echo "Running tests with verbose output..."
	go test ./... -v

run:
	@echo "Starting application on port 8080..."
	go run ./main.go

clean:
	@echo "Cleaning artifacts..."
	go clean
	rm -f coverage.out

.DEFAULT_GOAL := help

