GOCMD=go
APP_NAME=willserver
BINARY_DIR=bin

GOOS=linux
GOARCH=amd64
CGO_ENABLED=0
LDFLAGS=-ldflags="-s -w"

# Deploy settings
SSH_KEY=/home/will/.ssh/id_hetzner
SERVER=46.225.218.39
USER=root
REMOTE_PATH=/usr/local/bin

.PHONY: deps
deps: ## Install and tidy Go dependencies
	$(GOCMD) mod tidy

.PHONY: clean
clean: ## Clean build artifacts
	$(GOCMD) clean
	rm -rf $(BINARY_DIR) coverage.out

.PHONY: test
test: ## Run all tests
	$(GOCMD) test ./... -coverprofile=coverage.out

.PHONY: build
build: build-linux ## Build production binary

.PHONY: build-local
build-local: ## Build for local system
	mkdir -p $(BINARY_DIR)
	$(GOCMD) build -o $(BINARY_DIR)/$(APP_NAME) ./cmd/server/

.PHONY: build-linux
build-linux: ## Build Linux production binary
	mkdir -p $(BINARY_DIR)
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=$(CGO_ENABLED) \
	$(GOCMD) build $(LDFLAGS) -o $(BINARY_DIR)/$(APP_NAME) ./cmd/server/



.PHONY: deploy
deploy: build-linux ## Build and copy to server using SSH key
	scp -i $(SSH_KEY) $(BINARY_DIR)/$(APP_NAME) $(USER)@$(SERVER):$(REMOTE_PATH)/$(APP_NAME)

	scp -i $(SSH_KEY) -r cmd/server/config \
		$(USER)@$(SERVER):$(REMOTE_PATH)/

	scp -i $(SSH_KEY) -r ../web \
		$(USER)@$(SERVER):$(REMOTE_PATH)/


.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-25s\033[0m %s\n", $$1, $$2}'