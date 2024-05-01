# Directories
SRC_DIR := .
BIN_DIR := ./bin
BINARY_NAME := go-restapi-event-manager

# Build target
build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)/cmd/$(BINARY_NAME)/main.go

# Clean target
clean:
	@echo "Cleaning..."
	$(GO_CLEAN)
	rm -rf $(BIN_DIR)

.PHONY: build clean
