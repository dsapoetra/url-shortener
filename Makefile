# Database credentials
DB_HOST ?= 
DB_PORT ?= 
DB_USER ?= 
DB_PASSWORD ?= 
DB_NAME ?= 

# Database URL for migrations
DB_URL = postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable

.PHONY: migrate-create migrate-up migrate-down migrate-force

# Create a new migration file
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations -seq $$name

# Run all migrations
migrate-up:
	migrate -path migrations -database "${DB_URL}" up

# Rollback all migrations
migrate-down:
	migrate -path migrations -database "${DB_URL}" down

# Force set migration version
migrate-force:
	@read -p "Enter version: " version; \
	migrate -path migrations -database "${DB_URL}" force $$version

# Create initial migration
create-users-migration:
	migrate create -ext sql -dir migrations -seq create_users_table

# Install golang-migrate
install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# Run tests
test:
	go test -v ./...

# Format code
fmt:
	go fmt ./...

# Help
help:
	@echo "Available commands:"
	@echo "  make migrate-create    - Create a new migration file"
	@echo "  make migrate-up        - Run all migrations"
	@echo "  make migrate-down      - Rollback all migrations"
	@echo "  make migrate-force     - Force set migration version"
	@echo "  make install-migrate   - Install golang-migrate tool"
	@echo "  make test              - Run tests"
	@echo "  make fmt               - Format code"
	@echo "  make help              - Show this help message"
