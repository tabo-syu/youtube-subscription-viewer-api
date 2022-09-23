package interactors

import (
	"context"

	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
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

func (i *YoutubeAuthsInteractor) Login(ctx context.Context, code string) error {
	token, err := i.YoutubeAuthorization.Exchange(ctx, code)
	if err != nil {
		return err
	}

	client := i.YoutubeAuthorization.Client(ctx, token)
	youtube, err := youtube.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	res, err := youtube.Channels.List([]string{"id"}).Mine(true).Do()
	if err != nil {
		return err
	}

	return i.YoutubeAuthsOutputPort.Test(res.Items[0].Id)
}
