package interactors

import (
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

func (i *ListsInteractor) Create(list *entities.List) error {
	list, err := i.listsRepository.Create(list)
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.listsOutput.OutputList(list)
}

func (i *ListsInteractor) GetAll() error {
	lists, err := i.listsRepository.GetAll()
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.listsOutput.OutputLists(lists)
}

func (i *ListsInteractor) GetById(id entities.ListId) error {
	list, err := i.listsRepository.GetById(id)
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.listsOutput.OutputList(list)
}

func (i *ListsInteractor) Update(list *entities.List) error {
	list, err := i.listsRepository.Update(list)
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.listsOutput.OutputList(list)
}

func (i *ListsInteractor) DeleteById(id entities.ListId) error {
	err := i.listsRepository.DeleteById(id)
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.listsOutput.OutputList(&entities.List{})
}

func (i *ListsInteractor) AddChannel() error {
	channel, err := i.listsRepository.AddChannel(&entities.Channel{})
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.channelsOutput.OutputChannel(channel)
}

func (i *ListsInteractor) GetAllChannels() error {
	channels, err := i.listsRepository.GetAllChannels()
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.channelsOutput.OutputChannels(channels)
}

func (i *ListsInteractor) GetFeed() error {
	videos, err := i.listsRepository.GetFeed()
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.videosOutput.OutputVideos(videos)
}
