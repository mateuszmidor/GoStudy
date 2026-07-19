package main

import (
	"bank-account-persistent/events"
	"bank-account-persistent/slices/createaccount"
	"bank-account-persistent/slices/fundaccount"
	"bank-account-persistent/slices/getbalance"
	"bank-account-persistent/slices/listaccounts"
	"context"
	_ "embed"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/terraskye/eventsourcing"
	pgbus "github.com/terraskye/eventsourcing/eventbus/postgres"
	pgstore "github.com/terraskye/eventsourcing/eventstore/postgres"
)

//go:embed schema.sql
var schemaSQL string

func main() {
	ctx := context.Background()

	// initialize postgres connection
	pool, err := pgxpool.New(ctx, "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		slog.Error("failed to create connection pool", slog.Any("error", err))
		return
	}
	defer pool.Close()

	// initialize event store
	if _, err := pool.Exec(ctx, schemaSQL); err != nil {
		slog.Error("failed to apply schema", slog.Any("error", err))
		return
	}
	store := pgstore.NewEventStore(pool)
	// events must be registered for the event store to work
	eventsourcing.RegisterEvent(&events.AccountCreated{})
	eventsourcing.RegisterEvent(&events.AccountFunded{})

	// initialize event bus
	bus := pgbus.NewEventBus(pool, time.Second)
	projector := listaccounts.NewProjector()
	if err := bus.Subscribe(ctx, "list-accounts-projector", projector.EventHandlers()); err != nil {
		slog.Error("failed to add subscriber to bus", slog.Any("error", err))
		return
	}

	// initialize command handlers
	createAccountHandler := createaccount.NewHTTPHandler(createaccount.NewHandler(store))
	listAccountsHandler := listaccounts.NewHTTPHandler(listaccounts.NewQueryHandler(projector))
	fundAccountHandler := fundaccount.NewHTTPHandler(fundaccount.NewHandler(store))
	getBalanceHandler := getbalance.NewHTTPHandler(getbalance.NewQueryHandler(store))

	// initialize&run http server
	mux := http.NewServeMux()
	createAccountHandler.Register(mux)
	listAccountsHandler.Register(mux)
	fundAccountHandler.Register(mux)
	getBalanceHandler.Register(mux)
	server := http.Server{Addr: ":8080", Handler: mux}
	slog.Info("listening on " + server.Addr)
	slog.Error(server.ListenAndServe().Error())
}
