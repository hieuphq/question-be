
APP_NAME=question-be
DEFAULT_PORT=3000
.PHONY: init build dev test 

init:
	cp .env.sample .env && vim .env

build:
	go build -o bin/server $(shell pwd)/cmd/server/

dev:
	go run ./cmd/server/main.go

run: dev

test:
	go test -cover ./...

docker-build:
	docker build \
	--build-arg DEFAULT_PORT="${DEFAULT_PORT}" \
	-t ${APP_NAME}:latest .