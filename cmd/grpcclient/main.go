package main

import (
	"context"
	"fmt"
	"log"
	itemsv1 "riffwire/internal/pb/items/v1"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const URL = "localhost:50051"

func main() {

	conn, err := grpc.NewClient(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("GRPC Client Err: %v", err)
	}

	defer conn.Close()

	client := itemsv1.NewItemsServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.ListItems(ctx, &itemsv1.ListItemsRequest{})
	if err != nil {
		log.Fatalf("Response Err %v", err)
	}

	for _, v := range resp.Items {
		fmt.Printf("ID=%d Name=%s\n", v.Id, v.Name)
	}
}
