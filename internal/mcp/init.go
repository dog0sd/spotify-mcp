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
	toolPause := mcp.NewTool("spotify_pause",
		mcp.WithDescription("Pause playback on Spotify."),
	)
	toolNext := mcp.NewTool("spotify_next",
		mcp.WithDescription("Skip to the next track on Spotify."),
	)
	toolPrevious := mcp.NewTool("spotify_previous",
		mcp.WithDescription("Skip to the previous track on Spotify."),
	)
	toolSetVolume := mcp.NewTool("spotify_volume",
		mcp.WithDescription("Set volume in percents for current Spotify playback."),
		mcp.WithNumber("volume", mcp.Description("Integer value in percents between 0 and 100")),
	)
	toolCurrentlyPlaying := mcp.NewTool("spotify_currently_playing",
		mcp.WithDescription("Show what currently is playing in Spotify"),
	)
	toolListDevices := mcp.NewTool("spotify_list_devices",
		mcp.WithDescription("List existing user's Spotify devices"),
	)
	toolGetCover := mcp.NewTool("spotify_get_cover",
		mcp.WithDescription("Get the cover image URL for a track, album, or playlist by URI."),
		mcp.WithString("uri", mcp.Required(), mcp.Description("Spotify URI of the track, album, or playlist")),
	)
	toolGetPlaylists := mcp.NewTool("spotify_get_playlists",
		mcp.WithDescription("Get the list of user's playlists."),
	)
	toolGetTrackList := mcp.NewTool("spotify_get_track_list",
		mcp.WithDescription("Get the list of tracks in an album or playlist by URI."),
		mcp.WithString("uri", mcp.Required(), mcp.Description("Spotify URI of the album or playlist")),
	)
	s.AddTool(toolPlay, play)
	s.AddTool(toolSearch, search)
	s.AddTool(toolPause, pause)
	s.AddTool(toolNext, next)
	s.AddTool(toolPrevious, previous)
	s.AddTool(toolCurrentUser, currentUser)
	s.AddTool(toolSetVolume, setVolume)
	s.AddTool(toolListDevices, listDevices)
	s.AddTool(toolCurrentlyPlaying, currentlyPlaying)

	s.AddTool(toolGetCover, getCover)
	s.AddTool(toolGetPlaylists, getPlaylists)
	s.AddTool(toolGetTrackList, getTrackList)
}
