package interfaces

import (
	"context"
	"net/http"

	"google.golang.org/api/youtube/v3"
)

type YoutubeHandler interface {
	ListChannels(context.Context, *http.Client, []string) (*youtube.ChannelsListCall, error)
	ListSubscriptions(context.Context, *http.Client, []string) (*youtube.SubscriptionsListCall, error)
}
