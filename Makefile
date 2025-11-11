help: ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o ./bin/server ./cmd/main.go

run: ## Run the application
	./bin/server