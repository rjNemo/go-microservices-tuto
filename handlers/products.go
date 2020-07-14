package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rjNemo/go-micro/products/data"
)

// Products is a handler for Products API service
type Products struct {
	logger *log.Logger
}

// NewProducts creates a Products handler
func NewProducts(logger *log.Logger) *Products { return &Products{logger: logger} }

func (p *Products) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get resources
	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}
	// create one resource
	if r.Method == http.MethodPost {
		p.addProduct(w, r)
		return
	}
	// update one resource
	// if r.Method == http.MethodPut {
	// 	p.updateProduct(w,r)
	// 	return
	// }
	// catch all other HTTP requests
	w.WriteHeader(http.StatusMethodNotAllowed)
}

// getProducts writes all products to response in JSON format
func (p *Products) getProducts(w http.ResponseWriter, r *http.Request) {
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

func (p *Products) addProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle 'POST' request")
	// create a new product
	newProd := &data.Product{}
	// deserialize JSON to product
	err := newProd.FromJSON(r.Body)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to decode data: %s\n", err)
		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}
	p.logger.Printf("product: %#v", newProd)
	data.AddProduct(newProd)
}
