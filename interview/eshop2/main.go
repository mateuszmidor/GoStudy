package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Count uint   `json:"count"`
}

// var db = NewMemDB()
var db = NewSqliteDB()

func logging(next http.Handler) http.Handler {
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("### %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)

	}
	return http.HandlerFunc(handler)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/products", handlerCreateProduct)
	mux.HandleFunc("GET /api/products", handleGetAllProducts)
	mux.HandleFunc("GET /api/products/{id}", handleGetProduct)

	server := http.Server{Addr: ":9090", Handler: logging(mux)}
	slog.Info("listening on" + server.Addr)
	err := server.ListenAndServe()
	slog.Error(err.Error())
}

func handlerCreateProduct(w http.ResponseWriter, r *http.Request) {
	p := Product{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		respondError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	p, err = db.Create(p) //updates p.ID
	if err != nil {
		respondError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = respondJson(w, r, p)
	if err != nil {
		respondError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetAllProducts(w http.ResponseWriter, r *http.Request) {
	p, err := db.GetAll()
	if err != nil {
		respondError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	err = respondJson(w, r, p)
	if err != nil {
		respondError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		respondError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	p, err := db.Get(id)
	if err != nil {
		respondError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	err = respondJson(w, r, p)
	if err != nil {
		respondError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
}

func respondError(w http.ResponseWriter, r *http.Request, msg string, code int) {
	logmsg := r.Method + " " + r.URL.Path + " failure - " + msg
	slog.Error(logmsg)
	http.Error(w, msg, code)
}

func respondJson(w http.ResponseWriter, r *http.Request, v any) error {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		return err
	}

	logmsg := r.Method + " " + r.URL.Path + " success"
	slog.Info(logmsg, slog.Any("val", v))
	return nil
}
