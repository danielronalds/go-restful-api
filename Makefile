build:
	docker compose up --build api

build-all:
	docker compose up --build

run:
	docker compose start

stop:
	docker compose stop
