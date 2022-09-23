package infrastructures

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/controllers"
)

type Server struct {
	Authorizations *controllers.YoutubeAuthsController
	Users          *controllers.UsersController
	Lists          *controllers.ListsController
	Channels       *controllers.ChannelsController
}

func NewServer(db *SqlHandler, oauth2 *YoutubeOAuth2Handler, youtube *YoutubeHandler) (*Server, error) {
	return &Server{
		controllers.NewYoutubeAuthsController(db, oauth2, youtube),
		controllers.NewUsersController(db, youtube),
		controllers.NewListsController(db, youtube),
		controllers.NewChannelsController(db, youtube),
	}, nil
}

func (s *Server) Start(port string) {
	e := echo.New()
	e.Use(middleware.Logger())

	users := e.Group("/users")
	{
		authCSRF := middleware.CSRFWithConfig(
			middleware.CSRFConfig{
				TokenLookup:    "query:state",
				CookieHTTPOnly: true,
				CookieSameSite: http.SameSiteLaxMode,
			},
		)

		users.GET("/auth", s.Authorizations.Authorize(middleware.DefaultCSRFConfig.ContextKey), authCSRF)
		users.GET("/login", s.Authorizations.Login(), authCSRF)
		users.GET("/logout", s.Users.Logout())

		users.GET("/me", s.Users.GetMyself())
		users.GET("/me/subscriptions", s.Users.GetMySubscriptions())
	}

	lists := e.Group("/lists")
	{
		lists.POST("", s.Lists.Create())
		lists.GET("", s.Lists.GetAll())

		lists.GET("/:id", s.Lists.GetById())
		lists.PATCH("/:id", s.Lists.Update())
		lists.DELETE("/:id", s.Lists.DeleteById())

		lists.POST("/:id/channels", s.Lists.AddChannel())
		lists.GET("/:id/channels", s.Lists.GetAllChannels())

		lists.GET("/:id/feed", s.Lists.GetFeed())
	}

	channels := e.Group("/channels")
	{
		channels.GET("/:id/feed", s.Channels.GetFeed())
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}
