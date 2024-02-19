package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// for positive
func TestAddItemHandlerPositive(t *testing.T) {

	//create a item to be added , make a post request and check responses for positive case (201 & expcted Item in list )
	newItem := Item{
		Name:        "New Item",
		Description: "Description of New Item",
		Price:       15.23,
	}

	// Encode the new item as JSON
	jsonData, err := json.Marshal(newItem)
	if err != nil {
		t.Fatal(err)
	}

	// Create a POST request to add the item
	req, err := http.NewRequest("POST", "/post/items", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the POST request
	handleAddItem(rr, req)

	// Check if the status code is 201 (Created)
	if rr.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, got %d", http.StatusCreated, rr.Code)
	}

	// Decode the response body into an Item struct
	var responseItem Item
	if err := json.NewDecoder(rr.Body).Decode(&responseItem); err != nil {
		t.Fatal(err)
	}

	// Check if the returned item matches the added item
	if responseItem.Name != newItem.Name || responseItem.Description != newItem.Description || responseItem.Price != newItem.Price {
		t.Errorf("Expected item %+v, got %+v", newItem, responseItem)
	}

}

// for negative
func TestAddItemHandlerNegative(t *testing.T) {

	//create a item to be added with invalid json , make a post request and check responses for negative case (status code accordingly )
	invalidJSON := []byte(`{"name": "New Item", "description": "Description of New Item", "price": "invalid_price"}`)

	// Create a POST request with invalid JSON
	req, err := http.NewRequest("POST", "/post/items", bytes.NewBuffer(invalidJSON))
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the POST request
	handleAddItem(rr, req)

	// Check if the status code is 400 (Bad Request) for invalid JSON
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
	//Request is Invalid
	//400 Bad Request as a response
}

func TestGetItemHandlerPositive(t *testing.T) {
	// create get request and look for postive responses 200 status code & Item count or specfic item

	// Define some mock items to be returned by the handler
	mockItems := []Item{
		{ID: "1", Name: "Item 1", Description: "Description 1", Price: 20.99},
		{ID: "2", Name: "Item 2", Description: "Description 2", Price: 20.99},
	}

	// Mock the global items slice with the mock items
	items = mockItems

	// Create a GET request to the /get/items endpoint
	req, err := http.NewRequest("GET", "/get/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	rr := httptest.NewRecorder()

	// Call the handler function with the GET request
	handleGetItem(rr, req)

	// Check if the status code is 200 (OK)
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	// Decode the response body into a slice of Item structs
	var responseItems []Item
	if err := json.NewDecoder(rr.Body).Decode(&responseItems); err != nil {
		t.Fatal(err)
	}

	// Verify the correctness of the response
	if len(responseItems) != len(mockItems) {
		t.Errorf("Expected %d items, got %d", len(mockItems), len(responseItems))
	}

	// Compare each item in the response with the corresponding mock item
	for i, item := range mockItems {
		if responseItems[i].ID != item.ID || responseItems[i].Name != item.Name ||
			responseItems[i].Description != item.Description || responseItems[i].Price != item.Price {
			t.Errorf("Expected item %v, got %v", item, responseItems[i])
		}
	}
}

func TestGetItemHandlerNegative(t *testing.T) {
	// Create a GET request
	req, err := http.NewRequest("GET", "/get/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Call the handler function with the GET request
	handleGetItem(rr, req)

	// Check if the status code is 404 (Not Found) for a non-existent resource
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, got %d", http.StatusNotFound, rr.Code)
	}
}

func TestGetItemById(t *testing.T) {

	// Define some mock items to be returned by the handler
	mockItems := []Item{
		{ID: "1", Name: "Item 1", Description: "Description 1", Price: 20.99},
		{ID: "2", Name: "Item 2", Description: "Description 2", Price: 20.99},
	}

	// Mock the global items slice with the mock items
	items = mockItems

	// Create a mock HTTP server
	srv := httptest.NewServer(http.HandlerFunc(getItemById))
	defer srv.Close()

	// Prepare a request
	req, err := http.NewRequest("GET", srv.URL+"/get/itemById/1", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, resp.StatusCode)
	}

	// Decode the response body
	var item Item
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Validate the response
	expectedID := "1" // Assuming this is the ID of the item in your test scenario
	if item.ID != expectedID {
		t.Errorf("Expected item ID to be %s, got %s", expectedID, item.ID)
	}
}
