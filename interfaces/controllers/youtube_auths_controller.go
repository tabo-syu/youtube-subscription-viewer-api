package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type YoutubeAuthsController struct {
	users           *gateways.UsersRepository
	authorization   *gateways.YoutubeAuthorization
	youtubeChannels *gateways.YoutubeChannelsRepository
}

func NewYoutubeAuthsController(
	ur *gateways.UsersRepository,
	ya *gateways.YoutubeAuthorization,
	ycr *gateways.YoutubeChannelsRepository,
) *YoutubeAuthsController {
	return &YoutubeAuthsController{ur, ya, ycr}
}

func (yc *YoutubeAuthsController) interactor(echoCtx echo.Context) *interactors.YoutubeAuthsInteractor {
	return interactors.NewYoutubeAuthsInteractor(
		presenters.NewYoutubeAuthsPresenter(echoCtx),
		presenters.NewErrorsPresenter(echoCtx),
		yc.authorization,
		yc.users,
		yc.youtubeChannels,
	)
}

func (yc *YoutubeAuthsController) Authorize(stateKey string) echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return yc.interactor(echoCtx).Authorize(
			echoCtx.Get(stateKey).(string),
		)
	}
}

func (yc *YoutubeAuthsController) Login() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		code := echoCtx.QueryParam("code")
		if code == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing code token in the query string")
		}

		return yc.interactor(echoCtx).Login(
			echoCtx.Request().Context(),
			code,
		)
	}
}

func (yc *YoutubeAuthsController) Logout() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return yc.interactor(echoCtx).Logout(echoCtx.Request().Context())
	}
}
