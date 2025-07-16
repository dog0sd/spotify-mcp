package mcptools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)


func RegisterAllTools(s *server.MCPServer) {
	toolCreatePlaylist := mcp.NewTool("spotify_create_playlist",
		mcp.WithDescription("Create a new Spotify playlist for the current user. Returns Playlist URI"),
		mcp.WithString("name", mcp.Required(), mcp.Description("Name of the new playlist")),
		mcp.WithString("description", mcp.Description("Description of the playlist")),
		mcp.WithBoolean("public", mcp.Description("Whether the playlist is public (default: false)"), mcp.DefaultBool(false)),
	)
	toolAddTracksToPlaylist := mcp.NewTool("spotify_add_tracks_to_playlist",
		mcp.WithDescription("Add tracks to a Spotify playlist by playlist URI or ID and track URIs."),
		mcp.WithString("playlist_uri", mcp.Required(), mcp.Description("Spotify URI or ID of the playlist")),
		mcp.WithArray("track_uris", mcp.WithStringItems(), mcp.Required(), mcp.Description("List of Spotify track URIs to add")),
	)
	// Playback controls
	toolPlay := mcp.NewTool("spotify_play",
		mcp.WithDescription("Allows to play/resume playing the song in Spotify. If song name or artist are omited(they aren't required) last song will be played."),
		mcp.WithString("uri", mcp.Description("Spotify URI of the track, or album, or playlist")),
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
	// Information retrieval
	toolCurrentUser := mcp.NewTool("spotify_whoami",
		mcp.WithDescription("Allows to get information about current Spotify user(like email, display name)"),
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
	toolSearch := mcp.NewTool("spotify_search",
		mcp.WithDescription("Allows to search track/album/playlist/author on Spotify"),
		mcp.WithString("query", mcp.Required(), mcp.Description("search query")),
		mcp.WithString("type", mcp.Description("Search type"), mcp.DefaultString("track"), mcp.Enum("track", "album", "author", "playlist")),
	)
	// Registration
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
	s.AddTool(toolCreatePlaylist, createPlaylist)
	s.AddTool(toolAddTracksToPlaylist, addTracksToPlaylist)
	s.AddTool(toolGetTrackList, getTrackList)
}
