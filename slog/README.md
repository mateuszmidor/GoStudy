# slog

Structured logging introduced in go 1.21, supports default format, Text format and JSON format.

## invocation

```go
slog.Info("hello world!", "field1", 42) // field1=42
```

## default format

```text
2023/10/08 07:38:08.844741 main.go:14: [slog-demo] INFO hello world! field1=42
```

## text format

```text
time=2023-10-08T09:38:08.844+02:00 level=DEBUG msg="hello world!" field1=42
```

## json format

```json
{"time":"2023-10-08T09:42:43.184688767+02:00","level":"DEBUG","msg":"hello world!","field1":42}
```