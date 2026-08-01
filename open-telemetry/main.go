package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

func main() {
	// 1. setup shared exporters
	tracesOutputFile, err := os.Create("traces.json")
	if err != nil {
		log.Fatal(err)
	}
	defer tracesOutputFile.Close()
	fileExporter, err := newStdOutExporter(tracesOutputFile)
	if err != nil {
		log.Fatal(err)
	}
	otlpExporter, err := newOLTPExporter()
	if err != nil {
		log.Fatal(err)
	}

	// 2. create per-service TracerProviders (shared exporters, distinct Resources)
	constructionTp := newTracerProvider("construction-service", fileExporter, otlpExporter)
	concreteTp := newTracerProvider("concrete-service", fileExporter, otlpExporter)
	sandTp := newTracerProvider("sand-service", fileExporter, otlpExporter)

	// 3. set global defaults (fallback)
	otel.SetTracerProvider(constructionTp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// 4. shutdown all providers on exit
	defer func() {
		for _, tp := range []*trace.TracerProvider{constructionTp, concreteTp, sandTp} {
			if err := tp.Shutdown(context.Background()); err != nil {
				log.Printf("shutdown error: %v", err)
			}
		}
	}()

	// 5. start services (leaf-first to avoid startup races)
	go startSandService(sandTp)
	go startConcreteService(concreteTp)
	go startConstructionService(constructionTp)

	log.Println("Services up: :8080/build-house, :8081/provide-concrete, :8082/get-sand")
	select {}
}

func startConstructionService(tp *trace.TracerProvider) {
	otelTransport := otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp))
	client := &http.Client{Transport: otelTransport}

	mux := http.NewServeMux()
	mux.HandleFunc("/build-house", func(w http.ResponseWriter, r *http.Request) {
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
	})

	log.Panic(http.ListenAndServe(":8080", otelhttp.NewHandler(mux, "build-house", otelhttp.WithTracerProvider(tp))))
}

func startConcreteService(tp *trace.TracerProvider) {
	otelTransport := otelhttp.NewTransport(http.DefaultTransport, otelhttp.WithTracerProvider(tp))
	client := &http.Client{Transport: otelTransport}

	mux := http.NewServeMux()
	mux.HandleFunc("/provide-concrete", func(w http.ResponseWriter, r *http.Request) {
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
	})

	log.Panic(http.ListenAndServe(":8081", otelhttp.NewHandler(mux, "provide-concrete", otelhttp.WithTracerProvider(tp))))
}

func startSandService(tp *trace.TracerProvider) {
	mux := http.NewServeMux()
	mux.HandleFunc("/get-sand", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(750 * time.Millisecond)
		w.Write([]byte("sand delivered"))
	})

	log.Panic(http.ListenAndServe(":8082", otelhttp.NewHandler(mux, "get-sand", otelhttp.WithTracerProvider(tp))))
}

func newTracerProvider(name string, fileExp, otlpExp trace.SpanExporter) *trace.TracerProvider {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(name),
			semconv.ServiceVersion("v0.1.0"),
			attribute.String("environment", "demo"),
		),
	)
	return trace.NewTracerProvider(
		trace.WithBatcher(fileExp),
		trace.WithBatcher(otlpExp),
		trace.WithResource(r),
	)
}

func newStdOutExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
}

func newOLTPExporter() (trace.SpanExporter, error) {
	return otlptracehttp.New(context.Background(), otlptracehttp.WithInsecure())
}
