package interactors

import (
	"github.com/labstack/echo/v4"
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

func (i *UsersInteractor) GetMyself(ctx echo.Context) error {
	me, err := i.usersRepository.GetMyself()
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.usersOutput.OutputUser(ctx, me)
}

func (i *UsersInteractor) GetMySubscriptions(ctx echo.Context) error {
	videos, err := i.usersRepository.GetMySubscriptions()
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.videosOutput.OutputVideos(ctx, videos)
}
