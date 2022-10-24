
GOLANGCI_LINT_VER=v1.50.1

.PHONY: lint
lint: ## Run Go linters in a Docker container
	docker run \
		--rm \
		-v $(PWD):/go/src/$(PROJECT) \
		-w /go/src/$(PROJECT) \
		-e GO111MODULE=on \
		-e GOPROXY=https://proxy.golang.org \
		golangci/golangci-lint:$(GOLANGCI_LINT_VER) \
			golangci-lint run

