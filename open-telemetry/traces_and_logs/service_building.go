package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"
	apitrace "go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

// BuildingService represents both: http controller and business logic, for brevity.
type BuildingService struct {
	logger     *slog.Logger                // logger that prints logs (with trace ID and span ID) to stdout and appends them to open telemetry log batcher
	tp         *trace.TracerProvider       // tracer provider that sends traces to open telemetry collector
	client     *http.Client                // http client that automatically creates trace spans for outgoing requests
	logCleanup func(context.Context) error // flush logs from open telemetry log batcher
}

func NewBuildingService(ctx context.Context) *BuildingService {
	logger, logCleanup := newLogger("building-service") // or just set logger globally with: slog.SetDefault(logger)

	tp, err := newTracerProvider("building-service")
	if err != nil {
		log.Fatal(err)
	}

	// client := &http.Client{
	// 	Transport: otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp)),
	// }
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport,
			otelhttp.WithTracerProvider(noop.NewTracerProvider()),
		),
	}

	return &BuildingService{
		logger:     logger,
		tp:         tp,
		client:     client,
		logCleanup: logCleanup,
	}
}

// Start HTTP server
func (s *BuildingService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/build-house", s.handleBuildHouse)
	muxWithTracing := otelhttp.NewHandler(mux, "building-service", otelhttp.WithTracerProvider(s.tp)) // automatically creates trace spans for incoming requests
	log.Fatal(http.ListenAndServe(":8080", muxWithTracing))
}

// Shutdown gracefully shuts down the service, flushing logs and traces.
func (s *BuildingService) Shutdown(ctx context.Context) {
	s.logCleanup(ctx)
	s.tp.Shutdown(ctx)
}

// HTTP CONTROLLER
func (s *BuildingService) handleBuildHouse(w http.ResponseWriter, r *http.Request) {
	// log the request with context, so trace ID and span ID are included in the log output
	// note: trace span is automatically created by otelhttp middleware so nothing to do here
	s.logger.InfoContext(r.Context(), r.Method + " " + r.URL.RequestURI())

	// call business logic
	result, err := s.buildHouse(r.Context())

	// handle error
	if err != nil {
		s.logger.ErrorContext(r.Context(), err.Error())                           // log the error
		apitrace.SpanFromContext(r.Context()).SetStatus(codes.Error, err.Error()) // set the span status to error
		apitrace.SpanFromContext(r.Context()).RecordError(err)                    // attach the error to the span
		http.Error(w, err.Error(), http.StatusInternalServerError)                // return the error
		return
	}

	// handle success
	w.Write([]byte(result))
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
