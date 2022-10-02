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
	users *gateways.UsersRepository
}

func NewUsersController(u *gateways.UsersRepository) *UsersController {
	return &UsersController{u}
}

func (c *UsersController) interactor(ctx echo.Context) *interactors.UsersInteractor {
	return interactors.NewUsersInteractor(
		presenters.NewUsersPresenter(ctx),
		presenters.NewVideosPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.users,
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
		return c.interactor(ctx).GetMySubscriptions()
	}
}
