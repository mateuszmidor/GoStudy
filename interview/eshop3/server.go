package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type AddProductRequest struct {
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	store := NewStore()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/products", addProductHandler(store))
	mux.HandleFunc("GET /api/products", listProductsHandler(store))
	mux.HandleFunc("GET /api/products/{id}", getProductHandler(store))
	mux.HandleFunc("PUT /api/products/{id}", updateProductHandler(store))

	log.Println("Server starting on port 9090")
	log.Fatal(http.ListenAndServe(":9090", mux))
}

func addProductHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req AddProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		product, err := store.AddProduct(req.Name, req.Price)
		if err != nil {
			if err == ErrInvalidInput {
				writeError(w, http.StatusBadRequest, "name is required")
				return
			}
			log.Printf("error adding product: %v", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

func listProductsHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products := store.GetAllProducts()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

func getProductHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid product id")
			return
		}

		product, err := store.GetProduct(uint(id))
		if err != nil {
			if err == ErrNotFound {
				writeError(w, http.StatusNotFound, "product not found")
				return
			}
			log.Printf("error getting product: %v", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func updateProductHandler(store *Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			writeError(w, http.StatusBadRequest, "invalid product id")
			return
		}

		var req AddProductRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeError(w, http.StatusBadRequest, "invalid request body")
			return
		}

		product, err := store.UpdateProduct(uint(id), req.Name, req.Price)
		if err != nil {
			if err == ErrNotFound {
				writeError(w, http.StatusNotFound, "product not found")
				return
			}
			if err == ErrInvalidInput {
				writeError(w, http.StatusBadRequest, "name is required")
				return
			}
			log.Printf("error updating product: %v", err)
			writeError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
