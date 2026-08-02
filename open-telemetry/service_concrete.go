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
)

// ConcreteService represents both: http controller and business logic, for brevity.
type ConcreteService struct {
	logger     *slog.Logger                // logger that prints logs (with trace ID and span ID) to stdout and appends them to open telemetry log batcher
	tp         *trace.TracerProvider       // tracer provider that sends traces to open telemetry collector
	client     *http.Client                // http client that automatically creates trace spans for outgoing requests
	logCleanup func(context.Context) error // flush logs from open telemetry log batcher
}

func NewConcreteService(ctx context.Context) *ConcreteService {
	logger, logCleanup := newLogger("concrete-service")

	tp, err := newTracerProvider("concrete-service")
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp)),
	}

	return &ConcreteService{
		logger:     logger,
		tp:         tp,
		client:     client,
		logCleanup: logCleanup,
	}
}

// Start HTTP server
func (s *ConcreteService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/provide-concrete", s.handleProvideConcrete)
	muxWithTracing := otelhttp.NewHandler(mux, "concrete-service", otelhttp.WithTracerProvider(s.tp)) // automatically creates trace spans for incoming requests
	log.Fatal(http.ListenAndServe(":8081", muxWithTracing))
}

// Shutdown gracefully shuts down the service, flushing logs and traces.
func (s *ConcreteService) Shutdown(ctx context.Context) {
	s.logCleanup(ctx)
	s.tp.Shutdown(ctx)
}

// HTTP CONTROLLER
func (s *ConcreteService) handleProvideConcrete(w http.ResponseWriter, r *http.Request) {
	// log the request with context, so trace ID and span ID are included in the log output
	// note: trace span is automatically created by otelhttp middleware so nothing to do here
	s.logger.InfoContext(r.Context(), "request received", "url", r.URL.String())

	// call business logic
	result, err := s.produceConcrete(r.Context())

	// handle error
	if err != nil {
		s.logger.ErrorContext(r.Context(), err.Error())                           // log the error
		apitrace.SpanFromContext(r.Context()).SetStatus(codes.Error, err.Error()) // set the span status to error
		apitrace.SpanFromContext(r.Context()).RecordError(err)                    // trace the error
		http.Error(w, err.Error(), http.StatusInternalServerError)                // return the error
		return
	}

	// handle success
	w.Write([]byte(result))
}

// BUSINESS LOGIC
func (s *ConcreteService) produceConcrete(ctx context.Context) (result string, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8082/get-sand", nil)
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
		return "", fmt.Errorf("failed to produce concrete: %s", string(body))
	}
	return fmt.Sprintf("mixed 2m3 of concrete from %s", body), nil
}
