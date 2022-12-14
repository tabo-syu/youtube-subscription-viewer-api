package middlewares

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"golang.org/x/oauth2"
)

type AuthenticatorConfig struct {
	CookieName string
	Session    *sessions.Options
}

var (
	oneWeek                    = 60 * 60 * 24 * 7
	DefaultAuthenticatorConfig = AuthenticatorConfig{
		CookieName: "_sess",
		Session: &sessions.Options{
			Domain:   "localhost",
			Secure:   false,
			Path:     "/",
			MaxAge:   oneWeek,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		},
	}
	ErrUnauthorized = echo.NewHTTPError(http.StatusUnauthorized, "user is unauthorized")
)

type userTokenSource struct {
	src    oauth2.TokenSource
	users  *gateways.UsersRepository
	userID string
}

func newUserTokenSouce(src oauth2.TokenSource, users *gateways.UsersRepository, userID string) *userTokenSource {
	return &userTokenSource{src, users, userID}
}

// リフレッシュトークンを使い、アクセストークンを更新する。
// アクセストークンが更新された場合には、DB へキャッシュする。
func (s *userTokenSource) Token() (*oauth2.Token, error) {
	token, err := s.src.Token()
	if err != nil {
		return nil, err
	}

	return token, s.users.UpdateToken(s.userID, token)
}

type AuthenticatorFunc = echo.MiddlewareFunc

func Authenticator(
	users *gateways.UsersRepository,
	auth *gateways.YoutubeAuthorization,
	config AuthenticatorConfig,
) AuthenticatorFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(echoCtx echo.Context) error {
			ctx := echoCtx.Request().Context()

			sess, _ := session.Get(config.CookieName, echoCtx)
			sess.Options = config.Session

			rawUserID, success := sess.Values["user_id"]
			if !success {
				return ErrUnauthorized
			}

			userID, success := rawUserID.(string)
			if !success {
				return ErrUnauthorized
			}

			user, token, err := users.Get(userID)
			if err != nil {
				return ErrUnauthorized
			}

			// トークンの更新処理
			tokenSource := auth.TokenSource(ctx, token)
			userTokenSource := newUserTokenSouce(tokenSource, users, userID)
			client := oauth2.NewClient(ctx, oauth2.ReuseTokenSource(token, userTokenSource))

			echoCtx.Set("user", user)
			echoCtx.Set("client", client)

			return next(echoCtx)
		}
	}
}
