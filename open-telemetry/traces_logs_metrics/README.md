# open-telemetry

This demo showcases OpenTelemetry distributed tracing, logging, and metrics across a chain of Go microservices.
Traces, logs, and metrics are collected by the OpenTelemetry Collector and visualized in Grafana (Tempo for traces, Loki for logs, Prometheus for metrics).

## Architecture

```
Go App (building:8080 → concrete:8081 → sand:8082)
        │
        │ OTLP/HTTP :4318 (traces, logs, metrics)
        ▼
  OTel Collector :4318
   ├── traces   OTLP/HTTP :4318  ──► Tempo      :4318 (ingest)
   ├── logs     OTLP/HTTP :3100  ──► Loki       :3100/otlp
   └── metrics  Prometheus :8889 ──► Prometheus :9090 (scrapes /metrics)

                      Tempo      :3200 (HTTP queries)  ◄──┐
                      Tempo      :9095 (gRPC queries)  ◄──┤
                      Loki       :3100 (HTTP queries)  ◄──┤
                      Prometheus :9090 (HTTP queries)  ◄──┴── Grafana :3000 (web ui)
```

## Run

```bash
make up      # start Grafana + Tempo + Loki + Prometheus + OTel Collector
make run     # start the Go services
make house   # trigger a chain of service→service→service requests
```

1. Open http://localhost:3000 in your browser (no login required)
1. In the left sidebar click **Explore** (compass icon).
- **Explore → Loki** to view logs 
    - for all logs try query: `{service_name=~".+"}`
    - to filter by trace id: `{service_name=~".+"} | trace_id = "41df8208f8908c978d1bc2e96a0eecb5"`
- **Explore → Tempo** to view traces
- **Explore → Prometheus** to view metrics
    - request duration per service: 
        ```
        sum by (http_route, job) (rate(http_server_request_duration_seconds_sum[$__range]))
        /
        sum by (http_route, job) (rate(http_server_request_duration_seconds_count[$__range]))
        ```
    - raw metrics available at http://localhost:8889/metrics
```bash
make down    # tear down the observability stack
```

## What changes in case of microservice?
- set global log provider with `global.SetLoggerProvider(lp)`
- set global logger with `slog.SetDefault(logger)` instead of storing it in the service struct
- set global trace provider with `otel.SetTracerProvider(tp)` instead of storing it in the service struct
- set global meter provider with `otel.SetMeterProvider(mp)` instead of storing it in the service struct

## Custom spans

```go
ctx, span :=  s.tp.Tracer("sand-service").Start(r.Context(), "gather-sand") // tp is *trace.TracerProvider
result, err := gatherSand()
span.End()
```

## Custom metrics

```go
...
```

## How metrics work

The app **pushes** metrics to the OTel Collector via OTLP/HTTP (same endpoint as traces and logs).
The `MeterProvider` holds a `PeriodicReader` that fires every 60s and flushes accumulated data.

Prometheus does not receive a push — it **pulls** by scraping the collector's `/metrics` endpoint
(`:8889`) every 15s. The collector bridges the two by converting incoming OTLP metrics into
the Prometheus text exposition format.

```
App  ──OTLP/HTTP push──►  OTel Collector :4318
                               │
                               │ exposes GET :8889/metrics
                               ▼
                          Prometheus :9090  (scrapes every 15s)
```