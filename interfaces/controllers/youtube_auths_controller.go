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

func (yc *YoutubeAuthsController) interactor(c echo.Context) *interactors.YoutubeAuthsInteractor {
	return interactors.NewYoutubeAuthsInteractor(
		presenters.NewYoutubeAuthsPresenter(c),
		presenters.NewErrorsPresenter(c),
		yc.authorization,
		yc.users,
		yc.youtubeChannels,
	)
}

func (yc *YoutubeAuthsController) Authorize(stateKey string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return yc.interactor(c).Authorize(
			c.Get(stateKey).(string),
		)
	}
}

func (yc *YoutubeAuthsController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		code := c.QueryParam("code")
		if code == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing code token in the query string")
		}

		return yc.interactor(c).Login(
			c.Request().Context(),
			code,
		)
	}
}

func (yc *YoutubeAuthsController) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return yc.interactor(c).Logout(c.Request().Context())
	}
}
