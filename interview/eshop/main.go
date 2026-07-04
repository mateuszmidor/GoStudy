package main

import (
	"encoding/json"
	"log/slog"
	"maps"
	"net/http"
	"slices"
	"strconv"
	"sync"
)

type IdType = int
type Item struct {
	Id           IdType `json:"id,omitempty"`
	Name         string `json:"name"`
	Count        uint   `json:"count"`
	PriceInCents uint   `json:"price_cents"`
}

// storage abstraction/interface to be extracted later
type Storage struct {
	Items  map[IdType]Item
	freeID int
	mutex  *sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{Items: map[IdType]Item{}, mutex: &sync.Mutex{}, freeID: 1}
}
func (s *Storage) Add(item Item) Item {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	item.Id = s.freeID
	s.Items[s.freeID] = item
	s.freeID++
	return item
}

func (s *Storage) Delete(id IdType) {
	delete(s.Items, id)
}

func (s *Storage) GetAll() []Item {
	return slices.Collect(maps.Values(s.Items))
}
func (s *Storage) Get(id IdType) *Item {
	if item, exists := s.Items[id]; !exists {
		return nil
	} else {
		return &item
	}
}

// var storage *Storage = NewStorage()

var storage *SqliteStorage = NewSqliteStorage()

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/items", handlerCreateItems)
	mux.HandleFunc("GET /api/items", handleGetAllItems)
	mux.HandleFunc("GET /api/items/{id}", handleGetItem)
	mux.HandleFunc("DELETE /api/items/{id}", handleDeleteItem)

	server := http.Server{Addr: ":9090", Handler: mux}
	slog.Info("serving on :9090")
	slog.Error(server.ListenAndServe().Error())
}

func handlerCreateItems(w http.ResponseWriter, r *http.Request) {
	var items []Item
	if err := json.NewDecoder(r.Body).Decode(&items); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i := range items {
		items[i] = storage.Add(items[i])
	}
	slog.Debug("create item success", slog.Any("items", items))
}

func handleGetAllItems(w http.ResponseWriter, r *http.Request) {
	items := storage.GetAll()
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Debug("handleGetAllItems success", slog.Any("items", items))
}

func handleGetItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	item := storage.Get(id)
	if item == nil {
		http.Error(w, "item not found:"+idStr, http.StatusNotFound)
		slog.Debug("handleGetItem not found", slog.Any("id", idStr))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	slog.Debug("handleGetItem success", slog.Any("item", item))
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, _ := strconv.Atoi(idStr)
	item := storage.Get(id)
	if item == nil {
		http.Error(w, "item not found:"+idStr, http.StatusNotFound)
		slog.Debug("handleDeleteItem not found", slog.Any("id", idStr))
		return
	}
	storage.Delete(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	slog.Debug("handleDeleteItem success", slog.Any("id", id))
}
