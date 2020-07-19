package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

// RegisterRoutes associates path to controller
func (p *Products) RegisterRoutes(m *mux.Router, path string) {
	r := m.StrictSlash(true).PathPrefix(path).Subrouter()
	// GET
	getRouter := r.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", p.GetProducts)
	getRouter.HandleFunc("/{id:[0-9]+}", p.GetOneProduct)
	// POST
	postRouter := r.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", p.AddProduct)
	postRouter.Use(p.ProductValidationMiddleware)
	// PUT
	putRouter := r.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", p.UpdateProduct)
	putRouter.Use(p.ProductValidationMiddleware)
	// DELETE
	deleteRouter := r.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/{id:[0-9]+}", p.DeleteProduct)
	// swagger docs
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	swaggerHandler := middleware.Redoc(opts, nil)
	m.Handle("/docs", swaggerHandler).Methods(http.MethodGet)
	m.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))
}
