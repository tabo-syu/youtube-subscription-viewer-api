package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Serve(port string) {
	e := echo.New()
	e.Use(middleware.Logger())

	users := e.Group("/users")
	{
		users.GET("/me", hoge)
		users.GET("/me/subscriptions", hoge)
	}

	user := e.Group("/user")
	{
		user.GET("login", hoge)
		user.GET("logout", hoge)
	}

	channelLists := e.Group("/channelLists")
	{
		channelLists.POST("", hoge)
		channelLists.GET("", hoge)

		channelLists.GET("/:id", hoge)
		channelLists.PATCH("/:id", hoge)
		channelLists.DELETE("/:id", hoge)

		channelLists.POST("/:id/channels", hoge)
		channelLists.GET("/:id/channels", hoge)

		channelLists.GET("/:id/feed", hoge)
	}

	channels := e.Group("/channels")
	{
		channels.GET("/:id/feed", hoge)
	}

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

func hoge(c echo.Context) error {
	return c.String(http.StatusOK, "hoge\n")
}
