build:
	@go build -o build ./cmd/main.go

run: build
	@./build

migrate_status:
	@go run ./goose/main.go status

migrate_up:
	@go run ./goose/main.go up

