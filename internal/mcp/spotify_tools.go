package mcptools

import (
	"context"
	"fmt"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"

	"github.com/dog0sd/spotify-mcp/internal/spotify"

	s "github.com/zmb3/spotify/v2"
)

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
