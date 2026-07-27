# prettylogs

A command-line tool for formatting JSON logs from stdin into human-readable output with colors.

## Usage

```bash
prettylogs [flags]
```

### Flags

- `-e <pattern>` - Regex pattern to exclude fields from output

## Examples

### Basic usage (no filtering)

```bash
echo '{"level":"info","time":"2024-03-06T10:00:00Z","caller":"main.go:42","msg":"server started","port":8080}' | go run main.go
```

Output:
```
[INFO]  [2024-03-06 10:00:00] main.go:42 server started | port:8080
```

### Exclude a single field

```bash
echo '{"level":"info","time":"2024-03-06T10:00:00Z","caller":"main.go:42","msg":"server started","port":8080}' | go run main.go -e "time"
```

Output (without time field):
```
[INFO]  main.go:42 server started | port:8080
```

### Exclude multiple fields using regex OR

```bash
echo '{"level":"debug","time":"2024-03-06T10:00:00Z","caller":"main.go:42","msg":"processing request","user_id":123,"request_id":"abc-456"}' | go run main.go -e "time|caller"
```

Output (without time and caller fields):
```
[DEBUG] processing request | user_id:123, request_id:abc-456
```

### Reading from Stern (Kubernetes logs)

```bash
stern my-service | go run main.go -e "time|caller"
```

## Input Format

The tool supports two input formats:

1. **Pure JSON**: One JSON object per line
   ```json
   {"level":"info","msg":"hello"}
   ```

2. **Stern format**: Pod name, container name, followed by JSON
   ```
   my-pod-123 my-container {"level":"info","msg":"hello"}
   ```

## Field Handling

- **Standard fields** (level, time, caller, msg) are formatted with colors and special formatting
- **Additional fields** are displayed as key:value pairs after the `|` separator
- **stacktrace** field is automatically excluded
- All fields (including standard ones) can be filtered using the `-e` flag
