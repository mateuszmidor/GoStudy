// Command lambda packages the FindJobs MCP server as an AWS Lambda function.
//
// It speaks the MCP streamable HTTP protocol in stateless, non-streaming mode
// and is meant to be exposed through a Lambda Function URL.
//
// Build and package:
//
//	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap ./lambda
//	zip bootstrap.zip bootstrap
//
// Then upload bootstrap.zip in the AWS console: runtime "Amazon Linux 2023",
// handler "bootstrap" and create a Function URL for the function.
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
