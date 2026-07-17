package main

import (
	"bank-account/slices/createaccount"
	"bank-account/slices/fundaccount"
	"bank-account/slices/listaccounts"
	"log/slog"
	"net/http"

	memstore "github.com/terraskye/eventsourcing/eventstore/memory"
)

func main() {
	store := memstore.NewMemoryStore(100)
	defer store.Close()

	mux := http.NewServeMux()
	createAccountHandler := createaccount.NewHTTPHandler(createaccount.NewHandler(store))
	createAccountHandler.Register(mux)
	listAccountsHandler := listaccounts.NewHTTPHandler(listaccounts.NewQueryHandler(store))
	listAccountsHandler.Register(mux)
	fundAccountHandler := fundaccount.NewHTTPHandler(fundaccount.NewHandler(store))
	fundAccountHandler.Register(mux)

	server := http.Server{Addr: ":8080", Handler: mux}
	slog.Info("listening on " + server.Addr)
	slog.Error(server.ListenAndServe().Error())
}
