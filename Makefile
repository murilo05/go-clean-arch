run: ## Run the application locally
	@echo "Running application..."
	@go run ./app/main.go

test: ## Run tests
	@echo "Running tests..."
	@go test -v ./...

swagger: ## Generate Swagger documentation
	@echo "Generating Swagger docs..."
	@swag init -g app/main.go -o ./docs
	@echo "Swagger docs generated!"

docker-up: ## Start all services with docker-compose
	@echo "Starting services..."
	@docker compose up -d
	@echo "Services started! Application available at http://localhost:8080"

docker-down: ## Stop all services
	@echo "Stopping services..."
	@docker-compose down
	@echo "Services stopped!"

migrate-up:
	@echo "Running migrations UP..."
	@go run migrations/migrate.go -action=up

migrate-down:
	@echo "Running migrations DOWN..."
	@go run migrations/migrate.go -action=down

