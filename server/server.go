package server

import (
	"net/http"
	"time"
)

// New creates a server using given mux and port
func New(mux *http.ServeMux, port string) *http.Server {
	return &http.Server{
		Addr:         port,
		Handler:      mux,
		IdleTimeout:  120 * time.Second, // keep connection opened to prevent Ddos attacks
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
}
