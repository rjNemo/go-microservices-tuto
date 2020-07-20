package server

import (
	"net/http"
	"time"

	gohandlers "github.com/gorilla/handlers"
)

// New creates a server using given mux and port
func New(mux http.Handler, port string, origins []string) *http.Server {
	// CORS
	corsHandler := gohandlers.CORS(gohandlers.AllowedOrigins(origins))
	return &http.Server{
		Addr:         port,
		Handler:      corsHandler(mux),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second, // keep connection opened to prevent Ddos attacks
	}
}
