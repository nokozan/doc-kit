.PHONY: clean critic security lint test build run

APP_NAME = doc-kit
BUILD_DIR = $(PWD)/build
# MIGRATIONS_FOLDER = $(PWD)/platform/migrations
# DATABASE_URL = postgres://postgres:password@cgapp-postgres/postgres?sslmode=disable

run:
	go run main.go
	
build:
	CGO_ENABLED=0 GOOS=linux go build -o $(BUILD_DIR)/$(APP_NAME) -ldflags="-s -w" main.go

migrate:
	go run cmd/migrate/main.go

# db.reset:
# 	docker 

lint:
	golangci-lint run --timeout 5m

test:
	go test -v ./... -coverprofile=coverage.out