package gateways

import (
	"log"
	"time"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
	"golang.org/x/oauth2"
)

type UsersRepository struct {
	sql interfaces.SQLHandler
}

var _ ports.UsersRepository = (*UsersRepository)(nil)

func NewUsersRepository(s interfaces.SQLHandler) *UsersRepository {
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

	log.Println("user is registered", "user_id:", user.Id)

	return nil
}

func (r *UsersRepository) Get(userID string) (*entities.User, *oauth2.Token, error) {
	var (
		user  = entities.User{}
		token = oauth2.Token{
			TokenType:    "bearer",
			AccessToken:  "",
			RefreshToken: "",
			Expiry:       time.Time{},
		}
	)

	err := r.sql.QueryRow(`
		SELECT id, name, thumbnail, access_token, refresh_token, expiry FROM users WHERE id = $1`,
		userID,
	).Scan(&user.Id, &user.Name, &user.Thumbnail, &token.AccessToken, &token.RefreshToken, &token.Expiry)
	if err != nil {
		return nil, nil, err
	}

	return &user, &token, nil
}

func (r *UsersRepository) UpdateToken(userID string, token *oauth2.Token) error {
	_, err := r.sql.Exec(
		`UPDATE users SET access_token = $2, refresh_token = $3, expiry = $4, updated_at = $5 WHERE id = $1`,
		userID, token.AccessToken, token.RefreshToken, token.Expiry, time.Now(),
	)
	if err != nil {
		return err
	}

	log.Println("token is updated", "user_id:", userID)

	return nil
}

func (r *UsersRepository) GetMySubscriptions() ([]*entities.Channel, error) {
	return []*entities.Channel{}, nil
}
