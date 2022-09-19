package ports

import (
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type VideosOutputPort interface {
	OutputVideos([]*entities.Video) error
	OutputVideo(*entities.Video) error
}
