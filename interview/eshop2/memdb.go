package main

import (
	"fmt"
	"maps"
	"slices"
	"sync"
)

type MemDB struct {
	products map[int]Product
	nextID   int
	mutex    *sync.RWMutex
}

func NewMemDB() *MemDB {
	return &MemDB{products: map[int]Product{}, nextID: 1, mutex: &sync.RWMutex{}}
}

func (db *MemDB) Create(p Product) (Product, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	p.ID = db.nextID
	db.products[db.nextID] = p
	db.nextID++
	return p, nil
}

func (db *MemDB) GetAll() ([]Product, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	products := slices.Collect(maps.Values(db.products))
	return products, nil
}

func (db *MemDB) Get(id int) (Product, error) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	p, ok := db.products[id]
	if !ok {
		return Product{}, fmt.Errorf("product id=%d not found", id)
	}
	return p, nil
}
