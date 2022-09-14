package ports

import (
	"github.com/labstack/echo/v4"
	"github.com/tabo-syu/youtube-subscription-viewer-api/entities"
)

type ListsRepository interface {
	Create(*entities.List) (*entities.List, error)
	GetAll() ([]*entities.List, error)
	GetById(entities.ListId) (*entities.List, error)
	Update(*entities.List) (*entities.List, error)
	DeleteById(entities.ListId) error
	AddChannel(*entities.Channel) (*entities.Channel, error)
	GetAllChannels() ([]*entities.Channel, error)
	GetFeed() ([]*entities.Video, error)
}

type ListsInputPort interface {
	Create(echo.Context, *entities.List) error
	GetAll(echo.Context) error
	GetById(echo.Context, entities.ListId) error
	Update(echo.Context, *entities.List) error
	DeleteById(echo.Context, entities.ListId) error
}

type ListsOutputPort interface {
	OutputLists(echo.Context, []*entities.List) error
	OutputList(echo.Context, *entities.List) error
}
