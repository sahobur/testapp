help: ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o ./bin/server ./cmd/server/main.go
	go build -o ./bin/receiver ./cmd/receiver/main.go

docker-run: ## build and run in docker compose
	docker compose build
	docker compose up





