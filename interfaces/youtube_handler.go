package interfaces

import (
	"context"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

type YoutubeHandler interface {
	ListChannels([]string) (*youtube.ChannelsListCall, error)
	AddClient(context.Context, *http.Client) error
}
