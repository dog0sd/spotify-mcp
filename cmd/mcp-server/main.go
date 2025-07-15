package main

import (
	"fmt"

	"os"

	"github.com/mark3labs/mcp-go/server"

	mcp "github.com/dog0sd/spotify-mcp/internal/mcp"
	"github.com/dog0sd/spotify-mcp/internal/spotify"
)

func main() {
	
	if len(os.Args) == 2 && os.Args[1] == "login" {
		_, err := spotify.GetToken()
		if err != nil {
			fmt.Printf("error: %v", err)
		} else {
			fmt.Println("Token obtained from file.")
		}
		return
	}
	
	s := server.NewMCPServer("Spotify MCP", "0.1.0", server.WithToolCapabilities(false))
	
	mcp.RegisterAllTools(s)

	server.ServeStdio(s)
}
