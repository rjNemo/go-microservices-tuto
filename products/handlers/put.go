package handlers

import (
	"fmt"
	"net/http"

	"github.com/rjNemo/go-micro/products/data"
	"github.com/rjNemo/go-micro/products/models"
)

// swagger:route PUT /products/{id} products updateProduct
// Update a products details
//
// responses:
//	204: noContent
//  404: errorResponse
//  422: errorValidation

// UpdateProduct edit product identified by id
func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

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
	// write the no content success header
	w.WriteHeader(http.StatusNoContent)
}
