package gateways

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type YoutubeChannelsRepository struct {
	youtube interfaces.YoutubeHandler
}

var _ ports.YoutubeChannelsRepository = (*YoutubeChannelsRepository)(nil)

func NewYoutubeChannelsRepository(y interfaces.YoutubeHandler) *YoutubeChannelsRepository {
	return &YoutubeChannelsRepository{y}
}

func (r *YoutubeChannelsRepository) GetMyChannel(ctx context.Context, client *http.Client) (*entities.User, error) {
	channels, err := r.youtube.ListChannels(ctx, client, []string{"id", "snippet"})
	if err != nil {
		return nil, err
	}

	res, err := channels.Mine(true).MaxResults(1).Do()
	if err != nil {
		return nil, err
	}

	user := entities.User{
		Id:        res.Items[0].Id,
		Name:      res.Items[0].Snippet.Title,
		Thumbnail: res.Items[0].Snippet.Thumbnails.High.Url,
	}

	return &user, nil
}
