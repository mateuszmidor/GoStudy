package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
)

func newLogger(scopeName string) (*slog.Logger, func(context.Context) error) {
	exp, err := newOTLPLogExporter()
	if err != nil {
		log.Fatal(err)
	}
	lp := sdklog.NewLoggerProvider(sdklog.WithProcessor(sdklog.NewBatchProcessor(exp)))
	handler := newMultiHandler(
		slog.NewJSONHandler(os.Stdout, nil).WithAttrs([]slog.Attr{slog.String("service", scopeName)}),
		otelslog.NewHandler(scopeName, otelslog.WithLoggerProvider(lp)),
	)
	return slog.New(handler), lp.Shutdown
}

func newOTLPLogExporter() (sdklog.Exporter, error) {
	return otlploghttp.New(
		context.Background(), 
		otlploghttp.WithEndpoint("localhost:666"), // some log collector endpoint
		otlploghttp.WithInsecure(),
	)
}