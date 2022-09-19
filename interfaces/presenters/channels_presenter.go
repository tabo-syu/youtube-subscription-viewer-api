package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ChannelsPresenter struct {
	ctx echo.Context
}

var _ ports.ChannelsOutputPort = (*ChannelsPresenter)(nil)

func NewChannelsPresenter(ctx echo.Context) *ChannelsPresenter {
	return &ChannelsPresenter{ctx}
}

func (p *ChannelsPresenter) OutputChannels(channels []*entities.Channel) error {
	return nil
}

func (p *ChannelsPresenter) OutputChannel(channel *entities.Channel) error {
	return nil
}
