# mcp-local

Runs the FindJobs MCP server over **stdio** — no HTTP server, no network port.
OpenCode spawns this as a subprocess and communicates via stdin/stdout using the
JSON-RPC MCP protocol.

## Run it

```sh
go run ./cmd/mcp-local
```

## Add it to OpenCode

```sh
opencode mcp add

# Name: find-jobs
# Type: Command
# Command: go run ./cmd/mcp-local
```

Or add directly to `.opencode.json`:

```json
{
  "mcp": {
    "find-jobs": {
      "type": "command",
      "command": "go run ./cmd/mcp-local",
      "enabled": true
    }
  }
}
```

## Check it

```sh
opencode mcp list
```

## Use it

```sh
opencode run 'find Go developer jobs'
```

The server exposes a single tool:

- `find_jobs` — takes a `category` argument (e.g. `go`, `java`, `python`,
  `javascript`, `ai`) and returns all matching offers from JustJoin.it as raw
  JSON, including title, company, salary, work mode, location, technologies
  and a direct link.

## How it works

Unlike `mcp-http` (which serves an HTTP endpoint) or `mcp-aws-lambda` (which
runs as a Lambda function), this command uses the **stdio transport**:

1. OpenCode spawns the process (`go run ./cmd/mcp-local`)
2. OpenCode writes JSON-RPC messages to the process's **stdin**
3. The process writes JSON-RPC responses to its **stdout**
4. When OpenCode shuts down, it sends SIGTERM — the process exits gracefully

This is the simplest MCP transport: no ports, no CORS, no deployment. Ideal for
local development and personal use.
