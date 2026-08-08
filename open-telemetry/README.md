# open telemetry

## Highlights

### Architecture

1. App generates logs, traces, metrics and OTel SDK exports them to provided endpoint (OTel Collector)
1. OTel Collector ingest->process->export the logs, traces and metrics to configured backends (loki, tempo, prometheus)
1. Grafana has configured sources to fetch logs, traces and metrics from the backends

```
Go App (building:8080 → concrete:8081 → sand:8082)
        │
        │ OTLP/HTTP :4318 (traces, logs, metrics)
        ▼
  OTel Collector :4318
   ├── traces   OTLP/HTTP :4318  ──► Tempo      :4318 (ingest)
   ├── logs     OTLP/HTTP :3100  ──► Loki       :3100/otlp
   └── metrics  Prometheus :8889 ──► Prometheus :9090 (Prometheus scrapes :8889/metrics and exposes :9090/api)

                      Tempo      :3200 (HTTP queries)  ◄──┐
                      Tempo      :9095 (gRPC queries)  ◄──┤
                      Loki       :3100 (HTTP queries)  ◄──┤
                      Prometheus :9090 (HTTP queries)  ◄──┴── Grafana :3000 (web ui)
```
**ABOUT METRICS**: 
- your raw metrics exported from app are available at :8889/metrics (OTel collector port)
- what is available at :9090/metrics is some Prometheus internal metrics about itself
- Grafana uses Prometheus query api (like :9090/api/v1/query) and not the :9090/metrics enpoint.

### Logs

- logging in OTel is not built from ground up like traces and metrics
- app should use stdlib logger to generate logs and OTel "bridge" (implemented as log handler) to export logs to OTel Collector

### Traces

- automatic trace propagation in http headers: trace_id, span_id, parent_span_id
- can attach static Resource with data like service name, version
- can attach dynamic attributes like http method, http url
- should be sampled in production

### Metrics

- Params: name, description, unit, attributes (dimensions)
- Types:
    - counter
    - up-down counter
    - gauge
    - histogram
- can create View for a Metric to customize its name, exported dimensions, aggregation, bucket boundaries

