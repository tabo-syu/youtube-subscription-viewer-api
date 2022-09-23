package infrastructures

import (
	"context"
	"net/http"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeHandler struct {
	youtube *youtube.Service
}

func NewYoutubeHandler(ctx context.Context, client *http.Client) (*YoutubeHandler, error) {
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return &YoutubeHandler{service}, err
}

func (h *YoutubeHandler) ListChannels(part []string) *youtube.ChannelsListCall {
	return h.youtube.Channels.List(part)
}
