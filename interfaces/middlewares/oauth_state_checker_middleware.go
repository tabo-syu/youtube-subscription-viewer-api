package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/random"
)

type CheckerConfig struct {
	CookieName string
	CodeQuery  string
	StateQuery string
	StateKey   string
}

var (
	DefaultCheckerConfig = CheckerConfig{
		CookieName: "_state",
		CodeQuery:  "code",
		StateQuery: "state",
		StateKey:   "state",
	}
	ErrStateInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid state token")
)

func OAuthStateChecker(config CheckerConfig) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.QueryParam(config.CodeQuery) != "" {
				if !validateState(c, config, c.QueryParam(config.StateQuery)) {
					return ErrStateInvalid
				}
			} else {
				setCookie(c, config)
			}

			return next(c)
		}
	}
}

func validateState(c echo.Context, config CheckerConfig, state string) bool {
	cookie, err := c.Cookie(config.CookieName)
	// このエラーはクッキーが見つからないときに返却されるので、
	// 検証失敗として false を返す
	if err != nil {
		return false
	}

	return cookie.Value == state
}

func setCookie(c echo.Context, config CheckerConfig) {
	state := random.String(32)

	cookie := new(http.Cookie)
	cookie.Name = config.CookieName
	cookie.Value = state
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true

	c.SetCookie(cookie)
	c.Set(config.StateKey, state)
}
