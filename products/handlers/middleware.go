package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rjNemo/go-micro/products/models"
)

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
