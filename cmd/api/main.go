package main

import (
	"fmt"
	"log"
	"net/http"
	"riffwire/internal/httpapi"
	"riffwire/internal/models"
	"riffwire/internal/store"
)

var itemsDefault = []models.Item{
	{ID: 1, Name: "Lamp"},
	{ID: 2, Name: "Table"},
	{ID: 3, Name: "Cable"},
}

func main() {
	s := store.NewInMemoryStore(itemsDefault)
	// regiser handler
	http.HandleFunc("/", httpapi.RootHandler)
	http.HandleFunc("/health", httpapi.HealthHandler)
	http.HandleFunc("/items", httpapi.Logger(httpapi.ItemsHandler(s)))
	http.HandleFunc("/items/", httpapi.Logger(httpapi.ItemsByIDHandler(s)))

	fmt.Println("Starting a server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
