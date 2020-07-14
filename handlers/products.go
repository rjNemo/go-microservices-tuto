package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rjNemo/go-micro/products/data"
)

// Products is a handler for Products API service
type Products struct {
	logger *log.Logger
}

// NewProducts creates a Products handler
func NewProducts(logger *log.Logger) *Products { return &Products{logger: logger} }

// GetProducts writes all products to response in JSON format
func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle 'GET' request")
	// fetch products from the datastore
	productList := data.AllProducts()
	// serialize list to JSON
	err := productList.ToJSON(w)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to encode request: %s\n", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
}

// AddProduct reads request body and creates new product
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle 'POST' request")
	// get product from the request
	newProd := r.Context().Value(KeyProduct{}).(*data.Product) // cast into a Product

	p.logger.Printf("product: %#v", newProd)
	data.AddProduct(newProd)
}

// UpdateProduct edit product identified by id
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.logger.Println("Handle 'PUT' request", id)
	// get product from the request
	newProd := r.Context().Value(KeyProduct{}).(*data.Product) // cast into a Product

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
		newProd := &data.Product{}
		// deserialize JSON to product
		err := newProd.FromJSON(r.Body)
		if err != nil {
			errMsg := fmt.Sprintf("Unable to decode data: %s\n", err)
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
