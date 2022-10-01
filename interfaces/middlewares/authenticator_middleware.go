package middlewares

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
)

type AuthenticatorConfig struct {
	CookieName string
	Session    *sessions.Options
}

var (
	DefaultAuthenticatorConfig = AuthenticatorConfig{
		CookieName: "_sess",
		Session: &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
	}
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "user is unauthorized")
)

type AuthenticatorFunc = echo.MiddlewareFunc

func Authenticator(users *gateways.UsersRepository, config AuthenticatorConfig) AuthenticatorFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get(config.CookieName, c)
			sess.Options = config.Session
			userId, ok := sess.Values["user_id"]
			if !ok {
				return ErrUnauthorized
			}
			user, token, err := users.Get(userId.(string))
			if err != nil {
				return ErrUnauthorized
			}

			c.Set("user", user)
			c.Set("token", token)

			return next(c)
		}
	}
}
