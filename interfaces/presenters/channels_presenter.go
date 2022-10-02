package presenters

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ChannelsPresenter struct {
	c echo.Context
}

var _ ports.ChannelsOutputPort = (*ChannelsPresenter)(nil)

func NewChannelsPresenter(c echo.Context) *ChannelsPresenter {
	return &ChannelsPresenter{c}
}

func (p *ChannelsPresenter) OutputChannels(channels []*entities.Channel) error {
	return p.c.JSON(http.StatusOK, channels)
}

func (p *ChannelsPresenter) OutputChannel(channel *entities.Channel) error {
	return nil
}
