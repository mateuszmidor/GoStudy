# go-vcr

For testing your integration with some HTTP APIs using cached HTTP responses.

* go get -v gopkg.in/dnaeon/go-vcr.v3/recorder
* https://github.com/dnaeon/go-vcr

## Run

1. first run, it will hit the real API and cache the response in [./fixtures/api-nbp-pl.yaml](./fixtures/api-nbp-pl.yaml)
1. next run, it will read cached response from file without hitting the API


```bash
go test  -count=1 .
go test  -count=1 .
```