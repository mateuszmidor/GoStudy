# mcp-http

Runs the FindJobs MCP server as a regular HTTP server on your machine, speaking
the MCP streamable HTTP protocol. Point OpenCode at it to search JustJoin.it
job offers by category.

## Run it

```sh
go run ./cmd/mcp-http            # listens on localhost:8080
go run ./cmd/mcp-http -port 9000 # custom port
```

## Add it to OpenCode

```sh
opencode mcp add

# Name: find-jobs
# Type: Remote
# URL: http://localhost:8080/mcp
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
