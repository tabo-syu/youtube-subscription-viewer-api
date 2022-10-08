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
	channels             *gateways.ChannelsRepository
	youtubeSubscriptions *gateways.YoutubeSubscrptionsRepository
}

func NewUsersController(
	ur *gateways.UsersRepository,
	cr *gateways.ChannelsRepository,
	ysr *gateways.YoutubeSubscrptionsRepository,
) *UsersController {
	return &UsersController{ur, cr, ysr}
}

func (uc *UsersController) interactor(echoCtx echo.Context) *interactors.UsersInteractor {
	return interactors.NewUsersInteractor(
		presenters.NewUsersPresenter(echoCtx),
		presenters.NewChannelsPresenter(echoCtx),
		presenters.NewErrorsPresenter(echoCtx),
		uc.users,
		uc.channels,
		uc.youtubeSubscriptions,
	)
}

func (uc *UsersController) GetMyself() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		user, ok := echoCtx.Get("user").(*entities.User)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return uc.interactor(echoCtx).GetMyself(user)
	}
}

func (uc *UsersController) GetMySubscriptions() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		client, ok := echoCtx.Get("client").(*http.Client)
		if !ok {
			return echo.NewHTTPError(http.StatusInternalServerError)
		}

		return uc.interactor(echoCtx).GetMySubscriptions(
			echoCtx.Request().Context(),
			client,
		)
	}
}
