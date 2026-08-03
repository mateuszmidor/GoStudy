package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/sdk/trace"

	apitrace "go.opentelemetry.io/otel/trace"
)

// SandService represents both: http controller and business logic, for brevity.
type SandService struct {
	logger     *slog.Logger                // logger that prints logs (with trace ID and span ID) to stdout and appends them to open telemetry log batcher
	tp         *trace.TracerProvider       // tracer provider that sends traces to open telemetry collector
	logCleanup func(context.Context) error // flush logs from open telemetry log batcher
}

func NewSandService(ctx context.Context) *SandService {
	logger, logCleanup := newLogger("sand-service") // or just set logger globally with: slog.SetDefault(logger)

	tp, err := newTracerProvider("sand-service")
	if err != nil {
		log.Fatal(err)
	}

	return &SandService{
		logger:     logger,
		tp:         tp,
		logCleanup: logCleanup,
	}
}

// Start HTTP server
func (s *SandService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-sand", s.handleGetSand)
	muxWithTracing := otelhttp.NewHandler(mux, "sand-service", otelhttp.WithTracerProvider(s.tp)) // automatically creates trace spans for incoming requests
	log.Fatal(http.ListenAndServe(":8082", muxWithTracing))
}

// Shutdown gracefully shuts down the service, flushing logs and traces.
func (s *SandService) Shutdown(ctx context.Context) {
	s.logCleanup(ctx)
	s.tp.Shutdown(ctx)
}

// HTTP CONTROLLER
func (s *SandService) handleGetSand(w http.ResponseWriter, r *http.Request) {
	// log the request with context, so trace ID and span ID are included in the log output
	// note: trace span is automatically created by otelhttp middleware so nothing to do here
	s.logger.InfoContext(r.Context(), r.Method + " " + r.URL.RequestURI())
	
	// call business logic
	result, err := gatherSand()

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
func gatherSand() (result string, err error) {
	// simulate random failure 50% of the time
	if rand.Intn(2) == 0 {
		err := fmt.Errorf("failed to gather sand - dump track crashed")
		return "", err
	}
	return "1 ton of sand", nil
}
