package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rjNemo/go-micro/products/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//	204: noContent
//  404: errorResponse
//  501: errorResponse

// DeleteProduct delete product from datastore identified by id
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	p.logger.Println("Handle 'DELETE' request", id)

	err := data.DeleteProduct(id)
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
