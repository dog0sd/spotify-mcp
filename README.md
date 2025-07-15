# Spotify MCP server


This project provides a Model Context Protocol (MCP) server for controlling and querying Spotify via the [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go). It exposes a set of tools for playback, search, and user information and more in the future

## Configuration

To use this server, you must configure a Spotify App and provide credentials.  
**See the detailed setup guide here:** [SETUP.md](./SETUP.md)

**Required environment variables:**
- `SPOTIFY_CLIENT_ID`
- `SPOTIFY_CLIENT_SECRET`

Set these in your environment before running the server.

## Usage

1. **Install dependencies** (if needed):
   ```sh
   go mod tidy
   ```

2. **Run the server:**
   ```sh
   go run ./cmd/mcp-server
   ```
   The server will start and listen for MCP requests via stdio.

3. **Connect an MCP client** to interact with the server and use the available tools.

## Tools

The following tools are currently available:

- **spotify_play**
  - Play or resume playback on Spotify.
  - Parameters:
    - `uri` (optional): Spotify URI of the track or album.

- **spotify_pause**
  - Pause playback on Spotify..

- **spotify_next**
  - Skip to the next track on Spotify..

- **spotify_previous**
  - Skip to the previous track on Spotify..

- **spotify_search**
  - Search for tracks, albums, playlists, or artists.
  - Parameters:
    - `query` (required): Search query string.
    - `type` (optional): One of `track`, `album`, `author`, `playlist`. Default: `track`.

- **spotify_whoami**
  - Get information about the current Spotify user (e.g., email, display name).


## See Also

- [SETUP.md](./SETUP.md) — How to configure your Spotify App and credentials.

---

## TODO

- Add more tools
- Add tests and usage examples.
- Support for 
