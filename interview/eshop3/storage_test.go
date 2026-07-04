package main

import (
	"sync"
	"testing"
)

func TestStore_AddProduct(t *testing.T) {
	store := NewStore()

	product, err := store.AddProduct("Widget", 9999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if product.ID != 1 {
		t.Errorf("expected id 1, got %d", product.ID)
	}
	if product.Name != "Widget" {
		t.Errorf("expected name Widget, got %s", product.Name)
	}
	if product.Price != 9999 {
		t.Errorf("expected price 9999, got %d", product.Price)
	}
}

func TestStore_AddProduct_EmptyName(t *testing.T) {
	store := NewStore()

	_, err := store.AddProduct("", 9999)
	if err != ErrInvalidInput {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestStore_GetProduct(t *testing.T) {
	store := NewStore()

	store.AddProduct("Widget", 9999)
	product, err := store.GetProduct(1)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if product.Name != "Widget" {
		t.Errorf("expected name Widget, got %s", product.Name)
	}
}

func TestStore_GetProduct_NotFound(t *testing.T) {
	store := NewStore()

	_, err := store.GetProduct(1)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestStore_GetAllProducts(t *testing.T) {
	store := NewStore()

	store.AddProduct("Widget", 9999)
	store.AddProduct("Gadget", 14999)

	products := store.GetAllProducts()
	if len(products) != 2 {
		t.Errorf("expected 2 products, got %d", len(products))
	}
}

func TestStore_GetAllProducts_Empty(t *testing.T) {
	store := NewStore()

	products := store.GetAllProducts()
	if len(products) != 0 {
		t.Errorf("expected 0 products, got %d", len(products))
	}
}

func TestStore_UpdateProduct(t *testing.T) {
	store := NewStore()

	store.AddProduct("Widget", 9999)
	product, err := store.UpdateProduct(1, "Gadget", 14999)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if product.Name != "Gadget" {
		t.Errorf("expected name Gadget, got %s", product.Name)
	}
	if product.Price != 14999 {
		t.Errorf("expected price 14999, got %d", product.Price)
	}
}

func TestStore_UpdateProduct_NotFound(t *testing.T) {
	store := NewStore()

	_, err := store.UpdateProduct(1, "Gadget", 14999)
	if err != ErrNotFound {
		t.Errorf("expected ErrNotFound, got %v", err)
	}
}

func TestStore_UpdateProduct_EmptyName(t *testing.T) {
	store := NewStore()

	store.AddProduct("Widget", 9999)
	_, err := store.UpdateProduct(1, "", 14999)
	if err != ErrInvalidInput {
		t.Errorf("expected ErrInvalidInput, got %v", err)
	}
}

func TestStore_ThreadSafe(t *testing.T) {
	store := NewStore()
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				store.AddProduct("Product", 100)
			}
		}()
	}

	wg.Wait()

	products := store.GetAllProducts()
	if len(products) != 1000 {
		t.Errorf("expected 1000 products, got %d", len(products))
	}
}
