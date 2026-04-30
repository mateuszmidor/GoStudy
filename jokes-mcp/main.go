package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

const (
	serverName    = "jokes-mcp"
	serverVersion = "1.0.0"
	addr          = "localhost:8080"
	jokeAPIURL    = "https://v2.jokeapi.dev/joke/Programming?format=txt"
)

func main() {
	s := server.NewMCPServer(
		serverName,
		serverVersion,
		server.WithToolCapabilities(false),
	)

	tool := mcp.NewTool(
		"get_joke",
		mcp.WithDescription("Returns a random programming joke from jokeapi.dev"),
	)

	s.AddTool(tool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		log.Printf("[INFO] Tool call: %s", request.Params.Name)
		joke, err := fetchJoke()
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error: %v", err)), nil
		}
		log.Printf("[INFO] Joke: %s", joke)
		return mcp.NewToolResultText(joke), nil
	})

	httpServer := server.NewStreamableHTTPServer(s)

	log.Printf("Starting %s v%s on http://%s", serverName, serverVersion, addr)
	log.Printf("OpenCode config: add to .opencode.json:")
	log.Printf(`  "mcp": { "jokes": { "type": "remote", "url": "http://%s/mcp", "enabled": true } }`, addr)

	if err := httpServer.Start(addr); err != nil {
		log.Fatal(err)
	}
}

func fetchJoke() (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(jokeAPIURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
