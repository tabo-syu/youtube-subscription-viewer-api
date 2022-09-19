package interactors

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type YoutubeAuthsInteractor struct {
	YoutubeAuthsOutputPort ports.YoutubeAuthsOutputPort
	errorsOutput           ports.ErrorsOutputPort
	YoutubeAuthorization   ports.YoutubeAuthorization
}

var _ ports.YoutubeAuthsInputPort = (*YoutubeAuthsInteractor)(nil)

func NewYoutubeAuthsInteractor(
	ao ports.YoutubeAuthsOutputPort,
	eo ports.ErrorsOutputPort,
	ga ports.YoutubeAuthorization,
) *YoutubeAuthsInteractor {
	return &YoutubeAuthsInteractor{ao, eo, ga}
}

func (i *YoutubeAuthsInteractor) Authorize(state string) error {
	url := i.YoutubeAuthorization.AuthCodeUrl(state)

	return i.YoutubeAuthsOutputPort.OutputRedirectUrl(url)
}
