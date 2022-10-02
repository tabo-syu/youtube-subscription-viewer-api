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
	c echo.Context
}

var _ ports.YoutubeAuthsOutputPort = (*YoutubeAuthsPresenter)(nil)

func NewYoutubeAuthsPresenter(c echo.Context) *YoutubeAuthsPresenter {
	return &YoutubeAuthsPresenter{c}
}

func (p *YoutubeAuthsPresenter) OutputRedirectUrl(url string) error {
	return p.c.Redirect(http.StatusSeeOther, url)
}

func (p *YoutubeAuthsPresenter) Login(user *entities.User) error {
	sess, _ := session.Get(middlewares.DefaultAuthenticatorConfig.CookieName, p.c)
	sess.Options = middlewares.DefaultAuthenticatorConfig.Session
	sess.Values["user_id"] = user.Id
	sess.Save(p.c.Request(), p.c.Response())

	return p.c.JSON(http.StatusOK, user)
}
