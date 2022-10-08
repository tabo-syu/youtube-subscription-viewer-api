package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type ChannelsController struct {
	channels        *gateways.ChannelsRepository
	youtubeChannels *gateways.YoutubeChannelsRepository
}

func NewChannelsController(
	cr *gateways.ChannelsRepository,
	ycr *gateways.YoutubeChannelsRepository,
) *ChannelsController {
	return &ChannelsController{cr, ycr}
}

func (cc *ChannelsController) interactor(c echo.Context) *interactors.ChannelsInteractor {
	return interactors.NewChannelsInteractor(
		presenters.NewListsPresenter(c),
		presenters.NewChannelsPresenter(c),
		presenters.NewVideosPresenter(c),
		presenters.NewErrorsPresenter(c),
		cc.channels,
	)
}

func (cc *ChannelsController) GetFeed() echo.HandlerFunc {
	return func(c echo.Context) error {
		return cc.interactor(c).GetFeed()
	}
}
