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
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rjNemo/go-micro/products/data"
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

// New creates a Products handler
func New(logger *log.Logger) *Products {
	return &Products{logger: logger}
}

// AddProduct reads request body and creates new product
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle 'POST' request")
	// get product from the request
	newProd := r.Context().Value(KeyProduct{}).(*models.Product) // cast into a Product

	p.logger.Printf("product: %#v", newProd)
	data.AddProduct(newProd)
}

// UpdateProduct edit product identified by id
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.logger.Println("Handle 'PUT' request", id)
	// get product from the request
	newProd := r.Context().Value(KeyProduct{}).(*models.Product) // cast into a Product

	p.logger.Printf("product: %#v", newProd)
	err := data.UpdateProduct(id, newProd)
	if err == data.ErrorProductNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	if err != nil {
		errMsg := fmt.Sprintf("something went wrong: %s", err.Error())
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
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
