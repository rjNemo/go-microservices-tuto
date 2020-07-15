package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rjNemo/go-micro/products/data"
	"github.com/rjNemo/go-micro/products/models"
)

// swagger:route PUT /products/{id} products product
// Updates a product
// responses:
// 	204: productResponse

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
