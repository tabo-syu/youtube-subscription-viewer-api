package ports

import "github.com/tabo-syu/youtube-subscription-viewer-api/entities"

type UsersRepository interface {
	RegisterUser(*entities.User) (*entities.User, error)
	GetMyself() (*entities.User, error)
	GetMySubscriptions() ([]*entities.Video, error)
}

type UsersInputPort interface {
	GetMyself()
}

type UsersOutputPort interface{}
