create_migrate:
	goose -dir=assets/migrations create familytree sql

migrateup:
	goose -dir=assets/migrations sqlite3 database.db up

migratedown:
	goose -dir=assets/migrations sqlite3 database.db down

run:
	go run cmd/web/main.go

build:
	go install github.com/pressly/goose/cmd/goose@latest
	make migrateup
	go build -o app cmd/web/main.go
	./app


PHONY: create_migrate migrateup migratedown run
