package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

var itemsDefault = []Item{
	{ID: 1, Name: "Lamp"},
	{ID: 2, Name: "Table"},
	{ID: 3, Name: "Cable"},
}

func writeErrorJSON(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, ErrorResponse{Error: message})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func itemsByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	//extract id
	idStr := strings.TrimPrefix(r.URL.Path, "/items/")

	if len(idStr) == 0 {
		writeErrorJSON(w, http.StatusBadRequest, "id is empty")
		return
	}

	// parse str into int
	parsedId, err := strconv.Atoi(idStr)
	if err != nil {
		writeErrorJSON(w, http.StatusBadRequest, "invalid id")
		return
	}

	// helper find id in the items default
	found := findEntityInArray(parsedId, itemsDefault)
	if found == nil {
		writeErrorJSON(w, http.StatusNotFound, "item not found")
		return
	}

	writeJSON(w, http.StatusOK, found)
}

// pass entity id and array
func findEntityInArray(id int, items []Item) *Item {
	for i := range items {
		if items[i].ID == id {
			return &items[i]
		}
	}
	return nil
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeErrorJSON(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	writeJSON(w, http.StatusOK, itemsDefault)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	writeJSON(w, http.StatusOK, map[string]bool{"ok": true})
}

func rootHandler(w http.ResponseWriter, r *http.Request) {

	writeJSON(w, http.StatusOK, map[string]string{"message": "Hello"})

}

func logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next(w, r)
	}
}

func main() {
	fmt.Println("Hello")
	// regiser handler
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/items", logger(itemsHandler))
	http.HandleFunc("/items/", logger(itemsByIDHandler))

	fmt.Println("Starting a server at port 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
