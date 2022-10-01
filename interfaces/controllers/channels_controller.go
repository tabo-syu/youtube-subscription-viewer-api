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

func NewChannelsController(cr *gateways.ChannelsRepository, ycr *gateways.YoutubeChannelsRepository) *ChannelsController {
	return &ChannelsController{cr, ycr}
}

func (c *ChannelsController) interactor(ctx echo.Context) *interactors.ChannelsInteractor {
	return interactors.NewChannelsInteractor(
		presenters.NewListsPresenter(ctx),
		presenters.NewChannelsPresenter(ctx),
		presenters.NewVideosPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.channels,
	)
}

func (c *ChannelsController) GetFeed() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).GetFeed()
	}
}
