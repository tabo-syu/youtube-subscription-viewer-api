package gateways

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"golang.org/x/oauth2"
)

type YoutubeAuthorization struct {
	oAuth2 interfaces.YoutubeOAuth2Handler
}

func NewYoutubeAuthorization(o interfaces.YoutubeOAuth2Handler) *YoutubeAuthorization {
	return &YoutubeAuthorization{o}
}

func (a *YoutubeAuthorization) AuthCodeUrl(state string) string {
	return a.oAuth2.AuthCodeUrl(state)
}

func (a *YoutubeAuthorization) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return a.oAuth2.Exchange(ctx, code)
}

func (a *YoutubeAuthorization) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return a.oAuth2.Client(ctx, token)
}

func (a *YoutubeAuthorization) TokenSource(ctx context.Context, token *oauth2.Token) oauth2.TokenSource {
	return a.oAuth2.TokenSource(ctx, token)
}
