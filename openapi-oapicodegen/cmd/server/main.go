package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sort"

	server "github.com/mateuszmidor/GoStudy/openapi-oapicodegen/generated_server"
)

type FridgeServer struct {
	products map[string]float32
}

func NewFridgeServer() *FridgeServer {
	return &FridgeServer{
		products: make(map[string]float32),
	}
}

func (s *FridgeServer) GetProducts(w http.ResponseWriter, r *http.Request, params server.GetProductsParams) {
	products := make([]server.Product, 0, len(s.products))
	for name, quantity := range s.products {
		pname := server.ProductName(name)
		products = append(products, server.Product{
			Name:     &pname,
			Quantity: &quantity,
		})
	}

	if params.Sort != nil && *params.Sort {
		sortProducts(products)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (s *FridgeServer) PostProducts(w http.ResponseWriter, r *http.Request) {
	var product server.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if product.Name == nil || *product.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	quantity := float32(1.0)
	if product.Quantity != nil {
		quantity = *product.Quantity
	}
	s.products[*product.Name] += quantity

	w.WriteHeader(http.StatusCreated)
}

func (s *FridgeServer) GetProductsName(w http.ResponseWriter, r *http.Request, name server.ProductName) {
	quantity, ok := s.products[string(name)]
	if !ok {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	product := server.Product{
		Name:     &name,
		Quantity: &quantity,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (s *FridgeServer) PutProductsName(w http.ResponseWriter, r *http.Request, name server.ProductName) {
	var req server.PutProductsNameJSONBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	current, ok := s.products[string(name)]
	if !ok {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	withdraw := float32(1.0)
	if req.Quantity != nil {
		withdraw = *req.Quantity
	}

	newQuantity := current - withdraw
	if newQuantity < 0 {
		newQuantity = 0
	}
	s.products[string(name)] = newQuantity

	w.WriteHeader(http.StatusNoContent)
}

func sortProducts(products []server.Product) {
	sort.Slice(products, func(i, j int) bool {
		if products[i].Name == nil || products[j].Name == nil {
			return false
		}
		return *products[i].Name < *products[j].Name
	})
}

func main() {
	srv := NewFridgeServer()

	mux := http.NewServeMux()
	server.HandlerFromMux(srv, mux)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
