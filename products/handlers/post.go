package handlers

import (
	"net/http"

	"github.com/rjNemo/go-micro/products/data"
	"github.com/rjNemo/go-micro/products/models"
)

// swagger:route POST /products products product
// Creates a product
// responses:
// 	201: productResponse

// AddProduct reads request body and creates new product
func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request) {
	p.logger.Println("Handle 'POST' request")
	// get product from the request
	newProd := r.Context().Value(KeyProduct{}).(*models.Product) // cast into a Product

	p.logger.Printf("product: %#v", newProd)
	data.AddProduct(newProd)
}
