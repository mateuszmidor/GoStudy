# mcp-aws-lambda

Runs the FindJobs MCP server as an AWS Lambda function, exposed through a
Lambda Function URL. It speaks the MCP streamable HTTP protocol in
**stateless, non-streaming** mode: sessions don't survive between invocations
(Lambda instances are ephemeral) and GET/SSE notification streams are rejected
with 405 instead of hanging.

## Build the deployment package

```sh
./cmd/mcp-aws-lambda/build_aws_lambda.sh
```

or manually:

```sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bootstrap ./cmd/mcp-aws-lambda
zip bootstrap.zip bootstrap
```

The script produces `bootstrap.zip` in the current directory.

## Deploy from the AWS console

1. **Lambda → Create function → Author from scratch**
2. Name it (e.g. `find-jobs-mcp`), runtime **Amazon Linux 2023**, architecture
   **x86_64** (must match the `GOARCH` used to build), handler `bootstrap`
3. **Upload** the zip as the function code
4. **Configuration → General configuration**: memory 128 MB, timeout 10 s
5. **Configuration → Function URL → Create function URL**:
   Auth type `NONE` (anyone with the URL can invoke it - fine for a demo),
   CORS: allow origin `*`, methods `POST, GET, DELETE`, headers `*`
6. Copy the Function URL, e.g. `https://abcd1234.lambda-url.eu-central-1.on.aws`

## Add it to OpenCode

```sh
opencode mcp add

# Name: find-jobs
# Type: Remote
# URL: https://abcd1234.lambda-url.eu-central-1.on.aws/mcp
```

Each OpenCode request (initialize, tools/list, tools/call) is one Lambda
invocation and one fresh MCP session - the `find_jobs` tool is stateless, so
this is transparent to the client.
