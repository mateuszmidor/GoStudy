package main

import (
	"fmt"
	"log"

	"findjobsmcp/mcpserver"

	"github.com/mark3labs/mcp-go/server"
)

// runServer starts the MCP server exposing job search by category.
func runServer(port int) {
	s := mcpserver.New(serverName, serverVersion)

	httpServer := server.NewStreamableHTTPServer(s)
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Starting %s v%s on http://%s", serverName, serverVersion, addr)
	log.Printf("OpenCode config: add to .opencode.json:")
	log.Printf(`  "mcp": { "find-jobs": { "type": "remote", "url": "http://%s/mcp", "enabled": true } }`, addr)

	if err := httpServer.Start(addr); err != nil {
		log.Fatal(err)
	}
}
