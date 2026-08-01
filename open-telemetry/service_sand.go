package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/sdk/trace"
	apitrace "go.opentelemetry.io/otel/trace"
)

func startSandService() *trace.TracerProvider {
	otlpExp, err := newOLTPExporter()
	if err != nil {
		log.Fatal(err)
	}
	tp := newTracerProvider("sand-service", otlpExp)

	mux := http.NewServeMux()
	mux.HandleFunc("/get-sand", handleGetSand)

	go func() {
		log.Fatal(http.ListenAndServe(":8082",
			otelhttp.NewHandler(mux, "get-sand", otelhttp.WithTracerProvider(tp))))
	}()

	return tp
}

func handleGetSand(w http.ResponseWriter, r *http.Request) {
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
