package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type VideosPresenter struct{}

var _ ports.VideosOutputPort = (*VideosPresenter)(nil)

func NewVideosPresenter() *VideosPresenter {
	return &VideosPresenter{}
}

func (p *VideosPresenter) OutputVideos(ctx echo.Context, videos []*entities.Video) error {
	return nil
}

func (p *VideosPresenter) OutputVideo(ctx echo.Context, video *entities.Video) error {
	return nil
}
