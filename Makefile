build:
	@go build -o out cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./out

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations -seq $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/main.go up

migrate-down:
	@go run cmd/migrate/main.go down