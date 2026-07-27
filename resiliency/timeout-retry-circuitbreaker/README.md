# resiliency for HTTP requests client

Glossary: request is made of configured N-attempts

Means for http client resiliency:
- ensure Circuit Breaker (cross-request, so it counts against all failed requests from given http client)
- ensure Retry with exponential backoff (per-request, so every request has its own limit of attempts)
- set http client timeouts (per-attempt)
- ensure all attempts are executed within a deadline (overall timeout)