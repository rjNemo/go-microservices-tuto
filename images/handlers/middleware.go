package handlers

import (
	"compress/gzip"
	"net/http"
)

// GZipResponseMiddleware detects if the client can handle zipped content and if
// so returns the response in GZipped format
func GZipResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

type WrappedResponseWriter struct {
	w http.ResponseWriter
	g *gzip.Writer
}

func NewWrappedResponseWriter(w http.ResponseWriter) *WrappedResponseWriter {
	g := gzip.NewWriter(w)
	return &WrappedResponseWriter{w, g}
}
