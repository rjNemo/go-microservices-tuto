// Package classification Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// 	- application/json
//
// Produces:
// 	- application/json
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/rjNemo/go-micro/products/models"
)

// Products is a handler for Products API service
type Products struct {
	logger *log.Logger
}

// list of products in the response. For go-swagger
// swagger:response productsResponse
type productsResponse struct {
	// All products in the datastore
	// in: body
	Body []models.Product
}

// product in the response. For go-swagger
// swagger:response productResponse
type productResponse struct {
	// One product in the datastore
	// in: body
	Body models.Product
}

// New creates a Products handler
func New(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

// KeyProduct is a key used to pass validated product to handler
type KeyProduct struct{}

// ProductValidationMiddleware validates the data passed by the user
func (p *Products) ProductValidationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create a new product
		newProd := &models.Product{}
		// deserialize JSON to product
		err := newProd.FromJSON(r.Body)
		if err != nil {
			p.logger.Printf("Error deserializing %v", err)
			errMsg := fmt.Sprintf("Unable to decode data: %s\n", err)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}
		// validate the product
		err = newProd.Validate()
		if err != nil {
			p.logger.Printf("Error deserializing %v", err)
			errMsg := fmt.Sprintf("Validation error: %s\n", err)
			http.Error(w, errMsg, http.StatusBadRequest)
			return
		}

		// add product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, newProd)
		req := r.WithContext(ctx)

		// call the next handler
		next.ServeHTTP(w, req)
	})
}
