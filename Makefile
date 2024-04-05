init: migrate
	go build main.go

migrate:
	cp .env.example .env
	cd db
	goose sqlite3 ./big_bang_theory_quotes.db up
	cd ..

down:
	goose down

run:
	./main

