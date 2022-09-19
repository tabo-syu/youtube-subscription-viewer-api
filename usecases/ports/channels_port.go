package ports

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type ChannelsRepository interface {
	GetFeed() ([]*entities.Video, error)
}

type ChannelsInputPort interface {
	GetFeed() error
}

type ChannelsOutputPort interface {
	OutputChannels([]*entities.Channel) error
	OutputChannel(*entities.Channel) error
}
