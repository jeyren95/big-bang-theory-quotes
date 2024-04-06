build: migrate.up
	go build main.go

migrate.up:
	cp .env.example .env
	cd db
        include .env
	DB_MIGRATIONS_DIR = $(PWD)/db/migrations
	EXPORT SQLITE_DSN
	
	goose -dir  $(DB_MIGRATIONS_DIR) sqlite3 $(SQLITE_DSN) up
	cd ..

migrate.down:
	goose -dir $(DB_MIGRATIONS_DIR) sqlite3 $(SQLITE_DSN) down

run:
	./main

