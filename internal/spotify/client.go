package spotify

import (
	"context"
	"os"

	spotify "github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

const (
	redirectURL = "http://localhost:8821/callback"
	state       = "sldmflsdkfggelkrg"
)

func NewClient(ctx context.Context) (*spotify.Client, error) {
	clientId := os.Getenv("SPOTIFY_CLIENT_ID")
	clientSecret := os.Getenv("SPOTIFY_CLIENT_SECRET")
	config := &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     spotifyauth.TokenURL,
	}

	token, err := config.Token(ctx)
	if err != nil {
		return nil, err
	}

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient)
	return client, nil
}


func CurrentUser(ctx context.Context,c *spotify.Client) (*spotify.PrivateUser, error) {

	user, err := c.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	return user, err
}

