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

func (r *UsersRepository) RegisterUser(user *entities.User, token *oauth2.Token) error {
	_, err := r.sql.Exec(`
		INSERT INTO users (id, name, thumbnail, access_token, refresh_token, expiry)
		VALUES ($1, $2, $3, $4, $5, $6)`,
		user.Id, user.Name, user.Thumbnail, token.AccessToken, token.RefreshToken, token.Expiry,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) GetMyself() (*entities.User, error) {
	return &entities.User{}, nil
}

func (r *UsersRepository) GetMySubscriptions() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}
