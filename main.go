package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/rjNemo/go-micro/handlers"
	"github.com/rjNemo/go-micro/server"
)

const port = ":5000"

func main() {
	logger := log.New(os.Stdout, "Product API: ", log.LstdFlags|log.Lshortfile)

	// create the handlers
	productsHandler := handlers.NewProducts(logger)
	// create a server mux and register the handlers
	mux := http.NewServeMux()
	mux.Handle("/", productsHandler)

	// creates a new server
	srv := server.New(mux, port)

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
