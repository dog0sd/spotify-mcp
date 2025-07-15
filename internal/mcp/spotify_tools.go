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
			result = append(result, fmt.Sprintf("URI: \"%s\", Artist: \"%s\", Song: \"%s\"", track.URI, track.Artists[0], track.Name))
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

