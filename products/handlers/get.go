package handlers

import (
	"fmt"
	"net/http"

	"github.com/rjNemo/go-micro/products/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
// responses:
//	200: productsResponse

// GetProducts returns all products to response in JSON format
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

// swagger:route GET /products/{id} products getProduct
// Return a list of products from the database
// responses:
//	200: productResponse
//	404: errorResponse

// GetOneProduct handles GET requests
func (p *Products) GetOneProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.logger.Println("[DEBUG] get record id", id)
	prod, err := data.GetProductByID(id)

	switch err {
	case nil:

	case data.ErrorProductNotFound:
		p.logger.Println("[ERROR] fetching product", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return

	default:
		p.logger.Println("[ERROR] fetching product", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = prod.ToJSON(w)
	if err != nil {
		// we should never be here but log the error just incase
		p.logger.Println("[ERROR] serializing product", err)
	}
}
