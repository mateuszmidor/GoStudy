# open-telemetry

This demo showcases OpenTelemetry distributed tracing, logging, and metrics across a chain of Go microservices.
Traces, logs, and metrics are collected by the OpenTelemetry Collector and visualized in Grafana (Tempo for traces, Loki for logs, Prometheus for metrics).

## Run

```bash
make up      # start Grafana + Tempo + Loki + Prometheus + OTel Collector
make run     # start the Go services
make houses   # trigger a chain of service→service→service requests
```

1. Open http://localhost:3000 in your browser (no login required)
1. In the left sidebar click **Explore** (compass icon).
- **Explore → Loki** to view logs 
    - for all logs try query: `{service_name=~".+"}`
    - to filter by trace id: `{service_name=~".+"} | trace_id = "41df8208f8908c978d1bc2e96a0eecb5"`
- **Explore → Tempo** to view traces
- **Explore → Prometheus** to view metrics
    - num of http status codes from /build-house per-minute: 
        ```
        sum by (http_response_status_code) (
        increase(http_server_request_duration_seconds_count{http_route="/build-house"}[1m])
        )
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
meter := s.mp.Meter("sand-service") // mp is *sdkmetric.MeterProvider
sandGatheringFailures, _ := meter.Int64Counter("sand_gathering_failures", metric.WithDescription("Sand gather failed"))
sandGatheringFailures.Add(r.Context(), 1, metric.WithAttributes(attribute.String("result", "failure")))
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

## How tail sampling works

The collector does **not** forward every trace to Tempo. The `tail_sampling` processor buffers each trace
for 10s (`decision_wait`) so all its spans arrive, then keeps the whole trace if **any** policy matches
(policies are OR-ed):

- `errors-always` - any span in the trace has ERROR status
- `slow-traces` - the trace took longer than 500ms
- `baseline-sample` - random 50% of the remaining traces

Consequences: failed and slow requests are always visible in Tempo, healthy fast ones mostly are not,
and traces arrive up to ~10s later than logs/metrics because of the buffering window.

```
spans ──► otlp receiver ──► tail_sampling (drop or keep whole trace) ──► batch ──► Tempo
```

**Production note:** tail sampling only works correctly if *all* spans of a trace reach the **same collector
instance** - the decision is made per-trace on one node. This demo runs a single collector, so it just works.
In production with multiple collectors you must use a tiered setup: agent collectors forward spans to a
gateway tier through the `loadbalancing` exporter keyed by `traceID` (routing_table), and only the gateway
instances run `tail_sampling`.

```
app ──► agent collectors ──► loadbalancing exporter (hash by traceID) ──► gateway collectors (tail_sampling) ──► backends
```
