.PHONY: app-run tidy docker-up docker-down migrate-up migrate-down protoc-generate wire-generate gqlgen-generate

app-run:
	go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

tidy:
	go mod tidy

docker-up:
	docker compose up -d

docker-up-build:
	docker compose up --build -d

docker-down:
	docker compose down

migrate:
	docker run -v "$(pwd)/internal/infra/sql/migrations:/migrations" --network host migrate/migrate -path=/migrations/ -database "mysql://root:root@tcp(localhost:3306)/orders" up
	migrate -path=internal/infra/sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

protoc-generate:
	protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

wire-generate:
	wire gen ./cmd/ordersystem

gqlgen-generate:
	go run github.com/99designs/gqlgen generate