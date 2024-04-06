include .env
DB_MIGRATIONS_DIR = $(PWD)/db/migrations

migrate.up:
	goose -dir  $(DB_MIGRATIONS_DIR) sqlite3 $(SQLITE_DSN) up

migrate.down:
	goose -dir $(DB_MIGRATIONS_DIR) sqlite3 $(SQLITE_DSN) down

migrate.create:
	goose -dir $(DB_MIGRATIONS_DIR) create $(MIGRATION_NAME) sql

build:
	go build main.go

run: build
	./main