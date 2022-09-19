package interactors

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ChannelsInteractor struct {
	listsOutput        ports.ListsOutputPort
	channelsOutput     ports.ChannelsOutputPort
	videosOutput       ports.VideosOutputPort
	errorsOutput       ports.ErrorsOutputPort
	channelsRepository ports.ChannelsRepository
}

var _ ports.ChannelsInputPort = (*ChannelsInteractor)(nil)

func NewChannelsInteractor(
	lo ports.ListsOutputPort,
	co ports.ChannelsOutputPort,
	vo ports.VideosOutputPort,
	eo ports.ErrorsOutputPort,
	cr ports.ChannelsRepository,
) *ChannelsInteractor {
	return &ChannelsInteractor{lo, co, vo, eo, cr}
}

func (i *ChannelsInteractor) GetFeed() error {
	videos, err := i.channelsRepository.GetFeed()
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.videosOutput.OutputVideos(videos)
}
