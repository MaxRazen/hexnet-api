SHELL=/bin/bash

include .env
export

help: ## Print this help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

test: ## Test applications
	go test -v ./...

build: ## Build applications
	go build ./app.go
	go build -o ./cli/ ./cli/cli.go

run-app: ## Build and run App
	go run ./app.go

run-cli: ## Build and run App
	go run ./cli/cli.go

drop-builds: ## Remove all builds
	rm -f app
	rm -f ./cli/cli
