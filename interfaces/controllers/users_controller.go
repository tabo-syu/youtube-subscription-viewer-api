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

func (uc *UsersController) interactor(c echo.Context) *interactors.UsersInteractor {
	return interactors.NewUsersInteractor(
		presenters.NewUsersPresenter(c),
		presenters.NewChannelsPresenter(c),
		presenters.NewErrorsPresenter(c),
		uc.users,
		uc.youtubeSubscriptions,
	)
}

func (uc *UsersController) GetMyself() echo.HandlerFunc {
	return func(c echo.Context) error {
		user, ok := c.Get("user").(*entities.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return uc.interactor(c).GetMyself(user)
	}
}

func (uc *UsersController) GetMySubscriptions() echo.HandlerFunc {
	return func(c echo.Context) error {
		client, ok := c.Get("client").(*http.Client)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return uc.interactor(c).GetMySubscriptions(
			c.Request().Context(),
			client,
		)
	}
}
