package interactors

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type UsersInteractor struct {
	usersOutput     ports.UsersOutputPort
	videosOutput    ports.VideosOutputPort
	errorsOutput    ports.ErrorsOutputPort
	usersRepository ports.UsersRepository
}

var _ ports.UsersInputPort = (*UsersInteractor)(nil)

func NewUsersInteractor(
	uo ports.UsersOutputPort,
	vo ports.VideosOutputPort,
	eo ports.ErrorsOutputPort,
	ur ports.UsersRepository,
) *UsersInteractor {
	return &UsersInteractor{uo, vo, eo, ur}
}

func (i *UsersInteractor) GetMyself(user *entities.User) error {
	return i.usersOutput.OutputUser(user)
}

func (i *UsersInteractor) GetMySubscriptions() error {
	videos, err := i.usersRepository.GetMySubscriptions()
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.videosOutput.OutputVideos(videos)
}
