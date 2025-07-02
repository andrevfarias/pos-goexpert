start:
	docker compose up -d

stop:
	docker compose down

start-dev:
	docker compose --profile development run --rm -it dev bash

test:
	docker compose --profile development run --rm -it dev go test ./...