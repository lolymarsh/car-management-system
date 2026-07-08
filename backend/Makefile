.PHONY: run migrate-up migrate-down

run:
	go run main.go

# DATABASE_URL="postgres://postgres:postgres@localhost:5432/car_management?sslmode=disable" make migrate-up
migrate-up:
	goose -dir migrations postgres "$(DATABASE_URL)" up

migrate-down:
	goose -dir migrations postgres "$(DATABASE_URL)" down
