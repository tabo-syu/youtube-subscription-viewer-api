package gateways

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
	"golang.org/x/oauth2"
)

type UsersRepository struct {
	sql interfaces.SqlHandler
}

var _ ports.UsersRepository = (*UsersRepository)(nil)

func NewUsersRepository(s interfaces.SqlHandler) *UsersRepository {
	return &UsersRepository{s}
}

func (r *UsersRepository) RegisterUser(user *entities.User, token *oauth2.Token) (*entities.User, error) {
	return &entities.User{}, nil
}

func (r *UsersRepository) GetMyself() (*entities.User, error) {
	return &entities.User{}, nil
}

func (r *UsersRepository) GetMySubscriptions() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}
