package interactors

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ListsInteractor struct {
	listsOutput     ports.ListsOutputPort
	channelsOutput  ports.ChannelsOutputPort
	videosOutput    ports.VideosOutputPort
	errorsOutput    ports.ErrorsOutputPort
	listsRepository ports.ListsRepository
}

var _ ports.ListsInputPort = (*ListsInteractor)(nil)

func NewListsInteractor(
	lo ports.ListsOutputPort,
	co ports.ChannelsOutputPort,
	vo ports.VideosOutputPort,
	eo ports.ErrorsOutputPort,
	lr ports.ListsRepository,
) *ListsInteractor {
	return &ListsInteractor{lo, co, vo, eo, lr}
}

func (i *ListsInteractor) Create(ctx echo.Context, list *entities.List) error {
	list, err := i.listsRepository.Create(list)
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.listsOutput.OutputList(ctx, list)
}

func (i *ListsInteractor) GetAll(ctx echo.Context) error {
	lists, err := i.listsRepository.GetAll()
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.listsOutput.OutputLists(ctx, lists)
}

func (i *ListsInteractor) GetById(ctx echo.Context, id entities.ListId) error {
	list, err := i.listsRepository.GetById(id)
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.listsOutput.OutputList(ctx, list)
}

func (i *ListsInteractor) Update(ctx echo.Context, list *entities.List) error {
	list, err := i.listsRepository.Update(list)
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.listsOutput.OutputList(ctx, list)
}

func (i *ListsInteractor) DeleteById(ctx echo.Context, id entities.ListId) error {
	err := i.listsRepository.DeleteById(id)
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.listsOutput.OutputList(ctx, &entities.List{})
}

func (i *ListsInteractor) AddChannel(ctx echo.Context) error {
	channel, err := i.listsRepository.AddChannel(&entities.Channel{})
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.channelsOutput.OutputChannel(ctx, channel)
}

func (i *ListsInteractor) GetAllChannels(ctx echo.Context) error {
	channels, err := i.listsRepository.GetAllChannels()
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.channelsOutput.OutputChannels(ctx, channels)
}

func (i *ListsInteractor) GetFeed(ctx echo.Context) error {
	videos, err := i.listsRepository.GetFeed()
	if err != nil {
		return i.errorsOutput.OutputError(ctx)
	}

	return i.videosOutput.OutputVideos(ctx, videos)
}
