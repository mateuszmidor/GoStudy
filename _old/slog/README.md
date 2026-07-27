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

```sh
time=2023-10-08T10:18:08.170+02:00 level=DEBUG source=/home/user/SoftwareDevelopment/GoStudy/slog/main.go:22 msg="hello world!" field1=42
```

## json format

```json
{"time":"2023-10-08T10:18:08.170959432+02:00","level":"DEBUG","source":{"function":"main.main","file":"/home/user/SoftwareDevelopment/GoStudy/slog/main.go","line":28},"msg":"hello world!","field1":42}
```