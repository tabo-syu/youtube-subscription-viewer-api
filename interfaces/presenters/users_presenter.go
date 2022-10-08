package presenters

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type UsersPresenter struct {
	echoCtx echo.Context
}

var _ ports.UsersOutputPort = (*UsersPresenter)(nil)

func NewUsersPresenter(echoCtx echo.Context) *UsersPresenter {
	return &UsersPresenter{echoCtx}
}

func (p *UsersPresenter) OutputUsers(users []*entities.User) error {
	return nil
}

func (p *UsersPresenter) OutputUser(user *entities.User) error {
	return p.echoCtx.JSON(http.StatusOK, user)
}
