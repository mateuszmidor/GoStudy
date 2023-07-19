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
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

const name = "open-telemetry-demo"

func main() {
	// 1. Write telemetry data to a file
	tracesOutputFile, err := os.Create("traces.json")
	if err != nil {
		log.Fatal(err)
	}
	defer tracesOutputFile.Close()
	fileExporter, err := newStdOutExporter(tracesOutputFile)
	if err != nil {
		log.Fatal(err)
	}

	// 2. Also, send telemetry data to Jaeger
	oltpExporter, err := newOLTPExporter()
	if err != nil {
		log.Fatal(err)
	}

	// start tracing
	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(fileExporter),
		trace.WithBatcher(oltpExporter),
		trace.WithResource(newResource()),
	)
	otel.SetTracerProvider(traceProvider)

	// do actual work
	_, span := otel.Tracer(name).Start(context.Background(), "Run") // ctx returned here may be used to create child-spans for this root span
	span.SetAttributes(attribute.String("request.todays-weather", "semi-cloudy"))
	exampleProcessingError := fmt.Errorf("example error")
	span.RecordError(exampleProcessingError)
	span.SetStatus(codes.Error, exampleProcessingError.Error())
	log.Println("Running!")
	span.End()

	// finish tracing
	if err := traceProvider.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}

// newExporter returns a console exporter.
func newStdOutExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		stdouttrace.WithoutTimestamps(),
	)
}

// newExporter returns an OTLP exporter, it can directly push to Jaeger, port 4318
func newOLTPExporter() (trace.SpanExporter, error) {
	opts := []otlptracehttp.Option{
		otlptracehttp.WithInsecure(), // send to HTTP Jaeger server
	}
	return otlptracehttp.New(context.Background(), opts...)
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
