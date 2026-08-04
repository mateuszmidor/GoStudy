package main

import (
	"context"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
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
		otlptracehttp.WithEndpoint("localhost:4318"), // OTel collector endpoint (same as logs)
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}

	r := makeResource(serviceName)
	tp := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		trace.WithBatcher(exp),
		trace.WithResource(r),
	)
	// otel.SetTracerProvider(tp) // call this in microservice to set global trace provider
	return tp, nil
}

func newLoggerProvider(serviceName string) (*sdklog.LoggerProvider, *slog.Logger, error) {
	exp, err := otlploghttp.New(
		context.Background(),
		otlploghttp.WithEndpoint("localhost:4318"), // OTel Collector endpoint (same as traces)
		otlploghttp.WithInsecure(),
	)
	if err != nil {
		return nil, nil, err
	}
	r := makeResource(serviceName)
	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)),
		sdklog.WithResource(r),
	)

	otelLogHandler := otelslog.NewHandler(serviceName, otelslog.WithLoggerProvider(lp))
	stdoutLogHandler := slog.NewJSONHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("service", serviceName)})
	handler := slogmulti.Fanout(
		otelLogHandler,   // send to OTel collector
		stdoutLogHandler, // but also log to stdout
	)
	logger := slog.New(handler)
	// global.SetLoggerProvider(lp) // call this in microservice to set global log provider
	// slog.SetDefault(logger) // call this in microservice to set global logger
	return lp, logger, nil
}

func makeResource(serviceName string) *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}
