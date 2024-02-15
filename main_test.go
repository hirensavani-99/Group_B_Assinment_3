package main_test

import "testing"

//for positive
func TestAddItemHandlerPositive(t *testing.T) {

	//create a item to be added , make a post request and check responses for positive case (201 & expcted Item in list )

}

//for positive
func TestAddItemHandlerNagative(t *testing.T) {

	//create a item to be added with invalid json , make a post request and check responses for nagative case (status code accordingly )

}

func TestGetItemHandlerPositive(t *testing.T) {
	// create get request and look for postive responses 200 status code & Item count or specfic item
}

func TestGetItemHandlerNagative(t *testing.T) {
	// create get request with invalid url and look for nagative response from server
}