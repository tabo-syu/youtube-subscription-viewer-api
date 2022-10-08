package interactors

import (
	"context"

	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type YoutubeAuthsInteractor struct {
	youtubeAuthsOutputPort    ports.YoutubeAuthsOutputPort
	errorsOutput              ports.ErrorsOutputPort
	youtubeAuthorization      ports.YoutubeAuthorization
	youtubeChannelsRepository ports.YoutubeChannelsRepository
	usersRepository           ports.UsersRepository
}

var _ ports.YoutubeAuthsInputPort = (*YoutubeAuthsInteractor)(nil)

func NewYoutubeAuthsInteractor(
	ao ports.YoutubeAuthsOutputPort,
	eo ports.ErrorsOutputPort,
	ya ports.YoutubeAuthorization,
	ur ports.UsersRepository,
	yr ports.YoutubeChannelsRepository,
) *YoutubeAuthsInteractor {
	return &YoutubeAuthsInteractor{ao, eo, ya, yr, ur}
}

func (i *YoutubeAuthsInteractor) Authorize(state string) error {
	url := i.youtubeAuthorization.AuthCodeURL(state)

	return i.youtubeAuthsOutputPort.OutputRedirectURL(url)
}

func (i *YoutubeAuthsInteractor) Login(ctx context.Context, code string) error {
	token, err := i.youtubeAuthorization.Exchange(ctx, code)
	if err != nil {
		return err
	}

	client := i.youtubeAuthorization.Client(ctx, token)
	user, err := i.youtubeChannelsRepository.GetMyChannel(ctx, client)
	if err != nil {
		return err
	}

	if err := i.usersRepository.RegisterUser(user, token); err != nil {
		return err
	}

	return i.youtubeAuthsOutputPort.Login(user)
}

func (i *YoutubeAuthsInteractor) Logout(ctx context.Context) error {
	return nil
}
