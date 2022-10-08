package ports

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"golang.org/x/oauth2"
)

type YoutubeAuthorization interface {
	AuthCodeURL(string) string
	Exchange(context.Context, string) (*oauth2.Token, error)
	Client(context.Context, *oauth2.Token) *http.Client
	TokenSource(context.Context, *oauth2.Token) oauth2.TokenSource
}

type YoutubeAuthsInputPort interface {
	Authorize(string) error
}

type YoutubeAuthsOutputPort interface {
	OutputRedirectURL(string) error
	Login(*entities.User) error
}
