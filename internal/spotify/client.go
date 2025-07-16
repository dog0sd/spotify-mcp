package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

const (
	redirectURL = "http://localhost:8821/callback"
	tokenFile   = "spotify_token.json"
	state       = "state-token"
)

var (
	authenticator = spotifyauth.New(
		spotifyauth.WithClientID(os.Getenv("SPOTIFY_CLIENT_ID")),
		spotifyauth.WithClientSecret(os.Getenv("SPOTIFY_CLIENT_SECRET")),
		spotifyauth.WithRedirectURL(redirectURL),
		spotifyauth.WithScopes(
			spotifyauth.ScopeUserReadPlaybackState,
			spotifyauth.ScopeUserReadCurrentlyPlaying,
			spotifyauth.ScopeUserModifyPlaybackState,
			spotifyauth.ScopePlaylistModifyPrivate,
			spotifyauth.ScopePlaylistModifyPublic,
			spotifyauth.ScopePlaylistReadPrivate,
			spotifyauth.ScopeUserReadCurrentlyPlaying,
			spotifyauth.ScopeUserLibraryRead,
		),
	)
	tokenChan = make(chan *oauth2.Token)
)

func saveToken(tok *oauth2.Token) error {
	data, err := json.Marshal(tok)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(tokenFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(data); err != nil {
		return fmt.Errorf("failed to write token to file: %w", err)
	}
	return nil
}

func loadToken() (*oauth2.Token, error) {
	file, err := os.Open(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("could not open token file: %w", err)
	}
	defer file.Close()

	var token oauth2.Token
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&token); err != nil {
		return nil, fmt.Errorf("could not decode token: %w", err)
	}

	if !token.Valid() {
		return nil, fmt.Errorf("token is invalid or expired")
	}

	return &token, nil
}

func GetToken() (*oauth2.Token, error) {
	var token *oauth2.Token

	savedToken, err := loadToken()

	if err == nil {
		token = savedToken
		return token, nil
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := authenticator.AuthURL(state)
		http.Redirect(w, r, url, http.StatusFound)
	})

	mux.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		if r.FormValue("state") != state {
			http.Error(w, "state mismatch", http.StatusBadRequest)
			return
		}

		freshToken, err := authenticator.Token(r.Context(), state, r)
		if err != nil {
			http.Error(w, "token error: "+err.Error(), http.StatusForbidden)
			return
		}
		if err := saveToken(freshToken); err != nil {
			http.Error(w, "failure while saving token: "+err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Println("Authentication completed.")
		tokenChan <- freshToken
	})

	server := &http.Server{Addr: ":8821", Handler: mux}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	fmt.Println("1. Open http://localhost:8821/ in your browser")
	fmt.Println("2. Login into your Spotify account")

	token = <-tokenChan
	server.Shutdown(context.Background())
	return token, nil
}

func NewClient(ctx context.Context) (*spotify.Client, error) {
	var client *spotify.Client

	token, err := loadToken()
	if err == nil {
		client = spotify.New(authenticator.Client(ctx, token))
	}
	return client, err
}

func CurrentUser(ctx context.Context, c *spotify.Client) (*spotify.PrivateUser, error) {

	user, err := c.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, err
}
