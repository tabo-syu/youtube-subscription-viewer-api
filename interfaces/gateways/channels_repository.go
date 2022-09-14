package gateways

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ChannelsRepository struct {
	sql     interfaces.SqlHandler
	youtube interfaces.YoutubeHandler
}

var _ ports.ChannelsRepository = (*ChannelsRepository)(nil)

func NewChannelsRepository(s interfaces.SqlHandler, y interfaces.YoutubeHandler) *ChannelsRepository {
	return &ChannelsRepository{s, y}
}

func (r *ChannelsRepository) GetFeed() ([]*entities.Video, error) {
	return []*entities.Video{}, nil
}
