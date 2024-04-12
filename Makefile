build:
	@go build -o out cmd/main.go

test:
	@go test -v ./...
	
run: build
	@./out
