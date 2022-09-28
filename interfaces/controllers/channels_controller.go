package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type ChannelsController struct {
	channelsRepository        *gateways.ChannelsRepository
	youtubeChannelsRepository *gateways.YoutubeChannelsRepository
}

func NewChannelsController(s interfaces.SqlHandler, y interfaces.YoutubeHandler) *ChannelsController {
	return &ChannelsController{
		gateways.NewChannelsRepository(s),
		gateways.NewYoutubeChannelsRepository(y),
	}
}

func (c *ChannelsController) interactor(ctx echo.Context) *interactors.ChannelsInteractor {
	return interactors.NewChannelsInteractor(
		presenters.NewListsPresenter(ctx),
		presenters.NewChannelsPresenter(ctx),
		presenters.NewVideosPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.channelsRepository,
	)
}

func (c *ChannelsController) GetFeed() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).GetFeed()
	}
}
