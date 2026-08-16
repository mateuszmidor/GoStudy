package main

import (
	"context"
	"encoding/json"
	"findjobsmcp/justjoinit"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// runServer starts the MCP server exposing job search by category.
func runServer(port int) {
	s := server.NewMCPServer(
		serverName,
		serverVersion,
		server.WithToolCapabilities(false),
	)

	tool := mcp.NewTool(
		"find_jobs",
		mcp.WithDescription("Finds IT job offers on JustJoin.it by category (e.g. go, java, python, javascript, ai)"),
		mcp.WithString("category", mcp.Required(), mcp.Description("Job category, e.g. go, java, python")),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		category, _ := request.GetArguments()["category"].(string)
		log.Printf("[INFO] Tool call: find_jobs category=%q", category)

		if category == "" {
			return mcp.NewToolResultText(`{ "jobs": [] }`), nil
		}

		offers, err := justjoinit.FetchAllOffers([]string{category})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
		}
		log.Printf("[INFO] Fetched %d offers for category %q", len(offers.Jobs), category)

		body, err := json.MarshalIndent(offers, "", "  ")
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
		}
		return mcp.NewToolResultText(string(body)), nil
	})

	httpServer := server.NewStreamableHTTPServer(s)
	addr := fmt.Sprintf("localhost:%d", port)

	log.Printf("Starting %s v%s on http://%s", serverName, serverVersion, addr)
	log.Printf("OpenCode config: add to .opencode.json:")
	log.Printf(`  "mcp": { "find-jobs": { "type": "remote", "url": "http://%s/mcp", "enabled": true } }`, addr)

	if err := httpServer.Start(addr); err != nil {
		log.Fatal(err)
	}
}
