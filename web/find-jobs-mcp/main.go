package main

import (
	"flag"
	"fmt"
)

const (
	serverName    = "find-jobs-mcp"
	serverVersion = "1.0.0"
)

// main is the entry point - serves job offers over MCP by default,
// or prints them to stdout when -demo is given.
func main() {
	demoMode := flag.Bool("demo", false, "only print the job offers to stdout instead of running the MCP server")
	port := flag.Int("port", 8080, "port to listen on for the MCP server")
	flag.Parse()

	if *demoMode {
		justPrintOffers("go")
		return
	}

	fmt.Println("Running MCP server. To just print the jobs, pass -demo flag")
	runServer(*port)
}
