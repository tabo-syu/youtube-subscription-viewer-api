package presenters

import (
	"log"
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/middlewares"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type YoutubeAuthsPresenter struct {
	echoCtx echo.Context
}

var _ ports.YoutubeAuthsOutputPort = (*YoutubeAuthsPresenter)(nil)

func NewYoutubeAuthsPresenter(echoCtx echo.Context) *YoutubeAuthsPresenter {
	return &YoutubeAuthsPresenter{echoCtx}
}

func (p *YoutubeAuthsPresenter) OutputRedirectURL(url string) error {
	return p.echoCtx.Redirect(http.StatusSeeOther, url)
}

func (p *YoutubeAuthsPresenter) Login(user *entities.User) error {
	sess, _ := session.Get(middlewares.DefaultAuthenticatorConfig.CookieName, p.echoCtx)
	sess.Options = middlewares.DefaultAuthenticatorConfig.Session
	sess.Values["user_id"] = user.Id

	if err := sess.Save(p.echoCtx.Request(), p.echoCtx.Response()); err != nil {
		log.Println(err)
	}

	return p.echoCtx.JSON(http.StatusOK, user)
}
