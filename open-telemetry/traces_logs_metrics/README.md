# open-telemetry

This demo showcases OpenTelemetry distributed tracing and logging across a chain of Go microservices.
Traces and logs are collected by the OpenTelemetry Collector and visualized in Grafana (Tempo for traces, Loki for logs).

## Architecture

```
Go App (building:8080 → concrete:8081 → sand:8082)
        │
        │ OTLP/HTTP :4318 (both logs and traces)
        ▼
  OTel Collector :4318
   ├── traces  OTLP/HTTP :4318  ──► Tempo :4318 (ingest)
   └── logs    OTLP/HTTP :3100  ──► Loki  :3100/otlp

                      Tempo :3200 (HTTP queries)  ◄──┐
                      Tempo :9095 (gRPC queries)  ◄──┤
                      Loki  :3100 (HTTP queries)  ◄──┴── Grafana :3000 (web ui)
```

## Run

```bash
make up      # start Grafana + Tempo + Loki + OTel Collector
make run     # start the Go services
make house   # trigger a chain of service→service→service requests
```

1. Open http://localhost:3000 in your browser (no login required)
1. In the left sidebar click **Explore** (compass icon).
- **Explore → Loki** to view logs 
    - for all logs try query: `{service_name=~".+"}`
    - to filter by trace id: `{service_name=~".+"} | trace_id = "41df8208f8908c978d1bc2e96a0eecb5"`
- **Explore → Tempo** to view traces

```bash
make down    # tear down the observability stack
```

## What changes in case of microservice?
- set global log provider with `global.SetLoggerProvider(lp)`
- set global logger with `slog.SetDefault(logger)` instead of storing it in the service struct
- set global trace provider with `otel.SetTracerProvider(tp)` instead of storing it in the service struct

## Custom spans

```go
ctx, span :=  s.tp.Tracer("sand-service").Start(r.Context(), "gather-sand") // tp is *trace.TracerProvider
result, err := gatherSand()
span.End()
```