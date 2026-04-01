package main

import (
	"context"
	"log"
	"net/http"
	"sort"

	server "github.com/mateuszmidor/GoStudy/openapi-openapigenerator/generated_server/go"
)

type FridgeServer struct {
	products map[string]float32
}

func NewFridgeServer() *FridgeServer {
	return &FridgeServer{
		products: make(map[string]float32),
	}
}

func (s *FridgeServer) ProductsGet(ctx context.Context, sortParam bool) (server.ImplResponse, error) {
	products := make([]server.Product, 0, len(s.products))
	for name, quantity := range s.products {
		products = append(products, server.Product{
			Name:     name,
			Quantity: quantity,
		})
	}

	if sortParam {
		sortProducts(products)
	}

	return server.ImplResponse{Code: http.StatusOK, Body: products}, nil
}

func (s *FridgeServer) ProductsPost(ctx context.Context, product server.Product) (server.ImplResponse, error) {
	if product.Name == "" {
		return server.ImplResponse{Code: http.StatusBadRequest, Body: "name is required"}, nil
	}

	quantity := product.Quantity
	if quantity == 0 {
		quantity = 1
	}
	s.products[product.Name] += quantity

	return server.ImplResponse{Code: http.StatusCreated, Body: nil}, nil
}

func (s *FridgeServer) ProductsNameGet(ctx context.Context, name string) (server.ImplResponse, error) {
	quantity, ok := s.products[name]
	if !ok {
		return server.ImplResponse{Code: http.StatusNotFound, Body: "product not found"}, nil
	}

	product := server.Product{
		Name:     name,
		Quantity: quantity,
	}

	return server.ImplResponse{Code: http.StatusOK, Body: product}, nil
}

func (s *FridgeServer) ProductsNamePut(ctx context.Context, name string, req server.ProductsNamePutRequest) (server.ImplResponse, error) {
	current, ok := s.products[name]
	if !ok {
		return server.ImplResponse{Code: http.StatusNotFound, Body: "product not found"}, nil
	}

	withdraw := req.Quantity
	if withdraw == 0 {
		withdraw = 1
	}

	newQuantity := current - withdraw
	if newQuantity < 0 {
		newQuantity = 0
	}
	s.products[name] = newQuantity

	return server.ImplResponse{Code: http.StatusNoContent, Body: nil}, nil
}

func sortProducts(products []server.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Name < products[j].Name
	})
}

func main() {
	fridge := NewFridgeServer()

	service := &FridgeService{fridge}
	controller := server.NewDefaultAPIController(service)

	router := server.NewRouter(controller)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

type FridgeService struct {
	fridge *FridgeServer
}

func (s *FridgeService) ProductsGet(ctx context.Context, sortParam bool) (server.ImplResponse, error) {
	return s.fridge.ProductsGet(ctx, sortParam)
}

func (s *FridgeService) ProductsPost(ctx context.Context, product server.Product) (server.ImplResponse, error) {
	return s.fridge.ProductsPost(ctx, product)
}

func (s *FridgeService) ProductsNameGet(ctx context.Context, name string) (server.ImplResponse, error) {
	return s.fridge.ProductsNameGet(ctx, name)
}

func (s *FridgeService) ProductsNamePut(ctx context.Context, name string, req server.ProductsNamePutRequest) (server.ImplResponse, error) {
	return s.fridge.ProductsNamePut(ctx, name, req)
}
