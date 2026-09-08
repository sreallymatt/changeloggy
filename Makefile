default: build

build: ## Build the binary
	@go build -o changeloggy .

lint: ## Check source code with golangci-lint
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./...

lint-fix: ## Fix source code with golangci-lint
	@echo "==> Fixing source code with golangci-lint..."
	@golangci-lint run ./... --fix

.PHONY: build lint lint-fix
