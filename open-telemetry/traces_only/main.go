package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	// Log to stdout in JSON format
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	// Configure global propagator to extract and inject W3C trace context: trace id, parent span id, trace flags
	// on incoming and outgoing HTTP requests; otelhttp middleware will automatically use this:
	// otelhttp.NewTransport for http client, otelhttp.NewHandler for http server
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// Start services:
	// Sand service
	sandSvc := NewSandService(context.Background())
	go sandSvc.Start()
	defer sandSvc.Shutdown(context.Background())

	// Concrete service
	concreteSvc := NewConcreteService(context.Background())
	go concreteSvc.Start()
	defer concreteSvc.Shutdown(context.Background())

	// Building service
	buildingSvc := NewBuildingService(context.Background())
	go buildingSvc.Start()
	defer buildingSvc.Shutdown(context.Background())

	// Wait for CTRL+C
	log.Println("Services up: :8080/build-house, :8081/provide-concrete, :8082/get-sand")
	select {}
}
