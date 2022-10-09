package infrastructures

import (
	"fmt"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/controllers"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/middlewares"
)

type Controllers struct {
	Authorizations *controllers.YoutubeAuthsController
	Users          *controllers.UsersController
	Lists          *controllers.ListsController
	Channels       *controllers.ChannelsController
}

type Middlewares struct {
	OAuthStateChecker middlewares.OAuthStateCheckerFunc
	Authenticator     middlewares.AuthenticatorFunc
}

type Server struct {
	Controllers
	Middlewares
}

func NewServer(sql *SQLHandler, oauth2 *YoutubeOAuth2Handler, youtube *YoutubeHandler) (*Server, error) {
	usersRepository := gateways.NewUsersRepository(sql)
	youtubeAuthorization := gateways.NewYoutubeAuthorization(oauth2)
	youtubeChannelsRepository := gateways.NewYoutubeChannelsRepository(youtube)
	youtubeSubscriptionsRepository := gateways.NewYoutubeSubscriptionsRepository(youtube)
	channelsRepository := gateways.NewChannelsRepository(sql)
	listsRepository := gateways.NewListsRepository(sql)

	return &Server{
		Controllers{
			controllers.NewYoutubeAuthsController(usersRepository, youtubeAuthorization, youtubeChannelsRepository),
			controllers.NewUsersController(usersRepository, channelsRepository, youtubeSubscriptionsRepository),
			controllers.NewListsController(listsRepository),
			controllers.NewChannelsController(channelsRepository, youtubeChannelsRepository),
		},
		Middlewares{
			middlewares.OAuthStateChecker(middlewares.DefaultCheckerConfig),
			middlewares.Authenticator(usersRepository, youtubeAuthorization, middlewares.DefaultAuthenticatorConfig),
		},
	}, nil
}

func (s *Server) Start(port string) {
	echo := echo.New()
	echo.Use(
		middleware.Logger(),
		session.Middleware(
			sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_KEY"))),
		),
	)

	users := echo.Group("/users")
	{
		users.GET("/auth",
			s.Controllers.Authorizations.Authorize(middlewares.DefaultCheckerConfig.StateKey),
			s.Middlewares.OAuthStateChecker,
		)
		users.GET("/login", s.Controllers.Authorizations.Login(), s.Middlewares.OAuthStateChecker)

		users.GET("/logout", s.Controllers.Authorizations.Logout())
		users.GET("/me", s.Controllers.Users.GetMyself(), s.Middlewares.Authenticator)
		users.GET("/me/subscriptions", s.Controllers.Users.GetMySubscriptions(), s.Middlewares.Authenticator)
	}

	lists := echo.Group("/lists")
	{
		lists.POST("", s.Controllers.Lists.Create())
		lists.GET("", s.Controllers.Lists.GetAll())

		lists.GET("/:id", s.Controllers.Lists.GetByID())
		lists.PATCH("/:id", s.Controllers.Lists.Update())
		lists.DELETE("/:id", s.Controllers.Lists.DeleteByID())

		lists.POST("/:id/channels", s.Controllers.Lists.AddChannel())
		lists.GET("/:id/channels", s.Controllers.Lists.GetAllChannels())

		lists.GET("/:id/feed", s.Controllers.Lists.GetFeed())
	}

	channels := echo.Group("/channels")
	{
		channels.GET("/:id/feed", s.Controllers.Channels.GetFeed())
	}

	echo.Logger.Fatal(echo.Start(fmt.Sprintf(":%s", port)))
}
