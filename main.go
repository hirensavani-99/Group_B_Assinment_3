package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

// Item represents an item with ID, Name, Description, and Price
type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ItemStorage represents the in-memory storage for items

var items = []Item{}

const port = ":8080"

// common function for response writing
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	fmt.Print(w, message)
}

// Adding Items to storage
func handleAddItem(w http.ResponseWriter, r *http.Request) {

	//if method is not post request
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}
	//body not allowed
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		respondWithError(w, http.StatusBadRequest, "Malformed request body")
		return
	}

	newItem.ID = uuid.New().String() // Generate a unique ID for the item
	items = append(items, newItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// Get list of Items from storage
func handleGetItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Handle GET requests to /get/items
	if r.URL.Path != "/get/items" {
		http.NotFound(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)

}

// Get list of Items from storage
func getItemById(w http.ResponseWriter, r *http.Request) {

	id := strings.TrimPrefix(r.URL.Path, "/get/itemById/")

	var item Item

	if id == "0" || id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Find the item in the data map
	for _, v := range items {
		if v.ID == id {
			item = v
			break
		}
	}
	// If no item found send error
	if item.ID == "" {
		http.Error(w, "Item Not Found", http.StatusNotFound)
		return
	}

	// Serialize the item to JSON and write it to the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

func main() {
	http.HandleFunc("/post/items", handleAddItem)
	http.HandleFunc("/get/items", handleGetItem)
	http.HandleFunc("/get/itemById/", getItemById)

	fmt.Printf("The server is running at port #%v", port)

	log.Fatal(http.ListenAndServe(port, nil))
}
