package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/sdk/trace"
)

func startConstructionService() *trace.TracerProvider {
	otlpExp, err := newOLTPExporter()
	if err != nil {
		log.Fatal(err)
	}
	tp := newTracerProvider("construction-service", otlpExp)

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithTracerProvider(tp)),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/build-house", handleBuildHouse(client))

	go func() {
		log.Fatal(http.ListenAndServe(":8080",
			otelhttp.NewHandler(mux, "build-house", otelhttp.WithTracerProvider(tp))))
	}()

	return tp
}

func handleBuildHouse(client *http.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)
		req, err := http.NewRequestWithContext(r.Context(), "GET", "http://localhost:8081/provide-concrete", nil)
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
		fmt.Fprintf(w, "house built with %s", body)
	}
}
