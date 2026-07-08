.PHONY: setup-local run migrate-up migrate-down

DB_URL := "postgres://postgres:postgres@localhost:5432/car_management?sslmode=disable"

setup-local:
	docker compose up -d
	sleep 2
	docker compose exec -T postgres psql -U postgres -c "CREATE DATABASE car_management;" 2>/dev/null || true
	goose -dir migrations postgres $(DB_URL) up

run:
	go run main.go

migrate-up:
	goose -dir migrations postgres $(DB_URL) up

migrate-down:
	goose -dir migrations postgres $(DB_URL) down
