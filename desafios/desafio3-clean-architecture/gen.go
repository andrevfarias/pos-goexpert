package gen

//go:generate protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
//go:generate go run github.com/99designs/gqlgen generate
//go:generate wire gen ./cmd/ordersystem
