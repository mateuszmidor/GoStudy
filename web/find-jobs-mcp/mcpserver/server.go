// Package mcpserver builds the MCP server exposing the find_jobs tool.
// It is shared between the local HTTP server and the AWS Lambda function.
package mcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"findjobsmcp/api"
	"findjobsmcp/justjoinit"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// New builds the MCP server exposing the find_jobs tool.
func New(serverName, serverVersion string) *server.MCPServer {
	s := server.NewMCPServer(
		serverName,
		serverVersion,
		server.WithToolCapabilities(false),
	)

	tool := mcp.NewTool(
		"find_jobs",
		mcp.WithDescription("Finds IT job offers on JustJoin.it by category, e.g.: "+api.OfferCategoriesStr),
		mcp.WithString("category", mcp.Required(), mcp.Description("Job category. Allowed: "+api.OfferCategoriesStr)),
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

	return s
}
