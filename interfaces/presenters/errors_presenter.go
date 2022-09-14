package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ErrorsPresenter struct{}

var _ ports.ErrorsOutputPort = (*ErrorsPresenter)(nil)

func NewErrorsPresenter() *ErrorsPresenter {
	return &ErrorsPresenter{}
}

func (p *ErrorsPresenter) OutputError(ctx echo.Context) error {
	return nil
}
