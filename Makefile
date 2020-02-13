VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.Version=$(VERSION)' \
           -X 'main.Revision=$(REVISION)'
GOIMPORTS ?= goimports
GOCILINT ?= golangci-lint
GO ?= GO111MODULE=on go

.DEFAULT_GOAL := help

.PHONY: fmt
fmt: ## Formatting source codes.
	@$(GOIMPORTS) -w .

.PHONY: lint
lint: ## Run golint and go vet.
	@$(GOCILINT) run --no-config --disable-all --enable=goimports --enable=misspell ./...

.PHONY: test
test:  ## Run the tests.
	@$(GO) test ./...

.PHONY: build
build: ## Build a binary.
	$(GO) build -o gh-prompt -ldflags "$(LDFLAGS)" ./cmd/gh-prompt/main.go

.PHONY: help
help: ## Show help text
	@echo "Commands:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-20s\033[0m %s\n", $$1, $$2}'
