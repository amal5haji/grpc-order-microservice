package main

import (
	"context"
	"log"
	"net/http"
	"ordering-microservice/protogen/golang/orders"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	con, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error on connecting to grpc server")
	}

	defer con.Close()

	mux := runtime.NewServeMux()
	if err = orders.RegisterOrdersHandler(context.Background(), mux, con); err != nil {
		log.Fatal("Error in registering order service")
	}
	log.Println("Client Listening on port 3000")
	if err := http.ListenAndServe("localhost:3000", mux); err != nil {
		log.Fatal("Error in setting up http server")
	}

}
