package interactors

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type UsersInteractor struct {
	usersOutput                    ports.UsersOutputPort
	channelsOutput                 ports.ChannelsOutputPort
	errorsOutput                   ports.ErrorsOutputPort
	usersRepository                ports.UsersRepository
	youtubeSubscriptionsRepository ports.YoutubeSubscriptionsRepository
}

var _ ports.UsersInputPort = (*UsersInteractor)(nil)

func NewUsersInteractor(
	uo ports.UsersOutputPort,
	co ports.ChannelsOutputPort,
	eo ports.ErrorsOutputPort,
	ur ports.UsersRepository,
	ysr ports.YoutubeSubscriptionsRepository,
) *UsersInteractor {
	return &UsersInteractor{uo, co, eo, ur, ysr}
}

func (i *UsersInteractor) GetMyself(user *entities.User) error {
	return i.usersOutput.OutputUser(user)
}

func (i *UsersInteractor) GetMySubscriptions(ctx context.Context, client *http.Client) error {
	channels, err := i.youtubeSubscriptionsRepository.GetSubscriptions(ctx, client)
	if err != nil {
		return nil
	}

	return i.channelsOutput.OutputChannels(channels)
}
