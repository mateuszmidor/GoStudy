package main

import (
	"log/slog"
	"net/http"

	"bank-account/infrastructure"
	"bank-account/presentation"

	memstore "github.com/terraskye/eventsourcing/eventstore/memory"
)

func main() {
	store := memstore.NewMemoryStore(100)
	defer store.Close()

	repo := infrastructure.NewRepository(store)
	controller := presentation.NewHttpController(repo)

	mux := http.NewServeMux()
	controller.RegisterRoutes(mux)

	server := http.Server{Addr: ":8080", Handler: mux}
	slog.Info("listening on " + server.Addr)
	slog.Error(server.ListenAndServe().Error())
}
