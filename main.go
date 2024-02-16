package main

import "net/http"

// Item represents an item with ID, Name, Description, and Price
type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ItemStorage represents the in-memory storage for items
type ItemStorage struct {
	Items []Item
}

//Adding Items to storage
func handleAddItem(w http.ResponseWriter, r *http.Request) {

}

//Get list of Items from storage
func handleGetItem(w http.ResponseWriter, r *http.Request) {

}

func main() {
}
