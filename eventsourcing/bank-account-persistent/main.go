package main

import (
	"bank-account-persistent/events"
	"bank-account-persistent/slices/createaccount"
	"bank-account-persistent/slices/fundaccount"
	"bank-account-persistent/slices/listaccounts"
	"log/slog"
	"net/http"

	"github.com/kurrent-io/KurrentDB-Client-Go/kurrentdb"
	"github.com/terraskye/eventsourcing"
	eventstorekurrentdb "github.com/terraskye/eventsourcing/eventstore/kurrentdb"
)

func main() {
	config, err := kurrentdb.ParseConnectionString("esdb://localhost:2113?tls=false")
	if err != nil {
		slog.Error("failed to parse connection string", slog.Any("error", err))
		return
	}

	client, err := kurrentdb.NewClient(config)
	if err != nil {
		slog.Error("failed to create KurrentDB client", slog.Any("error", err))
		return
	}
	defer client.Close()

	store := eventstorekurrentdb.NewEventStore(client)

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
