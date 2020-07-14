package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

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
	if r.Method == http.MethodPut {
		// look for ID in the URI using regexp
		regx := regexp.MustCompile(`/([0-9]+)`)
		group := regx.FindAllStringSubmatch(r.URL.Path, -1)

		if len(group) != 1 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}
		if len(group[0]) != 2 {
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := group[0][1]
		id, _ := strconv.Atoi(idString)

		p.logger.Printf("Got ID: %d", id)
		p.updateProduct(id, w, r)
		return
	}
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

// addProduct reads request body and creates new product
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

// updateProduct edit product identified by id
func (p *Products) updateProduct(id int, w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle 'PUT' request")
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
	err = data.UpdateProduct(id, newProd)
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
