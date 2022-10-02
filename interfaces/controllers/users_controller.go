package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type UsersController struct {
	users                *gateways.UsersRepository
	youtubeSubscriptions *gateways.YoutubeSubscrptionsRepository
}

func NewUsersController(ur *gateways.UsersRepository, ysr *gateways.YoutubeSubscrptionsRepository) *UsersController {
	return &UsersController{ur, ysr}
}

func (c *UsersController) interactor(ctx echo.Context) *interactors.UsersInteractor {
	return interactors.NewUsersInteractor(
		presenters.NewUsersPresenter(ctx),
		presenters.NewChannelsPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.users,
		c.youtubeSubscriptions,
	)
}

func (c *UsersController) GetMyself() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user, ok := ctx.Get("user").(*entities.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.interactor(ctx).GetMyself(user)
	}
}

func (c *UsersController) GetMySubscriptions() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		client, ok := ctx.Get("client").(*http.Client)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return c.interactor(ctx).GetMySubscriptions(
			ctx.Request().Context(),
			client,
		)
	}
}
