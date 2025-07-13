package mcptools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func RegisterAllTools(s *server.MCPServer) {
	toolPlay := mcp.NewTool("spotify_play",
		mcp.WithDescription("Allows to play/resume playing the song in Spotify. If song name or artist are omited(they aren't required) last song will be played."),
		mcp.WithString("uri", mcp.Description("Spotify URI of the track, or album")),
	)
	toolCurrentUser := mcp.NewTool("spotify_whoami",
		mcp.WithDescription("Allows to get information about current Spotify user(like email, display name)"),
	)
	toolSearch := mcp.NewTool("spotify_search",
		mcp.WithDescription("Allows to search track/album/playlist/author on Spotify"),
		mcp.WithString("query", mcp.Required(), mcp.Description("search query")),
		mcp.WithString("type", mcp.Description("Search type"), mcp.DefaultString("track"), mcp.Enum("track", "album", "author", "playlist")),
    )
	s.AddTool(toolPlay, play)
	s.AddTool(toolCurrentUser, play)
	s.AddTool(toolSearch, search)
}

