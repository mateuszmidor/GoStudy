# errors - stack trace

Create error with attached stack trace, then extract the stack trace from the error.

## Run

```sh
go run .
```

```text
read failed: open failed
main.openFile /Users/mmidor/SoftwareDevelopment/GoStudy/errors-stacktrace/main.go:48
main.readFile /Users/mmidor/SoftwareDevelopment/GoStudy/errors-stacktrace/main.go:44
main.main /Users/mmidor/SoftwareDevelopment/GoStudy/errors-stacktrace/main.go:17
runtime.main /usr/local/Cellar/go/1.21.4/libexec/src/runtime/proc.go:267
runtime.goexit /usr/local/Cellar/go/1.21.4/libexec/src/runtime/asm_amd64.s:1650
```