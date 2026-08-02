# open-telemetry

This demo showcases the usage of Open Telemetry in service of tracing an application.
https://medium.com/jaegertracing/introducing-native-support-for-opentelemetry-in-jaeger-eb661be8183c

## Run

```bash
make jaeger-up
make run
make house # trigger a chain of service->service->service requests
firefox localhost:16686
make jaeger-down
```
