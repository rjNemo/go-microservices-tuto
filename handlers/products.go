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
	productList := data.AllProducts()
	err := productList.ToJSON(w)
	if err != nil {
		errMsg := fmt.Sprintf("Unable to encode request: %s\n", err)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}
}
