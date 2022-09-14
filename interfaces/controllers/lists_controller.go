package controllers

import "github.com/labstack/echo/v4"

type ListsController struct {
	sqlHandler     SqlHandler
	youtubeHandler YoutubeHandler
}

func NewListsController(s SqlHandler, y YoutubeHandler) *ListsController {
	return &ListsController{s, y}
}

func (c *ListsController) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) GetById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) UpdateById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) DeleteById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) AddChannel() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) GetAllChannels() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) GetFeed() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}
