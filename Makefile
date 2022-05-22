.DEFAULT_GOAL:=help

##@ Application
start: init down up dev ## Initialize local development environment
	echo "FINITO"

init: ## Load initial project setup
	cp .env.local .env
	go mod download

dev: ## Run app in development mode
	~/go/bin/air -c air.toml -d

##@ Containers
up: ## Run containers
	docker-compose up -d --build

down: ## Stop and delete containers and volumes
	docker-compose down -v

stop: ## Stops container for database
	docker-compose stop

##@ Helpers
.PHONY: help
help: ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
