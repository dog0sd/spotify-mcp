package mcptools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/dog0sd/spotify-mcp/internal/spotify"

	s "github.com/zmb3/spotify/v2"
)

// whoami
func currentUser(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	user, err := client.CurrentUser(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting current user: %v", err)
	}
	info := fmt.Sprintf("Name: %s\nEmail: %s", user.DisplayName, user.Email)
	return mcp.NewToolResultText(info), nil
}


// Play a track/playlist/album
func play(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	songURI := request.GetString("uri", "")
	if songURI != "" {
		err = client.PlayOpt(ctx, &s.PlayOptions{URIs: []s.URI{s.URI(songURI)}})
	} else {
		err = client.Play(ctx)
	}
	if err != nil {
		return nil, fmt.Errorf("error playing song: %v", err)
	}
	return mcp.NewToolResultText("done"), nil
}

// Search a track/author/playlist
func search(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	query, err := request.RequireString("query")
	if err != nil || query == "" {
		return nil, fmt.Errorf("query not provided")
	}
	var searchType s.SearchType
	switch request.GetString("type", "track") {
	case "album":
		searchType = s.SearchTypeAlbum
	case "playlist":
		searchType = s.SearchTypePlaylist
	case "artist":
		searchType = s.SearchTypeArtist
	default:
		searchType = s.SearchTypeTrack
	}
	searchResult, err := client.Search(ctx, query, searchType)
	if err != nil {
		return nil, fmt.Errorf("error searching: %v", err)
	}
	var result []string
	switch searchType {
	case s.SearchTypeAlbum:
		for _, album := range searchResult.Albums.Albums {
			result = append(result, fmt.Sprintf("URI: \"%s\", Artist: \"%s\", Album: \"%s\"", album.URI, album.Artists[0], album.Name))
		}
	case s.SearchTypePlaylist:
		for _, playlist := range searchResult.Playlists.Playlists {
			result = append(result, fmt.Sprintf("URI: \"%s\", Playlist: \"%s\"", playlist.URI, playlist.Name))
		}
	case s.SearchTypeArtist:
		for _, artist := range searchResult.Artists.Artists {
			result = append(result, fmt.Sprintf("URI: \"%s\", Artist: \"%s\"", artist.URI, artist.Name))
		}
	default:
		for _, track := range searchResult.Tracks.Tracks {
			artistNames := []string{}
			for _, a := range track.Artists {
				artistNames = append(artistNames, a.Name)
			}
			result = append(result, fmt.Sprintf("URI: \"%s\", Artist: \"%s\" Song: \"%s\"", track.URI, strings.Join(artistNames, ", "), track.Name))
		}
	}

	return mcp.NewToolResultText(strings.Join(result, "\n")), nil
}

// Pause playback
func pause(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	err = client.Pause(ctx)
	if err != nil {
		return nil, fmt.Errorf("error pausing playback: %v", err)
	}
	return mcp.NewToolResultText("done"), nil
}

// Play next track
func next(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	err = client.Next(ctx)
	if err != nil {
		return nil, fmt.Errorf("error skipping to next track: %v", err)
	}
	return mcp.NewToolResultText("done"), nil
}

// Play previous track
func previous(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	err = client.Previous(ctx)
	if err != nil {
		return nil, fmt.Errorf("error skipping to previous track: %v", err)
	}
	return mcp.NewToolResultText("done"), nil
}


// Get currently playing track
func currentlyPlaying(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	playing, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting currently playing track: %v", err)
	}
	if playing == nil || playing.Item == nil {
		return mcp.NewToolResultText("Nothing is currently playing."), nil
	}
	track := playing.Item
	artistNames := []string{}
	for _, a := range track.Artists {
		artistNames = append(artistNames, a.Name)
	}
	result := fmt.Sprintf("Now playing: \"%s\" by %s (URI: %s)", track.Name, strings.Join(artistNames, ", "), track.URI)
	return mcp.NewToolResultText(result), nil
}

// Set volume
func setVolume(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating spotify client: %v", err)
	}
	volume, err := request.RequireInt("volume")
	if err != nil {
		return nil, fmt.Errorf("missing or invalid 'volume' argument (should be 0-100)")
	}
	if volume < 0 || volume > 100 {
		return nil, fmt.Errorf("volume must be between 0 and 100")
	}
	err = client.Volume(ctx, volume)
	if err != nil {
		return nil, fmt.Errorf("error setting volume: %v", err)
	}
	return mcp.NewToolResultText(fmt.Sprintf("Volume set to %d%%", volume)), nil
}

// // Transfer playback to a different device
// func transferPlayback(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
// 	client, err := spotify.NewClient(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating spotify client: %v", err)
// 	}
// 	deviceID, ok := request.Args["device_id"].(string)
// 	if !ok || deviceID == "" {
// 		return nil, fmt.Errorf("missing or invalid 'device_id' argument")
// 	}
// 	err = client.TransferPlayback(ctx, deviceID)
// 	if err != nil {
// 		return nil, fmt.Errorf("error transferring playback: %v", err)
// 	}
// 	return mcp.NewToolResultText(fmt.Sprintf("Playback transferred to device: %s", deviceID)), nil
// }

// List available devices
func listDevices(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return mcp.NewToolResultError("error creating spotify client: " + err.Error()), nil
	}
	devices, err := client.PlayerDevices(ctx)
	if err != nil {
		return mcp.NewToolResultError("error getting devices: " + err.Error()), nil
	}
	if len(devices) == 0 {
		return mcp.NewToolResultText("No devices available."), nil
	}
	var result []string
	for _, d := range devices {
		active := ""
		if d.Active {
			active = " (active)"
		}
		result = append(result, fmt.Sprintf("Name: %s, ID: %s, Type: %s%s", d.Name, d.ID, d.Type, active))
	}
	return mcp.NewToolResultText(strings.Join(result, "\n")), nil
}

// Get the cover image URL for a track or album  by URI
func getCover(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return mcp.NewToolResultError("error creating spotify client: " + err.Error()), nil
	}
	uri, err := request.RequireString("uri")
	if err != nil || uri == "" {
		return mcp.NewToolResultError("missing or invalid 'uri' argument"), nil
	}
	var images []s.Image
	switch {
	case strings.HasPrefix(uri, "spotify:track:"):
		track, err := client.GetTrack(ctx, s.ID(strings.TrimPrefix(uri, "spotify:track:")))
		if err != nil {
			return mcp.NewToolResultError("error getting track: " + err.Error()), nil
		}
		images = track.Album.Images
	case strings.HasPrefix(uri, "spotify:album:"):
		album, err := client.GetAlbum(ctx, s.ID(strings.TrimPrefix(uri, "spotify:album:")))
		if err != nil {
			return mcp.NewToolResultError("error getting album: " + err.Error()), nil
		}
		images = album.Images
	case strings.HasPrefix(uri, "spotify:playlist:"):
		playlist, err := client.GetPlaylist(ctx, s.ID(strings.TrimPrefix(uri, "spotify:playlist:")))
		if err != nil {
			return mcp.NewToolResultError("error getting playlist: " + err.Error()), nil
		}
		images = playlist.Images
	default:
		return mcp.NewToolResultError("unsupported URI type for cover image"), nil
	}
	if len(images) == 0 {
		return mcp.NewToolResultText("No cover image found."), nil
	}
	return mcp.NewToolResultText(images[0].URL), nil
}

// Get the list of user's playlists
func getPlaylists(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return mcp.NewToolResultError("error creating spotify client: " + err.Error()), nil
	}
	playlists, err := client.CurrentUsersPlaylists(ctx)
	if err != nil {
		return mcp.NewToolResultError("error getting playlists: " + err.Error()), nil
	}
	if len(playlists.Playlists) == 0 {
		return mcp.NewToolResultText("No playlists found."), nil
	}
	var result []string
	for _, p := range playlists.Playlists {
		result = append(result, fmt.Sprintf("Name: %s, ID: %s, Tracks: %d", p.Name, p.ID, p.Tracks.Total))
	}
	return mcp.NewToolResultText(strings.Join(result, "\n")), nil
}

// Get the list of tracks in an album or playlist by URI
func getTrackList(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	client, err := spotify.NewClient(ctx)
	if err != nil {
		return mcp.NewToolResultError("error creating spotify client: " + err.Error()), nil
	}
	uri, err := request.RequireString("uri")
	if err != nil || uri == "" {
		return mcp.NewToolResultError("missing or invalid 'uri' argument"), nil
	}
	var tracks []s.SimpleTrack
	switch {
	case strings.HasPrefix(uri, "spotify:album:"):
		album, err := client.GetAlbum(ctx, s.ID(strings.TrimPrefix(uri, "spotify:album:")))
		if err != nil {
			return mcp.NewToolResultError("error getting album: " + err.Error()), nil
		}
		tracks = album.Tracks.Tracks
	case strings.HasPrefix(uri, "spotify:playlist:"):
		playlist, err := client.GetPlaylist(ctx, s.ID(strings.TrimPrefix(uri, "spotify:playlist:")))
		if err != nil {
			return mcp.NewToolResultError("error getting playlist: " + err.Error()), nil
		}
		for _, item := range playlist.Tracks.Tracks {
			tracks = append(tracks, item.Track.SimpleTrack)
		}
	default:
		return mcp.NewToolResultError("unsupported URI type for track list"), nil
	}
	if len(tracks) == 0 {
		return mcp.NewToolResultText("No tracks found."), nil
	}
	var result []string
	for i, t := range tracks {
		var artistNames []string
		for _, art := range t.Artists {
			artistNames = append(artistNames, art.Name)
		}
		result = append(result, fmt.Sprintf("%d. %s - %s", i+1, t.Name, strings.Join(artistNames, ", ")))
	}
	return mcp.NewToolResultText(strings.Join(result, "\n")), nil
}
