package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rjNemo/go-micro/handlers"
	"github.com/rjNemo/go-micro/server"
)

const port = ":5000"

func main() {
	logger := log.New(os.Stdout, "Product API: ", log.LstdFlags|log.Lshortfile)

	// create the handlers
	productsHandler := handlers.NewProducts(logger)
	// create a server mux and register the handlers
	router := mux.NewRouter()
	// GET
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", productsHandler.GetProducts)
	// POST
	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", productsHandler.AddProduct)
	postRouter.Use(productsHandler.ProductValidationMiddleware)
	// PUT
	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", productsHandler.UpdateProduct)
	putRouter.Use(productsHandler.ProductValidationMiddleware)

	// creates a new server
	srv := server.New(router, port)

	// non blocking application server
	go func() {
		logger.Printf("Server started at address http://localhost%s...", port)
		logger.Fatalf("Server failed: %v", srv.ListenAndServe())
	}()

	// catch sigterm or interrupt and gracefully terminates the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Printf("Received %v signal... graceful shutdown", sig)

	toCtx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(toCtx)
}
