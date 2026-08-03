package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
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

func newLogger(serviceName string) (*slog.Logger, func(context.Context) error) {
	exp, err := otlploghttp.New(
		context.Background(),
		otlploghttp.WithEndpoint("localhost:666"), // some log collector endpoint
		otlploghttp.WithInsecure(),
	)
	if err != nil {
		log.Fatal(err)
	}
	lp := sdklog.NewLoggerProvider(sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)))
	handler := newMultiHandler(
		slog.NewJSONHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("service", serviceName)}),
		otelslog.NewHandler(serviceName, otelslog.WithLoggerProvider(lp)),
	)
	return slog.New(handler), lp.Shutdown
}
