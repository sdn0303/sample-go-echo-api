.PHONY: build up rebuild setup migrate

up: build start logs
rebuild: down build start logs

build:
	COMPOSE_DOCKER_CLI_BUILD=1 DOCKER_BUILDKIT=1 docker compose build --no-cache

start:
	docker compose up -d

logs:
	docker compose logs -f -t api

down:
	docker compose down --rmi all --volumes --remove-orphans

exec-api:
	docker compose exec api sh

setup:
	pip install pre-commit \
	&& ./scripts/setup.sh \
  	&& pre-commit install

swagger:
	swag i -d ./cmd/smaple-go-echo-api

migrate:
	go generate ./pkg/ent \
	&& go run cmd/migration/main.go
