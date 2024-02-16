package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)

}

func main() {
	http.HandleFunc("/post/items", handleAddItem)
	http.HandleFunc("/get/items", handleGetItem)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
