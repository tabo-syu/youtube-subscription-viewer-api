package presenters

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/usecases/ports"
)

type ErrorsPresenter struct {
	echoCtx echo.Context
}

var _ ports.ErrorsOutputPort = (*ErrorsPresenter)(nil)

func NewErrorsPresenter(echoCtx echo.Context) *ErrorsPresenter {
	return &ErrorsPresenter{echoCtx}
}

func (p *ErrorsPresenter) OutputError() error {
	return nil
}
