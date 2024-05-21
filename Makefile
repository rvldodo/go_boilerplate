run:
	@go run ./cmd/main.go

migrate_status:
	@go run ./goose/main.go status

migrate_up:
	@go run ./goose/main.go up
