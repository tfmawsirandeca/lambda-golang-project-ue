# Project directories
CMD_DIR := cmd
BIN_DIR := bin
BUILD_DIR := $(BIN_DIR)/main

# Go build parameters
GO_FILES := $(shell find $(CMD_DIR) -type f -name '*.go')
GO_MOD := go.mod
GO_SUM := go.sum

# Build the Go binary and create the bootstrap file
build: $(BIN_DIR)/bootstrap

$(BUILD_DIR): $(GO_FILES) $(GO_MOD) $(GO_SUM)
	@echo "Building Go binary..."
	@mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o main $(CMD_DIR)/main.go
#	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR) $(CMD_DIR)/main.go

$(BIN_DIR)/bootstrap: $(BUILD_DIR)
	@echo "Creating bootstrap file..."
	@echo '#!/bin/sh' > $(BIN_DIR)/bootstrap
	@echo 'exec ./main' >> $(BIN_DIR)/bootstrap
	@chmod +x $(BIN_DIR)/bootstrap

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

.PHONY: all build clean

