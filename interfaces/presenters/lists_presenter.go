package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ListsPresenter struct {
	ctx echo.Context
}

var _ ports.ListsOutputPort = (*ListsPresenter)(nil)

func NewListsPresenter(ctx echo.Context) *ListsPresenter {
	return &ListsPresenter{ctx}
}

func (p *ListsPresenter) OutputLists(lists []*entities.List) error {
	return nil
}

func (p *ListsPresenter) OutputList(list *entities.List) error {
	return nil
}
