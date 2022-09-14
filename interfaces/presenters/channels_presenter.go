package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ChannelsPresenter struct{}

var _ ports.ChannelsOutputPort = (*ChannelsPresenter)(nil)

func NewChannelsPresenter() *ChannelsPresenter {
	return &ChannelsPresenter{}
}

func (p *ChannelsPresenter) OutputChannels(ctx echo.Context, channels []*entities.Channel) error {
	return nil
}

func (p *ChannelsPresenter) OutputChannel(ctx echo.Context, channel *entities.Channel) error {
	return nil
}
