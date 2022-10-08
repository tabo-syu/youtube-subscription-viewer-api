package gateways

import (
	"context"
	"net/http"

	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type YoutubeSubscrptionsRepository struct {
	youtube interfaces.YoutubeHandler
}

var _ ports.YoutubeSubscriptionsRepository = (*YoutubeSubscrptionsRepository)(nil)

func NewYoutubeSubscriptionsRepository(y interfaces.YoutubeHandler) *YoutubeSubscrptionsRepository {
	return &YoutubeSubscrptionsRepository{y}
}

func (r *YoutubeSubscrptionsRepository) GetSubscriptions(
	ctx context.Context,
	client *http.Client,
) ([]*entities.Channel, error) {
	subscriptions, err := r.youtube.ListSubscriptions(ctx, client, []string{"id", "snippet"})
	if err != nil {
		return nil, err
	}

	var (
		channels      []*entities.Channel
		nextPageToken string
		loopErr       error
	)

	for {
		subscriptionsCall := subscriptions.Mine(true).Order("alphabetical").
			MaxResults(50).PageToken(nextPageToken)

		res, err := subscriptionsCall.Do()
		if err != nil {
			loopErr = err

			break
		}

		for _, subscription := range res.Items {
			url := "https://www.youtube.com/channel/" + subscription.Snippet.ResourceId.ChannelId

			channels = append(channels, &entities.Channel{
				Id:        subscription.Snippet.ResourceId.ChannelId,
				Name:      &subscription.Snippet.Title,
				Thumbnail: &subscription.Snippet.Thumbnails.High.Url,
				Url:       &url,
			})
		}

		nextPageToken = res.NextPageToken
		if nextPageToken == "" {
			break
		}
	}

	if loopErr != nil {
		return nil, loopErr
	}

	return channels, nil
}
