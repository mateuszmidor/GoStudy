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

## Run it as an AWS Lambda function

The `lambda/` directory builds the same server as a Lambda function. It speaks
the MCP streamable HTTP protocol in **stateless, non-streaming** mode:
sessions don't survive between invocations (Lambda instances are ephemeral)
and GET/SSE notification streams are rejected with 405 instead of hanging.

### Build the deployment package

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bootstrap ./lambda
zip bootstrap.zip bootstrap
```

### Deploy from the AWS console

1. **Lambda → Create function → Author from scratch**
2. Name it (e.g. `find-jobs-mcp`), runtime **Amazon Linux 2023**, architecture
   **x86_64** (must match the `GOARCH` used to build), handler `bootstrap`
3. **Upload** the zip as the function code
4. **Configuration → General configuration**: memory 256 MB, timeout 30 s
5. **Configuration → Function URL → Create function URL**:
   Auth type `NONE` (anyone with the URL can invoke it - fine for a demo),
   CORS: allow origin `*`, methods `POST, GET, DELETE`, headers `*`
6. Copy the Function URL, e.g. `https://abcd1234.lambda-url.eu-central-1.on.aws`

### Add it to OpenCode

```sh
opencode mcp add

# Name: find-jobs
# Type: Remote
# URL: https://abcd1234.lambda-url.eu-central-1.on.aws/mcp
```

Each OpenCode request (initialize, tools/list, tools/call) is one Lambda
invocation and one fresh MCP session - the `find_jobs` tool is stateless, so
this is transparent to the client.
