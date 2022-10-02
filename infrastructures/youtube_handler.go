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

func NewYoutubeHandler() *YoutubeHandler {
	return &YoutubeHandler{}
}

func (h *YoutubeHandler) ListChannels(ctx context.Context, client *http.Client, part []string) (*youtube.ChannelsListCall, error) {
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return service.Channels.List(part), nil
}
