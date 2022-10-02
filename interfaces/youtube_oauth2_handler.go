package interfaces

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type YoutubeOAuth2Handler interface {
	AuthCodeUrl(string) string
	Exchange(context.Context, string) (*oauth2.Token, error)
	Client(context.Context, *oauth2.Token) *http.Client
	TokenSource(context.Context, *oauth2.Token) oauth2.TokenSource
}
