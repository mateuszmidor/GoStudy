package main

import (
	"errors"
	"sync"
)

type Product struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price uint   `json:"price"`
}

var ErrNotFound = errors.New("product not found")
var ErrInvalidInput = errors.New("invalid input")

type Store struct {
	mu       sync.Mutex
	products map[uint]Product
	nextID   uint
}

func NewStore() *Store {
	return &Store{
		products: make(map[uint]Product),
		nextID:   1,
	}
}

func (s *Store) AddProduct(name string, price uint) (*Product, error) {
	if name == "" {
		return nil, ErrInvalidInput
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	product := Product{
		ID:    s.nextID,
		Name:  name,
		Price: price,
	}
	s.products[product.ID] = product
	s.nextID++

	return &product, nil
}

func (s *Store) GetProduct(id uint) (*Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	product, ok := s.products[id]
	if !ok {
		return nil, ErrNotFound
	}

	return &product, nil
}

func (s *Store) GetAllProducts() []Product {
	s.mu.Lock()
	defer s.mu.Unlock()

	products := make([]Product, 0, len(s.products))
	for _, product := range s.products {
		products = append(products, product)
	}

	return products
}

func (s *Store) UpdateProduct(id uint, name string, price uint) (*Product, error) {
	if name == "" {
		return nil, ErrInvalidInput
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	product, ok := s.products[id]
	if !ok {
		return nil, ErrNotFound
	}

	product.Name = name
	product.Price = price
	s.products[id] = product

	return &product, nil
}
