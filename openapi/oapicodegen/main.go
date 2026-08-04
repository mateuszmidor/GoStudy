package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"
	"sort"

	httpSwagger "github.com/swaggo/http-swagger"

	generatedserver "github.com/mateuszmidor/GoStudy/openapi/oapicodegen/generated_server"
)

//go:embed fridge_api.yaml
var swaggerSpec []byte

type FridgeServer struct {
	products map[string]float32
}

func NewFridgeServer() *FridgeServer {
	return &FridgeServer{
		products: make(map[string]float32),
	}
}

func (s *FridgeServer) GetProducts(w http.ResponseWriter, r *http.Request, params generatedserver.GetProductsParams) {
	products := make([]generatedserver.Product, 0, len(s.products))
	for name, quantity := range s.products {
		pname := generatedserver.ProductName(name)
		products = append(products, generatedserver.Product{
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
	var product generatedserver.Product
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

func (s *FridgeServer) GetProductsName(w http.ResponseWriter, r *http.Request, name generatedserver.ProductName) {
	quantity, ok := s.products[string(name)]
	if !ok {
		http.Error(w, "product not found", http.StatusNotFound)
		return
	}

	product := generatedserver.Product{
		Name:     &name,
		Quantity: &quantity,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (s *FridgeServer) PutProductsName(w http.ResponseWriter, r *http.Request, name generatedserver.ProductName) {
	var req generatedserver.PutProductsNameJSONBody
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

func sortProducts(products []generatedserver.Product) {
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

	// 1. Register the generated fridge server handlers into the mux
	generatedserver.HandlerFromMux(srv, mux)

	// 2. Serve the raw OpenAPI spec at /swagger/swagger.yaml
	mux.HandleFunc("/swagger/swagger.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(swaggerSpec)
	})

	// 3. Serve Swagger UI at /swagger/
	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/swagger.yaml"),
	))

	// 4. Start the server
	log.Println("Server starting on :8080")
	log.Println("Swagger UI available at http://localhost:8080/swagger/")
	log.Fatal(http.ListenAndServe("localhost:8080", mux))
}
