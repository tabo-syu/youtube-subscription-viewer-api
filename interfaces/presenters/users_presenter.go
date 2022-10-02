package presenters

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type UsersPresenter struct {
	ctx echo.Context
}

var _ ports.UsersOutputPort = (*UsersPresenter)(nil)

func NewUsersPresenter(ctx echo.Context) *UsersPresenter {
	return &UsersPresenter{ctx}
}

func (p *UsersPresenter) OutputUsers(users []*entities.User) error {
	return nil
}

func (p *UsersPresenter) OutputUser(user *entities.User) error {
	return p.ctx.JSON(http.StatusOK, user)
}
