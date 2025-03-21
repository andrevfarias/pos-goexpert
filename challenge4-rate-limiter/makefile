start:
	docker compose up -d

stop:
	docker compose down

build:
	docker compose build

test:
	@docker compose run --rm k6 run /scripts/test.js

.PHONY: test start stop build
