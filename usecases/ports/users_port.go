package ports

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type UsersRepository interface {
	RegisterUser(*entities.User) (*entities.User, error)
	GetMyself() (*entities.User, error)
	GetMySubscriptions() ([]*entities.Video, error)
}

type UsersInputPort interface {
	GetMyself(echo.Context) error
	GetMySubscriptions(echo.Context) error
}

type UsersOutputPort interface {
	OutputUsers(echo.Context, []*entities.User) error
	OutputUser(echo.Context, *entities.User) error
}
