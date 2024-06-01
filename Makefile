build:
	docker compose up --build api

build-all:
	docker compose up --build

run:
	docker compose start

dev:
	go run ./src/

stop:
	docker compose stop

run-db:
	docker compose start postgres
