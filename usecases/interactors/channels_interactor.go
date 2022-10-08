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
	listsOutput ports.ListsOutputPort,
	channelsOutput ports.ChannelsOutputPort,
	videosOutput ports.VideosOutputPort,
	errorsOutput ports.ErrorsOutputPort,
	channelsRepository ports.ChannelsRepository,
) *ChannelsInteractor {
	return &ChannelsInteractor{listsOutput, channelsOutput, videosOutput, errorsOutput, channelsRepository}
}

func (i *ChannelsInteractor) GetFeed() error {
	videos, err := i.channelsRepository.GetFeed()
	if err != nil {
		return i.errorsOutput.OutputError()
	}

	return i.videosOutput.OutputVideos(videos)
}
