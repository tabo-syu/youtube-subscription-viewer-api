package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type VideosPresenter struct {
	echoCtx echo.Context
}

var _ ports.VideosOutputPort = (*VideosPresenter)(nil)

func NewVideosPresenter(echoCtx echo.Context) *VideosPresenter {
	return &VideosPresenter{echoCtx}
}

func (p *VideosPresenter) OutputVideos(videos []*entities.Video) error {
	return nil
}

func (p *VideosPresenter) OutputVideo(video *entities.Video) error {
	return nil
}
