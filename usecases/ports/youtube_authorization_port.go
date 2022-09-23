package ports

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
)

type YoutubeAuthorization interface {
	AuthCodeUrl(string) string
	Exchange(context.Context, string) (*oauth2.Token, error)
	Client(context.Context, *oauth2.Token) *http.Client
}

type YoutubeAuthsInputPort interface {
	Authorize(string) error
}

type YoutubeAuthsOutputPort interface {
	OutputRedirectUrl(string) error
	Test(string) error
}
