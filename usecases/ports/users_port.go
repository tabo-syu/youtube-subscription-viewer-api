package ports

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"golang.org/x/oauth2"
)

type YoutubeSubscriptionsRepository interface {
	GetSubscriptions(context.Context, *http.Client) ([]*entities.Channel, error)
}

type UsersRepository interface {
	RegisterUser(*entities.User, *oauth2.Token) error
	Get(id string) (*entities.User, *oauth2.Token, error)
	GetMySubscriptions() ([]*entities.Channel, error)
}

type UsersInputPort interface {
	GetMyself(*entities.User) error
	GetMySubscriptions(context.Context, *http.Client) error
}

type UsersOutputPort interface {
	OutputUsers([]*entities.User) error
	OutputUser(*entities.User) error
}
