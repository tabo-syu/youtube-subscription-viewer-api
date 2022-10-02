package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ErrorsPresenter struct {
	c echo.Context
}

var _ ports.ErrorsOutputPort = (*ErrorsPresenter)(nil)

func NewErrorsPresenter(c echo.Context) *ErrorsPresenter {
	return &ErrorsPresenter{c}
}

func (p *ErrorsPresenter) OutputError() error {
	return nil
}
