package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ErrorsPresenter struct {
	ctx echo.Context
}

var _ ports.ErrorsOutputPort = (*ErrorsPresenter)(nil)

func NewErrorsPresenter(ctx echo.Context) *ErrorsPresenter {
	return &ErrorsPresenter{ctx}
}

func (p *ErrorsPresenter) OutputError() error {
	return nil
}
