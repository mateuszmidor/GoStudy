// Command mcp-local runs the FindJobs MCP server over stdio.
//
// OpenCode spawns this process and communicates via stdin/stdout using
// the JSON-RPC MCP protocol — no HTTP server or port needed.
package main

import (
	"log"

	"findjobsmcp/mcpserver"

	"github.com/mark3labs/mcp-go/server"
)

const (
	serverName    = "find-jobs-mcp"
	serverVersion = "1.0.0"
)

func main() {
	s := mcpserver.New(serverName, serverVersion)

	log.Printf("Starting %s v%s (stdio)", serverName, serverVersion)
	if err := server.ServeStdio(s); err != nil {
		log.Fatal(err)
	}
}
