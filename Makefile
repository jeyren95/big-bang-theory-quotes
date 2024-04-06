include .env
DB_MIGRATIONS_DIR = $(PWD)/db/migrations
EXPORT SQLITE_DSN

run: migrate.up
	go run main.go

migrate.up:
	goose -dir  $(DB_MIGRATIONS_DIR) sqlite3 $(SQLITE_DSN) up

migrate.down:
	goose -dir $(DB_MIGRATIONS_DIR) sqlite3 $(SQLITE_DSN) down
