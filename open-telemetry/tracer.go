package main

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func newTracerProvider(name string, exp trace.SpanExporter) *trace.TracerProvider {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(name),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)
}

func newOLTPExporter() (trace.SpanExporter, error) {
	return otlptracehttp.New(context.Background(), otlptracehttp.WithInsecure())
}
