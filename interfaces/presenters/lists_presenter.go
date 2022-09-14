package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ListsPresenter struct{}

var _ ports.ListsOutputPort = (*ListsPresenter)(nil)

func NewListsPresenter() *ListsPresenter {
	return &ListsPresenter{}
}

func (p *ListsPresenter) OutputLists(ctx echo.Context, lists []*entities.List) error {
	return nil
}

func (p *ListsPresenter) OutputList(ctx echo.Context, list *entities.List) error {
	return nil
}
