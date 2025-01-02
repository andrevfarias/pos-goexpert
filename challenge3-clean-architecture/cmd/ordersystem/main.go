package main

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"

	graphql_handler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/configs"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/event/handler"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/infra/graph"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/infra/grpc/pb"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/infra/grpc/service"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/internal/infra/web/webserver"
	"github.com/andrevfarias/pos-goexpert/challenge3-clean-architecture/pkg/events"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Printf("Connecting to RabbitMQ on: %s:%s\n", configs.RabbitMQHost, configs.RabbitMQPort)
	conn, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s:%s/", configs.RabbitMQHost, configs.RabbitMQPort))
	if err != nil {
		panic(err)
	}
	rabbitMQChannel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	eventDispatcher := events.NewEventDispatcher()
	eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	})

	webserver := webserver.NewWebServer(fmt.Sprintf(":%s", configs.WebServerPort))
	webOrderHandler := NewWebOrderHandler(db, eventDispatcher)
	webserver.AddHandler(http.MethodPost, "/orders", webOrderHandler.Create)
	webserver.AddHandler(http.MethodGet, "/orders", webOrderHandler.FindAll)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	createOrderUseCase := NewCreateOrderUseCase(db, eventDispatcher)
	getOrdersUseCase := NewGetOrdersUseCase(db)

	grpcServer := grpc.NewServer()
	OrderService := service.NewOrderService(*createOrderUseCase, *getOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, OrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
		GetOrdersUseCase:   *getOrdersUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
	http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}
