# demo

Prints the latest **Go** job offers from JustJoin.it to stdout. The category is
hardcoded to `go` - it's a quick way to see what the `find_jobs` MCP tool
returns without starting a server.

## Run it

```sh
go run . -category=go
```

Sample output:

```
Fetched 81 offers

1. Senior Software Engineer, Platform Reliability
   - Company: Asana
   - Salary: 26765-36169 PLN
   - Mode: hybrid
   - Location: Warszawa
   - Technologies: Go, AWS, Kubernetes
   - Link: https://justjoin.it/job-offer/asana-senior-software-engineer-platform-reliability-warszawa-go-8a93939c
```
