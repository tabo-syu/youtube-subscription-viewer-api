package ports

import (
	"context"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type YoutubeAuthorization interface {
	AuthCodeUrl(string) string
	Exchange(context.Context, string) (*entities.User, error)
}

type YoutubeAuthsInputPort interface {
	Authorize(string) error
}

type YoutubeAuthsOutputPort interface {
	OutputRedirectUrl(string) error
}
