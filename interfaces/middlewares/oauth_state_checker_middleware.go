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

type OAuthStateCheckerFunc = echo.MiddlewareFunc

func OAuthStateChecker(config CheckerConfig) OAuthStateCheckerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			if echoCtx.QueryParam(config.CodeQuery) != "" {
				if !validateState(echoCtx, config, echoCtx.QueryParam(config.StateQuery)) {
					return ErrStateInvalid
				}
			} else {
				setCookie(echoCtx, config)
			}

			return next(echoCtx)
		}
	}
}

func validateState(echoCtx echo.Context, config CheckerConfig, state string) bool {
	cookie, err := echoCtx.Cookie(config.CookieName)
	// このエラーはクッキーが見つからないときに返却されるので、
	// 検証失敗として false を返す
	if err != nil {
		return false
	}

	return cookie.Value == state
}

func setCookie(echoCtx echo.Context, config CheckerConfig) {
	var length uint8 = 32
	state := random.String(length)

	cookie := new(http.Cookie)
	cookie.Name = config.CookieName
	cookie.Value = state
	cookie.SameSite = http.SameSiteLaxMode
	cookie.HttpOnly = true

	echoCtx.SetCookie(cookie)
	echoCtx.Set(config.StateKey, state)
}
