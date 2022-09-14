package gateways

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type UsersRepository struct {
	sql     interfaces.SqlHandler
	youtube interfaces.YoutubeHandler
}

var _ ports.UsersRepository = (*UsersRepository)(nil)

func NewUsersRepository(s interfaces.SqlHandler, y interfaces.YoutubeHandler) *UsersRepository {
	return &UsersRepository{s, y}
}

func (r *UsersRepository) RegisterUser(user *entities.User) (*entities.User, error) {
	return &entities.User{}, nil
}

func (r *UsersRepository) GetMyself() (*entities.User, error) {
	return &entities.User{}, nil
}

func (r *UsersRepository) GetMySubscriptions() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}
