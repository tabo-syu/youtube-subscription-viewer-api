package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type YoutubeAuthsController struct {
	authorization *gateways.YoutubeAuthorization
}

func NewYoutubeAuthsController(a interfaces.YoutubeOAuth2Handler) *YoutubeAuthsController {
	return &YoutubeAuthsController{
		gateways.NewYoutubeAuthorization(a),
	}
}

func (c *YoutubeAuthsController) interactor(ctx echo.Context) *interactors.YoutubeAuthsInteractor {
	return interactors.NewYoutubeAuthsInteractor(
		presenters.NewYoutubeAuthsPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.authorization,
	)
}

func (c *YoutubeAuthsController) Authorize(contextKey string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		state := ctx.Get(contextKey).(string)
		return c.interactor(ctx).Authorize(state)
	}
}
