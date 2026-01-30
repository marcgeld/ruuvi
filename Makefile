.PHONY: help test lint build tidy fmt vet clean install

# Default target
help:
	@echo "Available targets:"
	@echo "  make test       - Run all tests with race detection"
	@echo "  make lint       - Run golangci-lint"
	@echo "  make build      - Build all packages and the CLI"
	@echo "  make tidy       - Run go mod tidy"
	@echo "  make fmt        - Format code with gofmt and goimports"
	@echo "  make vet        - Run go vet"
	@echo "  make clean      - Remove build artifacts"
	@echo "  make install    - Install the ruuvi CLI tool"

# Run tests with race detection and coverage
test:
	go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...

# Run golangci-lint
lint:
	golangci-lint run --timeout=5m

# Build all packages and the CLI
build:
	go build -v ./...
	cd cmd/ruuvi && go build -v .

# Run go mod tidy
tidy:
	go mod tidy

# Format code
fmt:
	go fmt ./...
	goimports -w -local github.com/marcgeld/ruuvi .

# Run go vet
vet:
	go vet ./...

# Clean build artifacts
clean:
	rm -f ruuvi
	rm -f cmd/ruuvi/ruuvi
	rm -f coverage.txt
	rm -rf dist/

# Install the CLI tool
install:
	cd cmd/ruuvi && go install .
