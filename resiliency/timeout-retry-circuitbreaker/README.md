# resiliency for HTTP requests client

Glossary: request is made of configured N-attempts

Means for http client resiliency:
- ensure Circuit Breaker (cross-request, so it counts against all failed requests from given http client)
- ensure Retry with exponential backoff (per-request, so every request has its own limit of attempts)
- set http client timeouts (per-attempt)
- ensure all attempts are executed within a deadline (overall timeout)

## Run it
- the retry policy is configured to attempt max 3x before failure, and the circuit breaker policy is configured to open after first failure
- the first request tries to GET an unexistent domain, attempts 3x and fails
- second request tries to GET an existing domain, but circuit is already open, so it fails

```sh
go run .
```

, result:
```
### Requesting https://non-existent.com
checking retry policy: need retry: true
checking retry policy: need retry: true
checking retry policy: need retry: true
checking circuit breaker policy: request failed: true
Error Get "https://non-existent.com": retries exceeded. last result: <nil>, last error: dial tcp: lookup non-existent.com: no such host
### Requesting https://www.google.com
Error Get "https://www.google.com": circuit breaker open
```