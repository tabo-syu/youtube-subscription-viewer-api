package presenters

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
	"golang.org/x/oauth2"
)

type YoutubeAuthsPresenter struct {
	ctx echo.Context
}

var _ ports.YoutubeAuthsOutputPort = (*YoutubeAuthsPresenter)(nil)

func NewYoutubeAuthsPresenter(ctx echo.Context) *YoutubeAuthsPresenter {
	return &YoutubeAuthsPresenter{ctx}
}

func (p *YoutubeAuthsPresenter) OutputRedirectUrl(url string) error {
	return p.ctx.Redirect(http.StatusSeeOther, url)
}

func (p *YoutubeAuthsPresenter) Test(token *oauth2.Token) error {
	return p.ctx.String(http.StatusOK, token.Expiry.Format(time.RFC3339))
}
