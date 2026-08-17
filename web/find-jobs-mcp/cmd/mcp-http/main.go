// Command mcp-http runs the FindJobs MCP server as a local HTTP server.
package main

import (
	"flag"
	"fmt"
	"log"

	"findjobsmcp/mcpserver"

	"github.com/mark3labs/mcp-go/server"
)

const (
	serverName    = "find-jobs-mcp"
	serverVersion = "1.0.0"
)

// main starts the MCP server exposing job search by category.
func main() {
	port := flag.Int("port", 8080, "port to listen on for the MCP server")
	flag.Parse()

	s := mcpserver.New(serverName, serverVersion)

	httpServer := server.NewStreamableHTTPServer(s)
	addr := fmt.Sprintf("localhost:%d", *port)

	log.Printf("Starting %s v%s on http://%s", serverName, serverVersion, addr)
	log.Printf("OpenCode config: add to .opencode.json:")
	log.Printf(`  "mcp": { "find-jobs": { "type": "remote", "url": "http://%s/mcp", "enabled": true } }`, addr)

	if err := httpServer.Start(addr); err != nil {
		log.Fatal(err)
	}
}
