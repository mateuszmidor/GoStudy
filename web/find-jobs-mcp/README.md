# MCP Job Search Server for OpenCode

This MCP server exposes JustJoin.it job search by category (e.g. `go`, `java`, `python`) to OpenCode.

It can also run as a plain CLI that prints Go offers to stdout with the `-demo` flag.

## Run it

```sh
go run .           # MCP server (default)
go run . -demo     # just print the Go job offers
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

- `find_jobs` — takes a `category` argument (e.g. `go`, `java`, `python`, `javascript`, `ai`) and returns all matching offers from JustJoin.it as raw JSON, including title, company, salary, work mode, location, technologies and a direct link.
