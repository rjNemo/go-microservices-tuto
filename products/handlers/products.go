package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Products is a handler for Products API service
type Products struct {
	logger *log.Logger
}

// New creates a Products handler
func New(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

// KeyProduct is a key used to pass validated product to handler
type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
