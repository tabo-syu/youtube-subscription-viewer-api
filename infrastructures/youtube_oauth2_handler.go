package infrastructures

import (
	"context"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/youtube/v3"
)

type YoutubeOAuth2Handler struct {
	config *oauth2.Config
}

func NewYoutubeOAuth2Handler(clientSecret []byte) (*YoutubeOAuth2Handler, error) {
	config, err := google.ConfigFromJSON(clientSecret, youtube.YoutubeReadonlyScope)
	if err != nil {
		return nil, err
	}

	return &YoutubeOAuth2Handler{config}, nil
}

func (h *YoutubeOAuth2Handler) AuthCodeUrl(state string) string {
	return h.config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
}

func (h *YoutubeOAuth2Handler) Exchange(ctx context.Context, code string) (*oauth2.Token, error) {
	return h.config.Exchange(ctx, code)
}

func (h *YoutubeOAuth2Handler) Client(ctx context.Context, token *oauth2.Token) *http.Client {
	return h.config.Client(ctx, token)
}
