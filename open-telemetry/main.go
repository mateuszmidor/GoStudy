package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func main() {
	// Configure global propagator to extract and inject W3C trace context: trace id, parent span id, trace flags
	// on incoming and outgoing HTTP requests; otelhttp middleware will automatically use this:
	// otelhttp.NewTransport for http client, otelhttp.NewHandler for http server
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
	))

	sandSvc := NewSandService(context.Background())
	go sandSvc.Start()
	defer sandSvc.Shutdown(context.Background())
	concreteSvc := NewConcreteService(context.Background())
	go concreteSvc.Start()
	defer concreteSvc.Shutdown(context.Background())
	constructionSvc := NewConstructionService(context.Background())
	go constructionSvc.Start()
	defer constructionSvc.Shutdown(context.Background())

	log.Println("Services up: :8080/build-house, :8081/provide-concrete, :8082/get-sand")
	select {}
}
