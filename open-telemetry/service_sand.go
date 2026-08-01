package main

import (
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/sdk/trace"
	apitrace "go.opentelemetry.io/otel/trace"
)

func startSandService() *trace.TracerProvider {
	logger := newLogger("sand-service")

	otlpExp, err := newOLTPExporter()
	if err != nil {
		log.Fatal(err)
	}
	tp := newTracerProvider("sand-service", otlpExp)

	mux := http.NewServeMux()
	mux.HandleFunc("/get-sand", handleGetSand(logger))

	go func() {
		log.Fatal(http.ListenAndServe(":8082",
			otelhttp.NewHandler(mux, "get-sand", otelhttp.WithTracerProvider(tp))))
	}()

	return tp
}

func handleGetSand(logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.InfoContext(r.Context(), "request received", "url", r.URL.String())
		time.Sleep(750 * time.Millisecond)
		if rand.Intn(2) == 0 {
			apitrace.SpanFromContext(r.Context()).RecordError(
				fmt.Errorf("dump track crashed"),
			)
			http.Error(w, "dump track crashed", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("sand delivered"))
	}
}
