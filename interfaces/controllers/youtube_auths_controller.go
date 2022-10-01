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

func (c *YoutubeAuthsController) interactor(ctx echo.Context) *interactors.YoutubeAuthsInteractor {
	return interactors.NewYoutubeAuthsInteractor(
		presenters.NewYoutubeAuthsPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.authorization,
		c.users,
		c.youtubeChannels,
	)
}

func (c *YoutubeAuthsController) Authorize(stateKey string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).Authorize(
			ctx.Get(stateKey).(string),
		)
	}
}

func (c *YoutubeAuthsController) Login() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		code := ctx.QueryParam("code")
		if code == "" {
			return echo.NewHTTPError(http.StatusBadRequest, "missing code token in the query string")
		}

		return c.interactor(ctx).Login(
			ctx.Request().Context(),
			code,
		)
	}
}

func (c *YoutubeAuthsController) Logout() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).Logout(ctx.Request().Context())
	}
}
