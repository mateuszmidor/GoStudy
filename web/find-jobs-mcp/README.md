# find-jobs-mcp

An MCP server that exposes JustJoin.it job search by category (e.g. `go`,
`java`, `python`) to OpenCode. Each offer includes title, company, salary,
work mode, location, technologies and a direct link.

The project can be run in three ways, one command per directory under
[`cmd/`](cmd/):

| Command | Purpose |
| --- | --- |
| [`demo -json`](cmd/demo/README.md) | prints the latest Go job offers to stdout, no server needed |
| [`mcp-http`](cmd/mcp-http/README.md) | serves the MCP server as a local HTTP server for OpenCode |
| [`mcp-aws-lambda`](cmd/mcp-aws-lambda/README.md) | packages the MCP server as an AWS Lambda function with a Function URL |

Each command has its own README with run instructions, OpenCode setup and
deployment steps.
