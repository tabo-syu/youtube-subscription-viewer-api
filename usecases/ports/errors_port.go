package ports

import "github.com/labstack/echo/v4"

type ErrorsOutputPort interface {
	OutputError(echo.Context) error
}
