package controllers

import (
	"github.com/labstack/echo/v4"
)

type ChannelsController struct {
	sqlHandler     SqlHandler
	youtubeHandler YoutubeHandler
}

func NewChannelsController(s SqlHandler, y YoutubeHandler) *ChannelsController {
	return &ChannelsController{s, y}
}

func (c *ChannelsController) GetFeed() echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
