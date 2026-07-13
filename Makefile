.PHONY: dev build run test sqlc migrate-up migrate-down

dev:
	@go run ./cmd/server/main.go

build:
	@go build -o bin/server ./cmd/server/main.go

run:
	@./bin/server

test:
	@go test ./... -v -count=1

sqlc:
	@sqlc generate

include .env
export

migrate-up:
	goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DATABASE_URL)" down

migrate-create:
	@read -p "Migration name: " name; \
	goose -dir migrations create $$name sql

lint:
	@go vet ./...
	@staticcheck ./...
