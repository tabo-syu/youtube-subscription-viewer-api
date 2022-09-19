package gateways

import (
	"context"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
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

func (a *YoutubeAuthorization) Exchange(ctx context.Context, code string) (*entities.User, error) {
	_, err := a.oAuth2.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return &entities.User{}, nil
}
