package main

import (
	"bank-account-persistent/events"
	"bank-account-persistent/slices/createaccount"
	"bank-account-persistent/slices/fundaccount"
	"bank-account-persistent/slices/listaccounts"
	"log/slog"
	"net/http"

	"github.com/terraskye/eventsourcing"
	"github.com/terraskye/eventsourcing/eventstore/kurrentdb"
)

func main() {
	store := kurrentdb.NewEventStore("esdb://localhost:2113?tls=false")
	if store == nil {
		slog.Error("failed to initialize KurrentDB for event storage")
		return
	}
	defer store.Close()

	// for kurrentdb, events must be registered before they can be used
	eventsourcing.RegisterEvent(&events.AccountCreated{})
	eventsourcing.RegisterEvent(&events.AccountFunded{})

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
