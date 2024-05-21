
FROM golang:1.22-alpine AS builder

WORKDIR /app/

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .
RUN go build -o ./build ./cmd/main.go

FROM alpine:3.19

WORKDIR /app/

COPY --from=builder /app/build/ .

COPY .env .
COPY --from=builder /app/goose/migrations ./migrations

# Change permissions of the built binary
RUN chmod +x ./build

EXPOSE 1337

RUN ls -la /app/build
# Define the command to execute when the container starts
CMD ["sh", "-c", "./build migrate_up && ./build"]
