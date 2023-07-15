package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const name = "open-telemetry-demo"

func main() {
	// initialize tracing

	// Write telemetry data to a file.
	tracesOutputFile, err := os.Create("traces.json")
	if err != nil {
		log.Fatal(err)
	}
	defer tracesOutputFile.Close()

	exporter, err := newExporter(tracesOutputFile)
	if err != nil {
		log.Fatal(err)
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(newResource()),
	)
	defer func() {
		if err := traceProvider.Shutdown(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()
	otel.SetTracerProvider(traceProvider)

	// do actual work
	_, span := otel.Tracer(name).Start(context.Background(), "Run") // ctx returned here may be used to create child-spans for this root span
	span.SetAttributes(attribute.String("request.todays-weather", "semi-cloudy"))
	exampleProcessingError := fmt.Errorf("example error")
	span.RecordError(exampleProcessingError)
	span.SetStatus(codes.Error, exampleProcessingError.Error())
	log.Println("Running!")
	span.End()
}

// newExporter returns a console exporter.
func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newResource returns a resource describing this application.
func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(name),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
