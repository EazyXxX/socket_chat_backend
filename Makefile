.SILENT:

include .env

BINARY_NAME=socket_server
DB_PASS=$(DB_PASSWORD)

build: 
		docker compose build socket_chat

run:
		docker compose up socket_chat

migrate:
		migrate -path ./schema -database "postgres://postgres:$(DB_PASS)@0.0.0.0:6000/postgres?sslmode=disable" up

test:
		go test -v ./...

gobuild:
		go build -o $(BINARY_NAME) ./cmd/app

clean:
		rm -f $(BINARY_NAME)

dev: gobuild
		./$(BINARY_NAME)

.DEFAULT_GOAL := dev
