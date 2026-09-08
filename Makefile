default: build

build: ## Build the binary
	@go build -o changeloggy .

lint: ## Check source code with golangci-lint
	@echo "==> Checking source code with golangci-lint..."
	@golangci-lint run ./...

lint-fix: ## Fix source code with golangci-lint
	@echo "==> Fixing source code with golangci-lint..."
	@golangci-lint run ./... --fix

depscheck: ## Check that go.mod/go.sum and vendor/ are in sync
	@go mod tidy
	@git diff --exit-code -- go.mod go.sum || \
		(echo; echo "Unexpected difference in go.mod/go.sum files. Run 'go mod tidy' command or revert any go.mod/go.sum changes and commit."; exit 1)
	@go mod vendor
	@git diff --compact-summary --exit-code -- vendor || \
		(echo; echo "Unexpected difference in vendor/ directory. Run 'go mod vendor' command or revert any go.mod/go.sum/vendor changes and commit."; exit 1)

.PHONY: build lint lint-fix
