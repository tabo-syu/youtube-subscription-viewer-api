package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/gateways"
	"github.com/tabo-syu/youtube-subscription-viewer-api/interfaces/presenters"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/interactors"
)

type ListsController struct {
	lists *gateways.ListsRepository
}

func NewListsController(l *gateways.ListsRepository) *ListsController {
	return &ListsController{l}
}

func (c *ListsController) interactor(ctx echo.Context) *interactors.ListsInteractor {
	return interactors.NewListsInteractor(
		presenters.NewListsPresenter(ctx),
		presenters.NewChannelsPresenter(ctx),
		presenters.NewVideosPresenter(ctx),
		presenters.NewErrorsPresenter(ctx),
		c.lists,
	)
}

func (c *ListsController) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var list entities.List
		if err := ctx.Bind(&list); err != nil {
			return err
		}

		return c.interactor(ctx).Create(&list)
	}
}

func (c *ListsController) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return c.interactor(ctx).GetAll()
	}
}

func (c *ListsController) GetById() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		return nil
	}
}

func (c *ListsController) Update() echo.HandlerFunc {
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
