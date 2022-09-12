package controllers

import "github.com/labstack/echo/v4"

type UsersController struct {
	sqlHandler     SqlHandler
	youtubeHandler YoutubeHandler
}

func NewUsersController(s SqlHandler, y YoutubeHandler) *UsersController {
	return &UsersController{s, y}
}

func (c *UsersController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (c *UsersController) Logout() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (c *UsersController) GetMyself() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}

func (c *UsersController) GetMySubscriptions() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
