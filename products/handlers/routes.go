package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes associates path to controller
func (p *Products) RegisterRoutes(r *mux.Router) {
	// GET
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", p.GetProducts)
	// POST
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", p.AddProduct)
	postRouter.Use(p.ProductValidationMiddleware)
	// PUT
	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", p.UpdateProduct)
	putRouter.Use(p.ProductValidationMiddleware)
}
