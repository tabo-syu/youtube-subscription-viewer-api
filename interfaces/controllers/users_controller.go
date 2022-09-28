package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type UsersController struct {
	repository *gateways.UsersRepository
}

func NewUsersController(s interfaces.SqlHandler, y interfaces.YoutubeHandler) *UsersController {
	return &UsersController{
		gateways.NewUsersRepository(s),
	}
}

func (c *UsersController) interactor(ctx echo.Context) *interactors.UsersInteractor {
	return interactors.NewUsersInteractor(
		presenters.NewUsersPresenter(ctx),
		presenters.NewVideosPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.repository,
	)
}

func (c *UsersController) GetMyself() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).GetMyself()
	}
}

func (c *UsersController) GetMySubscriptions() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).GetMySubscriptions()
	}
}
