package ports

import (
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
	Create(*entities.List) error
	GetAll() error
	GetById(entities.ListId) error
	Update(*entities.List) error
	DeleteById(entities.ListId) error
}

type ListsOutputPort interface {
	OutputLists([]*entities.List) error
	OutputList(*entities.List) error
}
