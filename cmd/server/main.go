package main

import (
	"log"
	"net"
	"ordering-microservice/internal"
	"ordering-microservice/protogen/golang/orders"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatal("Failed to listen on port 8080")
	}

	server := grpc.NewServer()

	db := internal.NewDB()
	orderService := internal.NewOrderService(db)

	orders.RegisterOrdersServer(server, &orderService)

	log.Println("Server is listening!")

	if err = server.Serve(listener); err != nil {
		log.Fatal("Failed in starting server")
	}

}
