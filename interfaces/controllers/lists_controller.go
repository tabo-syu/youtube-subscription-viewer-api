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

func (lc *ListsController) interactor(echoCtx echo.Context) *interactors.ListsInteractor {
	return interactors.NewListsInteractor(
		presenters.NewListsPresenter(echoCtx),
		presenters.NewChannelsPresenter(echoCtx),
		presenters.NewVideosPresenter(echoCtx),
		presenters.NewErrorsPresenter(echoCtx),
		lc.lists,
	)
}

func (lc *ListsController) Create() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		var list entities.List
		if err := echoCtx.Bind(&list); err != nil {
			return err
		}

		return lc.interactor(echoCtx).Create(&list)
	}
}

func (lc *ListsController) GetAll() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return lc.interactor(echoCtx).GetAll()
	}
}

func (lc *ListsController) GetByID() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return nil
	}
}

func (lc *ListsController) Update() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return nil
	}
}

func (lc *ListsController) DeleteByID() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return nil
	}
}

func (lc *ListsController) AddChannel() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return nil
	}
}

func (c *ListsController) GetAllChannels() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return nil
	}
}

func (c *ListsController) GetFeed() echo.HandlerFunc {
	return func(echoCtx echo.Context) error {
		return nil
	}
}
