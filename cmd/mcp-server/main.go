package main

import (
	"log"

	"github.com/mark3labs/mcp-go/server"

	mcp "github.com/dog0sd/spotify-mcp/internal/mcp"
)

func main() {
	s := server.NewMCPServer("Spotify MCP", "0.1.0", server.WithToolCapabilities(false))

	mcp.RegisterAllTools(s)

	if err := server.ServeStdio(s); err != nil {
		log.Fatal("server error: ", err)
	}
}
