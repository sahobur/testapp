help: ## Show this help
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

build: ## Build the application
	go build -o ./bin/server ./main.go

run: ## Run the application
	./bin/server

docker-build:
	docker build . -t server


