package middlewares

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
)

type AuthenticatorConfig struct {
	SessionName string
}

var DefaultAuthenticatorConfig = AuthenticatorConfig{
	SessionName: "_sess",
}

type AuthenticatorFunc = echo.MiddlewareFunc

func Authenticator(users *gateways.UsersRepository, config AuthenticatorConfig) AuthenticatorFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get(config.SessionName, c)
			sess.Options = &sessions.Options{
				Path:   "/",
				MaxAge: 86400 * 7,
			}

			// if userId, ok := sess.Values["user_id"]; !ok {
			// 	return
			// }

			return next(c)
		}
	}
}
