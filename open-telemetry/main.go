package main

import (
	"context"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	sandTp := startSandService()
	concreteTp := startConcreteService()
	constructionTp := startConstructionService()

	defer shutdown(sandTp, concreteTp, constructionTp)
	
	log.Println("Services up: :8080/build-house, :8081/provide-concrete, :8082/get-sand")
	select {}
}


func shutdown(providers ...*trace.TracerProvider) {
	for _, tp := range providers {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("shutdown error: %v", err)
		}
	}
}
