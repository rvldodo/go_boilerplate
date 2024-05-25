FROM golang:1.22-alpine

RUN apk update && apk add --no-cache make git

WORKDIR /usr/src

COPY go.mod go.sum ./
COPY .env ./

RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

COPY . .

RUN make migrate_status
RUN make migrate_up


CMD ["air"]
