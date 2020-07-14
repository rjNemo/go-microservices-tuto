package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/rjNemo/go-micro/products/handlers"
	"github.com/rjNemo/go-micro/server"
)

const port = ":5000"

func main() {
	// create a logger to control application wide logging
	logger := log.New(os.Stdout, "Product API: ", log.LstdFlags|log.Lshortfile)

	// create a router
	router := mux.NewRouter()

	// create the handler
	productsHandler := handlers.New(logger)

	// register the handler method to the router
	productsHandler.RegisterRoutes(router)

	// creates a production-ready server using the handler
	srv := server.New(router, port)

	// start a non blocking application server
	go func() {
		logger.Printf("Server started at address http://localhost%s ...", port)
		logger.Fatalf("Server failed: %v", srv.ListenAndServe()) // TODO: use ListenAndServeTLS in production
	}()

	// catch sigterm or interrupt and gracefully terminates the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt) // interrupt
	signal.Notify(sigChan, os.Kill)      // sigterm
	// log received signal
	sig := <-sigChan
	logger.Printf("Received %v signal... graceful shutdown", sig)

	toCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	srv.Shutdown(toCtx)
}
