package infrastructures

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

type YoutubeHandler struct {
	youtube *youtube.Service
}

func NewYoutubeHandler() *YoutubeHandler {
	return &YoutubeHandler{nil}
}

func (h *YoutubeHandler) AddClient(ctx context.Context, client *http.Client) error {
	service, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	h.youtube = service

	return nil
}

func (h *YoutubeHandler) ListChannels(part []string) (*youtube.ChannelsListCall, error) {
	if h.youtube == nil {
		return nil, entities.ErrYoutubeHandlerInitError
	}

	return h.youtube.Channels.List(part), nil
}
