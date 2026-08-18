// Package mcpserver builds the MCP server exposing the find_jobs tool.
// It is shared between the local HTTP server and the AWS Lambda function.
package mcpserver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

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
		mcp.WithDescription("Finds IT job offers on JustJoin.it by category. Allowed categories: "+api.OfferCategoriesStr),
		mcp.WithString("category", mcp.Required(), mcp.Description("Job category. Allowed: "+api.OfferCategoriesStr)),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// decode request
		category, _ := request.GetArguments()["category"].(string)
		slog.Info("mcp tool call: find_jobs", slog.String("category", category))

		// validate request
		if category == "" {
			err := errors.New("offer category parameter is required. Allowed: " + api.OfferCategoriesStr)
			slog.Error(err.Error())
			return nil, err
		}

		// fetch
		fetchStart := time.Now()
		offers, err := justjoinit.FetchAllOffers([]string{category})
		if err != nil {
			slog.Error("failed to fetch offers", slog.Any("error", err))
			return mcp.NewToolResultText(fmt.Sprintf("failed to fetch offers: %v", err)), nil
		}
		slog.Info("fetched offers", slog.String("category", category), slog.Int("count", len(offers.Jobs)), slog.Duration("took", time.Since(fetchStart)))

		// format json
		body, err := json.MarshalIndent(offers, "", "  ")
		if err != nil {
			slog.Error(err.Error())
			return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
		}

		// respond
		return mcp.NewToolResultText(string(body)), nil
	})

	return s
}
