help: ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o ./bin/server ./main.go

run: ## Run the application
	./bin/server

docker-run: ## build and run in docker compose
	docker compose build
	docker compose up





