package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/codes"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"

	apitrace "go.opentelemetry.io/otel/trace"
)

// SandService represents both: http controller and business logic, for brevity.
type SandService struct {
	logger *slog.Logger             // logger that prints logs (with trace ID and span ID) to stdout and appends them to open telemetry log batcher
	lp     *sdklog.LoggerProvider   // logger provider that sends logs to open telemetry collector
	tp     *trace.TracerProvider    // tracer provider that sends traces to open telemetry collector
	mp     *sdkmetric.MeterProvider // meter provider that sends metrics to open telemetry collector
}

func NewSandService(ctx context.Context) *SandService {
	lp, logger, err := newLoggerProvider("sand-service")
	if err != nil {
		log.Fatal(err)
	}

	tp, err := newTracerProvider("sand-service")
	if err != nil {
		log.Fatal(err)
	}

	mp, err := newMeterProvider("sand-service")
	if err != nil {
		log.Fatal(err)
	}

	return &SandService{
		logger: logger, // or in actual microservice just set logger globally with: slog.SetDefault(logger)
		tp:     tp,     // or in actual microservice just set tp globally with: otel.SetTracerProvider(tp)
		mp:     mp,     // or in actual microservice just set mp globally with: otel.SetMeterProvider(mp)
		lp:     lp,     // or in actual microservice just set lp globally with: global.SetLoggerProvider(lp)
	}
}

// Start HTTP server
func (s *SandService) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-sand", s.handleGetSand)
	muxWithTracing := otelhttp.NewHandler(mux, "sand-service",
		otelhttp.WithTracerProvider(s.tp),
		otelhttp.WithMeterProvider(s.mp), // automatically records http.server.request.duration histogram
	)
	log.Fatal(http.ListenAndServe(":8082", muxWithTracing))
}

// Shutdown gracefully shuts down the service, flushing logs, traces, and metrics.
func (s *SandService) Shutdown(ctx context.Context) {
	s.lp.Shutdown(ctx)
	s.tp.Shutdown(ctx)
	s.mp.Shutdown(ctx)
}

// HTTP CONTROLLER
func (s *SandService) handleGetSand(w http.ResponseWriter, r *http.Request) {
	// log the request with context, so trace ID and span ID are included in the OTel log output
	// note: trace span is automatically created by otelhttp middleware so nothing to do here
	s.logger.InfoContext(r.Context(), r.Method+" "+r.URL.RequestURI())

	// call business logic
	result, err := s.gatherSand(r.Context())

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
func (s *SandService) gatherSand(ctx context.Context) (result string, err error) {
	// simulate random lag
	if rand.Intn(5) == 0 {
		s.logger.WarnContext(ctx, "unexpected delay of 750ms")
		time.Sleep(time.Millisecond * 750)
	}

	// simulate random failure
	if rand.Intn(5) == 0 {
		err := fmt.Errorf("failed to gather sand - dump track crashed")
		return "", err
	}
	return "1 ton of sand", nil
}
