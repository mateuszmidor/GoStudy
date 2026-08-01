package main

import (
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func startConcreteService() *trace.TracerProvider {
	logger := newLogger("concrete-service")

	otlpExp, err := newOLTPExporter()
	if err != nil {
		log.Fatal(err)
	}
	tp := newTracerProvider("concrete-service", otlpExp)

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithTracerProvider(tp)),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/provide-concrete", handleProvideConcrete(logger, client))

	go func() {
		log.Fatal(http.ListenAndServe(":8081",
			otelhttp.NewHandler(mux, "provide-concrete", otelhttp.WithTracerProvider(tp))))
	}()

	return tp
}

func handleProvideConcrete(logger *slog.Logger, client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(r.Context(), "request received", "url", r.URL.String())
		time.Sleep(250 * time.Millisecond)
		req, err := http.NewRequestWithContext(r.Context(), "GET", "http://localhost:8082/get-sand", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		fmt.Fprintf(w, "concrete + %s", body)
	}
}
