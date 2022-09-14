package ports

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type VideosOutputPort interface {
	OutputVideos(echo.Context, []*entities.Video) error
	OutputVideo(echo.Context, *entities.Video) error
}
