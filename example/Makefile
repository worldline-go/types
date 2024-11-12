.DEFAULT_GOAL := help

.PHONY: postgres
postgres: ## Start a postgres container
	docker run -d --name postgres -e POSTGRES_HOST_AUTH_METHOD=trust -p 5432:5432 -d postgres:15-alpine

.PHONY: postgres-remove
postgres-remove: ## Remove the postgres container
	docker rm -f postgres

.PHONY: run
run: ## Run the application
	go run main.go

.PHONY: help
help: ## Display this help screen
	@printf "First init the database with the following command:\n  make postgres\nThen run the application with the following command:\n  make run\n__________________________________________________\n"
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
