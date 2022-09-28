package ports

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type ChannelsRepository interface {
	GetFeed() ([]*entities.Video, error)
}

type YoutubeChannelsRepository interface {
	GetMyChannel() (*entities.User, error)
	AddClient(context.Context, *http.Client) error
}

type ChannelsInputPort interface {
	GetFeed() error
}

type ChannelsOutputPort interface {
	OutputChannels([]*entities.Channel) error
	OutputChannel(*entities.Channel) error
}
