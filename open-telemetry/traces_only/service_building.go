package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"strings"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	apitrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var buildingLogger = slog.With("service", "building-service")

// BuildingService represents both: http controller and business logic, for brevity.
type BuildingService struct {
	tp     *trace.TracerProvider // tracer provider that sends traces to open telemetry collector
	client *http.Client          // http client that automatically creates trace spans for outgoing requests
}

func NewBuildingService(ctx context.Context) *BuildingService {
	tp, err := newTracerProvider("building-service")
	if err != nil {
		log.Fatal(err)
	}

	// option A: create new trace span for outgoing requests
	// client := &http.Client{
	// 	Transport: otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp)),
	// }
	// option B: use noop tracer provider to avoid creating new trace spans for outgoing requests, they will be created by the upstream service anyway
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithTracerProvider(noop.NewTracerProvider()),
		),
	}

	return &BuildingService{
		tp:     tp, // or in actual microservice just set tp globally with: otel.SetTracerProvider(tp)
		client: client,
	}
}

// Start HTTP server
func (s *BuildingService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/build-house", s.handleBuildHouse)
	muxWithTracing := otelhttp.NewHandler(mux, "building-service", otelhttp.WithTracerProvider(s.tp)) // automatically creates trace spans for incoming requests
	log.Fatal(http.ListenAndServe(":8080", muxWithTracing))
}

// Shutdown gracefully shuts down the service, flushing traces.
func (s *BuildingService) Shutdown(ctx context.Context) {
	s.tp.Shutdown(ctx)
}

// HTTP CONTROLLER
func (s *BuildingService) handleBuildHouse(w http.ResponseWriter, r *http.Request) {
	// log the request with context, so trace ID and span ID are included in the OTel log output
	// note: trace span is automatically created by otelhttp middleware so nothing to do here
	buildingLogger.InfoContext(r.Context(), r.Method+" "+r.URL.RequestURI())

	// call business logic
	result, err := s.buildHouse(r.Context())

	// handle error
	if err != nil {
		buildingLogger.ErrorContext(r.Context(), err.Error())                     // log the error
		apitrace.SpanFromContext(r.Context()).SetStatus(codes.Error, err.Error()) // set the span status to error
		apitrace.SpanFromContext(r.Context()).RecordError(err)                    // attach the error to the span
		http.Error(w, strings.ReplaceAll(err.Error(), "\n", ""), http.StatusInternalServerError)
		return
	}

	// handle success
	fmt.Fprintln(w, result)
}

// BUSINESS LOGIC
func (s *BuildingService) buildHouse(ctx context.Context) (string, error) {
	// call downstream service
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8081/provide-concrete", nil)
	if err != nil {
		return "", err
	}
	resp, err := s.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	// handle response
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to build a house: %s", string(body))
	}
	return fmt.Sprintf("built 1 house from %s", body), nil
}
