# open-telemetry/traces_only

This demo showcases OpenTelemetry distributed tracing across a chain of Go microservices.
Traces are sent to Jaeger for visualization. Logs are written as JSON to stdout only.

## Architecture

```
Go App (building:8080 → concrete:8081 → sand:8082)
        │
        │ OTLP/HTTP :4318 (traces only)
        ▼
  Jaeger all-in-one :4318 (ingest)
  Jaeger UI         :16686 (web ui)

stdout ← logs (JSON, per-service name)
```

## Prerequisites

- Go 1.25+
- Docker

## Run

```bash
make up      # start Jaeger
make run     # start the Go services
make house   # trigger a chain of service→service→service requests
```

1. Open http://localhost:16686 in your browser
2. Select a service from the **Service** dropdown and click **Find Traces**

```bash
make down    # stop Jaeger
```

## Reference

https://medium.com/jaegertracing/introducing-native-support-for-opentelemetry-in-jaeger-eb661be8183c
