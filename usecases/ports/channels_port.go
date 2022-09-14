package ports

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type ChannelsRepository interface {
	GetFeed() ([]*entities.Video, error)
}

type ChannelsInputPort interface {
	GetFeed() error
}

type ChannelsOutputPort interface {
	OutputChannels(echo.Context, []*entities.Channel) error
	OutputChannel(echo.Context, *entities.Channel) error
}
