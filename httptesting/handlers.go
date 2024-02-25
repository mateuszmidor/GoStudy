package handlers

import "net/http"

// NewMyServer creates url multiplexer
func NewMyServer() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/about", http.HandlerFunc(HandleAbout))
	mux.Handle("/admin", http.HandlerFunc(HandleAdmin))
	http.Handle("/", http.HandlerFunc(http.NotFound))
	return mux
}

// HandleAbout handles /about request
func HandleAbout(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// HandleAdmin handles /admin request
func HandleAdmin(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
}
