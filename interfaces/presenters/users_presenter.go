package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type UsersPresenter struct{}

var _ ports.UsersOutputPort = (*UsersPresenter)(nil)

func NewUsersPresenter() *UsersPresenter {
	return &UsersPresenter{}
}

func (p *UsersPresenter) OutputUsers(ctx echo.Context, users []*entities.User) error {
	return nil
}

func (p *UsersPresenter) OutputUser(ctx echo.Context, user *entities.User) error {
	return nil
}
