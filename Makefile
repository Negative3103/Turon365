
APP_NAME=turon365
DOCKER_COMPOSE_FILE=docker-compose.yml

build:
	go build -o $(APP_NAME) ./cmd/main.go

run: build
	./$(APP_NAME)

up:
	docker-compose up -d

down:
	docker-compose down

clean:
	rm -f $(APP_NAME)

rebuild: clean up build run
.PHONY: build run up down clean rebuild
