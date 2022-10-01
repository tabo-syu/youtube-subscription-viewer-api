package presenters

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/middlewares"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
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

func (p *YoutubeAuthsPresenter) Login(user *entities.User) error {
	sess, _ := session.Get(middlewares.DefaultAuthenticatorConfig.CookieName, p.ctx)
	sess.Options = middlewares.DefaultAuthenticatorConfig.Session
	sess.Values["user_id"] = user.Id
	sess.Save(p.ctx.Request(), p.ctx.Response())

	return p.ctx.JSON(http.StatusOK, user)
}
