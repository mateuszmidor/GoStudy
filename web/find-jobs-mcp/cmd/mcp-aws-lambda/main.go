// Command lambda packages the FindJobs MCP server as an AWS Lambda function.
//
// It speaks the MCP streamable HTTP protocol in stateless, non-streaming mode
// and is meant to be exposed through a Lambda Function URL.
package main

import (
	"findjobsmcp/mcpserver"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	serverName    = "find-jobs-mcp"
	serverVersion = "1.0.0"
)

func main() {
	adapter := newMCPAdapter(mcpserver.New(serverName, serverVersion))
	lambda.Start(adapter.handle)
}
