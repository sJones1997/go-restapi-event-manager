# Directories
SRC_DIR := .
BIN_DIR := ./bin
BINARY_NAME := go-restapi-event-manager

run:
	go run cmd/$(BINARY_NAME)/main.go

# Build target
build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)/cmd/$(BINARY_NAME)/main.go

# Clean target
clean:
	@echo "Cleaning..."
	$(GO_CLEAN)
	rm -rf $(BIN_DIR)

test:
	go test ./tests/internal/...

.PHONY: build clean
