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

func (r *UsersRepository) Get(id string) (*entities.User, *oauth2.Token, error) {
	var (
		user  = entities.User{}
		token = oauth2.Token{
			TokenType: "bearer",
		}
	)
	err := r.sql.QueryRow(`
		SELECT id, name, thumbnail, access_token, refresh_token, expiry FROM users WHERE id = $1`,
		id,
	).Scan(&user.Id, &user.Name, &user.Thumbnail, &token.AccessToken, &token.RefreshToken, &token.Expiry)
	if err != nil {
		return nil, nil, err
	}

	return &user, &token, nil
}

func (r *UsersRepository) GetMySubscriptions() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}
