# Spotify MCP Server

![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)
![License](https://img.shields.io/github/license/dog0sd/spotify-mcp)

> **Control Spotify with the power of Model Context Protocol (MCP)!**  
> Seamlessly integrate Spotify playback, search, and device management into your AI workflows.

---

## 🚀 Overview

**Spotify MCP Server** bridges the gap between AI agents and your Spotify account.  
It exposes a suite of tools for playback, search, and device management, making it easy to control Spotify from any MCP-compatible client.

- 🎵 Play, pause, skip, and control playback
- 🔍 Search tracks, albums, playlists, and artists
- 📱 List and manage playback devices
- 🔊 Adjust volume
- 👤 Query user info and current playback
- 🖼️ Get album/playlist cover image

---

## 🛠️ Quick Start

1. **Download binary:**
   Download binary from [releases](/releases) and store it in e.g. `~/.mcp/`

2. **Configure your Spotify App credentials:**  
   See [SETUP.md](./SETUP.md) for a step-by-step guide.

3. **Login to Spotify to save credentials**
   **Important:** run `spotify-mcp` always in the same directory, to save and load the same `spotify_token.json` file.  
   1. Set `SPOTIFY_CLIENT_ID` and `SPOTIFY_CLIENT_SECRET` env variables
   2. `./spotify-mcp login`
   3. Go to [http://localhost:8821/](http://localhost:8821/) in your browser
   4. Log in into your account

3. **Configure Cursor or Claude desktop:**
   ```json
   "spotify-mcp": {
      "type": "stdio",
      "command": "~/.mcp/spotify-mcp.sh",
      "autoApprove": ["spotify_search", "spotify_whoami", "spotify_currently_playing", "spotify_volume", "spotify_list_devices"]
    }
   ```
   Add provided block to `mcpServers` in `mcp.json`.  
   Content of the `spotify-mcp.sh` can be like this:  
   ```shell
   #!/usr/bin/env bash
   cd ~/.mcp
   ./spotify-mcp $@
   ```

4. **Connect your MCP client** and start controlling Spotify!

---

## 🧰 Available Tools

| Tool                        | Description                                      | Parameters                |
|-----------------------------|--------------------------------------------------|---------------------------|
| `spotify_play`              | Play or resume playback                          | `uri` (optional)          |
| `spotify_pause`             | Pause playback                                   |                           |
| `spotify_next`              | Skip to next track                               |                           |
| `spotify_previous`          | Skip to previous track                           |                           |
| `spotify_search`            | Search tracks, albums, playlists, or artists     | `query` (required), `type` (optional) |
| `spotify_whoami`            | Get current user info                            |                           |
| `spotify_currently_playing` | Get info about the currently playing track       |                           |
| `spotify_volume`            | Set playback volume                              | `volume` (0-100, required)|
| `spotify_list_devices`      | List available playback devices                  |                           |
| `spotify_get_cover`         | Get the cover image URL for a track, album, or playlist by URI | `uri` (required)          |
| `spotify_get_playlists`     | Get the list of user's playlists                |                           |
| `spotify_get_track_list`    | Get the list of tracks in an album or playlist by URI | `uri` (required)          |

---

## ⚙️ Configuration

To use this server, you must configure a Spotify App and provide credentials.  
**See the detailed setup guide here:** [SETUP.md](./SETUP.md)

**Required environment variables for login:**
- `SPOTIFY_CLIENT_ID`
- `SPOTIFY_CLIENT_SECRET`

Set these in your environment before running the server.

---

## 📚 Resources

- [SETUP.md](./SETUP.md) — Configure your Spotify App and credentials
- [Spotify for Developers](https://developer.spotify.com/documentation/web-api/)
- [Model Context Protocol (MCP)](https://github.com/mark3labs/mcp-go)

---

## 📝 TODO

- [ ] Add more tools
- [ ] Add logging
- [ ] Add tests and usage examples
- [ ] Do something with required re-login every hour🫠

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!  
Feel free to open an issue or submit a pull request.
