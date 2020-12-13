# Prometheus client demo

This is simplest example, no dedicated prometheus module is used

## Metrics source

Prometheus pulls metrics from targets defined in scrape_configs in config.yml:
```yaml
scrape_configs:
- job_name: test_target
  honor_timestamps: true
  scrape_interval: 1000ms
  scrape_timeout: 500ms
  metrics_path: /metrics
  scheme: http
  static_configs:
  - targets:
    - localhost:8080
```

## Metrics format

The metrics are in key-value format:  
```
wave 0.5
```