package main

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func newTracerProvider(serviceName string) (*trace.TracerProvider, error) {
	exp, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint("localhost:4318"), // jaeger collector endpoint
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(r),
	), nil
}
