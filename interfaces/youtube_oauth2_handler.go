package interfaces

import (
	"context"

	"golang.org/x/oauth2"
)

type YoutubeOAuth2Handler interface {
	AuthCodeUrl(string) string
	Exchange(context.Context, string) (*oauth2.Token, error)
}
