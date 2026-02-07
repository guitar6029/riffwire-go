package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"riffwire/internal/models"
	"riffwire/internal/store"

	"riffwire/internal/grpcapi"

	itemsv1 "riffwire/internal/pb/items/v1"
)

func main() {
	fmt.Println("Starting grpc on port 50051")
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("GRPC ERROR : %v", err)
	}

	grpcServer := grpc.NewServer()

	grpcServer.Serve(listener)

	itemsDefault := []models.Item{
		{ID: 1, Name: "Lamp"},
		{ID: 2, Name: "Table"},
		{ID: 3, Name: "Cable"},
	}

	s := store.NewInMemoryStore(itemsDefault)
	itemsSrv := grpcapi.NewItemsServer(s)

	itemsv1.RegisterItemsServiceServer(grpcServer, itemsSrv)

}
