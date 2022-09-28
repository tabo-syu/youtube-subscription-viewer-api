package ports

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"golang.org/x/oauth2"
)

type UsersRepository interface {
	RegisterUser(*entities.User, *oauth2.Token) (*entities.User, error)
	GetMyself() (*entities.User, error)
	GetMySubscriptions() ([]*entities.Video, error)
}

type UsersInputPort interface {
	GetMyself() error
	GetMySubscriptions() error
}

type UsersOutputPort interface {
	OutputUsers([]*entities.User) error
	OutputUser(*entities.User) error
}
