// Package mcpserver builds the MCP server exposing the find_jobs tool.
// It is shared between the local HTTP server and the AWS Lambda function.
package mcpserver

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
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
		mcp.WithDescription("Finds IT job offers on JustJoin.it. All parameters are optional; combine them to narrow results."),
		mcp.WithString("category", mcp.Description("Job category; prefer this to 'keywords' if possible. Allowed: "+api.OfferCategoriesStr)),
		mcp.WithString("keywords", mcp.Description("Full-text search phrase matched against job titles and skills, e.g. \"golang\" or \"react developer\"")),
		mcp.WithString("city", mcp.Description("Filter by city name, e.g. \"Warszawa\", \"Kraków\", \"Berlin\". Diacritics optional")),
		mcp.WithString("experienceLevels", mcp.Description("Comma-separated seniority levels. Allowed: intern, junior, mid, senior, manager, c_level")),
		mcp.WithString("employmentTypes", mcp.Description("Comma-separated contract types. Allowed: b2b, permanent, uoz (mandate contract), internship")),
		mcp.WithString("remoteWorkOptions", mcp.Description("Comma-separated workplace modes. Allowed: remote, hybrid, office")),
		mcp.WithBoolean("withSalary", mcp.Description("When true, only return offers that disclose salary")),
		mcp.WithNumber("minSalary", mcp.Description("Minimum salary threshold in PLN, e.g. 15000")),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args := request.GetArguments()

		params := justjoinit.SearchParams{}
		if v, _ := args["category"].(string); v != "" {
			params.Categories = []string{v}
		}
		if v, _ := args["keywords"].(string); v != "" {
			params.Keywords = v
		}
		if v, _ := args["city"].(string); v != "" {
			params.City = v
		}
		if v, _ := args["experienceLevels"].(string); v != "" {
			params.ExperienceLevels = splitTrim(v)
		}
		if v, _ := args["employmentTypes"].(string); v != "" {
			params.EmploymentTypes = splitTrim(v)
		}
		if v, _ := args["remoteWorkOptions"].(string); v != "" {
			params.RemoteWorkOptions = splitTrim(v)
		}
		if v, ok := args["withSalary"].(bool); ok {
			params.WithSalary = &v
		}
		if v, ok := args["minSalary"].(float64); ok {
			n := int(v)
			params.MinSalary = &n
		}

		slog.Info("mcp tool call: find_jobs", slog.Any("params", params))

		// fetch
		fetchStart := time.Now()
		offers, err := justjoinit.FetchAllOffers(params)
		if err != nil {
			slog.Error("failed to fetch offers", slog.Any("error", err))
			return mcp.NewToolResultText(fmt.Sprintf("failed to fetch offers: %v", err)), nil
		}
		slog.Info("fetched offers", slog.Int("count", len(offers.Jobs)), slog.Duration("took", time.Since(fetchStart)))

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

// splitTrim splits a comma-separated string and trims whitespace from each part,
// filtering out empty strings.
func splitTrim(s string) []string {
	parts := strings.Split(s, ",")
	result := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			result = append(result, p)
		}
	}
	return result
}
