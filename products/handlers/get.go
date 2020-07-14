package handlers

import (
	"fmt"
	"net/http"

	"github.com/rjNemo/go-micro/products/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
// 	200: productsResponse

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
