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
	OutputChannels([]*entities.Video) error
	OutputChannel(*entities.Video) error
	OutputError() error
}
