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

func NewYoutubeAuthorization(g interfaces.YoutubeOAuth2Handler) *YoutubeAuthorization {
	return &YoutubeAuthorization{g}
}

func (a *YoutubeAuthorization) AuthCodeUrl(state string) string {
	return a.oAuth2.AuthCodeUrl(state)
}

func (a *YoutubeAuthorization) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	token, err := a.oAuth2.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (a *YoutubeAuthorization) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return a.oAuth2.Client(ctx, token)
}
