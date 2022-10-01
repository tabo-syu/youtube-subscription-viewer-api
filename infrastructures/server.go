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

func NewServer(db *SqlHandler, oauth2 *YoutubeOAuth2Handler, youtube *YoutubeHandler) (*Server, error) {
	ur := gateways.NewUsersRepository(db)
	ya := gateways.NewYoutubeAuthorization(oauth2)
	ycr := gateways.NewYoutubeChannelsRepository(youtube)
	cr := gateways.NewChannelsRepository(db)
	lr := gateways.NewListsRepository(db)

	return &Server{
		Controllers{
			controllers.NewYoutubeAuthsController(ur, ya, ycr),
			controllers.NewUsersController(ur),
			controllers.NewListsController(lr),
			controllers.NewChannelsController(cr, ycr),
		},
		Middlewares{
			middlewares.OAuthStateChecker(middlewares.DefaultCheckerConfig),
			middlewares.Authenticator(ur, middlewares.DefaultAuthenticatorConfig),
		},
	}, nil
}

func (s *Server) Start(port string) {
	e := echo.New()
	e.Use(
		middleware.Logger(),
		session.Middleware(
			sessions.NewFilesystemStore("", []byte(os.Getenv("SESSION_KEY"))),
		),
		s.Middlewares.Authenticator,
	)

	users := e.Group("/users")
	{
		users.GET("/auth",
			s.Controllers.Authorizations.Authorize(middlewares.DefaultCheckerConfig.StateKey),
			s.Middlewares.OAuthStateChecker,
		)
		users.GET("/login", s.Controllers.Authorizations.Login(), s.Middlewares.OAuthStateChecker)

		users.GET("/logout", s.Controllers.Authorizations.Logout())
		users.GET("/me", s.Controllers.Users.GetMyself())
		users.GET("/me/subscriptions", s.Controllers.Users.GetMySubscriptions())
	}

	lists := e.Group("/lists")
	{
		lists.POST("", s.Controllers.Lists.Create())
		lists.GET("", s.Controllers.Lists.GetAll())

		lists.GET("/:id", s.Controllers.Lists.GetById())
		lists.PATCH("/:id", s.Controllers.Lists.Update())
		lists.DELETE("/:id", s.Controllers.Lists.DeleteById())

		lists.POST("/:id/channels", s.Controllers.Lists.AddChannel())
		lists.GET("/:id/channels", s.Controllers.Lists.GetAllChannels())

		lists.GET("/:id/feed", s.Controllers.Lists.GetFeed())
	}

	channels := e.Group("/channels")
	{
		channels.GET("/:id/feed", s.Controllers.Channels.GetFeed())
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
