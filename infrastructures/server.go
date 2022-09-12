package infrastructures

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/controllers"
)

type Server struct {
	Users    *controllers.UsersController
	Lists    *controllers.ListsController
	Channels *controllers.ChannelsController
}

func NewServer(config *Config) *Server {
	s := NewSqlHandler(config)
	y := NewYoutubeHandler(config)

	return &Server{
		controllers.NewUsersController(s, y),
		controllers.NewListsController(s, y),
		controllers.NewChannelsController(s, y),
	}
}

func (s *Server) Start(port string) {
	e := echo.New()
	e.Use(middleware.Logger())

	users := e.Group("/users")
	{
		users.GET("/login", s.Users.Login())
		users.GET("/logout", s.Users.Logout())

		users.GET("/me", s.Users.GetMyself())
		users.GET("/me/subscriptions", s.Users.GetMySubscriptions())
	}

	lists := e.Group("/lists")
	{
		lists.POST("", s.Lists.Create())
		lists.GET("", s.Lists.GetAll())

		lists.GET("/:id", s.Lists.GetById())
		lists.PATCH("/:id", s.Lists.UpdateById())
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
